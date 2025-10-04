package language

import (
	"image"
	"image/color"

	"github.com/toxyl/math"
)

// blendWithAlpha applies the blend mode only where the top alpha is not zero,
// and interpolates for partial alpha, matching professional software behavior.
func blendWithAlpha(
	r1, g1, b1, a1, r2, g2, b2, a2 uint32,
	blendFunc func(uint32, uint32, uint32, uint32, uint32, uint32) (uint32, uint32, uint32),
) (r, g, b, a uint32) {
	if a2 == 0 {
		// Top is fully transparent, keep bottom pixel
		return r1, g1, b1, a1
	}
	// Calculate blend mode result
	rBlend, gBlend, bBlend := blendFunc(r1, g1, b1, r2, g2, b2)
	if a2 == 0xffff {
		// Top is fully opaque, use blend result
		r, g, b = rBlend, gBlend, bBlend
	} else {
		// Top is partially transparent, interpolate
		alpha := float64(a2) / 65535.0
		r = uint32(float64(rBlend)*alpha + float64(r1)*(1.0-alpha))
		g = uint32(float64(gBlend)*alpha + float64(g1)*(1.0-alpha))
		b = uint32(float64(bBlend)*alpha + float64(b1)*(1.0-alpha))
	}
	// Always use Porter-Duff alpha for output alpha
	a = porterDuffAlpha(a1, a2)
	return
}

// Helper function for Porter-Duff alpha compositing
func porterDuffAlpha(a1, a2 uint32) uint32 {
	return a1 + a2 - (a1 * a2 / 0xffff)
}

// hueToRGB helper function for HSL to RGB conversion
func hueToRGB(p, q, t float64) float64 {
	if t < 0 {
		t += 1
	}
	if t > 1 {
		t -= 1
	}
	if t < 1.0/6.0 {
		return p + (q-p)*6*t
	}
	if t < 1.0/2.0 {
		return q
	}
	if t < 2.0/3.0 {
		return p + (q-p)*(2.0/3.0-t)*6
	}
	return p
}

// Helper function to calculate saturation from RGB
func calculateSaturation(r, g, b float64) float64 {
	max := math.Max(math.Max(r, g), b)
	min := math.Min(math.Min(r, g), b)

	if max == min {
		return 0
	}

	l := (max + min) / 2
	if l > 0.5 {
		return (max - min) / (2 - max - min)
	}
	return (max - min) / (max + min)
}

// Helper function to convert RGB to HSL
func rgbToHsl(r, g, b float64) (h, s, l float64) {
	max := math.Max(math.Max(r, g), b)
	min := math.Min(math.Min(r, g), b)

	l = (max + min) / 2

	if max == min {
		h = 0
		s = 0
	} else {
		d := max - min
		if l > 0.5 {
			s = d / (2 - max - min)
		} else {
			s = d / (max + min)
		}

		switch max {
		case r:
			h = (g - b) / d
			if g < b {
				h += 6
			}
		case g:
			h = (b-r)/d + 2
		case b:
			h = (r-g)/d + 4
		}
		h /= 6
	}

	return h, s, l
}

// Helper function to convert HSL to RGB
func hslToRgb(h, s, l float64) (r, g, b float64) {
	if s == 0 {
		r, g, b = l, l, l
	} else {
		var q float64
		if l < 0.5 {
			q = l * (1 + s)
		} else {
			q = l + s - l*s
		}
		p := 2*l - q

		r = hueToRgb(p, q, h+1.0/3.0)
		g = hueToRgb(p, q, h)
		b = hueToRgb(p, q, h-1.0/3.0)
	}
	return r, g, b
}

// Helper function for HSL to RGB conversion
func hueToRgb(p, q, t float64) float64 {
	if t < 0 {
		t += 1
	}
	if t > 1 {
		t -= 1
	}
	if t < 1.0/6.0 {
		return p + (q-p)*6*t
	}
	if t < 1.0/2.0 {
		return q
	}
	if t < 2.0/3.0 {
		return p + (q-p)*(2.0/3.0-t)*6
	}
	return p
}

// Helper function for absolute difference between uint32 values
func absDiff(a, b uint32) uint32 {
	if a > b {
		return a - b
	}
	return b - a
}

// Helper function to convert RGB to HSL
func convertRGBToHSL(r, g, b uint32) (h, s, l float64) {
	// Convert to 0-1 range
	rf := float64(r) / 0xffff
	gf := float64(g) / 0xffff
	bf := float64(b) / 0xffff

	max := math.Max(math.Max(rf, gf), bf)
	min := math.Min(math.Min(rf, gf), bf)

	l = (max + min) / 2

	if max == min {
		h = 0
		s = 0
	} else {
		d := max - min
		s = d / (2 - max - min)
		if l > 0.5 {
			s = d / (2 - (max + min))
		}

		switch max {
		case rf:
			h = (gf - bf) / d
			if gf < bf {
				h += 6
			}
		case gf:
			h = 2 + (bf-rf)/d
		case bf:
			h = 4 + (rf-gf)/d
		}
		h /= 6
	}

	return h, s, l
}

// Helper function to convert HSL to RGB
func convertHSLToRGB(h, s, l float64) (r, g, b uint32) {
	var rf, gf, bf float64

	if s == 0 {
		rf = l
		gf = l
		bf = l
	} else {
		var q float64
		if l < 0.5 {
			q = l * (1 + s)
		} else {
			q = l + s - l*s
		}
		p := 2*l - q

		rf = convertHueToRGB(p, q, h+1.0/3.0)
		gf = convertHueToRGB(p, q, h)
		bf = convertHueToRGB(p, q, h-1.0/3.0)
	}

	// Convert back to uint32 range
	r = uint32(rf * 0xffff)
	g = uint32(gf * 0xffff)
	b = uint32(bf * 0xffff)

	return r, g, b
}

// Helper function for HSL to RGB conversion
func convertHueToRGB(p, q, t float64) float64 {
	if t < 0 {
		t += 1
	}
	if t > 1 {
		t -= 1
	}
	if t < 1.0/6.0 {
		return p + (q-p)*6*t
	}
	if t < 1.0/2.0 {
		return q
	}
	if t < 2.0/3.0 {
		return p + (q-p)*(2.0/3.0-t)*6
	}
	return p
}

// Helper function for bilinear interpolation of NRGBA64 images
func bilinearInterpolate(img *image.NRGBA64, x, y float64) color.Color {
	bounds := img.Bounds()
	x0 := math.Floor(x)
	y0 := math.Floor(y)
	x1 := x0 + 1
	y1 := y0 + 1

	ix0 := int(x0)
	iy0 := int(y0)
	ix1 := int(x1)
	iy1 := int(y1)

	// Check bounds strictly
	if ix0 < bounds.Min.X || ix1 > bounds.Max.X || iy0 < bounds.Min.Y || iy1 > bounds.Max.Y {
		// If interpolation points go out of bounds, return the nearest valid pixel color.
		// Clamp coordinates to valid range.
		clampX := math.Clamp(x, float64(bounds.Min.X), float64(bounds.Max.X-1))
		clampY := math.Clamp(y, float64(bounds.Min.Y), float64(bounds.Max.Y-1))
		return img.NRGBA64At(int(clampX), int(clampY))
	}

	c00 := img.NRGBA64At(ix0, iy0)
	c10 := img.NRGBA64At(ix1, iy0)
	c01 := img.NRGBA64At(ix0, iy1)
	c11 := img.NRGBA64At(ix1, iy1)

	tx := x - x0
	ty := y - y0

	r := lerp(lerp(float64(c00.R), float64(c10.R), tx), lerp(float64(c01.R), float64(c11.R), tx), ty)
	g := lerp(lerp(float64(c00.G), float64(c10.G), tx), lerp(float64(c01.G), float64(c11.G), tx), ty)
	b := lerp(lerp(float64(c00.B), float64(c10.B), tx), lerp(float64(c01.B), float64(c11.B), tx), ty)
	// Interpolate alpha as well? Or keep center pixel alpha? Let's interpolate.
	a := lerp(lerp(float64(c00.A), float64(c10.A), tx), lerp(float64(c01.A), float64(c11.A), tx), ty)

	return color.NRGBA64{
		R: uint16(math.Clamp(r, 0, 65535)),
		G: uint16(math.Clamp(g, 0, 65535)),
		B: uint16(math.Clamp(b, 0, 65535)),
		A: uint16(math.Clamp(a, 0, 65535)),
	}
}

// Linear interpolation helper
func lerp(a, b, t float64) float64 {
	return a*(1-t) + b*t
}

// clampU16 clamps v to the range [min, max] (inclusive)
func clampU16(v uint32, min, max uint16) uint16 {
	if v < uint32(min) {
		return min
	}
	if v > uint32(max) {
		return max
	}
	return uint16(v)
}

func calcKernelSize(radius float64) (float64, int, int) {
	sigma := radius / 2.0
	kernelSize := int(math.Ceil(sigma * 6)) // Cover 3 standard deviations
	if kernelSize%2 == 0 {
		kernelSize++
	}
	if kernelSize < 3 {
		kernelSize = 3
	}
	halfSize := kernelSize / 2
	return sigma, kernelSize, halfSize
}

func makeKernelGaussian1D(kernelSize int, sigma float64, halfSize int) []float64 {
	kernel := make([]float64, kernelSize)
	sum := 0.0

	// Calculate 1D Gaussian kernel
	sigma2 := 2 * sigma * sigma
	for i := range kernelSize {
		x := float64(i - halfSize)
		kernel[i] = math.Exp(-(x * x) / sigma2)
		sum += kernel[i]
	}

	// Normalize kernel
	for i := range kernelSize {
		kernel[i] /= sum
	}
	return kernel
}
