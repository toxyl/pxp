package language

import (
	"image"
	"sync"

	"github.com/toxyl/math"
)

type PixelBufferProcessor func(r1, g1, b1, a1 uint32) (res float64)

func parallelPixelBuffer(img *image.NRGBA64, processor PixelBufferProcessor) (buf []float64) {
	numWorkers := NumColorConversionWorkers
	bounds := img.Bounds()
	width := bounds.Dx()
	height := bounds.Dy()
	minX := bounds.Min.X
	maxX := bounds.Max.X
	minY := bounds.Min.Y
	maxY := bounds.Max.Y
	buf = make([]float64, width*height)

	if height == 0 || width == 0 {
		return // Nothing to process
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
					buf[y*width+x] = processor(dsl.getColor(img, x, y))
				}
			}

		}(startY, endY)
	}

	wg.Wait() // Wait for all workers to finish
	return
}
