package fonts

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"strings"

	"github.com/toxyl/flo"
	"gopkg.in/yaml.v3"
)

type BitmapFont struct {
	GlyphWidth  int             `yaml:"height"`
	GlyphHeight int             `yaml:"width"`
	Palette     *Palette        `yaml:"palette"`
	Glyphs      map[rune]*Glyph `yaml:"glyphs"`
}

func (b *BitmapFont) MeasureText(s string) (columns, lines, runes, width, height int) {
	columns, lines, runes = 0, 0, 0
	for line := range strings.SplitSeq(s, "\n") {
		r := len([]rune(line))
		runes += r
		columns = max(columns, r)
		lines++
	}
	width = columns * b.GlyphWidth
	height = lines * b.GlyphHeight
	return
}

func (b *BitmapFont) Save(file *flo.FileObj) error {
	return file.StoreYAML(b)
}

func (b *BitmapFont) RenderToFile(file, str string, textColor, outlineColor color.RGBA64, additionalColors ...color.RGBA64) error {
	img := b.Render(str, textColor, outlineColor, additionalColors...)
	f, err := os.Create(file)
	if err != nil {
		return err
	}
	defer f.Close()
	return png.Encode(f, img)
}

func (b *BitmapFont) Render(s string, textColor, outlineColor color.RGBA64, additionalColors ...color.RGBA64) *image.NRGBA64 {
	_, _, _, w, h := b.MeasureText(s)
	x, y := 0, 0
	img := image.NewNRGBA64(image.Rect(0, 0, w, h))
	for _, r := range []rune(s) {
		if r == '\n' {
			y += b.GlyphHeight
			x = 0
			continue
		}
		if g, ok := b.Glyphs[r]; ok {
			repl := map[uint8]color.RGBA64{
				0: TRANSPARENT,
				1: textColor,
				2: outlineColor,
			}
			for i, ar := range additionalColors {
				repl[uint8(i)+2] = ar
			}
			if len(repl) != len(*b.Palette) {
				fmt.Printf("Number of colors does not match font palette. It expects %d colors, but got %d.\n", len(*b.Palette), len(repl))
				return nil
			}
			gr := g.render(b.GlyphWidth, b.GlyphHeight, repl)
			gw := gr.Bounds().Dx()
			gh := gr.Bounds().Dy()
			for gy := range gh {
				for gx := range gw {
					img.SetRGBA64(x+gx, y+gy, gr.RGBA64At(gx, gy))
				}
			}
		}
		x += b.GlyphWidth
	}
	return img
}

func LoadBitmapFont(file *flo.FileObj) (*BitmapFont, error) {
	f := &BitmapFont{
		GlyphWidth:  0,
		GlyphHeight: 0,
		Glyphs:      map[rune]*Glyph{},
		Palette:     &Palette{},
	}
	err := file.LoadYAML(f)
	return f, err
}

func LoadBitmapFontFromBytes(data []byte) (*BitmapFont, error) {
	f := &BitmapFont{
		GlyphWidth:  0,
		GlyphHeight: 0,
		Glyphs:      map[rune]*Glyph{},
		Palette:     &Palette{},
	}
	err := yaml.Unmarshal(data, f)
	return f, err
}

// MakeBitmapFont creates a bitmap font from the given charset and spritesheet.
// The palette hardcodes the first three items: 0 - transparent (background), 1 - white (text color), 2 - black (outline color).
// If provided the given palette will be appended, i.e. its indices start at 3.
func MakeBitmapFont(fileOut, fileCharset, fileSprites *flo.FileObj, additionalColors ...color.RGBA) error {
	const (
		gutter = 1
	)
	var (
		charset     = fileCharset.AsString()
		spritesheet = fileSprites.AsBytes()
	)

	b := BitmapFont{
		GlyphWidth:  0,
		GlyphHeight: 0,
		Glyphs:      map[rune]*Glyph{},
		Palette:     NewPalette(additionalColors...),
	}

	// load the image
	fontImg, err := png.Decode(bytes.NewReader(spritesheet))
	if err != nil {
		return err
	}

	// extract and calculate dimensions
	fb := fontImg.Bounds()
	w, h := fb.Dx(), fb.Dy()
	cols, rows, _, _, _ := b.MeasureText(charset)
	b.GlyphWidth, b.GlyphHeight = (w/cols)-1, (h/rows)-1

	// and convert it to RGBA
	img := image.NewRGBA(fb)
	for y := range h {
		for x := range w {
			r, g, b, a := fontImg.At(x, y).RGBA()
			var nrgba color.NRGBA
			if a > 0 {
				nrgba = color.NRGBA{
					R: uint8((r * 255) / a),
					G: uint8((g * 255) / a),
					B: uint8((b * 255) / a),
					A: uint8(a >> 8),
				}
			}
			img.Set(x, y, nrgba)
		}
	}

	// calculate offsets for each rune and generate glyph from it
	done := map[rune]struct{}{} // to avoid processing runes more than once; first seen, first used
	for row, line := range strings.Split(charset, "\n") {
		for col, char := range []rune(line) {
			if _, ok := done[char]; !ok {
				x, y := col*(b.GlyphWidth+gutter), gutter+row*(b.GlyphHeight+gutter)
				glyphCols := Glyph{}
				for gy := range b.GlyphHeight {
					cols := []GlyphPixel{}
					for gx := range b.GlyphWidth {
						orig := img.RGBAAt(x+gx, y+gy)
						idx, alpha := b.Palette.Find(orig)
						cols = append(cols, GlyphPixel{idx, alpha})
					}
					glyphCols = append(glyphCols, cols)
				}
				b.Glyphs[char] = &glyphCols
				done[char] = struct{}{}
			}
		}
	}

	return b.Save(fileOut)
}
