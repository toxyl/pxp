package language

import (
	"image"

	"github.com/toxyl/math"
)

// @Name: blend-normal
// @Desc: Blends the two images using the normal blend mode (alpha compositing)
// @Param:      imgA     - -   	-   The bottom image
// @Param:      imgB     - -   	-   The top image
// @Returns:    result  - -   	-   The blended image
func blendNormal(imgA *image.NRGBA64, imgB *image.NRGBA64) (*image.NRGBA64, error) {
	return parallelBlend(imgA, imgB, func(r1, g1, b1, a1, r2, g2, b2, a2 uint32) (r, g, b, a uint32) {
		// Normal blend mode for un-premultiplied colors
		// Standard alpha compositing (Porter-Duff over)
		a3 := 0xffff - a2
		r = ((r2 * a2) + (r1 * a3)) / 0xffff
		g = ((g2 * a2) + (g1 * a3)) / 0xffff
		b = ((b2 * a2) + (b1 * a3)) / 0xffff
		a = porterDuffAlpha(a1, a2)
		return
	})
}

// @Name: blend-erase
// @Desc: Erases the bottom image wherever the top image is present (destination out)
// @Param:      imgA     - -   	-   The bottom image
// @Param:      imgB     - -   	-   The top image
// @Returns:    result  - -   	-   The blended image
func blendErase(imgA *image.NRGBA64, imgB *image.NRGBA64) (*image.NRGBA64, error) {
	return parallelBlend(imgA, imgB, func(r1, g1, b1, a1, r2, g2, b2, a2 uint32) (r, g, b, a uint32) {
		// Erase blend mode for un-premultiplied colors
		// Formula: destination out (bottom * (1 - top alpha))
		invA2 := 0xffff - a2
		r = (r1 * invA2) / 0xffff
		g = (g1 * invA2) / 0xffff
		b = (b1 * invA2) / 0xffff
		a = (a1 * invA2) / 0xffff
		return
	})
}

// @Name: blend-multiply
// @Desc: Blends the two images using the multiply blend mode
// @Param:      imgA     - -   	-   The bottom image
// @Param:      imgB     - -   	-   The top image
// @Returns:    result  - -   	-   The blended image
func blendMultiply(imgA *image.NRGBA64, imgB *image.NRGBA64) (*image.NRGBA64, error) {
	return parallelBlend(imgA, imgB, func(r1, g1, b1, a1, r2, g2, b2, a2 uint32) (r, g, b, a uint32) {
		return blendWithAlpha(r1, g1, b1, a1, r2, g2, b2, a2, func(r1, g1, b1, r2, g2, b2 uint32) (uint32, uint32, uint32) {
			return (r1 * r2) / 0xffff, (g1 * g2) / 0xffff, (b1 * b2) / 0xffff
		})
	})
}

// @Name: blend-screen
// @Desc: Blends the two images using the screen blend mode
// @Param:      imgA     - -   	-   The bottom image
// @Param:      imgB     - -   	-   The top image
// @Returns:    result  - -   	-   The blended image
func blendScreen(imgA *image.NRGBA64, imgB *image.NRGBA64) (*image.NRGBA64, error) {
	return parallelBlend(imgA, imgB, func(r1, g1, b1, a1, r2, g2, b2, a2 uint32) (r, g, b, a uint32) {
		return blendWithAlpha(r1, g1, b1, a1, r2, g2, b2, a2, func(r1, g1, b1, r2, g2, b2 uint32) (uint32, uint32, uint32) {
			// Screen blend mode for un-premultiplied colors
			// Formula: 1 - (1 - a) * (1 - b)
			r := 0xffff - ((0xffff - r1) * (0xffff - r2) / 0xffff)
			g := 0xffff - ((0xffff - g1) * (0xffff - g2) / 0xffff)
			b := 0xffff - ((0xffff - b1) * (0xffff - b2) / 0xffff)
			return r, g, b
		})
	})
}

// @Name: blend-exclusion
// @Desc: Blends the two images using the exclusion blend mode
// @Param:      imgA     - -   	-   The bottom image
// @Param:      imgB     - -   	-   The top image
// @Returns:    result  - -   	-   The blended image
func blendExclusion(imgA *image.NRGBA64, imgB *image.NRGBA64) (*image.NRGBA64, error) {
	return parallelBlend(imgA, imgB, func(r1, g1, b1, a1, r2, g2, b2, a2 uint32) (r, g, b, a uint32) {
		return blendWithAlpha(r1, g1, b1, a1, r2, g2, b2, a2, func(r1, g1, b1, r2, g2, b2 uint32) (uint32, uint32, uint32) {
			// Exclusion blend mode for un-premultiplied colors
			// Formula: a + b - 2ab
			r := r1 + r2 - ((r1 * r2) >> 15)
			g := g1 + g2 - ((g1 * g2) >> 15)
			b := b1 + b2 - ((b1 * b2) >> 15)
			return r, g, b
		})
	})
}

// @Name: blend-overlay
// @Desc: Blends the two images using the overlay blend mode
// @Param:      imgA     - -   	-   The bottom image
// @Param:      imgB     - -   	-   The top image
// @Returns:    result  - -   	-   The blended image
func blendOverlay(imgA *image.NRGBA64, imgB *image.NRGBA64) (*image.NRGBA64, error) {
	return parallelBlend(imgA, imgB, func(r1, g1, b1, a1, r2, g2, b2, a2 uint32) (r, g, b, a uint32) {
		return blendWithAlpha(r1, g1, b1, a1, r2, g2, b2, a2, func(r1, g1, b1, r2, g2, b2 uint32) (uint32, uint32, uint32) {
			// Overlay blend mode for un-premultiplied colors
			// Formula: if bottom < 0.5: 2 * bottom * top
			//          else: 1 - 2 * (1 - bottom) * (1 - top)

			// Red channel
			var r uint32
			if r1 < 0x8000 {
				r = (2 * r1 * r2) / 0xffff
			} else {
				r = 0xffff - ((2 * (0xffff - r1) * (0xffff - r2)) / 0xffff)
			}

			// Green channel
			var g uint32
			if g1 < 0x8000 {
				g = (2 * g1 * g2) / 0xffff
			} else {
				g = 0xffff - ((2 * (0xffff - g1) * (0xffff - g2)) / 0xffff)
			}

			// Blue channel
			var b uint32
			if b1 < 0x8000 {
				b = (2 * b1 * b2) / 0xffff
			} else {
				b = 0xffff - ((2 * (0xffff - b1) * (0xffff - b2)) / 0xffff)
			}

			return r, g, b
		})
	})
}

// @Name: blend-color-burn
// @Desc: Blends the two images using the color burn blend mode
// @Param:      imgA     - -   	-   The bottom image
// @Param:      imgB     - -   	-   The top image
// @Returns:    result  - -   	-   The blended image
func blendColorBurn(imgA *image.NRGBA64, imgB *image.NRGBA64) (*image.NRGBA64, error) {
	return parallelBlend(imgA, imgB, func(r1, g1, b1, a1, r2, g2, b2, a2 uint32) (r, g, b, a uint32) {
		return blendWithAlpha(r1, g1, b1, a1, r2, g2, b2, a2, func(r1, g1, b1, r2, g2, b2 uint32) (uint32, uint32, uint32) {
			// Color burn blend mode for un-premultiplied colors
			// Formula: 1 - (1 - bottom) / top

			// Red channel
			var r uint32
			if r2 == 0xffff {
				r = r1
			} else if r2 > 0 {
				r = 0xffff - ((0xffff-r1)*0xffff)/r2
				if r > 0xffff {
					r = 0
				}
			}

			// Green channel
			var g uint32
			if g2 == 0xffff {
				g = g1
			} else if g2 > 0 {
				g = 0xffff - ((0xffff-g1)*0xffff)/g2
				if g > 0xffff {
					g = 0
				}
			}

			// Blue channel
			var b uint32
			if b2 == 0xffff {
				b = b1
			} else if b2 > 0 {
				b = 0xffff - ((0xffff-b1)*0xffff)/b2
				if b > 0xffff {
					b = 0
				}
			}

			return r, g, b
		})
	})
}

// @Name: blend-color-dodge
// @Desc: Blends the two images using the color dodge blend mode
// @Param:      imgA     - -   	-   The bottom image
// @Param:      imgB     - -   	-   The top image
// @Returns:    result  - -   	-   The blended image
func blendColorDodge(imgA *image.NRGBA64, imgB *image.NRGBA64) (*image.NRGBA64, error) {
	return parallelBlend(imgA, imgB, func(r1, g1, b1, a1, r2, g2, b2, a2 uint32) (r, g, b, a uint32) {
		return blendWithAlpha(r1, g1, b1, a1, r2, g2, b2, a2, func(r1, g1, b1, r2, g2, b2 uint32) (uint32, uint32, uint32) {
			// Color dodge blend mode for un-premultiplied colors
			// Formula: bottom / (1 - top)

			// Red channel
			var r uint32
			if r2 == 0 {
				r = r1
			} else if r2 == 0xffff {
				r = 0xffff
			} else {
				r = min((r1*0xffff)/(0xffff-r2), 0xffff)
			}

			// Green channel
			var g uint32
			if g2 == 0 {
				g = g1
			} else if g2 == 0xffff {
				g = 0xffff
			} else {
				g = min((g1*0xffff)/(0xffff-g2), 0xffff)
			}

			// Blue channel
			var b uint32
			if b2 == 0 {
				b = b1
			} else if b2 == 0xffff {
				b = 0xffff
			} else {
				b = min((b1*0xffff)/(0xffff-b2), 0xffff)
			}

			return r, g, b
		})
	})
}

// @Name: blend-soft-light
// @Desc: Blends the two images using the soft light blend mode
// @Param:      imgA     - -   	-   The bottom image
// @Param:      imgB     - -   	-   The top image
// @Returns:    result  - -   	-   The blended image
func blendSoftLight(imgA *image.NRGBA64, imgB *image.NRGBA64) (*image.NRGBA64, error) {
	return parallelBlend(imgA, imgB, func(r1, g1, b1, a1, r2, g2, b2, a2 uint32) (r, g, b, a uint32) {
		return blendWithAlpha(r1, g1, b1, a1, r2, g2, b2, a2, func(r1, g1, b1, r2, g2, b2 uint32) (uint32, uint32, uint32) {
			// Soft light blend mode for un-premultiplied colors
			// Formula: if top < 0.5: bottom - (1-2*top)*bottom*(1-bottom)
			//          else: bottom + (2*top-1)*(sqrt(bottom)-bottom)

			softLightChannel := func(b, t float64) float64 {
				if t < 0.5 {
					return b - (1-2*t)*b*(1-b)
				}
				return b + (2*t-1)*(math.Sqrt(b)-b)
			}

			// Convert to float64 in range [0,1]
			rf := float64(r1) / 65535.0
			gf := float64(g1) / 65535.0
			bf := float64(b1) / 65535.0
			rf2 := float64(r2) / 65535.0
			gf2 := float64(g2) / 65535.0
			bf2 := float64(b2) / 65535.0

			// Apply soft light blend
			r := uint32(math.Clamp(softLightChannel(rf, rf2)*65535.0, 0, 65535))
			g := uint32(math.Clamp(softLightChannel(gf, gf2)*65535.0, 0, 65535))
			b := uint32(math.Clamp(softLightChannel(bf, bf2)*65535.0, 0, 65535))
			return r, g, b
		})
	})
}

// @Name: blend-hard-light
// @Desc: Blends the two images using the hard light blend mode
// @Param:      imgA     - -   	-   The bottom image
// @Param:      imgB     - -   	-   The top image
// @Returns:    result  - -   	-   The blended image
func blendHardLight(imgA *image.NRGBA64, imgB *image.NRGBA64) (*image.NRGBA64, error) {
	return parallelBlend(imgA, imgB, func(r1, g1, b1, a1, r2, g2, b2, a2 uint32) (r, g, b, a uint32) {
		return blendWithAlpha(r1, g1, b1, a1, r2, g2, b2, a2, func(r1, g1, b1, r2, g2, b2 uint32) (uint32, uint32, uint32) {
			// Hard light blend mode for un-premultiplied colors
			// Formula: if top < 0.5: 2 * bottom * top
			//          else: 1 - 2 * (1 - bottom) * (1 - top)

			// Red channel
			var r uint32
			if r2 < 0x8000 {
				r = (2 * r1 * r2) / 0xffff
			} else {
				r = 0xffff - (2*(0xffff-r1)*(0xffff-r2))/0xffff
			}

			// Green channel
			var g uint32
			if g2 < 0x8000 {
				g = (2 * g1 * g2) / 0xffff
			} else {
				g = 0xffff - (2*(0xffff-g1)*(0xffff-g2))/0xffff
			}

			// Blue channel
			var b uint32
			if b2 < 0x8000 {
				b = (2 * b1 * b2) / 0xffff
			} else {
				b = 0xffff - (2*(0xffff-b1)*(0xffff-b2))/0xffff
			}

			return r, g, b
		})
	})
}

// @Name: blend-difference
// @Desc: Blends the two images using the difference blend mode
// @Param:      imgA     - -   	-   The bottom image
// @Param:      imgB     - -   	-   The top image
// @Returns:    result  - -   	-   The blended image
func blendDifference(imgA *image.NRGBA64, imgB *image.NRGBA64) (*image.NRGBA64, error) {
	return parallelBlend(imgA, imgB, func(r1, g1, b1, a1, r2, g2, b2, a2 uint32) (r, g, b, a uint32) {
		return blendWithAlpha(r1, g1, b1, a1, r2, g2, b2, a2, func(r1, g1, b1, r2, g2, b2 uint32) (uint32, uint32, uint32) {
			// Difference blend mode for un-premultiplied colors
			// Formula: |bottom - top|
			r := absDiff(r1, r2)
			g := absDiff(g1, g2)
			b := absDiff(b1, b2)
			return r, g, b
		})
	})
}

// @Name: blend-subtract
// @Desc: Blends the two images using the subtract blend mode
// @Param:      imgA     - -   	-   The bottom image
// @Param:      imgB     - -   	-   The top image
// @Returns:    result  - -   	-   The blended image
func blendSubtract(imgA *image.NRGBA64, imgB *image.NRGBA64) (*image.NRGBA64, error) {
	return parallelBlend(imgA, imgB, func(r1, g1, b1, a1, r2, g2, b2, a2 uint32) (r, g, b, a uint32) {
		return blendWithAlpha(r1, g1, b1, a1, r2, g2, b2, a2, func(r1, g1, b1, r2, g2, b2 uint32) (uint32, uint32, uint32) {
			// Subtract blend mode for un-premultiplied colors
			// Formula: bottom - top, clamp at 0
			r := math.Clamp(r1-r2, 0, 0xffff)
			g := math.Clamp(g1-g2, 0, 0xffff)
			b := math.Clamp(b1-b2, 0, 0xffff)
			return r, g, b
		})
	})
}

// @Name: blend-divide
// @Desc: Blends the two images using the divide blend mode
// @Param:      imgA     - -   	-   The bottom image
// @Param:      imgB     - -   	-   The top image
// @Returns:    result  - -   	-   The blended image
func blendDivide(imgA *image.NRGBA64, imgB *image.NRGBA64) (*image.NRGBA64, error) {
	return parallelBlend(imgA, imgB, func(r1, g1, b1, a1, r2, g2, b2, a2 uint32) (r, g, b, a uint32) {
		return blendWithAlpha(r1, g1, b1, a1, r2, g2, b2, a2, func(r1, g1, b1, r2, g2, b2 uint32) (uint32, uint32, uint32) {
			// Divide blend mode for un-premultiplied colors
			// Formula: bottom / top

			// Red channel
			var r uint32
			if r2 > 0 {
				r = min((r1*0xffff)/r2, 0xffff)
			}

			// Green channel
			var g uint32
			if g2 > 0 {
				g = min((g1*0xffff)/g2, 0xffff)
			}

			// Blue channel
			var b uint32
			if b2 > 0 {
				b = min((b1*0xffff)/b2, 0xffff)
			}

			return r, g, b
		})
	})
}

// @Name: blend-hue
// @Desc: Blends the two images using the hue blend mode
// @Param:      imgA     - -   	-   The bottom image
// @Param:      imgB     - -   	-   The top image
// @Returns:    result  - -   	-   The blended image
func blendHue(imgA *image.NRGBA64, imgB *image.NRGBA64) (*image.NRGBA64, error) {
	return parallelBlend(imgA, imgB, func(r1, g1, b1, a1, r2, g2, b2, a2 uint32) (r, g, b, a uint32) {
		return blendWithAlpha(r1, g1, b1, a1, r2, g2, b2, a2, func(r1, g1, b1, r2, g2, b2 uint32) (uint32, uint32, uint32) {
			// Hue blend mode for un-premultiplied colors
			// Take hue from top layer, saturation and luminosity from bottom

			// Convert to HSL
			_, s1, l1 := convertRGBToHSL(r1, g1, b1)
			h2, _, _ := convertRGBToHSL(r2, g2, b2)

			// Take hue from top layer, saturation and luminosity from bottom
			r, g, b := convertHSLToRGB(h2, s1, l1)
			return r, g, b
		})
	})
}

// @Name: blend-saturation
// @Desc: Blends the two images using the saturation blend mode
// @Param:      imgA     - -   	-   The bottom image
// @Param:      imgB     - -   	-   The top image
// @Returns:    result  - -   	-   The blended image
func blendSaturation(imgA *image.NRGBA64, imgB *image.NRGBA64) (*image.NRGBA64, error) {
	return parallelBlend(imgA, imgB, func(r1, g1, b1, a1, r2, g2, b2, a2 uint32) (r, g, b, a uint32) {
		return blendWithAlpha(r1, g1, b1, a1, r2, g2, b2, a2, func(r1, g1, b1, r2, g2, b2 uint32) (uint32, uint32, uint32) {
			// Saturation blend mode for un-premultiplied colors
			// Take hue and luminosity from bottom layer, saturation from top

			// Convert to HSL
			h1, _, l1 := convertRGBToHSL(r1, g1, b1)
			_, s2, _ := convertRGBToHSL(r2, g2, b2)

			// Take hue and luminosity from bottom layer, saturation from top
			r, g, b := convertHSLToRGB(h1, s2, l1)
			return r, g, b
		})
	})
}

// @Name: blend-color
// @Desc: Blends the two images using the color blend mode
// @Param:      imgA     - -   	-   The bottom image
// @Param:      imgB     - -   	-   The top image
// @Returns:    result  - -   	-   The blended image
func blendColor(imgA *image.NRGBA64, imgB *image.NRGBA64) (*image.NRGBA64, error) {
	return parallelBlend(imgA, imgB, func(r1, g1, b1, a1, r2, g2, b2, a2 uint32) (r, g, b, a uint32) {
		return blendWithAlpha(r1, g1, b1, a1, r2, g2, b2, a2, func(r1, g1, b1, r2, g2, b2 uint32) (uint32, uint32, uint32) {
			// Color blend mode for un-premultiplied colors
			// Take hue and saturation from top layer, luminosity from bottom

			// Convert to HSL
			_, _, l1 := convertRGBToHSL(r1, g1, b1)
			h2, s2, _ := convertRGBToHSL(r2, g2, b2)

			// Take hue and saturation from top layer, luminosity from bottom
			r, g, b := convertHSLToRGB(h2, s2, l1)
			return r, g, b
		})
	})
}

// @Name: blend-luminosity
// @Desc: Blends the two images using the luminosity blend mode
// @Param:      imgA     - -   	-   The bottom image
// @Param:      imgB     - -   	-   The top image
// @Returns:    result  - -   	-   The blended image
func blendLuminosity(imgA *image.NRGBA64, imgB *image.NRGBA64) (*image.NRGBA64, error) {
	return parallelBlend(imgA, imgB, func(r1, g1, b1, a1, r2, g2, b2, a2 uint32) (r, g, b, a uint32) {
		return blendWithAlpha(r1, g1, b1, a1, r2, g2, b2, a2, func(r1, g1, b1, r2, g2, b2 uint32) (uint32, uint32, uint32) {
			// Luminosity blend mode for un-premultiplied colors
			// Take hue and saturation from bottom layer, luminosity from top

			// Convert to HSL
			h1, s1, _ := convertRGBToHSL(r1, g1, b1)
			_, _, l2 := convertRGBToHSL(r2, g2, b2)

			// Take hue and saturation from bottom layer, luminosity from top
			r, g, b := convertHSLToRGB(h1, s1, l2)
			return r, g, b
		})
	})
}

// @Name: blend-average
// @Desc: Blends the two images using the average blend mode
// @Param:      imgA     - -   	-   The bottom image
// @Param:      imgB     - -   	-   The top image
// @Returns:    result  - -   	-   The blended image
func blendAverage(imgA *image.NRGBA64, imgB *image.NRGBA64) (*image.NRGBA64, error) {
	return parallelBlend(imgA, imgB, func(r1, g1, b1, a1, r2, g2, b2, a2 uint32) (r, g, b, a uint32) {
		return blendWithAlpha(r1, g1, b1, a1, r2, g2, b2, a2, func(r1, g1, b1, r2, g2, b2 uint32) (uint32, uint32, uint32) {
			// Average blend mode for un-premultiplied colors
			// Simple average of the two colors
			r := (r1 + r2) / 2
			g := (g1 + g2) / 2
			b := (b1 + b2) / 2
			return r, g, b
		})
	})
}

// @Name: blend-negation
// @Desc: Blends the two images using the negation blend mode
// @Param:      imgA     - -   	-   The bottom image
// @Param:      imgB     - -   	-   The top image
// @Returns:    result  - -   	-   The blended image
func blendNegation(imgA *image.NRGBA64, imgB *image.NRGBA64) (*image.NRGBA64, error) {
	return parallelBlend(imgA, imgB, func(r1, g1, b1, a1, r2, g2, b2, a2 uint32) (r, g, b, a uint32) {
		return blendWithAlpha(r1, g1, b1, a1, r2, g2, b2, a2, func(r1, g1, b1, r2, g2, b2 uint32) (uint32, uint32, uint32) {
			// Negation blend mode for un-premultiplied colors
			// Formula: 1 - |1 - bottom - top|
			r := 0xffff - absDiff(0xffff-r1, r2)
			g := 0xffff - absDiff(0xffff-g1, g2)
			b := 0xffff - absDiff(0xffff-b1, b2)
			return r, g, b
		})
	})
}

// @Name: blend-reflect
// @Desc: Blends the two images using the reflect blend mode
// @Param:      imgA     - -   	-   The bottom image
// @Param:      imgB     - -   	-   The top image
// @Returns:    result  - -   	-   The blended image
func blendReflect(imgA *image.NRGBA64, imgB *image.NRGBA64) (*image.NRGBA64, error) {
	return parallelBlend(imgA, imgB, func(r1, g1, b1, a1, r2, g2, b2, a2 uint32) (r, g, b, a uint32) {
		return blendWithAlpha(r1, g1, b1, a1, r2, g2, b2, a2, func(r1, g1, b1, r2, g2, b2 uint32) (uint32, uint32, uint32) {
			// Reflect blend mode for un-premultiplied colors
			// Formula: bottom^2 / (1 - top)

			// Red channel
			var r uint32
			if r2 == 0xffff {
				r = 0xffff
			} else if r2 < 0xffff {
				r = min((r1*r1)/(0xffff-r2), 0xffff)
			}

			// Green channel
			var g uint32
			if g2 == 0xffff {
				g = 0xffff
			} else if g2 < 0xffff {
				g = min((g1*g1)/(0xffff-g2), 0xffff)
			}

			// Blue channel
			var b uint32
			if b2 == 0xffff {
				b = 0xffff
			} else if b2 < 0xffff {
				b = min((b1*b1)/(0xffff-b2), 0xffff)
			}

			return r, g, b
		})
	})
}

// @Name: blend-glow
// @Desc: Blends the two images using the glow blend mode
// @Param:      imgA     - -   	-   The bottom image
// @Param:      imgB     - -   	-   The top image
// @Returns:    result  - -   	-   The blended image
func blendGlow(imgA *image.NRGBA64, imgB *image.NRGBA64) (*image.NRGBA64, error) {
	return parallelBlend(imgA, imgB, func(r1, g1, b1, a1, r2, g2, b2, a2 uint32) (r, g, b, a uint32) {
		return blendWithAlpha(r1, g1, b1, a1, r2, g2, b2, a2, func(r1, g1, b1, r2, g2, b2 uint32) (uint32, uint32, uint32) {
			// Glow blend mode for un-premultiplied colors
			// Formula: top^2 / (1 - bottom)

			// Red channel
			var r uint32
			if r1 == 0xffff {
				r = 0xffff
			} else if r1 < 0xffff {
				r = min((r2*r2)/(0xffff-r1), 0xffff)
			}

			// Green channel
			var g uint32
			if g1 == 0xffff {
				g = 0xffff
			} else if g1 < 0xffff {
				g = min((g2*g2)/(0xffff-g1), 0xffff)
			}

			// Blue channel
			var b uint32
			if b1 == 0xffff {
				b = 0xffff
			} else if b1 < 0xffff {
				b = min((b2*b2)/(0xffff-b1), 0xffff)
			}

			return r, g, b
		})
	})
}

// @Name: blend-contrast-negate
// @Desc: Blends the two images using the contrast negate blend mode
// @Param:      imgA     - -   	-   The bottom image
// @Param:      imgB     - -   	-   The top image
// @Returns:    result  - -   	-   The blended image
func blendContrastNegate(imgA *image.NRGBA64, imgB *image.NRGBA64) (*image.NRGBA64, error) {
	return parallelBlend(imgA, imgB, func(r1, g1, b1, a1, r2, g2, b2, a2 uint32) (r, g, b, a uint32) {
		return blendWithAlpha(r1, g1, b1, a1, r2, g2, b2, a2, func(r1, g1, b1, r2, g2, b2 uint32) (uint32, uint32, uint32) {
			// Contrast negate blend mode for un-premultiplied colors
			// Formula: Inverts colors based on relative brightness

			// Calculate luminosity of blend color (simple average)
			blendLum := (r2 + g2 + b2) / 3
			bottomLum := (r1 + g1 + b1) / 3

			const mid = 0x8000

			var r, g, b uint32
			if blendLum > mid {
				if bottomLum < mid {
					r, g, b = r2, g2, b2
				} else {
					r, g, b = 0xffff-r2, 0xffff-g2, 0xffff-b2
				}
			} else {
				if bottomLum > mid {
					r, g, b = r2, g2, b2
				} else {
					r, g, b = 0xffff-r2, 0xffff-g2, 0xffff-b2
				}
			}

			return r, g, b
		})
	})
}

// @Name: blend-vivid-light
// @Desc: Blends the two images using the vivid light blend mode
// @Param:      imgA     - -   	-   The bottom image
// @Param:      imgB     - -   	-   The top image
// @Returns:    result  - -   	-   The blended image
func blendVividLight(imgA *image.NRGBA64, imgB *image.NRGBA64) (*image.NRGBA64, error) {
	return parallelBlend(imgA, imgB, func(r1, g1, b1, a1, r2, g2, b2, a2 uint32) (r, g, b, a uint32) {
		return blendWithAlpha(r1, g1, b1, a1, r2, g2, b2, a2, func(r1, g1, b1, r2, g2, b2 uint32) (uint32, uint32, uint32) {
			// Vivid Light blend mode for un-premultiplied colors
			// Combination of Color Dodge and Color Burn
			// If top color > 50%, use color dodge, otherwise use color burn

			// Red channel
			var r uint32
			if r2 < 0x8000 {
				// Color burn formula
				if r2 == 0 {
					r = 0
				} else {
					r = 0xffff - ((0xffff-r1)*0x8000)/r2
					if r > 0xffff {
						r = 0
					}
				}
			} else {
				// Color dodge formula
				if r2 == 0xffff {
					r = 0xffff
				} else {
					r = min((r1*0x8000)/(0xffff-r2), 0xffff)
				}
			}

			// Green channel
			var g uint32
			if g2 < 0x8000 {
				if g2 == 0 {
					g = 0
				} else {
					g = 0xffff - ((0xffff-g1)*0x8000)/g2
					if g > 0xffff {
						g = 0
					}
				}
			} else {
				if g2 == 0xffff {
					g = 0xffff
				} else {
					g = min((g1*0x8000)/(0xffff-g2), 0xffff)
				}
			}

			// Blue channel
			var b uint32
			if b2 < 0x8000 {
				if b2 == 0 {
					b = 0
				} else {
					b = 0xffff - ((0xffff-b1)*0x8000)/b2
					if b > 0xffff {
						b = 0
					}
				}
			} else {
				if b2 == 0xffff {
					b = 0xffff
				} else {
					b = min((b1*0x8000)/(0xffff-b2), 0xffff)
				}
			}

			return r, g, b
		})
	})
}

// @Name: blend-linear-light
// @Desc: Blends the two images using the linear light blend mode
// @Param:      imgA     - -   	-   The bottom image
// @Param:      imgB     - -   	-   The top image
// @Returns:    result  - -   	-   The blended image
func blendLinearLight(imgA *image.NRGBA64, imgB *image.NRGBA64) (*image.NRGBA64, error) {
	return parallelBlend(imgA, imgB, func(r1, g1, b1, a1, r2, g2, b2, a2 uint32) (r, g, b, a uint32) {
		return blendWithAlpha(r1, g1, b1, a1, r2, g2, b2, a2, func(r1, g1, b1, r2, g2, b2 uint32) (uint32, uint32, uint32) {
			// Linear Light blend mode for un-premultiplied colors
			// Formula: Bottom + 2 * (Top - 128)
			// If result < 0, clamp to 0; if > 255, clamp to 255
			r := uint32(math.Clamp(int64(r1)+2*int64(r2)-0xffff, 0, 0xffff))
			g := uint32(math.Clamp(int64(g1)+2*int64(g2)-0xffff, 0, 0xffff))
			b := uint32(math.Clamp(int64(b1)+2*int64(b2)-0xffff, 0, 0xffff))
			return r, g, b
		})
	})
}

// @Name: blend-pin-light
// @Desc: Blends the two images using the pin light blend mode
// @Param:      imgA     - -   	-   The bottom image
// @Param:      imgB     - -   	-   The top image
// @Returns:    result  - -   	-   The blended image
func blendPinLight(imgA *image.NRGBA64, imgB *image.NRGBA64) (*image.NRGBA64, error) {
	return parallelBlend(imgA, imgB, func(r1, g1, b1, a1, r2, g2, b2, a2 uint32) (r, g, b, a uint32) {
		return blendWithAlpha(r1, g1, b1, a1, r2, g2, b2, a2, func(r1, g1, b1, r2, g2, b2 uint32) (uint32, uint32, uint32) {
			// Pin Light blend mode for un-premultiplied colors
			// If top color > 50%, use Lighten blend mode
			// If top color < 50%, use Darken blend mode

			// Red channel
			var r uint32
			if r2 < 0x8000 {
				// Darken
				r = min(r1, 2*r2)
			} else {
				// Lighten
				r = max(r1, 2*(r2-0x8000))
			}

			// Green channel
			var g uint32
			if g2 < 0x8000 {
				g = min(g1, 2*g2)
			} else {
				g = max(g1, 2*(g2-0x8000))
			}

			// Blue channel
			var b uint32
			if b2 < 0x8000 {
				b = min(b1, 2*b2)
			} else {
				b = max(b1, 2*(b2-0x8000))
			}

			return r, g, b
		})
	})
}

// @Name: blend-darken
// @Desc: Blends the two images using the darken blend mode
// @Param:      imgA     - -   	-   The bottom image
// @Param:      imgB     - -   	-   The top image
// @Returns:    result  - -   	-   The blended image
func blendDarken(imgA *image.NRGBA64, imgB *image.NRGBA64) (*image.NRGBA64, error) {
	return parallelBlend(imgA, imgB, func(r1, g1, b1, a1, r2, g2, b2, a2 uint32) (r, g, b, a uint32) {
		return blendWithAlpha(r1, g1, b1, a1, r2, g2, b2, a2, func(r1, g1, b1, r2, g2, b2 uint32) (uint32, uint32, uint32) {
			// Darken blend mode for un-premultiplied colors
			// Formula: min(bottom, top)
			r := min(r1, r2)
			g := min(g1, g2)
			b := min(b1, b2)
			return r, g, b
		})
	})
}

// @Name: blend-darker-color
// @Desc: Blends the two images using the darker color blend mode
// @Param:      imgA     - -   	-   The bottom image
// @Param:      imgB     - -   	-   The top image
// @Returns:    result  - -   	-   The blended image
func blendDarkerColor(imgA *image.NRGBA64, imgB *image.NRGBA64) (*image.NRGBA64, error) {
	return parallelBlend(imgA, imgB, func(r1, g1, b1, a1, r2, g2, b2, a2 uint32) (r, g, b, a uint32) {
		return blendWithAlpha(r1, g1, b1, a1, r2, g2, b2, a2, func(r1, g1, b1, r2, g2, b2 uint32) (uint32, uint32, uint32) {
			// Darker color blend mode for un-premultiplied colors
			// Formula: Choose the color with lower luminosity

			// Calculate luminosity for both pixels using standard weights
			lum1 := (r1*299 + g1*587 + b1*114) / 1000
			lum2 := (r2*299 + g2*587 + b2*114) / 1000

			// Choose the darker color based on luminosity
			if lum1 < lum2 {
				return r1, g1, b1
			}
			return r2, g2, b2
		})
	})
}

// @Name: blend-lighten
// @Desc: Blends the two images using the lighten blend mode
// @Param:      imgA     - -   	-   The bottom image
// @Param:      imgB     - -   	-   The top image
// @Returns:    result  - -   	-   The blended image
func blendLighten(imgA *image.NRGBA64, imgB *image.NRGBA64) (*image.NRGBA64, error) {
	return parallelBlend(imgA, imgB, func(r1, g1, b1, a1, r2, g2, b2, a2 uint32) (r, g, b, a uint32) {
		return blendWithAlpha(r1, g1, b1, a1, r2, g2, b2, a2, func(r1, g1, b1, r2, g2, b2 uint32) (uint32, uint32, uint32) {
			// Lighten blend mode for un-premultiplied colors
			// Formula: max(bottom, top)
			r := max(r1, r2)
			g := max(g1, g2)
			b := max(b1, b2)
			return r, g, b
		})
	})
}

// @Name: blend-lighter-color
// @Desc: Blends the two images using the lighter color blend mode
// @Param:      imgA     - -   	-   The bottom image
// @Param:      imgB     - -   	-   The top image
// @Returns:    result  - -   	-   The blended image
func blendLighterColor(imgA *image.NRGBA64, imgB *image.NRGBA64) (*image.NRGBA64, error) {
	return parallelBlend(imgA, imgB, func(r1, g1, b1, a1, r2, g2, b2, a2 uint32) (r, g, b, a uint32) {
		return blendWithAlpha(r1, g1, b1, a1, r2, g2, b2, a2, func(r1, g1, b1, r2, g2, b2 uint32) (uint32, uint32, uint32) {
			// Lighter color blend mode for un-premultiplied colors
			// Formula: Choose the color with higher luminosity

			// Calculate luminosity for both pixels using standard weights
			lum1 := (r1*299 + g1*587 + b1*114) / 1000
			lum2 := (r2*299 + g2*587 + b2*114) / 1000

			// Choose the lighter color based on luminosity
			if lum1 > lum2 {
				return r1, g1, b1
			}
			return r2, g2, b2
		})
	})
}

// @Name: blend-hard-mix
// @Desc: Blends the two images using the hard mix blend mode
// @Param:      imgA     - -   	-   The bottom image
// @Param:      imgB     - -   	-   The top image
// @Returns:    result  - -   	-   The blended image
func blendHardMix(imgA *image.NRGBA64, imgB *image.NRGBA64) (*image.NRGBA64, error) {
	return parallelBlend(imgA, imgB, func(r1, g1, b1, a1, r2, g2, b2, a2 uint32) (r, g, b, a uint32) {
		return blendWithAlpha(r1, g1, b1, a1, r2, g2, b2, a2, func(r1, g1, b1, r2, g2, b2 uint32) (uint32, uint32, uint32) {
			// Hard mix blend mode for un-premultiplied colors
			// Formula: Apply vivid light and then threshold to pure black or white

			// Red channel (vivid light)
			var vR uint32
			if r2 < 0x8000 {
				if r2 == 0 {
					vR = 0
				} else {
					vR = 0xffff - ((0xffff-r1)*0x8000)/r2
					if vR > 0xffff {
						vR = 0
					}
				}
			} else {
				if r2 == 0xffff {
					vR = 0xffff
				} else {
					vR = min((r1*0x8000)/(0xffff-r2), 0xffff)
				}
			}
			r := uint32(0)
			if vR >= 0x8000 {
				r = 0xffff
			}

			// Green channel (vivid light)
			var vG uint32
			if g2 < 0x8000 {
				if g2 == 0 {
					vG = 0
				} else {
					vG = 0xffff - ((0xffff-g1)*0x8000)/g2
					if vG > 0xffff {
						vG = 0
					}
				}
			} else {
				if g2 == 0xffff {
					vG = 0xffff
				} else {
					vG = min((g1*0x8000)/(0xffff-g2), 0xffff)
				}
			}
			g := uint32(0)
			if vG >= 0x8000 {
				g = 0xffff
			}

			// Blue channel (vivid light)
			var vB uint32
			if b2 < 0x8000 {
				if b2 == 0 {
					vB = 0
				} else {
					vB = 0xffff - ((0xffff-b1)*0x8000)/b2
					if vB > 0xffff {
						vB = 0
					}
				}
			} else {
				if b2 == 0xffff {
					vB = 0xffff
				} else {
					vB = min((b1*0x8000)/(0xffff-b2), 0xffff)
				}
			}
			b := uint32(0)
			if vB >= 0x8000 {
				b = 0xffff
			}

			return r, g, b
		})
	})
}
