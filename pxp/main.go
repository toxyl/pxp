package pxp

import (
	"fmt"
	"image"
	"runtime"
	"strings"
	"time"

	"github.com/toxyl/pxp/language"
	"github.com/toxyl/safe"
)

var (
	MAX_CONCURRENCY = 16
	activeRenders   = safe.NewSafeInt(0)
)

func init() {
	MAX_CONCURRENCY = max(1, runtime.NumCPU()>>1)
	language.NumColorConversionWorkers = 2
}

// RenderToFile processes the given `script` and stores the result in `path`.
//
// The script must use the variable `img` to store the final result.
//
// When `maxW` and `maxH` are greater than zero, the output image will be resized to fit within the given dimensions.
func RenderToFile(script, baseDir, path string, maxW, maxH int, replacements map[string]string) (err error) {
	sb := strings.Builder{}
	sb.WriteString(script + "\n")
	sb.WriteString(`save(`)
	if maxW > 0 && maxH > 0 {
		sb.WriteString(fmt.Sprintf("resize-fit(img %d %d)", maxW, maxH))
	} else {
		sb.WriteString(`img`)
	}
	sb.WriteString(` "` + path + `")`)
	_, err = New().Script(sb.String()).Render(baseDir, replacements)
	return
}

func DocMarkdown() string                { return language.DocMarkdown() }
func DocHTML() string                    { return language.DocHTML() }
func DocText() string                    { return language.DocText() }
func ExportToVSIX(vsixFile string) error { return language.ExportToVSIX(vsixFile) }

type PXP struct {
	lang   *language.Language
	err    error
	script string
}

func New() *PXP {
	return &PXP{
		lang:   language.New(),
		script: "",
	}
}

func (p *PXP) Script(script string) *PXP {
	if p.err != nil {
		return p
	}
	p.script = script
	return p
}

func (p *PXP) Render(baseDir string, replacements map[string]string) (*image.NRGBA, error) {
	if p.err != nil {
		return nil, p.err
	}
	for activeRenders.Get() >= MAX_CONCURRENCY {
		time.Sleep(10 * time.Second)
	}
	activeRenders.Inc()
	res, err := p.lang.Run(p.script, baseDir, replacements, []any{"dummy"}...)
	activeRenders.Dec()
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

	if res, ok := img.(*image.NRGBA); ok {
		return res, nil
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
