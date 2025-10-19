package pxp

import (
	"bytes"
	"fmt"
	"image"
	"image/png"
	"os"
	"path/filepath"
	"strings"

	"github.com/toxyl/pxp/language"
)

// RenderFile loads `pathIn` into the variable `in`, processes it with the given `script` and stores the result in `pathOut`.
//
// The script must use the variable `in` as input image and store the final result in the `img` variable.
//
// The `pathIn` variable can be a URL or a local file path.
//
// When `maxW` and `maxH` are greater than zero, the output image will be resized to fit within the given sizes.
func RenderFile(script, pathIn, pathOut string, maxW, maxH int) (*image.NRGBA, error) {
	sb := strings.Builder{}
	sb.WriteString(`in: load("` + pathIn + `")` + "\n")
	sb.WriteString(script + "\n")
	sb.WriteString(`save(`)
	if maxW > 0 && maxH > 0 {
		sb.WriteString(fmt.Sprintf("resize-fit(img %d %d)", maxW, maxH))
	} else {
		sb.WriteString(`img`)
	}
	sb.WriteString(` "` + pathOut + `")`)
	return New().Script(sb.String()).render("dummy")
}

func RenderWithPXPFile(script string, args []any, files []string) ([]*image.NRGBA, error) {
	return New().ScriptFromFile(script).Args(args...).Files(files...).RenderImages()
}

func RenderWithPXPScript(script string, args []any, files []string) ([]*image.NRGBA, error) {
	return New().Script(script).Args(args...).Files(files...).RenderImages()
}

func DocMarkdown() string                { return language.DocMarkdown() }
func DocHTML() string                    { return language.DocHTML() }
func DocText() string                    { return language.DocText() }
func ExportToVSIX(vsixFile string) error { return language.ExportToVSIX(vsixFile) }

type PXP struct {
	err    error
	script string
	args   []any
	files  []string
}

func New() *PXP {
	return &PXP{
		script: "",
		args:   nil,
		files:  []string{},
	}
}

func (p *PXP) ScriptFromFile(file string) *PXP {
	if p.err != nil {
		return p
	}
	data, err := os.ReadFile(file)
	if err != nil {
		p.err = err
		return p
	}
	p.script = string(data)
	return p
}

func (p *PXP) Script(script string) *PXP {
	if p.err != nil {
		return p
	}
	p.script = script
	return p
}

func (p *PXP) Args(args ...any) *PXP {
	if p.err != nil {
		return p
	}
	p.args = args
	return p
}

func (p *PXP) Files(files ...string) *PXP {
	if p.err != nil {
		return p
	}
	p.files = files
	return p
}

func (p *PXP) RenderImages() ([]*image.NRGBA, error) {
	images := []*image.NRGBA{}
	if p.err != nil {
		return images, p.err
	}

	for _, file := range p.files {
		file = strings.TrimSpace(file)
		if file == "" {
			continue
		}
		img, err := p.render(file)
		if err != nil {
			p.err = fmt.Errorf("render error: %s", err.Error())
			continue
		}
		if img == nil {
			p.err = fmt.Errorf("render error: %s", "no image data returned")
			continue
		}
		images = append(images, img)
	}

	return images, p.err
}

func (p *PXP) RenderFiles(outputDir string) ([]string, error) {
	files := []string{}
	if p.err != nil {
		return files, p.err
	}

	if err := os.MkdirAll(outputDir, 0755); err != nil {
		p.err = fmt.Errorf("error creating output directory: %s", err.Error())
		return files, p.err
	}

	for _, file := range p.files {
		file = strings.TrimSpace(file)
		if file == "" {
			continue
		}
		img, err := p.render(file)
		if err != nil {
			p.err = fmt.Errorf("render error: %s", err.Error())
			continue
		}
		if img == nil {
			p.err = fmt.Errorf("render error: %s", "no image data returned")
			continue
		}

		var buf bytes.Buffer
		if err := png.Encode(&buf, img); err != nil {
			p.err = fmt.Errorf("encoding error: %s", err.Error())
			continue
		}

		baseName := filepath.Base(file)
		nameWithoutExt := strings.TrimSuffix(baseName, filepath.Ext(baseName))
		outputFile := filepath.Join(outputDir, nameWithoutExt+".png")

		if err := os.WriteFile(outputFile, buf.Bytes(), 0644); err != nil {
			p.err = fmt.Errorf("write error: %s", err.Error())
			continue
		}
		files = append(files, outputFile)
	}

	return files, p.err
}

func (p *PXP) render(file string) (*image.NRGBA, error) {
	if p.err != nil {
		return nil, p.err
	}

	res, err := language.Run(p.script, append([]any{file}, p.args...)...)
	if err != nil {
		return nil, fmt.Errorf("script execution error: %w", err)
	}
	if res == nil {
		return nil, fmt.Errorf("no result returned from script")
	}

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
		return nil, fmt.Errorf("unexpected result type: %T", res.Value())
	}

	if img == nil {
		return nil, fmt.Errorf("no image data returned")
	}

	if nrgba, ok := img.(*image.NRGBA); ok {
		return nrgba, nil
	}

	bounds := img.Bounds()
	nrgba := image.NewNRGBA(bounds)
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			nrgba.Set(x, y, img.At(x, y))
		}
	}

	return nrgba, nil
}
