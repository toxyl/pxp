package fonts

import (
	"fmt"
	"image"
	"image/color"
	"math"
)

type GlyphPixel = [2]uint8

type Glyph [][]GlyphPixel

func (g *Glyph) render(w, h int, replacements map[uint8]color.RGBA64) *image.RGBA64 {
	img := image.NewRGBA64(image.Rect(0, 0, w, h))
	for y := range h {
		for x := range w {
			gp := (*g)[y][x]
			idx := gp[0]
			if idx == 0 {
				// background
				continue
			}
			c, ok := replacements[idx]
			if !ok {
				fmt.Printf("ERROR replacement color for index %d is not defined\n", idx)
				continue
			}
			alpha := gp[1]
			c.A = uint16(math.Round(float64(c.A) * (float64(alpha) / 255.0)))
			img.SetRGBA64(x, y, c)
		}
	}
	return img
}
