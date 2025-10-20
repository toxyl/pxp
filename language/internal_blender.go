package language

import (
	"image"
	"image/color"
	"sync"

	"github.com/toxyl/math"
)

type PixelBlendProcessor func(r1, g1, b1, a1, r2, g2, b2, a2 uint32) (r, g, b, a uint32)

type Blender struct {
	name    string
	process PixelBlendProcessor
}

func (bl *Blender) Color(colA, colB color.RGBA64) color.RGBA64 {
	r1, g1, b1, a1 := colA.RGBA()
	r2, g2, b2, a2 := colB.RGBA()
	r, g, b, a := bl.process(r1, g1, b1, a1, r2, g2, b2, a2)
	return color.RGBA64{
		R: uint16(r),
		G: uint16(g),
		B: uint16(b),
		A: uint16(a),
	}
}

// Pixel blends the `Pixel` from `imgB` onto `imgA`.
// Unlike other functions this does not create a new image but instead
// modifies `imgA` directly.
func (bl *Blender) Pixel(pixel Point, imgA, imgB *image.NRGBA64) *image.NRGBA64 {
	x, y := int(pixel.X), int(pixel.Y)
	r1, g1, b1, a1 := dsl.getColor(imgA, x, y)
	r2, g2, b2, a2 := dsl.getColor(imgB, x, y)
	r, g, b, a := bl.process(r1, g1, b1, a1, r2, g2, b2, a2)
	dsl.setColor(imgA, x, y, r, g, b, a)
	return imgA
}

// Images blends `imgB` onto `imgA`.
// Unlike other functions this does not create a new image but instead
// modifies `imgA` directly.
func (bl *Blender) Images(imgA, imgB *image.NRGBA64) *image.NRGBA64 {
	bounds := imgA.Bounds()
	width := bounds.Dx()
	height := bounds.Dy()
	minX := bounds.Min.X
	maxX := bounds.Max.X
	minY := bounds.Min.Y
	maxY := bounds.Max.Y

	if height == 0 || width == 0 {
		return imgA // Nothing to process
	}

	rowsPerWorker := (height + numWorkers - 1) / numWorkers

	var wg sync.WaitGroup

	for i := range numWorkers {
		startY := minY + i*rowsPerWorker
		endY := math.Min(startY+rowsPerWorker, maxY)

		if startY >= endY {
			continue
		}

		wg.Add(1)
		go func(startY, endY int) {
			defer wg.Done()
			for y := startY; y < endY; y++ {
				for x := minX; x < maxX; x++ {
					bl.Pixel(*P(float64(x), float64(y)), imgA, imgB)
				}
			}
		}(startY, endY)
	}

	wg.Wait()
	return imgA
}

func NewBlender(name string, fnProcess PixelBlendProcessor) *Blender {
	b := Blender{
		name:    name,
		process: fnProcess,
	}
	return &b
}
