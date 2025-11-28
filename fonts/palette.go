package fonts

import (
	"image/color"
)

type Palette map[uint8]color.RGBA

func (p *Palette) Get(index uint8) (*color.RGBA, bool) {
	c, ok := (*p)[index]
	return &c, ok
}

func (p *Palette) Append(colors ...color.RGBA) {
	idx := uint8(len(*p))
	for i, c := range colors {
		(*p)[idx+uint8(i)] = c
	}
}

func (p *Palette) Find(col color.RGBA) (uint8, uint8) {
	if col.A == 0 {
		return 0, 0 // background
	}
	oR := uint8((uint16(col.R) * 0xFF) / uint16(col.A))
	oG := uint8((uint16(col.G) * 0xFF) / uint16(col.A))
	oB := uint8((uint16(col.B) * 0xFF) / uint16(col.A))
	d := uint8(0)

	fInRange := func(a, b uint8) bool {
		if a == b {
			return true
		}
		l, u := uint8(max(0, int(b)-int(d))), uint8(min(255, int(b)+int(d)))
		ok := a >= l && a <= u
		return ok
	}

	for i, c := range *p {
		if i > 0 && fInRange(oR, c.R) && fInRange(oG, c.G) && fInRange(oB, c.B) { // i > 0 so we skip the background color, we already checked for it at the start
			return i, col.A
		}
	}
	return 0, 0 // background
}

func NewPalette(colors ...color.RGBA) *Palette {
	p := Palette{}
	p.Append(color.RGBA{R: 0x00, G: 0x00, B: 0x00, A: 0x00},
		color.RGBA{R: 0xFF, G: 0xFF, B: 0xFF, A: 0xFF},
		color.RGBA{R: 0x00, G: 0x00, B: 0x00, A: 0xFF})
	p.Append(colors...)
	return &p
}
