package main

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"image"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"image/jpeg"
	"image/png"

	"github.com/gen2brain/heic"
	"github.com/toxyl/flo"
	"github.com/toxyl/math"
	"github.com/toxyl/pxp/language"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"golang.org/x/image/draw"
)

// FileResult represents the result of a file operation
type FileResult struct {
	Name    string `json:"name"`
	Content string `json:"content"`
	Path    string `json:"path"`
}

// OpResult represents the result of a  operation
type OpResult struct {
	Data  any   `json:"data"`
	Error error `json:"error"`
}

// App struct
type App struct {
	ctx             context.Context
	lastRunImage    image.Image // Store the last run image
	cancelRendering bool
	batchProgress   float64 // used for visual feedback during batch processing
	batchLen        int
	reviewQueue     [][2]string // used to store files the user has to review
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{
		ctx:             nil,
		lastRunImage:    nil,
		cancelRendering: false,
		batchProgress:   0,
		batchLen:        0,
		reviewQueue:     [][2]string{},
	}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// OpenFile opens a file dialog and returns the file contents
func (a *App) OpenFile(filter string) OpResult {
	// Open file dialog
	file, err := runtime.OpenFileDialog(a.ctx, runtime.OpenDialogOptions{
		Filters: []runtime.FileFilter{
			{
				DisplayName: filter,
				Pattern:     filter,
			},
		},
	})
	if err != nil {
		return OpResult{nil, fmt.Errorf("failed to open file dialog: %v", err)}
	}

	// Read file contents
	content, err := os.ReadFile(file)
	if err != nil {
		return OpResult{nil, fmt.Errorf("failed to read file: %v", err)}
	}

	return OpResult{&FileResult{
		Name:    filepath.Base(file),
		Content: string(content),
		Path:    file,
	}, nil}
}

// SaveFile saves content to a file
func (a *App) SaveFile(filename string, content string) OpResult {
	// Open save file dialog
	file, err := runtime.SaveFileDialog(a.ctx, runtime.SaveDialogOptions{
		DefaultFilename: filename,
		Filters: []runtime.FileFilter{
			{
				DisplayName: "PixelPipeline Script (*.pxp)",
				Pattern:     "*.pxp",
			},
		},
	})
	if err != nil {
		return OpResult{
			Data:  "",
			Error: fmt.Errorf("failed to open save dialog: %v", err),
		}
	}

	// Write content to file
	err = os.WriteFile(file, []byte(content), 0644)
	if err != nil {
		return OpResult{
			Data:  "",
			Error: fmt.Errorf("failed to write file: %v", err),
		}
	}

	return OpResult{
		Data:  file,
		Error: nil,
	}
}

// LoadImage loads an image file and returns its base64 encoded data
func (a *App) FileToBase64(file string, maxWidth, maxHeight int) OpResult {
	// Read image file
	data, err := os.ReadFile(file)
	if err != nil {
		return OpResult{"", fmt.Errorf("failed to read image: %v", err)}
	}
	ext := filepath.Ext(file)
	if language.DetectImageType(data) == "heic" {
		// our viewers do not support HEIC, so we first need to convert them to PNG
		img, err := heic.Decode(bytes.NewReader(data))
		if err != nil {
			return OpResult{"", err}
		}
		buf := bytes.NewBuffer([]byte{})
		if err := png.Encode(buf, img); err != nil {
			return OpResult{"", err}
		}
		data = buf.Bytes()
		ext = ".png"
	}

	// Scale image if dimensions are provided
	if maxWidth > 0 && maxHeight > 0 {
		// Decode the image
		img, _, err := image.Decode(bytes.NewReader(data))
		if err != nil {
			return OpResult{"", fmt.Errorf("failed to decode image: %v", err)}
		}

		img = language.ImageTo8Bit(img)

		// Get original dimensions
		bounds := img.Bounds()
		origWidth := bounds.Max.X - bounds.Min.X
		origHeight := bounds.Max.Y - bounds.Min.Y

		// Calculate scale ratios
		widthRatio := float64(maxWidth) / float64(origWidth)
		heightRatio := float64(maxHeight) / float64(origHeight)

		// Use the smaller ratio to maintain aspect ratio
		ratio := math.Min(widthRatio, heightRatio)
		if ratio < 1 {
			newWidth := int(float64(origWidth) * ratio)
			newHeight := int(float64(origHeight) * ratio)

			// Create a new scaled image
			scaled := image.NewRGBA(image.Rect(0, 0, newWidth, newHeight))
			draw.CatmullRom.Scale(scaled, scaled.Bounds(), img, bounds, draw.Over, nil)

			// Encode the scaled image
			buf := new(bytes.Buffer)
			if err := png.Encode(buf, scaled); err != nil {
				return OpResult{"", fmt.Errorf("failed to encode scaled image: %v", err)}
			}
			data = buf.Bytes()
			ext = ".png"
		}
	}

	return OpResult{fmt.Sprintf("data:image/%s;base64,%s", ext[1:], base64.StdEncoding.EncodeToString(data)), nil}
}

// LoadMultipleImages opens a file dialog for selecting multiple images and returns their data
func (a *App) LoadMultipleImages() OpResult {
	// Open file dialog for multiple images
	files, err := runtime.OpenMultipleFilesDialog(a.ctx, runtime.OpenDialogOptions{
		Filters: []runtime.FileFilter{
			{
				DisplayName: "Images (*.png;*.jpg;*.jpeg;*.heic)",
				Pattern:     "*.png;*.jpg;*.jpeg;*.heic;*.PNG;*.JPG;*.JPEG;*.HEIC",
			}, {
				DisplayName: "PNG Images (*.png)",
				Pattern:     "*.png;*.PNG",
			}, {
				DisplayName: "JPEG Images (*.jpg;*.jpeg)",
				Pattern:     "*.jpg;*.jpeg;*.JPG;*.JPEG",
			}, {
				DisplayName: "HEIC Images (*.heic)",
				Pattern:     "*.heic;*.HEIC",
			},
		},
	})
	if err != nil {
		return OpResult{nil, fmt.Errorf("failed to open file dialog: %v", err)}
	}

	results := make([]map[string]string, 0, len(files))
	for _, file := range files {
		// Read image file
		data, err := os.ReadFile(file)
		if err != nil {
			return OpResult{nil, fmt.Errorf("failed to read image: %v", err)}
		}
		ext := filepath.Base(file)
		if language.DetectImageType(data) == "heic" {
			// our viewers do not support HEIC, so we first need to convert them to PNG
			img, err := heic.Decode(bytes.NewReader(data))
			if err != nil {
				return OpResult{nil, err}
			}
			buf := bytes.NewBuffer([]byte{})
			if err := png.Encode(buf, img); err != nil {
				return OpResult{nil, err}
			}
			data = buf.Bytes()
			ext = ".png"
		}

		// Create response with full path and base64 data
		result := map[string]string{
			"name": filepath.Base(file),
			"path": file,
			"data": fmt.Sprintf("data:image/%s;base64,%s",
				ext[1:], // Use file extension for MIME type
				base64.StdEncoding.EncodeToString(data)),
		}
		results = append(results, result)
	}

	return OpResult{results, nil}
}

// LoadMultipleImages opens a file dialog for selecting multiple images and returns their data
func (a *App) LoadMultipleInputImages() OpResult {
	// Open file dialog for multiple images
	res, err := runtime.OpenMultipleFilesDialog(a.ctx, runtime.OpenDialogOptions{
		Filters: []runtime.FileFilter{
			{
				DisplayName: "Images (*.png;*.jpg;*.jpeg;*.heic)",
				Pattern:     "*.png;*.jpg;*.jpeg;*.heic;*.PNG;*.JPG;*.JPEG;*.HEIC",
			}, {
				DisplayName: "PNG Images (*.png)",
				Pattern:     "*.png;*.PNG",
			}, {
				DisplayName: "JPEG Images (*.jpg;*.jpeg)",
				Pattern:     "*.jpg;*.jpeg;*.JPG;*.JPEG",
			}, {
				DisplayName: "HEIC Images (*.heic)",
				Pattern:     "*.heic;*.HEIC",
			},
		},
	})
	return OpResult{res, err}
}

// Run runs the script and returns the result image as base64 data
func (a *App) Run(script string, filePaths []string) OpResult {
	args := make([]any, len(filePaths))
	for i, v := range filePaths {
		args[i] = v
	}
	res, err := language.New().Run(string(script), "", nil, args...)
	if err != nil {
		return OpResult{map[string]string{
			"error": err.Error(),
		}, nil}
	}

	// Get the image data from the result
	if res == nil {
		return OpResult{map[string]string{
			"error": "No result returned",
		}, nil}
	}

	// Try to get the image from the result
	var img image.Image
	switch val := res.Value().(type) {
	case *image.RGBA:
		img = val
	case *image.NRGBA:
		img = val
	case *image.RGBA64:
		img = val
	case *image.NRGBA64:
		img = val
	default:
		return OpResult{map[string]string{
			"error": fmt.Sprintf("Unexpected result type: %T", res.Value()),
		}, nil}
	}

	if img == nil {
		return OpResult{map[string]string{
			"error": "No image data returned",
		}, nil}
	}

	// Store the image
	a.lastRunImage = img

	// Get image dimensions
	bounds := img.Bounds()
	width := bounds.Max.X - bounds.Min.X
	height := bounds.Max.Y - bounds.Min.Y
	dimensions := fmt.Sprintf("%dx%d", width, height)

	// Encode the image to PNG and then to base64 for preview
	var buf bytes.Buffer
	if err := png.Encode(&buf, img); err != nil {
		return OpResult{map[string]string{
			"error": fmt.Sprintf("Error encoding image: %v", err),
		}, nil}
	}
	encoded := base64.StdEncoding.EncodeToString(buf.Bytes())
	return OpResult{map[string]string{
		"data":       fmt.Sprintf("data:image/png;base64,%s", encoded),
		"dimensions": dimensions,
	}, nil}
}

func (a *App) CancelRendering() {
	a.cancelRendering = true
	a.batchLen = 0
	a.reviewQueue = [][2]string{}
}

func (a *App) BatchProgress() float64 {
	return a.batchProgress
}

func (a *App) NextReviewQueueItem() OpResult {
	if a.batchLen <= 0 {
		return OpResult{"", fmt.Errorf("no more items to review")}
	}
	for len(a.reviewQueue) == 0 {
		time.Sleep(time.Second) // wait for the item to be processed and added to the queue
	}
	item := a.reviewQueue[0]
	return OpResult{item[0], nil}
}

func (a *App) ApproveReviewQueueItem() bool {
	defer func() { a.batchLen-- }()
	if len(a.reviewQueue) == 0 {
		return false
	}
	item := a.reviewQueue[0]
	a.reviewQueue = a.reviewQueue[1:]
	flo.File(item[0]).Copy(flo.File(item[1]).Path())
	return a.batchLen > 0
}

func (a *App) RejectReviewQueueItem() bool {
	defer func() { a.batchLen-- }()

	if len(a.reviewQueue) == 0 {
		return false
	}
	a.reviewQueue = a.reviewQueue[1:]
	return a.batchLen > 0
}

// RunBatch runs the script for all possible combinations of the given filepaths and returns a map[string]string with all errors encountered
// during the batch process, keyed on filenames.
func (a *App) RunBatch(script, outputDir string, filePaths [][]string, reviewEnabled bool) (errors map[string]string) {
	n := len(filePaths) // determines how many fields each combination must have
	if n == 0 {
		return map[string]string{
			"error": "No input files provided",
		}
	}

	// Generate all possible combinations
	combinations := [][]string{}

	// Helper function to generate combinations recursively
	var generateCombinations func(current []string, depth int)
	generateCombinations = func(current []string, depth int) {
		if depth == n {
			// Make a copy of the current combination
			combination := make([]string, n)
			copy(combination, current)
			combinations = append(combinations, combination)
			return
		}

		// For each file in the current depth's file list
		for _, file := range filePaths[depth] {
			current[depth] = file
			generateCombinations(current, depth+1)
		}
	}

	// Start the combination generation
	current := make([]string, n)
	generateCombinations(current, 0)

	errors = make(map[string]string)
	reExt := regexp.MustCompile(`(?i)(.+?)\.(?:png|jpg|jpeg|heic)`)

	// We use this to store all generated images, so we can remove them if the user cancels
	processed := []string{}
	tempDir, err := os.MkdirTemp("", "pxp-*")
	if err != nil {
		return map[string]string{"all": err.Error()}
	}
	defer flo.Dir(tempDir).Remove()

	a.batchLen = len(combinations)
	a.batchProgress = 0

	lang := language.New()

	l := float64(a.batchLen)
	// Convert []string to []any for each combination and process them
	for i, combo := range combinations {
		if a.cancelRendering {
			break // the user decided to cancel
		}
		suffix := make([]string, len(combo))
		args := make([]any, len(combo))
		for j, v := range combo {
			args[j] = v
			suffix[j] = reExt.ReplaceAllString(filepath.Base(v), "$1")
		}
		fname := fmt.Sprintf("%s.png", strings.Join(suffix, "_-_"))

		res, err := lang.Run(script, "", nil, args...)
		if err != nil {
			errors[fname] = err.Error()
			continue
		}
		a.batchProgress = float64(i) / l
		processed = append(processed, fname)

		// Process the result image
		var img image.Image
		switch val := res.Value().(type) {
		case *image.RGBA:
			img = val
		case *image.NRGBA:
			img = val
		case *image.RGBA64:
			img = val
		case *image.NRGBA64:
			img = val
		default:
			errors[fname] = fmt.Sprintf("Unexpected result type: %T", res.Value())
			continue
		}

		var buf bytes.Buffer
		if err := png.Encode(&buf, img); err != nil {
			errors[fname] = fmt.Sprintf("Error encoding image: %v", err)
			continue
		}
		src := flo.Dir(tempDir).File(fname)
		dst := flo.Dir(outputDir).File(fname)

		if err := src.StoreBytes(buf.Bytes()); err != nil {
			errors[fname] = fmt.Sprintf("Error storing image: %v", err)
			continue
		}
		a.reviewQueue = append(a.reviewQueue, [2]string{src.Path(), dst.Path()})
	}
	if reviewEnabled && a.batchLen > 0 {
		for a.batchLen > 0 {
			// the user is not done reviewing, let's wait
			time.Sleep(time.Second)
		}
	} else {
		// everything rendered
		if !a.cancelRendering {
			tDir, oDir := flo.Dir(tempDir), flo.Dir(outputDir)
			for _, f := range processed {
				tDir.File(f).Copy(oDir.File(f).Path())
			}
		}
	}

	a.cancelRendering = false
	a.batchProgress = 0
	a.batchLen = 0
	a.reviewQueue = [][2]string{}
	return errors
}

// SaveOutput saves the last run image to a file
func (a *App) SaveOutput() OpResult {
	if a.lastRunImage == nil {
		return OpResult{false, fmt.Errorf("no image to save")}
	}

	// Open save file dialog
	file, err := runtime.SaveFileDialog(a.ctx, runtime.SaveDialogOptions{
		Filters: []runtime.FileFilter{
			{
				DisplayName: "Images (*.png;*.jpg;*.jpeg)",
				Pattern:     "*.png;*.jpg;*.jpeg;*.PNG;*.JPG;*.JPEG",
			}, {
				DisplayName: "PNG Images (*.png)",
				Pattern:     "*.png;*.PNG",
			}, {
				DisplayName: "JPEG Images (*.jpg;*.jpeg)",
				Pattern:     "*.jpg;*.jpeg;*.JPG;*.JPEG",
			},
		},
	})
	if err != nil {
		return OpResult{false, fmt.Errorf("failed to open save dialog: %v", err)}
	}

	// Create the output file
	outFile, err := os.Create(file)
	if err != nil {
		return OpResult{false, fmt.Errorf("failed to create output file: %v", err)}
	}
	defer outFile.Close()

	// Save based on file extension
	ext := strings.ToLower(filepath.Ext(file))
	switch ext {
	case ".png":
		if err := png.Encode(outFile, a.lastRunImage); err != nil {
			return OpResult{false, fmt.Errorf("failed to encode PNG: %v", err)}
		}
	case ".jpg", ".jpeg":
		if err := jpeg.Encode(outFile, a.lastRunImage, &jpeg.Options{Quality: 100}); err != nil {
			return OpResult{false, fmt.Errorf("failed to encode JPEG: %v", err)}
		}
	default:
		return OpResult{false, fmt.Errorf("unsupported file format: %s", ext)}
	}

	return OpResult{true, nil}
}

// SelectDirectory opens a directory selection dialog
func (a *App) SelectDirectory() OpResult {
	dir, err := runtime.OpenDirectoryDialog(a.ctx, runtime.OpenDialogOptions{})
	if err != nil {
		return OpResult{"", fmt.Errorf("failed to open directory dialog: %v", err)}
	}
	return OpResult{dir, nil}
}

// ChangeDirectory changes the current working directory
func (a *App) ChangeDirectory(path string) OpResult {
	if err := os.Chdir(path); err != nil {
		return OpResult{false, fmt.Errorf("failed to change directory: %v", err)}
	}
	return OpResult{true, nil}
}
