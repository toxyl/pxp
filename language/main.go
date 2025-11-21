package language

import (
	"image"
)

// Language represents the PixelPipeline Studio language functionality
type Language struct {
	dsl *dslCollection
}

func Shell() {
	dsl.shell()
}

func New() *Language {
	return &Language{
		dsl: NewLanguage(), // Use the generated constructor
	}
}

func (l *Language) Run(script, baseDir string, replacements map[string]string, args ...any) (*dslResult, error) {
	l.dsl.storeState()
	res, err := l.dsl.run(script, baseDir, replacements, false, args...)
	l.dsl.restoreState()
	if err != nil {
		return nil, err
	}
	switch t := res.value.(type) {
	case *image.RGBA64:
		res.value = l.dsl.convertRGBA64ToNRGBA(t)
	case *image.NRGBA64:
		res.value = l.dsl.convertNRGBA64ToNRGBA(t)
	}
	return res, err
}

func DocMarkdown() string                { return dsl.docMarkdown() }
func DocHTML() string                    { return dsl.docHTML() }
func DocText() string                    { return dsl.docText() }
func ExportToVSIX(vsixFile string) error { return dsl.exportVSCodeExtension(vsixFile) }

// GetLanguageDefinition returns the complete language definition including grammar, theme, and snippets
func (l *Language) GetLanguageDefinition() (map[string]interface{}, error) {
	return dsl.GetLanguageDefinition()
}

func ImageTo8Bit(img image.Image) image.Image {
	switch t := img.(type) {
	case *image.RGBA64:
		img = dsl.convertRGBA64ToNRGBA(t)
	case *image.NRGBA64:
		img = dsl.convertNRGBA64ToNRGBA(t)
	}
	return img
}
