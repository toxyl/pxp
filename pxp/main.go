package pxp

import (
	"bytes"
	"fmt"
	"image"
	"image/png"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/toxyl/pxp/language"
)

type renderLock struct {
	mu     *sync.Mutex
	locked bool
}

func (l *renderLock) lock() {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.locked = true
}

func (l *renderLock) unlock() {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.locked = false
}

func (l *renderLock) isLocked() bool {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.locked
}

func (l *renderLock) exec(fn func()) {
	for l.isLocked() {
		time.Sleep(1 * time.Second)
	}
	l.lock()
	defer l.unlock()
	fn()
}

var (
	rlock = &renderLock{
		mu:     &sync.Mutex{},
		locked: false,
	}
)

// RenderToFile processes the given `script` and stores the result in `path`.
//
// The script must use the variable `img` to store the final result.
//
// When `maxW` and `maxH` are greater than zero, the output image will be resized to fit within the given dimensions.
func RenderToFile(script, path string, maxW, maxH int) (err error) {
	sb := strings.Builder{}
	sb.WriteString(script + "\n")
	sb.WriteString(`save(`)
	if maxW > 0 && maxH > 0 {
		sb.WriteString(fmt.Sprintf("resize-fit(img %d %d)", maxW, maxH))
	} else {
		sb.WriteString(`img`)
	}
	sb.WriteString(` "` + path + `")`)
	_, err = New().Script(sb.String()).render(nil)
	return
}

// RenderFile loads `pathIn` into the variable `in`, processes it with the given `script` and stores the result in `pathOut`.
//
// The script must use the variable `in` as input image and store the final result in the `img` variable.
//
// The `pathIn` variable can be a URL or a local file path.
//
// When `maxW` and `maxH` are greater than zero, the output image will be resized to fit within the given dimensions.
func RenderFile(script, pathIn, pathOut string, maxW, maxH int) (img *image.NRGBA, err error) {
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
	return New().Script(sb.String()).render(nil)
}

func RenderWithPXPFile(script string, args []any, files []string) (images []*image.NRGBA, err error) {
	return New().ScriptFromFile(script).Args(args...).Files(files...).RenderImages()
}

func RenderWithPXPScript(script string, args []any, files []string) (images []*image.NRGBA, err error) {
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
		img, err := p.render(&file)
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
		img, err := p.render(&file)
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

func (p *PXP) render(file *string) (*image.NRGBA, error) {
	if p.err != nil {
		return nil, p.err
	}
	var nrgba *image.NRGBA
	var err error
	rlock.exec(func() {
		f := "dummy"
		if file != nil {
			f = *file
		}
		res, err2 := language.Run(p.script, append([]any{f}, p.args...)...)
		if err2 != nil {
			err = fmt.Errorf("script execution error: %w", err2)
			return
		}
		if res == nil {
			err = fmt.Errorf("no result returned from script")
			return
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
			err = fmt.Errorf("unexpected result type: %T", res.Value())
			return
		}

		if img == nil {
			err = fmt.Errorf("no image data returned")
			return
		}

		if res, ok := img.(*image.NRGBA); ok {
			nrgba = res
			return
		}

		bounds := img.Bounds()
		nrgba := image.NewNRGBA(bounds)
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			for x := bounds.Min.X; x < bounds.Max.X; x++ {
				nrgba.Set(x, y, img.At(x, y))
			}
		}
	})

	return nrgba, err
}
