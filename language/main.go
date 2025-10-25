package language

import "image"

// Language represents the PixelPipeline Studio language functionality
type Language struct{}

func Shell() {
	dsl.shell()
}

func Run(script string, args ...any) (*dslResult, error) {
	dsl.storeState()
	defer dsl.restoreState()
	res, err := dsl.run(script, false, args...)
	if err != nil {
		return nil, err
	}
	switch t := res.value.(type) {
	case *image.RGBA64:
		res.value = dsl.convertRGBA64ToNRGBA(t)
	case *image.NRGBA64:
		res.value = dsl.convertNRGBA64ToNRGBA(t)
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
