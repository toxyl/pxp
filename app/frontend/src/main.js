import './style.css';
import './app.css';
import * as monaco from 'monaco-editor';
import {Run, OpenFile, SaveFile, SaveOutput, RunBatch, SelectDirectory, LoadMultipleInputImages, LoadMultipleImages, CancelRendering, BatchProgress, NextReviewQueueItem, ApproveReviewQueueItem, RejectReviewQueueItem, FileToBase64, ChangeDirectory } from '../wailsjs/go/main/App';
import {GetLanguageDefinition} from '../wailsjs/go/language/Language';
import 'viewerjs/dist/viewer.css';
import Viewer from 'viewerjs';

// Initialize state
const state = {
    tabs: [{ id: 'demo', name: 'new', content: '', savedContent: '' }],
    imageFiles: {},
    activeTab: 'demo',
    imageCounter: 1,
    tabCounter: 1,
    editor: null,
};

function createTabId() {
    return `script_${state.tabCounter++}`;
}

function createTab(name = null, content = '', filePath = null) {
    const id = createTabId();
    const tabName = name || `Script${state.tabCounter}.pxp`;
    return { id, name: tabName, content, savedContent: content, filePath };
}

function hasUnsavedChanges(tab) {
    return tab.content !== tab.savedContent;
}

// Initialize Monaco Editor with custom language
async function initializeMonaco() {
    try {
        // Get the complete language definition from the backend
        const langDef = await GetLanguageDefinition();
        
        // Register the custom language
        monaco.languages.register({ 
            id: 'pixelpipeline-studio',
            extensions: ['.pxp'],
            aliases: ['PixelPipeline Studio', 'pixelpipeline-studio']
        });
        
        // Set the language configuration
        monaco.languages.setLanguageConfiguration('pixelpipeline-studio', langDef.configuration);

        // Convert TextMate grammar to Monaco's format
        const convertedTokenizer = {
            defaultToken: '',
            tokenPostfix: '.pixelpipeline-studio',
            tokenizer: {
                root: [
                    // Comments
                    [/#/, { token: 'comment.block', next: '@comment' }],
                    
                    // Strings
                    [/"/, { token: 'string.quoted.double', next: '@string' }],
                    
                    // Numbers
                    [/\b\d+(\.\d+)?\b/, 'constant.numeric'],
                    
                    // Booleans
                    [/\b(true|false)\b/, 'constant.language.boolean'],
                    
                    // Null
                    [/\bnil\b/, 'constant.language.null'],
                    
                    // Keywords
                    [/\bfor\b/, 'keyword.control'],
                    [/\bdone\b/, 'keyword.control'],
                    [/\binclude\b/, 'keyword.control'],
                    [/\bmacro\b/, 'keyword.control'],
                    [/\bif\b/, 'keyword.control'],
                    [/\belse\b/, 'keyword.control'],
                    [/\bend\b/, 'keyword.control'],
                    
                    // Argument references ($1, $2, etc)
                    [/\$\d+/, 'variable.parameter'],
                    
                    // Variable assignments (left side of :)
                    [/\b[a-zA-Z_][a-zA-Z0-9_]*(?=\s*:)/, 'variable.assign'],
                    
                    // Functions
                    [/\b[a-zA-Z_][a-zA-Z0-9_-]*(?=\s*\()/, 'entity.name.function'],
                    
                    // Named arguments
                    [/[a-zA-Z_][a-zA-Z0-9_]*(?=\s*=)/, 'variable.parameter'],
                    
                    // Operators
                    [/[:=]/, 'keyword.operator.assignment'],
                    
                    // Punctuation
                    [/[(),]/, 'punctuation'],
                    
                    // Variables (used in expressions)
                    [/[a-zA-Z_]\w*/, 'variable.other']
                ],
                comment: [
                    [/#/, { token: 'comment.block', next: '@pop' }],
                    [/[^#]+/, 'comment.block'],
                    [/#/, 'comment.block']
                ],
                string: [
                    [/[^\\"]+/, 'string.quoted.double'],
                    [/\\./, 'constant.character.escape'],
                    [/"/, { token: 'string.quoted.double', next: '@pop' }]
                ]
            }
        };

        // Set the tokenizer rules
        monaco.languages.setMonarchTokensProvider('pixelpipeline-studio', convertedTokenizer);

        // Create editor instance
        state.editor = monaco.editor.create(document.getElementById('codeEditor'), {
            value: '',
            language: 'pixelpipeline-studio',
            theme: 'pixelpipeline-studio-theme',
            automaticLayout: true,
            minimap: {
                enabled: false
            },
            scrollBeyondLastLine: false,
            fontSize: 12,
            tabSize: 4,
            insertSpaces: true,
            wordWrap: 'on',
            renderWhitespace: 'selection',
            rulers: [],
            bracketPairColorization: {
                enabled: true
            },
            formatOnPaste: true,
            formatOnType: true,
            suggestOnTriggerCharacters: true,
            quickSuggestions: {
                other: true,
                comments: true,
                strings: true
            }
        });

        // Get theme colors from the language definition
        const getThemeColor = (scope) => {
            const token = langDef.theme.tokenColors.find(t => t.scope === scope);
            return token ? token.settings.foreground.replace('#', '') : null;
        };

        // Define theme using colors from the language definition
        monaco.editor.defineTheme('pixelpipeline-studio-theme', {
            base: 'vs-dark',
            inherit: true,
            rules: [
                { token: 'comment.block', foreground: getThemeColor('comment.block') || '808080' },
                { token: 'string.quoted.double', foreground: getThemeColor('string.quoted.double') || 'FFD4D4' },
                { token: 'constant.numeric', foreground: getThemeColor('constant.numeric') || 'B5CEA8' },
                { token: 'constant.language.boolean', foreground: getThemeColor('constant.language') || '569CD6' },
                { token: 'constant.language.null', foreground: getThemeColor('constant.language') || '569CD6' },
                { token: 'variable.parameter', foreground: getThemeColor('variable.parameter') || 'B5DCFE' },
                { token: 'variable.assign', foreground: getThemeColor('variable.assign') || 'D7BA7D' },
                { token: 'variable.other', foreground: getThemeColor('variable.other') || 'D4E6FF' },
                { token: 'entity.name.function', foreground: getThemeColor('entity.name.function') || 'FFD700' },
                { token: 'keyword.control', foreground: getThemeColor('keyword.control') || 'C586C0' },
                { token: 'keyword.operator.assignment', foreground: getThemeColor('keyword.operator.assignment') || 'D4D4D4' },
                { token: 'punctuation', foreground: getThemeColor('punctuation') || 'D4D4D4' },
                { token: 'constant.character.escape', foreground: getThemeColor('constant.character.escape') || '808080' }
            ],
            colors: {
                'editor.background': langDef.theme.colors['editor.background'],
                'editor.foreground': langDef.theme.colors['editor.foreground']
            }
        });

        // Set the theme
        monaco.editor.setTheme('pixelpipeline-studio-theme');

        // Register completions provider
        monaco.languages.registerCompletionItemProvider('pixelpipeline-studio', {
            provideCompletionItems: (model, position) => {
                const word = model.getWordUntilPosition(position);
                const range = {
                    startLineNumber: position.lineNumber,
                    endLineNumber: position.lineNumber,
                    startColumn: word.startColumn,
                    endColumn: word.endColumn
                };

                const suggestions = [];

                // Parse functions from the string format
                if (langDef?.completions?.functions) {
                    const functionStr = langDef.completions.functions;
                    const functionCalls = functionStr.match(/this\.addFunction\((.*?)\);/g) || [];
                    
                    functionCalls.forEach(call => {
                        try {
                            // Extract the parameters from the function call
                            const params = call.match(/this\.addFunction\("([^"]+)", "([^"]+)", \[(.*?)\], "([^"]+)"\)/);
                            if (!params) return;

                            const [_, name, description, paramsStr, returnType] = params;
                            
                            // Parse the parameters array
                            const paramList = paramsStr.split('}, {').map(p => {
                                const paramMatch = p.match(/name: "([^"]+)", type: "([^"]+)", description: "([^"]+)"/);
                                if (!paramMatch) return null;
                                const [__, pName, pType, pDesc] = paramMatch;
                                return { name: pName, type: pType, description: pDesc };
                            }).filter(Boolean);

                            suggestions.push({
                                label: {
                                    label: name,
                                    description: paramList.map(p => p.name).join(' '),
                                    detail: ` - ${description}`
                                },
                                kind: monaco.languages.CompletionItemKind.Function,
                                documentation: {
                                    value: [
                                        description,
                                        '',
                                        'Parameters:',
                                        ...paramList.map(p => `- ${p.name} (${p.type}): ${p.description}`),
                                        '',
                                        `Returns: ${returnType}`
                                    ].join('\n')
                                },
                                insertText: `${name}(${paramList.map((p, i) => `\${${i + 1}:${p.name}}`).join(' ')})`,
                                insertTextRules: monaco.languages.CompletionItemInsertTextRule.InsertAsSnippet,
                                range: range
                            });
                        } catch (e) {
                            console.warn('Error parsing function:', e);
                        }
                    });
                }

                // Handle snippets
                if (langDef?.snippets) {
                    Object.entries(langDef.snippets).forEach(([name, snippet]) => {
                        if (!snippet?.prefix || !snippet?.body) return;

                        const body = Array.isArray(snippet.body) ? snippet.body.join('\n') : snippet.body;

                        suggestions.push({
                            label: {
                                label: snippet.prefix,
                                description: name,
                                detail: ` - ${snippet.description || ''}`
                            },
                            kind: monaco.languages.CompletionItemKind.Snippet,
                            documentation: snippet.description || '',
                            insertText: body,
                            insertTextRules: monaco.languages.CompletionItemInsertTextRule.InsertAsSnippet,
                            range: range
                        });
                    });
                }

                return { suggestions };
            },
            triggerCharacters: [':', ' ', '(', '$']
        });

        // Add change listener
        state.editor.onDidChangeModelContent(() => {
            const activeTab = state.tabs.find(tab => tab.id === state.activeTab);
            if (activeTab) {
                activeTab.content = state.editor.getValue();
                renderTabs(); // Update unsaved indicator
            }
        });

    } catch (err) {
        console.error('Error initializing Monaco Editor:', err);
    }
}

function switchTab(tabId) {
    // Save current tab content
    const currentTab = state.tabs.find(tab => tab.id === state.activeTab);
    if (currentTab && state.editor) {
        currentTab.content = state.editor.getValue();
    }

    // Update active tab
    state.activeTab = tabId;
    const newTab = state.tabs.find(tab => tab.id === tabId);
    if (newTab && state.editor) {
        state.editor.setValue(newTab.content);
        
        // Change working directory if file path is available
        if (newTab.filePath) {
            const dir = newTab.filePath.substring(0, newTab.filePath.lastIndexOf('/') || newTab.filePath.lastIndexOf('\\'));
            ChangeDirectory(dir);
        }
    }

    // Update tab UI
    document.querySelectorAll('.script-tabs .tab').forEach(tab => {
        tab.classList.remove('active');
        if (tab.dataset.tabId === tabId) {
            tab.classList.add('active');
        }
    });
}

function renderTabs() {
    const tabsContainer = document.querySelector('.script-tabs');
    tabsContainer.innerHTML = state.tabs.map(tab => {
        const unsavedIndicator = hasUnsavedChanges(tab) ? '<span class="unsaved-indicator">●</span>' : '';
        const showCloseButton = state.tabs.length > 1; // Show close button if there's more than one tab
        return `
            <div class="tab ${tab.id === state.activeTab ? 'active' : ''}" 
                 data-tab-id="${tab.id}">
                ${unsavedIndicator}${tab.name}
                ${showCloseButton ? '<span class="close-tab">×</span>' : ''}
            </div>
        `;
    }).join('') + '<div class="new-tab-btn">+</div>';

    // Add event listeners to tabs
    tabsContainer.querySelectorAll('.tab').forEach(tab => {
        tab.addEventListener('click', (e) => {
            if (e.target.classList.contains('close-tab')) {
                const tabId = tab.dataset.tabId;
                closeTab(tabId);
            } else {
                switchTab(tab.dataset.tabId);
            }
        });
    });

    // Add new tab button listener
    tabsContainer.querySelector('.new-tab-btn').addEventListener('click', addNewTab);
}

function closeTab(tabId) {
    const index = state.tabs.findIndex(tab => tab.id === tabId);
    if (index === -1) return;

    const tab = state.tabs[index];
    if (hasUnsavedChanges(tab)) {
        if (!confirm(`${tab.name} has unsaved changes. Close anyway?`)) {
            return;
        }
    }

    // Remove the tab
    state.tabs.splice(index, 1);

    // If it was the last tab, create a new one
    if (state.tabs.length === 0) {
        const newTab = createTab();
        state.tabs.push(newTab);
        state.activeTab = newTab.id;
    }
    // If closing active tab, switch to another tab
    else if (tabId === state.activeTab) {
        // Try to switch to the tab at the same index, or the last tab if we're at the end
        const newIndex = Math.min(index, state.tabs.length - 1);
        state.activeTab = state.tabs[newIndex].id;
    }

    renderTabs();
    // Explicitly update editor content after tab switch
    const newActiveTab = state.tabs.find(tab => tab.id === state.activeTab);
    if (newActiveTab && state.editor) {
        state.editor.setValue(newActiveTab.content);
    }
}

function addNewTab(name = null, content = '', filePath = null) {
    // If name is an object (from click event) or not provided, create a new file
    const isNewFile = typeof name !== 'string' || !name;
    const tabName = isNewFile ? `Script${state.tabCounter}.pxp` : name;
    const tabContent = isNewFile ? '' : content;
    
    const newTab = createTab(tabName, tabContent, filePath);
    state.tabs.push(newTab);
    state.activeTab = newTab.id;  // Set as active tab immediately
    
    // Change working directory if file path is provided
    if (filePath) {
        const dir = filePath.substring(0, filePath.lastIndexOf('/') || filePath.lastIndexOf('\\'));
        ChangeDirectory(dir);
    }
    
    // Update the editor with the new content
    if (state.editor) {
        state.editor.setValue(tabContent);
    }
    
    renderTabs();
    return newTab;
}


// Update the HTML template
document.querySelector('#app').innerHTML = `
    <div class="app-container">
        <div class="main-content">
            <div class="left-section">
                <div class="editor-section">
                    <div class="editor-header">
                        <span>SCRIPT</span>
                        <div class="editor-header-buttons">
                            <button class="btn" id="openBtn">OPEN</button>
                            <button class="btn" id="saveBtn">SAVE</button>
                            <button class="btn" id="renderBtn">RENDER</button>
                            <button class="btn" id="batchBtn">BATCH</button>
                        </div>
                    </div>
                    <div class="script-tabs"></div>
                    <div class="editor-container" id="codeEditor"></div>
                </div>
                <div class="splitter horizontal"></div>
                <div class="input-images-section">
                    <div class="preview-header">
                        <span>INPUT IMAGES</span>
                        <div class="preview-header-buttons">
                            <div class="loading-indicator" id="addImageLoadingIndicator">Loading...</div>
                            <button class="btn" id="addImageBtn">ADD</button>
                        </div>
                    </div>
                    <div class="preview-toolbar">
                        <div class="image-tabs" id="imageTabs"></div>
                    </div>
                    <div class="preview-container">
                        <div class="image-preview" id="imagePreview"></div>
                    </div>
                </div>
            </div>
            <div class="splitter vertical"></div>
            <div class="output-section">
                <div class="output-container">
                    <div class="output-header">
                        <span>OUTPUT</span>
                        <div class="output-header-buttons">
                            <div class="loading-indicator" id="loadingIndicator">Rendering...</div>
                            <button class="btn" id="saveOutputBtn">SAVE</button>
                        </div>
                    </div>
                    <div class="output-preview" id="outputPreview"></div>
                </div>
            </div>
        </div>
    </div>
`;

// Initialize elements and editor
document.addEventListener('DOMContentLoaded', async () => {
    // Initialize Monaco Editor
    await initializeMonaco();

    // Initialize tabs
    renderTabs();

    // Add horizontal scroll with mouse wheel
    const scriptTabs = document.querySelector('.script-tabs');
    scriptTabs.addEventListener('wheel', (e) => {
        if (e.deltaY !== 0) {
            e.preventDefault();
            scriptTabs.scrollLeft += e.deltaY;
        }
    });

    // Initialize other elements
    const renderBtn = document.getElementById('renderBtn');
    const batchBtn = document.getElementById('batchBtn');
    const saveBtn = document.getElementById('saveBtn');
    const openBtn = document.getElementById('openBtn');
    const addImageBtn = document.getElementById('addImageBtn');
    const addImageLoadingIndicator = document.getElementById('addImageLoadingIndicator');
    const imagePreview = document.getElementById('imagePreview');
    const imageTabs = document.getElementById('imageTabs');
    const loadingIndicator = document.getElementById('loadingIndicator');
    const saveOutputBtn = document.getElementById('saveOutputBtn');

    // Set initial button states
    updateButtonStates();

    // Add change listener to editor for batch button state
    state.editor.onDidChangeModelContent(() => {
        updateButtonStates();
    });

    // Add save functionality for output
    saveOutputBtn.addEventListener('click', async () => {
        const res = await SaveOutput();
        if (res.error) {
            console.error('Error saving output:', res.error);
        }
        if (!res.data) {
            console.error('Failed to save output');
        }
    });

    // Initialize splitters
    const verticalSplitter = document.querySelector('.splitter.vertical');
    const horizontalSplitter = document.querySelector('.splitter.horizontal');
    const leftSection = document.querySelector('.left-section');
    const inputImagesSection = document.querySelector('.input-images-section');

    // Helper function for vertical splitting
    function handleVerticalSplit(e, startX, startWidth) {
        const deltaX = e.clientX - startX;
        return startWidth + deltaX;  // Moving right increases width
    }

    // Helper function for horizontal splitting
    function handleHorizontalSplit(e, startY, startHeight) {
        const deltaY = e.clientY - startY;
        return startHeight - deltaY;  // Moving down decreases height
    }

    // Vertical splitter
    verticalSplitter.addEventListener('mousedown', (e) => {
        e.preventDefault();
        const startX = e.clientX;
        const startWidth = leftSection.offsetWidth;

        const onMouseMove = (e) => {
            const width = handleVerticalSplit(e, startX, startWidth);
            leftSection.style.width = `${width}px`;
            if (state.editor) {
                state.editor.layout();
            }
            // Trigger window resize for ViewerJS
            window.dispatchEvent(new Event('resize'));
        };

        const onMouseUp = () => {
            document.removeEventListener('mousemove', onMouseMove);
            document.removeEventListener('mouseup', onMouseUp);
        };

        document.addEventListener('mousemove', onMouseMove);
        document.addEventListener('mouseup', onMouseUp);
    });

    // Horizontal splitter
    horizontalSplitter.addEventListener('mousedown', (e) => {
        e.preventDefault();
        const startY = e.clientY;
        const startHeight = inputImagesSection.offsetHeight;

        const onMouseMove = (e) => {
            const height = handleHorizontalSplit(e, startY, startHeight);
            inputImagesSection.style.height = `${height}px`;
            if (state.editor) {
                state.editor.layout();
            }
            // Trigger window resize for ViewerJS
            window.dispatchEvent(new Event('resize'));
        };

        const onMouseUp = () => {
            document.removeEventListener('mousemove', onMouseMove);
            document.removeEventListener('mouseup', onMouseUp);
        };

        document.addEventListener('mousemove', onMouseMove);
        document.addEventListener('mouseup', onMouseUp);
    });

    // Setup event listeners
    openBtn.addEventListener('click', async () => {
        const res = await OpenFile("*.pxp");
        if (res.error) {
            console.error('Error opening file:', res.error);
        } else if (res.data && res.data.content) {
            // Create new tab with the file name, content, and full path
            addNewTab(res.data.name, res.data.content, res.data.path);
        }
    });

    function getDirectory(path) {
        const lastSlash = Math.max(path.lastIndexOf('/'), path.lastIndexOf('\\'));
        return lastSlash > 0 ? path.substring(0, lastSlash) : path;
    }

    saveBtn.addEventListener('click', async () => {
        const activeTab = state.tabs.find(tab => tab.id === state.activeTab);
        if (!activeTab || !state.editor) return;

        const content = state.editor.getValue();
        const res = await SaveFile(activeTab.name, content);
        if (res.error) {
            console.error('Error saving file:', res.error);
        } else if (res.data) { 
            const pathParts = res.data.split(/[/\\]/);
            activeTab.name = pathParts[pathParts.length - 1];
            activeTab.filePath = res.data;
            activeTab.content = content;
            activeTab.savedContent = content;
            // Change working directory to the saved file's directory
            const dir = getDirectory(res.data);
            ChangeDirectory(dir);
            renderTabs(); // Update unsaved indicator
        }
    });

    renderBtn.addEventListener('click', async () => {
        if (!state.editor) return;
        
        const script = state.editor.getValue();
        if (!script) return;

        try {
            // Disable render button and show loading indicator
            renderBtn.disabled = true;
            loadingIndicator.classList.add('active');

            // Create an array of absolute file paths from the image files
            const filePaths = Object.entries(state.imageFiles).map(([index, file]) => file.name);
            
            // Update the output preview with the result
            var res = await Run(script, filePaths);
            updateOutputPreview(res.data);
            if (res.error) {
                console.error('Error running script:', res.error);
                updateOutputPreview(null, `Runtime Error: ${res.error.toString()}`);
            }
        } catch (err) {
            console.error('Error running script:', err);
            updateOutputPreview(null, `Runtime Error: ${err.toString()}`);
        } finally {
            // Re-enable render button and hide loading indicator
            renderBtn.disabled = false;
            loadingIndicator.classList.remove('active');
        }
    });

    batchBtn.addEventListener('click', async () => {
        if (!state.editor) return;
        
        const script = state.editor.getValue();
        if (!script) return;

        try {
            // Disable render batch button and show loading indicator
            batchBtn.disabled = true;
            loadingIndicator.classList.add('active');

            // Create dialog element
            const dialog = document.createElement('dialog');
            dialog.innerHTML = `
                <form method="dialog">
                    <div>
                        <label for="outputDir">Output Directory</label>
                        <div style="display: flex; gap: 8px;">
                            <input type="text" id="outputDir" placeholder="Select output directory" required readonly>
                            <button type="button" id="selectDir">Select...</button>
                        </div>
                    </div>
                    <label>Image Inputs</label>
                    <div id="placeholderInputs"></div>
                    <div class="batch-errors"></div>
                    <div class="button-bar">
                        <input type="checkbox">Review results
                        <button type="button" value="cancel">Cancel</button>
                        <button type="submit" value="confirm">Start</button>
                        <div class="processing-overlay">
                            <div class="spinner"></div>
                            <div class="progress-label">0% processed</div>
                        </div>
                    </div>
                </form>
            `;
            document.body.appendChild(dialog);

            // Find all $n placeholders in script
            const placeholders = [...script.matchAll(/\$(\d+)/g)]
                .map(match => parseInt(match[1]))
                .filter((value, index, self) => self.indexOf(value) === index)
                .sort((a, b) => a - b);

            // Create file inputs for each placeholder
            const inputsDiv = dialog.querySelector('#placeholderInputs');
            placeholders.forEach(n => {
                const div = document.createElement('div');
                div.innerHTML = `
                <div class="placeholderFiles">
                    <label for="files${n}">$${n}</label>
                    <input type="text" id="files${n}" placeholder="Select input files" readonly>
                    <button type="button" class="select-files" data-placeholder="${n}">Select...</button>
                </div>
                <div class="selected-files" style="display: none" id="selected-files-${n}"></div>
                `;
                inputsDiv.appendChild(div);
            });

            // Store selected files for each placeholder
            const selectedFiles = {};

            // Add click handlers for file selection buttons 
            dialog.querySelectorAll('.select-files').forEach(button => {
                button.addEventListener('click', async () => {
                    const placeholder = button.dataset.placeholder;
                    const res = await LoadMultipleInputImages();
                    if (res.error) {
                        console.error('Error selecting files:', err);
                    } else if (res.data && res.data.length > 0) {
                        selectedFiles[placeholder] = res.data;
                        const input = dialog.querySelector(`#files${placeholder}`);
                        const selectedFilesDiv = dialog.querySelector(`#selected-files-${placeholder}`);
                        input.value = `${res.data.length} files selected`;
                        selectedFilesDiv.innerHTML = res.data.map(f => f.split(/[\\/]/).pop()).join('<br>');
                    }
                });
            });

            let outputDir = '';
            const dirInput = dialog.querySelector('#outputDir');
            const dirBtn = dialog.querySelector('#selectDir');
            
            // Add click handler for directory selection
            dirBtn.addEventListener('click', async () => {
                const res = await SelectDirectory();
                if (res.error) {
                    console.error('Error selecting directory:', err);
                } else {
                    outputDir = res.data;
                    dirInput.value = res.data;
                }
            });

            // Show dialog
            dialog.showModal();

            // Handle form submission and cancellation
            const form = dialog.querySelector('form');
            const cancelBtn = dialog.querySelector('button[value="cancel"]');
            const startBtn = dialog.querySelector('button[value="confirm"]');               
            startBtn.disabled = false;
            startBtn.style = '';
            
            // Handle cancel button click
            cancelBtn.addEventListener('click', () => {
                CancelRendering();
                document.body.removeChild(dialog);
                throw new Error('Batch render cancelled');
            });

            // Handle form submission
            form.addEventListener('submit', async (e) => {
                e.preventDefault(); // Prevent form from closing dialog
                
                // Get selected files for each placeholder
                const filePaths = placeholders.map(n => selectedFiles[n] || []);
                
                if (!outputDir) {
                    outputDir = dirInput.value; // Get final value before checking
                    if (!outputDir) {
                        throw new Error('Output directory path is required');
                    }
                }
                
                // Show processing state
                dialog.classList.add('processing');
                const errorsContainer = dialog.querySelector('.batch-errors');
                errorsContainer.classList.remove('has-errors');
                errorsContainer.innerHTML = '';
                startBtn.disabled = true;
                startBtn.style = 'display:none';
                let reviewDialog = null;

                // Function to load next review item
                const loadNextReviewItem = async () => {
                    const resNext = await NextReviewQueueItem();
                    if (resNext.error) {
                        // All done, close dialogs
                        document.body.removeChild(reviewDialog);
                        document.body.removeChild(dialog);
                        return;
                    }                    
                    // Load images
                    const container = reviewDialog.querySelector('#reviewImage');
                    let res = await FileToBase64(resNext.data, reviewDialog.clientWidth, reviewDialog.clientHeight);
                    if (!res.error) {
                        container.innerHTML = `<img src="${res.data}" style="max-width: 100%; max-height: 100%; object-fit: contain;">`;
                    } else {
                        container.innerHTML = `<div style="max-width: 100%; max-height: 100%; object-fit: contain;"><h2>Error</h2>${res.error}</div>`;
                    }
                };
                
                try {
                    // Create review dialog if review is enabled
                    const reviewCheckbox = dialog.querySelector('input[type="checkbox"]');
                    if (reviewCheckbox && reviewCheckbox.checked) {
                        reviewDialog = document.createElement('dialog');            
                        reviewDialog.className = "review-dialog";
                        reviewDialog.innerHTML = ` 
                            <div class="review-container">
                                <div class="review-images">
                                    <div class="review-image">
                                        <div class="image-container" id="reviewImage"><div style="max-width: 100%; max-height: 100%; object-fit: contain;">Waiting for the next image to review...</div></div>
                                    </div>
                                </div>
                                <div class="review-buttons">
                                    <button type="button" id="cancelBtn">Cancel</button>
                                    <button type="button" id="rejectBtn">Reject</button>
                                    <button type="button" id="acceptBtn">Accept</button>
                                </div>
                            </div>
                        `;
                        document.body.appendChild(reviewDialog);

                        
                        // Add button handlers
                        reviewDialog.querySelector('#cancelBtn').addEventListener('click', async () => {
                            await CancelRendering();
                            document.body.removeChild(reviewDialog);
                            document.body.removeChild(dialog);
                        });
                        
                        reviewDialog.querySelector('#acceptBtn').addEventListener('click', async () => {
                            let hasMore = await ApproveReviewQueueItem();
                            if (!hasMore) {
                                document.body.removeChild(reviewDialog);
                                document.body.removeChild(dialog);
                                return;
                            }
                            const container = reviewDialog.querySelector('#reviewImage');
                            container.innerHTML = `<div style="max-width: 100%; max-height: 100%; object-fit: contain;">Waiting for the next image to review...</div>`;
                            await loadNextReviewItem();
                        });
                        
                        reviewDialog.querySelector('#rejectBtn').addEventListener('click', async () => {
                            let hasMore = await RejectReviewQueueItem();
                            if (!hasMore) {
                                document.body.removeChild(reviewDialog);
                                document.body.removeChild(dialog);
                                return;
                            }
                            const container = reviewDialog.querySelector('#reviewImage');
                            container.innerHTML = `<div style="max-width: 100%; max-height: 100%; object-fit: contain;">Waiting for the next image to review...</div>`;
                            await loadNextReviewItem();
                        });

                    }

                    // Run the batch process in the background
                    const batchPromise = RunBatch(script, outputDir, filePaths, reviewDialog != null);
                    var showingReviewDialog = false;

                    // Start polling progress
                    const progressInterval = setInterval(async () => {
                        const progress = await BatchProgress();
                        const progressLabel = dialog.querySelector('.progress-label');
                        progressLabel.textContent = `${Math.round(progress * 100)}% processed`;
                        if (!showingReviewDialog) {
                            showingReviewDialog = true;
                            reviewDialog.showModal();
                            await loadNextReviewItem();
                        }
                    }, 1000); // Poll every second
                    
                    // Wait for batch to complete
                    const response = await batchPromise;
                    clearInterval(progressInterval); // Stop polling
                    
                    // Check for errors in the result
                    if (response && response.result && Object.keys(response.result).length > 0) {
                        errorsContainer.classList.add('has-errors');
                        errorsContainer.innerHTML = Object.entries(response.result)
                            .map(([file, error]) => `
                                <div class="batch-error-item">
                                    <strong>${file}:</strong> ${error}
                                </div>
                            `).join('');
                        
                        // Add a close button at the bottom
                        errorsContainer.innerHTML += `
                            <div style="text-align: right; margin-top: 16px;">
                                <button type="button" onclick="this.closest('dialog').remove()">Close</button>
                            </div>
                        `;
                        
                        // Remove processing state but keep dialog open to show errors
                        dialog.classList.remove('processing');
                    } else if (!reviewDialog) {
                        // No errors and no review dialog, close the dialog
                        document.body.removeChild(dialog);
                    }
                } catch (err) {
                    // Show error in the errors container
                    errorsContainer.classList.add('has-errors');
                    errorsContainer.innerHTML = `
                        <div class="batch-error-item">
                            ${err.toString()}
                        </div>
                        <div style="text-align: right; margin-top: 16px;">
                            <button type="button" onclick="this.closest('dialog').remove()">Close</button>
                        </div>
                    `;
                    dialog.classList.remove('processing');
                }
            });
        } catch (err) {
            console.error('Error running script:', err);
            updateOutputPreview(null, `Runtime Error: ${err.toString()}`);
        } finally {
            // Re-enable render batch button and hide loading indicator
            batchBtn.disabled = false;
            loadingIndicator.classList.remove('active');
        }
    });

    // Helper function to create image tabs
    function createImageTab(name, index, imageData) {
        const tab = document.createElement('div');
        tab.className = 'tab';
        tab.title = name; // This will show the filename in the tooltip
        tab.dataset.imageId = index;
        tab.draggable = true;

        // Create thumbnail image
        const img = document.createElement('img');
        img.src = imageData;
        img.draggable = false; // Prevent image dragging
        tab.appendChild(img);

        // Create ID label
        const label = document.createElement('div');
        label.className = 'tab-id';
        label.textContent = `$${index}`;
        tab.appendChild(label);

        // Create close button
        const closeBtn = document.createElement('span');
        closeBtn.className = 'close-tab';
        closeBtn.textContent = '×';
        closeBtn.onclick = (e) => {
            e.stopPropagation();
            removeImageTab(index);
        };
        tab.appendChild(closeBtn);

        // Add click handler for selection
        tab.onclick = (e) => {
            if (!isDragging) {
                const currentId = parseInt(tab.dataset.imageId);
                const imageData = state.imageFiles[currentId].data;
                updateImagePreview(imageData);
                
                // Update active state
                document.querySelectorAll('.image-tabs .tab').forEach(t => t.classList.remove('active'));
                tab.classList.add('active');
            }
        };

        // Add drag and drop event listeners
        tab.addEventListener('dragstart', handleDragStart);
        tab.addEventListener('dragend', handleDragEnd);
        tab.addEventListener('dragover', handleDragOver);
        tab.addEventListener('dragenter', handleDragEnter);
        tab.addEventListener('dragleave', handleDragLeave);
        tab.addEventListener('drop', handleDrop);

        return tab;
    }

    // Drag and drop handlers
    let draggedTab = null;
    let dragStartX = 0;
    let isDragging = false;

    function handleDragStart(e) {
        if (!isDragging) {
            draggedTab = this;
            dragStartX = e.clientX;
            isDragging = true;
            this.classList.add('dragging');
            e.dataTransfer.effectAllowed = 'move';
            e.dataTransfer.setData('text/plain', this.dataset.imageId);
            
            // Remove any lingering visual states
            document.querySelectorAll('.image-tabs .tab').forEach(tab => {
                tab.classList.remove('drag-over');
                tab.classList.remove('dragging');
            });
        }
    }

    function handleDragEnd(e) {
        if (isDragging) {
            // Clean up all drag-related classes
            document.querySelectorAll('.image-tabs .tab').forEach(tab => {
                tab.classList.remove('drag-over');
                tab.classList.remove('dragging');
            });
            draggedTab = null;
            isDragging = false;
        }
    }

    function handleDragOver(e) {
        if (e.preventDefault) {
            e.preventDefault();
        }
        e.dataTransfer.dropEffect = 'move';
        return false;
    }

    function handleDragEnter(e) {
        if (isDragging && this !== draggedTab) {
            this.classList.add('drag-over');
        }
    }

    function handleDragLeave(e) {
        this.classList.remove('drag-over');
    }

    function handleDrop(e) {
        e.stopPropagation();
        e.preventDefault();
        
        if (draggedTab !== this && isDragging) {
            // Get all tabs and convert to array for easier manipulation
            const tabs = Array.from(document.querySelectorAll('.image-tabs .tab'));
            const draggedIndex = tabs.indexOf(draggedTab);
            const droppedIndex = tabs.indexOf(this);
            
            // Only reorder if actually moved
            if (draggedIndex !== droppedIndex) {
                // Reorder the tabs in the DOM
                if (draggedIndex < droppedIndex) {
                    this.parentNode.insertBefore(draggedTab, this.nextSibling);
                } else {
                    this.parentNode.insertBefore(draggedTab, this);
                }
                
                // Update the tab IDs
                updateTabIds();
            }
        }
        
        // Clean up all drag-related classes
        document.querySelectorAll('.image-tabs .tab').forEach(tab => {
            tab.classList.remove('drag-over');
            tab.classList.remove('dragging');
        });
        
        isDragging = false;
        draggedTab = null;
        return false;
    }

    // Function to update tab IDs after reordering
    function updateTabIds() {
        const tabs = document.querySelectorAll('.image-tabs .tab');
        const newImageFiles = {};
        
        tabs.forEach((tab, index) => {
            const oldId = parseInt(tab.dataset.imageId);
            const newId = index + 1;
            
            // Update the tab's ID and label
            tab.dataset.imageId = newId;
            const label = tab.querySelector('.tab-id');
            label.textContent = `$${newId}`;
            
            // Update the imageFiles object
            newImageFiles[newId] = state.imageFiles[oldId];
            
            // Update the close button event listener
            const closeBtn = tab.querySelector('.close-tab');
            closeBtn.onclick = (e) => {
                e.stopPropagation();
                removeImageTab(newId);
            };
        });
        
        // Replace the old imageFiles with the reordered one
        state.imageFiles = newImageFiles;
        
        // Update the preview for the active tab
        const activeTab = document.querySelector('.image-tabs .tab.active');
        if (activeTab) {
            const activeId = parseInt(activeTab.dataset.imageId);
            const imageData = state.imageFiles[activeId]?.data;
            if (imageData) {
                updateImagePreview(imageData);
            }
        }
    }

    // Helper function to update image preview
    function updateImagePreview(imageData) {
        const allTabs = document.querySelectorAll('.image-tabs .tab');
        allTabs.forEach(tab => tab.classList.remove('active'));
        
        if (!imageData) {
            imagePreview.innerHTML = '<div class="no-image">No image selected</div>';
            return;
        }
        
        // Find and activate the corresponding tab
        const activeTab = Array.from(allTabs).find(tab => {
            const imageId = parseInt(tab.dataset.imageId);
            return state.imageFiles[imageId]?.data === imageData;
        });
        if (activeTab) {
            activeTab.classList.add('active');
        }
        
        const img = document.createElement('img');
        img.src = imageData;
        img.style.maxWidth = '100%';
        img.style.maxHeight = '100%';
        img.style.objectFit = 'contain';
        
        imagePreview.innerHTML = '';
        imagePreview.appendChild(img);
    }

    // Update the addImageBtn click handler
    addImageBtn.addEventListener('click', async () => {
        try {
            // Disable button and show loading indicator
            addImageBtn.disabled = true;
            addImageLoadingIndicator.classList.add('active');
            
            const res = await LoadMultipleImages();
            if (res.data && res.data.length > 0) {
                // Get the next available index based on existing tabs
                const existingTabs = document.querySelectorAll('.image-tabs .tab');
                let index = existingTabs.length + 1;
                
                // Process each selected file
                for (const result of res.data) {
                    state.imageFiles[index] = {
                        name: result.path,
                        data: result.data,
                        displayName: result.name
                    };
                    
                    const tab = createImageTab(result.name, index, result.data);
                    imageTabs.appendChild(tab);
                    
                    // Remove active class from all tabs and activate the new one
                    document.querySelectorAll('.image-tabs .tab').forEach(t => t.classList.remove('active'));
                    tab.classList.add('active');
                    
                    // Show the newly added image
                    updateImagePreview(result.data);
                    index++;
                }
                
                updateButtonStates(); // Update button states after adding images
            }
        } catch (err) {
            console.error('Error loading images:', err);
        } finally {
            // Re-enable button and hide loading indicator
            addImageBtn.disabled = false;
            addImageLoadingIndicator.classList.remove('active');
        }
    });

    // Initialize with empty preview
    updateImagePreview(null);
    // Initialize with empty output
    updateOutputPreview(null);
});

// Helper function to update output preview
function updateOutputPreview(result, error = null) {
    const outputPreview = document.getElementById('outputPreview');
    const saveOutputBtn = document.getElementById('saveOutputBtn');
    
    // Proper cleanup of any existing viewer
    if (outputPreview._viewer) {
        outputPreview._viewer.destroy();
        outputPreview._viewer = null;
    }
    outputPreview.innerHTML = '';
    
    // Disable save button initially
    saveOutputBtn.disabled = true;
    
    if (error) {
        outputPreview.innerHTML = `
            <div class="error-message">
                <div class="error-title">Error</div>
                ${error}
            </div>`;
        return;
    }
    
    if (!result) {
        outputPreview.innerHTML = '<div class="no-image">No output yet</div>';
        return;
    }

    if (result.error) {
        outputPreview.innerHTML = `<div class="error-message"><div class="error-title">Render Error</div>${result.error}</div>`;
        return;
    }
    
    // Create a dedicated viewer container
    const viewerContainer = document.createElement('div');
    viewerContainer.className = 'viewer-container';
    viewerContainer.style.width = '100%';
    viewerContainer.style.height = 'calc(100% - 40px)'; // Subtract header height
    viewerContainer.style.position = 'absolute';
    viewerContainer.style.top = '40px';
    viewerContainer.style.bottom = '0';
    viewerContainer.style.zIndex = '1'; // Lower than header's z-index
    outputPreview.appendChild(viewerContainer);

    // Create dimensions tooltip
    const dimensionsTooltip = document.createElement('div');
    dimensionsTooltip.className = 'dimensions-tooltip';
    dimensionsTooltip.textContent = result.dimensions || '';
    dimensionsTooltip.style.display = 'none';
    outputPreview.appendChild(dimensionsTooltip);

    // Create the image element inside the viewer container
    const img = document.createElement('img');
    img.src = result.data;
    img.style.display = 'none'; // Hide the original image
    viewerContainer.appendChild(img);

    // Initialize viewer on the container
    outputPreview._viewer = new Viewer(img, {
        inline: true,
        container: viewerContainer,
        navbar: false,
        title: false,
        backdrop: false,
        button: false,
        toolbar: {
            zoomIn: true,
            zoomOut: true,
            oneToOne: true,
            reset: true,
            rotateLeft: false,
            rotateRight: false,
            flipHorizontal: false,
            flipVertical: false,
        },
        viewed() {
            img.style.display = 'none'; // Ensure original image stays hidden
            // Enable save button when viewer is ready
            saveOutputBtn.disabled = false;
        },
    });

    // Add hover event handlers for dimensions tooltip
    outputPreview.addEventListener('mouseenter', () => {
        if (dimensionsTooltip.textContent) {
            dimensionsTooltip.style.display = 'block';
        }
    });

    outputPreview.addEventListener('mouseleave', () => {
        dimensionsTooltip.style.display = 'none';
    });
}

// Function to remove an image tab
function removeImageTab(index) {
    // Remove from state
    delete state.imageFiles[index];
    
    // Remove from DOM
    const tab = document.querySelector(`.image-tabs .tab[data-image-id="${index}"]`);
    if (tab) {
        tab.remove();
    }
    
    // Update remaining tab IDs and labels
    const tabs = document.querySelectorAll('.image-tabs .tab');
    const newImageFiles = {};
    
    tabs.forEach((tab, newIndex) => {
        const oldId = parseInt(tab.dataset.imageId);
        const newId = newIndex + 1;
        
        // Update the tab's ID and label
        tab.dataset.imageId = newId;
        const label = tab.querySelector('.tab-id');
        label.textContent = `$${newId}`;
        
        // Update the imageFiles object
        newImageFiles[newId] = state.imageFiles[oldId];
        
        // Update the close button event listener
        const closeBtn = tab.querySelector('.close-tab');
        closeBtn.onclick = (e) => {
            e.stopPropagation();
            removeImageTab(newId);
        };
    });
    
    // Replace the old imageFiles with the reordered one
    state.imageFiles = newImageFiles;
    
    // Update preview
    const activeTab = document.querySelector('.image-tabs .tab.active');
    if (activeTab) {
        const activeId = parseInt(activeTab.dataset.imageId);
        const imageData = state.imageFiles[activeId]?.data;
        updateImagePreview(imageData);
    } else {
        updateImagePreview(null);
    }
    
    // Update button states
    updateButtonStates();
}

// Function to update button states
function updateButtonStates() {
    const renderBtn = document.getElementById('renderBtn');
    const batchBtn = document.getElementById('batchBtn');
    const saveOutputBtn = document.getElementById('saveOutputBtn');
    const outputPreview = document.getElementById('outputPreview');
    
    // Update RENDER button state
    const script = state.editor ? state.editor.getValue().trim() : '';
    const hasImageFiles = Object.keys(state.imageFiles).length > 0;
    const hasLoadFunction = /load\s*\(/.test(script);
    renderBtn.disabled = !state.editor || script === '' || (!hasImageFiles && !hasLoadFunction);
    
    // Update BATCH button state
    batchBtn.disabled = !state.editor || state.editor.getValue().trim() === '';
    
    // Update SAVE button state
    saveOutputBtn.disabled = !outputPreview._viewer;
}
