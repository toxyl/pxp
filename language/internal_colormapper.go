package language

import (
	"fmt"
	"image"

	"github.com/toxyl/math"
)

func colorMapper(img *image.NRGBA64, sourceStops [][]any, targetStops [][]any, tolerance float64, precision float64) (*image.NRGBA64, error) {
	if len(sourceStops) == 0 {
		return nil, fmt.Errorf("no source color stops provided")
	}

	// Find source range from stops
	srcMin := math.MaxFloat64
	srcMax := -math.MaxFloat64
	for _, s := range sourceStops {
		t := s[0].(float64)
		if t < srcMin {
			srcMin = t
		}
		if t > srcMax {
			srcMax = t
		}
	}

	trgMin := math.MaxFloat64
	trgMax := -math.MaxFloat64
	if len(targetStops) == 0 {
		trgMin = srcMin
		trgMax = srcMax
		targetStops = [][]any{
			{srcMax, 0.0, 0.0, 1.0, 1.0},
			{srcMin, 0.0, 0.0, 0.0, 1.0},
		}
	} else {
		// Find target range from stops
		for _, s := range targetStops {
			t := s[0].(float64)
			if t < trgMin {
				trgMin = t
			}
			if t > trgMax {
				trgMax = t
			}
		}
	}

	// Create color bar slices - N is max of image dimensions multiplied by precision
	bounds := img.Bounds()
	n := bounds.Dx()
	if bounds.Dy() > n {
		n = bounds.Dy()
	}
	n = int(float64(n) * precision)

	type labColor struct {
		l, a, b, alpha float64
	}

	// Build source color bar using LAB
	srcBar := make([]labColor, n)
	for i := 0; i < n; i++ {
		value := srcMin + float64(i)*(srcMax-srcMin)/float64(n-1)
		if c, err := mapColor(value, srcMin, srcMax, sourceStops); err == nil {
			rf := float64(c.R) / 65535.0
			gf := float64(c.G) / 65535.0
			bf := float64(c.B) / 65535.0
			l, a, b := convertRGBToLAB(rf, gf, bf)
			srcBar[i] = labColor{
				l:     l,
				a:     a,
				b:     b,
				alpha: float64(c.A) / 65535.0,
			}
		}
	}

	// Build target color bar using LAB
	trgBar := make([]labColor, n)
	for i := 0; i < n; i++ {
		value := trgMin + float64(i)*(trgMax-trgMin)/float64(n-1)
		if c, err := mapColor(value, trgMin, trgMax, targetStops); err == nil {
			rf := float64(c.R) / 65535.0
			gf := float64(c.G) / 65535.0
			bf := float64(c.B) / 65535.0
			l, a, b := convertRGBToLAB(rf, gf, bf)
			trgBar[i] = labColor{
				l:     l,
				a:     a,
				b:     b,
				alpha: float64(c.A) / 65535.0,
			}
		}
	}

	// Process image
	return dsl.parallelProcessNRGBA64(img, func(r1, g1, b1, a1 uint32) (r, g, b, a uint32) {
		// Convert pixel RGB to LAB
		rf := float64(r1) / 65535.0
		gf := float64(g1) / 65535.0
		bf := float64(b1) / 65535.0
		l, aLab, bLab := convertRGBToLAB(rf, gf, bf)

		// Find all matches within tolerance using LAB distance (perceptually uniform)
		matches := []int{}
		for i, srcLAB := range srcBar {
			// Delta E in LAB space (Euclidean distance is perceptually uniform)
			dl := l - srcLAB.l
			da := aLab - srcLAB.a
			db := bLab - srcLAB.b
			dist := math.Sqrt(dl*dl + da*da + db*db)
			if dist <= tolerance {
				matches = append(matches, i)
			}
		}

		var bestIdx int
		if len(matches) > 0 {
			// Use median index for stability
			bestIdx = matches[len(matches)/2]
		} else {
			// Fallback to nearest neighbor
			bestDist := math.MaxFloat64
			for i, srcLAB := range srcBar {
				dl := l - srcLAB.l
				da := aLab - srcLAB.a
				db := bLab - srcLAB.b
				dist := math.Sqrt(dl*dl + da*da + db*db)
				if dist < bestDist {
					bestDist = dist
					bestIdx = i
				}
			}
		}

		// Use that index in target bar and convert back to RGB
		trgLAB := trgBar[bestIdx]
		tr, tg, tb := convertLABToRGB(trgLAB.l, trgLAB.a, trgLAB.b)

		r = uint32(math.Max(0.0, math.Min(65535.0, tr*65535.0)))
		g = uint32(math.Max(0.0, math.Min(65535.0, tg*65535.0)))
		b = uint32(math.Max(0.0, math.Min(65535.0, tb*65535.0)))
		a = uint32(math.Max(0.0, math.Min(65535.0, trgLAB.alpha*65535.0)))
		return
	}, NumColorConversionWorkers), nil
}
