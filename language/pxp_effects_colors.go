package language

import (
	"image"
	"image/color"

	"github.com/toxyl/math"
)

// @Name: invert
// @Desc: Inverts an image
// @Param:      img     - -   -   The image to invert
// @Returns:    result  - -   -   The inverted image
func colorInvert(img *image.NRGBA64) (*image.NRGBA64, error) {
	return dsl.parallelProcessNRGBA64(img, func(r1, g1, b1, a1 uint32) (r, g, b, a uint32) {
		r = 0xFFFF - r1
		g = 0xFFFF - g1
		b = 0xFFFF - b1
		a = a1 // Keep original alpha
		return
	}, NumColorConversionWorkers), nil
}

// @Name: grayscale
// @Desc: Grayscales an image
// @Param:      img     - -   -   The image to grayscale
// @Returns:    result  - -   -   The grayscaled image
func colorGrayscale(img *image.NRGBA64) (*image.NRGBA64, error) {
	return dsl.parallelProcessNRGBA64(img, func(r1, g1, b1, a1 uint32) (r, g, b, a uint32) {
		// Using luminosity method: 0.21 R + 0.72 G + 0.07 B
		gray := uint32(float64(r1)*0.21 + float64(g1)*0.72 + float64(b1)*0.07)
		r = gray
		g = gray
		b = gray
		a = a1 // Keep original alpha
		return
	}, NumColorConversionWorkers), nil
}

// @Name: sepia
// @Desc: Changes the tone of an image to sepia
// @Param:      img     - -   -   The image to change to sepia tone
// @Returns:    result  - -   -   The sepia-toned image
func colorSepia(img *image.NRGBA64) (*image.NRGBA64, error) {
	return dsl.parallelProcessNRGBA64(img, func(r1, g1, b1, a1 uint32) (r, g, b, a uint32) {
		rf := float64(r1)
		gf := float64(g1)
		bf := float64(b1)
		r = uint32(math.Min((rf*0.393)+(gf*0.769)+(bf*0.189), 65535))
		g = uint32(math.Min((rf*0.349)+(gf*0.686)+(bf*0.168), 65535))
		b = uint32(math.Min((rf*0.272)+(gf*0.534)+(bf*0.131), 65535))
		a = a1 // Keep original alpha
		return
	}, NumColorConversionWorkers), nil
}

// @Name: brightness
// @Desc: Changes the brightness of an image
// @Param:      img     - -   	-   The image to change brightness of
// @Param:      factor  - 0..2  0   The change factor
// @Returns:    result  - -   	-   The image with brightness changed
func colorBrightness(img *image.NRGBA64, factor float64) (*image.NRGBA64, error) {
	return dsl.parallelProcessNRGBA64(img, func(r1, g1, b1, a1 uint32) (r, g, b, a uint32) {
		r = uint32(math.Min(float64(r1)*factor, 0xFFFF))
		g = uint32(math.Min(float64(g1)*factor, 0xFFFF))
		b = uint32(math.Min(float64(b1)*factor, 0xFFFF))
		a = a1 // Keep original alpha
		return
	}, NumColorConversionWorkers), nil
}

// @Name: colorize
// @Desc: Colorizes the image
// @Param:      img     - - -   The image to colorize
// @Param:      col  	- - -   The color that determines the hue to use for colorization
// @Returns:    result  - - -	The colorized image
func colorColorize(img *image.NRGBA64, col color.RGBA64) (*image.NRGBA64, error) {
	// Convert target color to normalized RGB and get alpha
	targetR := float64(col.R) / 65535.0
	targetG := float64(col.G) / 65535.0
	targetB := float64(col.B) / 65535.0
	alpha := float64(col.A) / 65535.0

	// Convert target color to HSL to get hue and saturation
	targetH, targetS, targetL := rgbToHsl(targetR, targetG, targetB)

	return dsl.parallelProcessNRGBA64(img, func(r1, g1, b1, a1 uint32) (r, g, b, a uint32) {
		// Convert pixel to normalized RGB
		rf := float64(r1) / 65535.0
		gf := float64(g1) / 65535.0
		bf := float64(b1) / 65535.0

		// Convert original pixel to HSL
		_, _, originalL := rgbToHsl(rf, gf, bf)

		// Calculate new luminance by blending original and target luminance
		// This preserves the image's contrast while allowing some influence from target luminance
		newL := originalL*(1-alpha*0.5) + targetL*(alpha*0.5)

		// Calculate new saturation by blending original and target saturation
		// Original saturation is calculated from the RGB values
		originalS := calculateSaturation(rf, gf, bf)
		newS := originalS*(1-alpha) + targetS*alpha

		// Convert back to RGB using the new HSL values
		newR, newG, newB := hslToRgb(targetH, newS, newL)

		// Blend with original color based on alpha
		r = uint32(math.Clamp(rf*(1-alpha)+newR*alpha, 0, 1) * 65535.0)
		g = uint32(math.Clamp(gf*(1-alpha)+newG*alpha, 0, 1) * 65535.0)
		b = uint32(math.Clamp(bf*(1-alpha)+newB*alpha, 0, 1) * 65535.0)
		a = a1
		return
	}, NumColorConversionWorkers), nil
}

// @Name: contrast
// @Desc: Adjusts the contrast of an image
// @Param:      img     - -   	-   The image to adjust contrast of
// @Param:      factor  - 0..2  1   The contrast factor (0 = gray, 1 = unchanged, 2 = maximum)
// @Returns:    result  - -   	-   The contrast-adjusted image
func colorContrast(img *image.NRGBA64, factor float64) (*image.NRGBA64, error) {
	mid := float64(0x7FFF)
	return dsl.parallelProcessNRGBA64(img, func(r1, g1, b1, a1 uint32) (r, g, b, a uint32) {
		// Apply contrast formula: ((color - mid) * factor) + mid
		r = uint32(math.Clamp(((float64(r1)-mid)*factor)+mid, 0, 0xFFFF))
		g = uint32(math.Clamp(((float64(g1)-mid)*factor)+mid, 0, 0xFFFF))
		b = uint32(math.Clamp(((float64(b1)-mid)*factor)+mid, 0, 0xFFFF))
		a = a1
		return
	}, NumColorConversionWorkers), nil
}

// @Name: saturation
// @Desc: Adjusts the color saturation of an image
// @Param:      img     - -   	-   The image to adjust saturation of
// @Param:      factor  - 0..2  1   The saturation factor (0 = grayscale, 1 = unchanged, 2 = super saturated)
// @Returns:    result  - -   	-   The saturation-adjusted image
func colorSaturation(img *image.NRGBA64, factor float64) (*image.NRGBA64, error) {
	return dsl.parallelProcessNRGBA64(img, func(r1, g1, b1, a1 uint32) (r, g, b, a uint32) {
		// Convert to HSL
		h, s, l := rgbToHsl(float64(r1)/65535.0, float64(g1)/65535.0, float64(b1)/65535.0)

		// Adjust saturation
		s = math.Min(s*factor, 1.0)

		// Convert back to RGB
		rf, gf, bf := hslToRgb(h, s, l)

		r = uint32(rf * 65535.0)
		g = uint32(gf * 65535.0)
		b = uint32(bf * 65535.0)
		a = a1
		return
	}, NumColorConversionWorkers), nil
}

// @Name: opacity
// @Desc: Adjusts the overall opacity/transparency of an image
// @Param:      img      - -   		-   The image to adjust opacity of
// @Param:      amount   - 0..1  	1   The opacity amount (0 = fully transparent, 1 = unchanged)
// @Returns:    result   - -   		-   The opacity-adjusted image
func colorOpacity(img *image.NRGBA64, amount float64) (*image.NRGBA64, error) {
	alphaMultiplier := uint32(amount * 65535.0)
	return dsl.parallelProcessNRGBA64(img, func(r1, g1, b1, a1 uint32) (r, g, b, a uint32) {
		// Multiply existing alpha by the opacity amount while preserving color
		return r1, g1, b1, uint32((a1 * alphaMultiplier) / 65535)
	}, NumColorConversionWorkers), nil
}

// @Name: chromatic-aberration
// @Desc: Creates a chromatic aberration effect by offsetting color channels
// @Param:      img      - -   	-   The image to apply the effect to
// @Param:      amount   - 0..20 5   The amount of color channel separation
// @Returns:    result   - -   	-   The image with chromatic aberration
func colorChromaticAberration(img *image.NRGBA64, amount float64) (*image.NRGBA64, error) {
	bounds := img.Bounds()
	width := bounds.Dx()
	height := bounds.Dy()
	result := IFromBounds(bounds)

	// Create offset lookup tables for optimization
	offsetX := int(amount)

	for y := range height {
		for x := range width {
			var r, g, b uint16
			a := uint16(0xFFFF)

			// Get red from offset left
			if x-offsetX >= 0 {
				r = img.NRGBA64At(x-offsetX+bounds.Min.X, y+bounds.Min.Y).R
			}

			// Get green from original position
			g = img.NRGBA64At(x+bounds.Min.X, y+bounds.Min.Y).G

			// Get blue from offset right
			if x+offsetX < width {
				b = img.NRGBA64At(x+offsetX+bounds.Min.X, y+bounds.Min.Y).B
			}

			result.Set(x+bounds.Min.X, y+bounds.Min.Y, color.NRGBA64{R: r, G: g, B: b, A: a})
		}
	}

	return result, nil
}

// @Name: hue-rotate
// @Desc: Rotates the hue of image colors
// @Param:      img     - 	-   	-   The image to rotate hue of
// @Param:      angle   "°" 0..360 	0  The angle in degrees (0-360)
// @Returns:    result  - 	-   	-   The hue-rotated image
func colorHueRotate(img *image.NRGBA64, angle float64) (*image.NRGBA64, error) {
	hueShift := angle / 360.0 // Normalize angle to 0.0-1.0 range for HSL calculation
	return dsl.parallelProcessNRGBA64(img, func(r1, g1, b1, a1 uint32) (r, g, b, a uint32) {
		h, s, l := rgbToHsl(float64(r1)/65535.0, float64(g1)/65535.0, float64(b1)/65535.0)

		// Rotate hue
		h += hueShift
		if h > 1.0 {
			h -= 1.0
		} else if h < 0.0 {
			h += 1.0
		}

		rf, gf, bf := hslToRgb(h, s, l)

		r = uint32(rf * 65535.0)
		g = uint32(gf * 65535.0)
		b = uint32(bf * 65535.0)
		a = a1
		return
	}, NumColorConversionWorkers), nil
}

// @Name: color-balance
// @Desc: Adjusts the balance of Red, Green, and Blue channels
// @Param:      img       - -   	-   The image to adjust color balance of
// @Param:      rFactor   - 0..2  	1   Red channel adjustment factor
// @Param:      gFactor   - 0..2  	1   Green channel adjustment factor
// @Param:      bFactor   - 0..2  	1   Blue channel adjustment factor
// @Returns:    result    - -   	-   The color-balanced image
func colorBalance(img *image.NRGBA64, rFactor, gFactor, bFactor float64) (*image.NRGBA64, error) {
	return dsl.parallelProcessNRGBA64(img, func(r1, g1, b1, a1 uint32) (r, g, b, a uint32) {
		r = uint32(math.Clamp(float64(r1)*rFactor, 0, 0xFFFF))
		g = uint32(math.Clamp(float64(g1)*gFactor, 0, 0xFFFF))
		b = uint32(math.Clamp(float64(b1)*bFactor, 0, 0xFFFF))
		a = a1
		return
	}, NumColorConversionWorkers), nil
}

// @Name: posterize
// @Desc: Reduces the number of color levels in the image
// @Param:      img     - -   	-   The image to posterize
// @Param:      levels  - 2..16 4   Number of color levels per channel (2-16)
// @Returns:    result  - -   	-   The posterized image
func colorPosterize(img *image.NRGBA64, levels int) (*image.NRGBA64, error) {
	numLevels := float64(levels)
	levelStep := 65535.0 / (numLevels - 1) // Size of each color step
	return dsl.parallelProcessNRGBA64(img, func(r1, g1, b1, a1 uint32) (r, g, b, a uint32) {
		// Quantize each channel
		r = uint32(math.Round(float64(r1)/levelStep) * levelStep)
		g = uint32(math.Round(float64(g1)/levelStep) * levelStep)
		b = uint32(math.Round(float64(b1)/levelStep) * levelStep)
		a = a1
		return
	}, NumColorConversionWorkers), nil
}

// @Name: threshold
// @Desc: Converts image to black and white based on a brightness threshold
// @Param:      img     - -   	-   The image to apply thresholding to
// @Param:      level   - 0..1 	0.5 The brightness threshold
// @Returns:    result  - -   	-   The thresholded (black and white) image
func colorThreshold(img *image.NRGBA64, level float64) (*image.NRGBA64, error) {
	thresholdLevel := float64(level) * 65535.0
	return dsl.parallelProcessNRGBA64(img, func(r1, g1, b1, a1 uint32) (r, g, b, a uint32) {
		// Using luminosity method (same as grayscale)
		if float64(r1)*0.21+float64(g1)*0.72+float64(b1)*0.07 > thresholdLevel {
			r, g, b, a = 0xFFFF, 0xFFFF, 0xFFFF, a1
		} else {
			r, g, b, a = 0, 0, 0, a1
		}
		return
	}, NumColorConversionWorkers), nil
}

// @Name: edge-detect
// @Desc: Detects edges in the image using the Sobel operator
// @Param:      img     - -   	-   The image to detect edges in
// @Returns:    result  - -   	-   An image highlighting the edges
func colorEdgeDetect(img *image.NRGBA64) (*image.NRGBA64, error) {
	bounds := img.Bounds()
	edges := IFromBounds(bounds)

	// Sobel kernels
	sobelX := [][]float64{
		{-1, 0, 1},
		{-2, 0, 2},
		{-1, 0, 1},
	}
	sobelY := [][]float64{
		{-1, -2, -1},
		{0, 0, 0},
		{1, 2, 1},
	}

	// Helper function to get grayscale value (luminosity method)
	getGray := func(x, y int) float64 {
		if x < bounds.Min.X || x >= bounds.Max.X || y < bounds.Min.Y || y >= bounds.Max.Y {
			return 0 // Handle boundaries
		}
		c := img.NRGBA64At(x, y)
		return float64(c.R)*0.21 + float64(c.G)*0.72 + float64(c.B)*0.07
	}

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			var gx, gy float64

			// Apply Sobel kernels
			for ky := -1; ky <= 1; ky++ {
				for kx := -1; kx <= 1; kx++ {
					grayVal := getGray(x+kx, y+ky)
					gx += grayVal * sobelX[ky+1][kx+1]
					gy += grayVal * sobelY[ky+1][kx+1]
				}
			}

			// Calculate gradient magnitude
			magnitude := math.Sqrt(gx*gx + gy*gy)
			// Normalize magnitude to 0-65535 range
			edgeVal := uint16(math.Clamp(magnitude, 0, 65535))

			// Get original alpha
			originalA := img.NRGBA64At(x, y).A

			edges.Set(x, y, color.NRGBA64{
				R: edgeVal,
				G: edgeVal,
				B: edgeVal,
				A: originalA, // Preserve original alpha
			})
		}
	}

	return edges, nil
}

// @Name: vignette
// @Desc: Adds a vignette effect (darkens/lightens edges)
// @Param:      img       - -   		-   The image to apply vignette to
// @Param:      strength  - 0.0..1.0  	0.5 Darkness/Lightness intensity (0 to 1)
// @Param:      falloff   - 0.1..2.0 	0.8 How quickly the effect fades (0.1 to 2.0)
// @Returns:    result    - -   		-   The image with vignette effect
func colorVignette(img *image.NRGBA64, strength float64, falloff float64) (*image.NRGBA64, error) {
	bounds := img.Bounds()
	centerX := float64(bounds.Dx()) / 2.0
	centerY := float64(bounds.Dy()) / 2.0
	result := IFromBounds(bounds)

	// Calculate max distance from center (approx to corner)
	maxDist := math.Sqrt(centerX*centerX + centerY*centerY)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			c := img.NRGBA64At(x, y)

			// Calculate distance from center for the current pixel
			dx := centerX - (float64(x-bounds.Min.X) + 0.5)
			dy := centerY - (float64(y-bounds.Min.Y) + 0.5)
			dist := math.Sqrt(dx*dx+dy*dy) / maxDist // Normalized distance (0 center, 1 corner)

			// Calculate vignette factor (0 = max effect, 1 = no effect)
			// factor := math.Pow(math.Cos(dist*math.Pi/2.0), falloff) // Cosine based falloff
			factor := 1.0 - strength*math.Pow(dist, falloff) // Power based falloff
			factor = math.Clamp(factor, 0.0, 1.0)            // Clamp factor [0, 1]

			// Apply factor to R, G, B
			r := float64(c.R) * factor
			g := float64(c.G) * factor
			b := float64(c.B) * factor

			result.Set(x, y, color.NRGBA64{
				R: uint16(math.Clamp(r, 0, 65535)),
				G: uint16(math.Clamp(g, 0, 65535)),
				B: uint16(math.Clamp(b, 0, 65535)),
				A: c.A,
			})
		}
	}

	return result, nil
}

// @Name: vibrance
// @Desc: Adjusts the saturation of an image, protecting already saturated colors and skin tones.
// @Param:      img     - -   	-   The image to adjust vibrance of
// @Param:      factor  - -1..1 0   The vibrance adjustment factor (-1 = less vibrant, 0 = unchanged, 1 = more vibrant)
// @Returns:    result  - -   	-   The vibrance-adjusted image
func colorVibrance(img *image.NRGBA64, factor float64) (*image.NRGBA64, error) {
	return dsl.parallelProcessNRGBA64(img, func(r1, g1, b1, a1 uint32) (r, g, b, a uint32) {
		h, s, l := rgbToHsl(float64(r1)/65535.0, float64(g1)/65535.0, float64(b1)/65535.0)

		// Calculate saturation adjustment - more effect on less saturated colors
		// Max increase/decrease is scaled by how far current saturation is from max (1.0)
		// or min (0.0) respectively.
		var satAdjust float64
		if factor > 0 {
			// Increase vibrance: more effect on lower saturation
			satAdjust = factor * (1.0 - s) // Scale factor by how unsaturated it is
		} else {
			// Decrease vibrance: more effect on higher saturation
			satAdjust = factor * s // Scale factor by how saturated it is
		}

		newS := s + satAdjust
		newS = math.Clamp(newS, 0.0, 1.0) // Clamp saturation [0, 1]

		// Convert back to RGB
		rf, gf, bf := hslToRgb(h, newS, l)

		r = uint32(rf * 65535.0)
		g = uint32(gf * 65535.0)
		b = uint32(bf * 65535.0)
		a = a1
		return
	}, NumColorConversionWorkers), nil
}

// @Name: exposure
// @Desc: Adjusts the overall lightness or darkness of the image, simulating photographic exposure.
// @Param:      img     - -   	-   The image to adjust exposure of
// @Param:      level   - -2..2 0   The exposure level adjustment (-2 = much darker, 0 = unchanged, 2 = much brighter)
// @Returns:    result  - -   	-   The exposure-adjusted image
func colorExposure(img *image.NRGBA64, level float64) (*image.NRGBA64, error) {
	factor := math.Pow(2, level) // Exponential adjustment factor
	return dsl.parallelProcessNRGBA64(img, func(r1, g1, b1, a1 uint32) (r, g, b, a uint32) {
		r = uint32(math.Clamp(float64(r1)*factor, 0, 0xFFFF))
		g = uint32(math.Clamp(float64(g1)*factor, 0, 0xFFFF))
		b = uint32(math.Clamp(float64(b1)*factor, 0, 0xFFFF))
		a = a1
		return
	}, NumColorConversionWorkers), nil
}

// @Name: select-hue
// @Desc: Selects a specific hue from the image and makes everything else transparent
// @Param:      img            - -   		-   The image to process (16-bit)
// @Param:      hue            "°" 0..360 	0   The target hue to keep (in degrees)
// @Param:      toleranceLeft  "°" 0..180 	30  How much to include to the left (lower hue, in degrees)
// @Param:      toleranceRight "°" 0..180 	30  How much to include to the right (higher hue, in degrees)
// @Param:      softness       - 0..1   	0.5 How soft the transition should be (0 = hard cut, 1 = very soft)
// @Param:      minSaturation  - 0..1 0.15 The minimum saturation required to keep a pixel (smooth fade below)
// @Returns:    result         - -   		-   The image with only the selected hue visible (16-bit)
func colorSelectHue(img *image.NRGBA64, hue, toleranceLeft, toleranceRight, softness, minSaturation float64) (*image.NRGBA64, error) {
	// Convert hue and tolerances to 0-1 range for HSL calculations
	targetHue := hue / 360.0
	tolL := toleranceLeft / 360.0
	tolR := toleranceRight / 360.0
	softnessL := tolL * softness
	softnessR := tolR * softness
	satFade := 0.10 // Range below minSaturation for smooth fade

	bounds := img.Bounds()
	result := IFromBounds(bounds)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			c := img.NRGBA64At(x, y)
			// Normalize to 0-1
			rf := float64(c.R) / 65535.0
			gf := float64(c.G) / 65535.0
			bf := float64(c.B) / 65535.0
			af := float64(c.A) / 65535.0

			// Convert to HSL
			h, s, _ := rgbToHsl(rf, gf, bf)

			// Calculate hue difference, handling circular hue space
			diff := h - targetHue
			if diff < -0.5 {
				diff += 1.0
			} else if diff > 0.5 {
				diff -= 1.0
			}

			var alpha float64
			if diff >= -tolL && diff <= tolR {
				if diff >= -(tolL-softnessL) && diff <= (tolR-softnessR) {
					// Full opacity in the core range
					alpha = 1.0
				} else if diff < -(tolL - softnessL) {
					// Left softness fade
					alpha = 1.0 - (-(tolL-softnessL)-diff)/softnessL
				} else {
					// Right softness fade
					alpha = 1.0 - (diff-(tolR-softnessR))/softnessR
				}
				// Clamp alpha to [0,1]
				if alpha < 0.0 {
					alpha = 0.0
				} else if alpha > 1.0 {
					alpha = 1.0
				}
			} else {
				alpha = 0.0
			}

			// Fade alpha for low saturation (smoothly)
			satAlpha := 1.0
			if s < minSaturation {
				if s < minSaturation-satFade {
					satAlpha = 0.0
				} else {
					satAlpha = (s - (minSaturation - satFade)) / satFade
				}
			}
			alpha *= satAlpha

			// Apply the alpha to the original color, preserve original alpha channel
			finalAlpha := alpha * af

			result.SetNRGBA64(x, y, color.NRGBA64{
				R: uint16(rf * 65535.0),
				G: uint16(gf * 65535.0),
				B: uint16(bf * 65535.0),
				A: uint16(finalAlpha * 65535.0),
			})
		}
	}

	return result, nil
}

// @Name: select-hsl
// @Desc: Selects pixels based on their hue, saturation, and luminance, making pixels outside the specified ranges transparent
// @Param:      img            - -   		-   The image to process (16-bit)
// @Param:      lowerHue       "°" 0..360 	0   The lower hue threshold (below this becomes transparent)
// @Param:      minHue         "°" 0..360 	30  The minimum hue for full opacity (fade from 0% to 100% between lowerHue and this)
// @Param:      maxHue         "°" 0..360 	330 The maximum hue for full opacity (fade from 100% to 0% between this and upperHue)
// @Param:      upperHue       "°" 0..360 	360 The upper hue threshold (above this becomes transparent)
// @Param:      lowerSat       - 0..1 	0.1  The lower saturation threshold (below this becomes transparent)
// @Param:      minSat         - 0..1 	0.2  The minimum saturation for full opacity (fade from 0% to 100% between lowerSat and this)
// @Param:      maxSat         - 0..1 	0.8  The maximum saturation for full opacity (fade from 100% to 0% between this and upperSat)
// @Param:      upperSat       - 0..1 	0.9  The upper saturation threshold (above this becomes transparent)
// @Param:      lowerLum       - 0..1 	0.1  The lower luminance threshold (below this becomes transparent)
// @Param:      minLum         - 0..1 	0.2  The minimum luminance for full opacity (fade from 0% to 100% between lowerLum and this)
// @Param:      maxLum         - 0..1 	0.8  The maximum luminance for full opacity (fade from 100% to 0% between this and upperLum)
// @Param:      upperLum       - 0..1 	0.9  The upper luminance threshold (above this becomes transparent)
// @Returns:    result         - -   		-   The image with only pixels in all specified ranges visible (16-bit)
func colorSelectHSL(img *image.NRGBA64,
	lowerHue, minHue, maxHue, upperHue float64,
	lowerSat, minSat, maxSat, upperSat float64,
	lowerLum, minLum, maxLum, upperLum float64,
) (*image.NRGBA64, error) {

	// Convert hue values to 0-1 range for HSL calculations
	lowerH := lowerHue / 360.0
	minH := minHue / 360.0
	maxH := maxHue / 360.0
	upperH := upperHue / 360.0

	bounds := img.Bounds()
	result := IFromBounds(bounds)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			c := img.NRGBA64At(x, y)
			// Normalize to 0-1
			rf := float64(c.R) / 65535.0
			gf := float64(c.G) / 65535.0
			bf := float64(c.B) / 65535.0
			af := float64(c.A) / 65535.0

			// Convert to HSL
			h, s, l := rgbToHsl(rf, gf, bf)

			// Calculate alpha for each component
			var alphaH, alphaS, alphaL float64

			// Hue alpha
			if h < lowerH || h > upperH {
				alphaH = 0.0
			} else if h < minH {
				alphaH = (h - lowerH) / (minH - lowerH)
			} else if h > maxH {
				alphaH = 1.0 - (h-maxH)/(upperH-maxH)
			} else {
				alphaH = 1.0
			}

			// Saturation alpha
			if s < lowerSat || s > upperSat {
				alphaS = 0.0
			} else if s < minSat {
				alphaS = (s - lowerSat) / (minSat - lowerSat)
			} else if s > maxSat {
				alphaS = 1.0 - (s-maxSat)/(upperSat-maxSat)
			} else {
				alphaS = 1.0
			}

			// Luminance alpha
			if l < lowerLum || l > upperLum {
				alphaL = 0.0
			} else if l < minLum {
				alphaL = (l - lowerLum) / (minLum - lowerLum)
			} else if l > maxLum {
				alphaL = 1.0 - (l-maxLum)/(upperLum-maxLum)
			} else {
				alphaL = 1.0
			}

			// Combine alphas (multiply them together)
			alpha := alphaH * alphaS * alphaL
			alpha = math.Clamp(alpha, 0.0, 1.0)

			// Apply the alpha to the original color, preserve original alpha channel
			finalAlpha := alpha * af

			result.SetNRGBA64(x, y, color.NRGBA64{
				R: uint16(rf * 65535.0),
				G: uint16(gf * 65535.0),
				B: uint16(bf * 65535.0),
				A: uint16(finalAlpha * 65535.0),
			})
		}
	}

	return result, nil
}

// @Name: remove-hsl
// @Desc: Removes pixels based on their hue, saturation, and luminance, making pixels inside the specified ranges transparent
// @Param:      img            - -   		-   The image to process (16-bit)
// @Param:      lowerHue       "°" 0..360 	0   The lower hue threshold (below this becomes transparent)
// @Param:      minHue         "°" 0..360 	30  The minimum hue for full opacity (fade from 0% to 100% between lowerHue and this)
// @Param:      maxHue         "°" 0..360 	330 The maximum hue for full opacity (fade from 100% to 0% between this and upperHue)
// @Param:      upperHue       "°" 0..360 	360 The upper hue threshold (above this becomes transparent)
// @Param:      lowerSat       - 0..1 	0.1  The lower saturation threshold (below this becomes transparent)
// @Param:      minSat         - 0..1 	0.2  The minimum saturation for full opacity (fade from 0% to 100% between lowerSat and this)
// @Param:      maxSat         - 0..1 	0.8  The maximum saturation for full opacity (fade from 100% to 0% between this and upperSat)
// @Param:      upperSat       - 0..1 	0.9  The upper saturation threshold (above this becomes transparent)
// @Param:      lowerLum       - 0..1 	0.1  The lower luminance threshold (below this becomes transparent)
// @Param:      minLum         - 0..1 	0.2  The minimum luminance for full opacity (fade from 0% to 100% between lowerLum and this)
// @Param:      maxLum         - 0..1 	0.8  The maximum luminance for full opacity (fade from 100% to 0% between this and upperLum)
// @Param:      upperLum       - 0..1 	0.9  The upper luminance threshold (above this becomes transparent)
// @Returns:    result         - -   		-   The image with only pixels in all specified ranges visible (16-bit)
func colorRemoveHSL(img *image.NRGBA64,
	lowerHue, minHue, maxHue, upperHue float64,
	lowerSat, minSat, maxSat, upperSat float64,
	lowerLum, minLum, maxLum, upperLum float64,
) (*image.NRGBA64, error) {
	hsl, err := colorSelectHSL(
		img,
		lowerHue, minHue, maxHue, upperHue,
		lowerSat, minSat, maxSat, upperSat,
		lowerLum, minLum, maxLum, upperLum,
	)
	if err != nil {
		return nil, err
	}
	return blendErase(img, hsl)

}

// @Name: invert-hsl
// @Desc: Inverts pixels based on their hue, saturation, and luminance, inverting pixels inside the specified ranges
// @Param:      img            - -   		-   The image to process (16-bit)
// @Param:      lowerHue       "°" 0..360 	0   The lower hue threshold (below this becomes transparent)
// @Param:      minHue         "°" 0..360 	30  The minimum hue for full opacity (fade from 0% to 100% between lowerHue and this)
// @Param:      maxHue         "°" 0..360 	330 The maximum hue for full opacity (fade from 100% to 0% between this and upperHue)
// @Param:      upperHue       "°" 0..360 	360 The upper hue threshold (above this becomes transparent)
// @Param:      lowerSat       - 0..1 	0.1  The lower saturation threshold (below this becomes transparent)
// @Param:      minSat         - 0..1 	0.2  The minimum saturation for full opacity (fade from 0% to 100% between lowerSat and this)
// @Param:      maxSat         - 0..1 	0.8  The maximum saturation for full opacity (fade from 100% to 0% between this and upperSat)
// @Param:      upperSat       - 0..1 	0.9  The upper saturation threshold (above this becomes transparent)
// @Param:      lowerLum       - 0..1 	0.1  The lower luminance threshold (below this becomes transparent)
// @Param:      minLum         - 0..1 	0.2  The minimum luminance for full opacity (fade from 0% to 100% between lowerLum and this)
// @Param:      maxLum         - 0..1 	0.8  The maximum luminance for full opacity (fade from 100% to 0% between this and upperLum)
// @Param:      upperLum       - 0..1 	0.9  The upper luminance threshold (above this becomes transparent)
// @Returns:    result         - -   		-   The image with only pixels in all specified ranges visible (16-bit)
func colorInvertHSL(img *image.NRGBA64,
	lowerHue, minHue, maxHue, upperHue float64,
	lowerSat, minSat, maxSat, upperSat float64,
	lowerLum, minLum, maxLum, upperLum float64,
) (*image.NRGBA64, error) {
	hsl, err := colorSelectHSL(
		img,
		lowerHue, minHue, maxHue, upperHue,
		lowerSat, minSat, maxSat, upperSat,
		lowerLum, minLum, maxLum, upperLum,
	)
	if err != nil {
		return nil, err
	}
	hsl, err = colorInvert(hsl)
	if err != nil {
		return nil, err
	}
	return blendNormal(img, hsl)

}

// @Name: rotate-hsl
// @Desc: Rotates the hue of pixels based on their hue, saturation, and luminance
// @Param:      img            	- 	-   	-   The image to process (16-bit)
// @Param:      rotate       	"°" 0..360 	0   The lower hue threshold (below this becomes transparent)
// @Param:      lowerHue       	"°" 0..360 	0   The lower hue threshold (below this becomes transparent)
// @Param:      minHue         	"°" 0..360 	30  The minimum hue for full opacity (fade from 0% to 100% between lowerHue and this)
// @Param:      maxHue         	"°" 0..360 	330 The maximum hue for full opacity (fade from 100% to 0% between this and upperHue)
// @Param:      upperHue       	"°" 0..360 	360 The upper hue threshold (above this becomes transparent)
// @Param:      lowerSat       	- 	0..1 	0.1 The lower saturation threshold (below this becomes transparent)
// @Param:      minSat         	- 	0..1 	0.2 The minimum saturation for full opacity (fade from 0% to 100% between lowerSat and this)
// @Param:      maxSat         	- 	0..1 	0.8 The maximum saturation for full opacity (fade from 100% to 0% between this and upperSat)
// @Param:      upperSat       	- 	0..1 	0.9 The upper saturation threshold (above this becomes transparent)
// @Param:      lowerLum       	- 	0..1 	0.1 The lower luminance threshold (below this becomes transparent)
// @Param:      minLum         	- 	0..1 	0.2 The minimum luminance for full opacity (fade from 0% to 100% between lowerLum and this)
// @Param:      maxLum         	- 	0..1 	0.8 The maximum luminance for full opacity (fade from 100% to 0% between this and upperLum)
// @Param:      upperLum       	- 	0..1 	0.9 The upper luminance threshold (above this becomes transparent)
// @Returns:    result         	- 	-   	-   The image with only pixels in all specified ranges visible (16-bit)
func colorRotateHSL(img *image.NRGBA64,
	rotate float64,
	lowerHue, minHue, maxHue, upperHue float64,
	lowerSat, minSat, maxSat, upperSat float64,
	lowerLum, minLum, maxLum, upperLum float64,
) (*image.NRGBA64, error) {
	hsl, err := colorSelectHSL(
		img,
		lowerHue, minHue, maxHue, upperHue,
		lowerSat, minSat, maxSat, upperSat,
		lowerLum, minLum, maxLum, upperLum,
	)
	if err != nil {
		return nil, err
	}
	hsl, err = colorHueRotate(hsl, rotate)
	if err != nil {
		return nil, err
	}
	return blendNormal(img, hsl)

}

// @Name: auto-levels
// @Desc: Automatically adjusts the contrast and brightness of an image by stretching the histogram to use the full range of values, ignoring outliers using percentiles
// @Param:      img            - -  -     	The image to auto-level
// @Param:      lowPercentile  % - 	0.05   	The lower percentile to ignore (e.g., 0.5)
// @Param:      highPercentile % - 	0.995 	The upper percentile to ignore (e.g., 99.5)
// @Param:      adjustAlpha    		false 	Whether to adjust alpha channel (false = preserve original alpha)
// @Returns:    result         - -  -     	The auto-leveled image
func colorAutoLevels(img *image.NRGBA64, lowPercentile, highPercentile float64, adjustAlpha bool) (*image.NRGBA64, error) {
	bounds := img.Bounds()
	pixelCount := (bounds.Max.X - bounds.Min.X) * (bounds.Max.Y - bounds.Min.Y)

	// Build histograms for each channel (16-bit, so 65536 bins)
	histR := make([]int, 65536)
	histG := make([]int, 65536)
	histB := make([]int, 65536)
	histA := make([]int, 65536)

	// Count non-transparent pixels for RGB histograms
	nonTransparentCount := 0

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			c := img.NRGBA64At(x, y)
			// Always build alpha histogram
			histA[c.A]++

			// Only include non-transparent pixels in RGB histograms
			if c.A > 0 {
				histR[c.R]++
				histG[c.G]++
				histB[c.B]++
				nonTransparentCount++
			}
		}
	}

	findPercentile := func(hist []int, percentile float64, total int) uint16 {
		target := int(float64(total) * percentile)
		sum := 0
		for i, v := range hist {
			sum += v
			if sum >= target {
				return uint16(i)
			}
		}
		return 65535
	}

	// Use nonTransparentCount for RGB percentiles
	minR := findPercentile(histR, lowPercentile, nonTransparentCount)
	maxR := findPercentile(histR, highPercentile, nonTransparentCount)
	minG := findPercentile(histG, lowPercentile, nonTransparentCount)
	maxG := findPercentile(histG, highPercentile, nonTransparentCount)
	minB := findPercentile(histB, lowPercentile, nonTransparentCount)
	maxB := findPercentile(histB, highPercentile, nonTransparentCount)

	// Calculate alpha percentiles if adjusting alpha
	var minA, maxA uint16
	if adjustAlpha {
		minA = findPercentile(histA, lowPercentile, pixelCount)
		maxA = findPercentile(histA, highPercentile, pixelCount)
	}

	// Avoid division by zero
	if maxR == minR {
		maxR = minR + 1
	}
	if maxG == minG {
		maxG = minG + 1
	}
	if maxB == minB {
		maxB = minB + 1
	}
	if adjustAlpha && maxA == minA {
		maxA = minA + 1
	}

	scaleR := float64(0xFFFF) / float64(maxR-minR)
	scaleG := float64(0xFFFF) / float64(maxG-minG)
	scaleB := float64(0xFFFF) / float64(maxB-minB)
	var scaleA float64
	if adjustAlpha {
		scaleA = float64(0xFFFF) / float64(maxA-minA)
	}

	return dsl.parallelProcessNRGBA64(img, func(r1, g1, b1, a1 uint32) (r, g, b, a uint32) {
		// For transparent pixels, keep original values
		if a1 == 0 {
			return r1, g1, b1, a1
		}

		// Clamp to percentile range, then scale
		r = uint32(math.Clamp((float64(clampU16(r1, minR, maxR))-float64(minR))*scaleR, 0, 0xFFFF))
		g = uint32(math.Clamp((float64(clampU16(g1, minG, maxG))-float64(minG))*scaleG, 0, 0xFFFF))
		b = uint32(math.Clamp((float64(clampU16(b1, minB, maxB))-float64(minB))*scaleB, 0, 0xFFFF))

		// Handle alpha based on adjustAlpha parameter
		if adjustAlpha {
			a = uint32(math.Clamp((float64(clampU16(a1, minA, maxA))-float64(minA))*scaleA, 0, 0xFFFF))
		} else {
			a = a1 // Keep original alpha
		}
		return
	}, NumColorConversionWorkers), nil
}

// @Name: auto-white-balance
// @Desc: Automatically adjusts the white balance of an image by finding bright areas and making them neutral
// @Param:      img            - -  -     	The image to auto-white-balance
// @Param:      threshold      % - 	0.95   	The brightness threshold to consider as white (0-1)
// @Param:      strength       - 0..1 	1.0   	How strongly to apply the white balance (0 = no change, 1 = full correction)
// @Returns:    result         - -  -     	The white-balanced image
func colorAutoWhiteBalance(img *image.NRGBA64, threshold, strength float64) (*image.NRGBA64, error) {
	bounds := img.Bounds()

	// Calculate brightness threshold in 16-bit range
	brightnessThreshold := uint32(threshold * 0xFFFF)

	// Variables to accumulate color values of bright pixels
	var sumR, sumG, sumB uint64
	var brightPixelCount uint64

	// First pass: find bright pixels and calculate their average color
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			c := img.NRGBA64At(x, y)

			// Skip transparent pixels
			if c.A == 0 {
				continue
			}

			// Calculate pixel brightness using luminosity method
			brightness := uint32(float64(c.R)*0.21 + float64(c.G)*0.72 + float64(c.B)*0.07)

			// If pixel is bright enough, include it in the average
			if brightness >= brightnessThreshold {
				sumR += uint64(c.R)
				sumG += uint64(c.G)
				sumB += uint64(c.B)
				brightPixelCount++
			}
		}
	}

	// If no bright pixels found, return original image
	if brightPixelCount == 0 {
		return img, nil
	}

	// Calculate average color of bright pixels
	avgR := float64(sumR) / float64(brightPixelCount)
	avgG := float64(sumG) / float64(brightPixelCount)
	avgB := float64(sumB) / float64(brightPixelCount)

	// Calculate maximum value to normalize against
	maxAvg := math.Max(avgR, math.Max(avgG, avgB))

	// Calculate scaling factors to make bright areas neutral
	scaleR := maxAvg / avgR
	scaleG := maxAvg / avgG
	scaleB := maxAvg / avgB

	// Apply white balance correction
	return dsl.parallelProcessNRGBA64(img, func(r1, g1, b1, a1 uint32) (r, g, b, a uint32) {
		// For transparent pixels, keep original values
		if a1 == 0 {
			return r1, g1, b1, a1
		}

		// Apply white balance correction with strength factor
		r = uint32(math.Clamp(float64(r1)*((scaleR-1.0)*strength+1.0), 0, 0xFFFF))
		g = uint32(math.Clamp(float64(g1)*((scaleG-1.0)*strength+1.0), 0, 0xFFFF))
		b = uint32(math.Clamp(float64(b1)*((scaleB-1.0)*strength+1.0), 0, 0xFFFF))
		a = a1 // Keep original alpha
		return
	}, NumColorConversionWorkers), nil
}

// @Name: auto-contrast
// @Desc: Automatically adjusts the contrast of an image by stretching the histogram to use the full range of values
// @Param:      img            - -  -     	The image to auto-contrast
// @Param:      threshold      % - 	0.01   	The percentage of pixels to ignore at both ends of the histogram (0-0.5)
// @Param:      strength       - 0..1 	1.0   	How strongly to apply the contrast adjustment (0 = no change, 1 = full correction)
// @Returns:    result         - -  -     	The contrast-adjusted image
func colorAutoContrast(img *image.NRGBA64, threshold, strength float64) (*image.NRGBA64, error) {
	bounds := img.Bounds()

	// Build histogram for brightness values (16-bit, so 65536 bins)
	hist := make([]int, 65536)

	// Count non-transparent pixels for histogram
	nonTransparentCount := 0

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			c := img.NRGBA64At(x, y)
			// Skip transparent pixels
			if c.A == 0 {
				continue
			}

			// Calculate pixel brightness using luminosity method
			brightness := uint32(float64(c.R)*0.21 + float64(c.G)*0.72 + float64(c.B)*0.07)
			hist[brightness]++
			nonTransparentCount++
		}
	}

	// Find minimum and maximum brightness values that exclude threshold% of pixels
	findPercentile := func(hist []int, percentile float64, total int) uint32 {
		target := int(float64(total) * percentile)
		sum := 0
		for i, v := range hist {
			sum += v
			if sum >= target {
				return uint32(i)
			}
		}
		return 65535
	}

	minBrightness := findPercentile(hist, threshold, nonTransparentCount)
	maxBrightness := findPercentile(hist, 1.0-threshold, nonTransparentCount)

	// Avoid division by zero
	if maxBrightness == minBrightness {
		maxBrightness = minBrightness + 1
	}

	// Calculate scaling factor
	scale := float64(0xFFFF) / float64(maxBrightness-minBrightness)

	return dsl.parallelProcessNRGBA64(img, func(r1, g1, b1, a1 uint32) (r, g, b, a uint32) {
		// For transparent pixels, keep original values
		if a1 == 0 {
			return r1, g1, b1, a1
		}

		// Calculate pixel brightness
		brightness := float64(r1)*0.21 + float64(g1)*0.72 + float64(b1)*0.07

		// Clamp brightness to the found range
		brightness = math.Clamp(brightness, float64(minBrightness), float64(maxBrightness))

		// Calculate contrast adjustment factor
		contrastFactor := ((brightness-float64(minBrightness))*scale)/brightness - 1.0

		// Apply contrast adjustment with strength factor
		r = uint32(math.Clamp(float64(r1)*(contrastFactor*strength+1.0), 0, 0xFFFF))
		g = uint32(math.Clamp(float64(g1)*(contrastFactor*strength+1.0), 0, 0xFFFF))
		b = uint32(math.Clamp(float64(b1)*(contrastFactor*strength+1.0), 0, 0xFFFF))
		a = a1 // Keep original alpha
		return
	}, NumColorConversionWorkers), nil
}

// @Name: auto-tone
// @Desc: Automatically enhances the image by applying auto-levels, auto-white-balance, and auto-contrast in sequence
// @Param:      img            		- -  	-     	The image to auto-tone
// @Param:      levelsLow      		% - 	0.005  	Lower percentile for auto-levels (e.g., 0.5)
// @Param:      levelsHigh     		% - 	0.9995 	Upper percentile for auto-levels (e.g., 99.5)
// @Param:      whiteThresh    		% - 	0.99   	Brightness threshold for auto-white-balance (0-1)
// @Param:      whiteStrength  		- 0..1 	0.8  	Strength for auto-white-balance (0-1)
// @Param:      contrastThresh 		% - 	0.01   	Percentile for auto-contrast (0-0.5)
// @Param:      contrastStrength 	- 0..1 	0.8   	Strength for auto-contrast (0-1)
// @Returns:    result         		- -  	-     	The auto-toned image
func colorAutoTone(
	img *image.NRGBA64,
	levelsLow, levelsHigh float64,
	whiteThresh, whiteStrength float64,
	contrastThresh, contrastStrength float64,
) (*image.NRGBA64, error) {
	// Step 1: Auto-Levels
	autoLeveled, err := colorAutoLevels(img, levelsLow, levelsHigh, false)
	if err != nil {
		return nil, err
	}
	// Step 2: Auto-White-Balance
	autoWhite, err := colorAutoWhiteBalance(autoLeveled, whiteThresh, whiteStrength)
	if err != nil {
		return nil, err
	}
	// Step 3: Auto-Contrast
	autoContrasted, err := colorAutoContrast(autoWhite, contrastThresh, contrastStrength)
	if err != nil {
		return nil, err
	}
	return autoContrasted, nil
}

// @Name: select-brightness
// @Desc: Selects pixels based on their brightness, making pixels outside the specified range transparent
// @Param:      img            - -   		-   The image to process (16-bit)
// @Param:      lowerBright    - 0..1 	0.1  The lower brightness threshold (below this becomes transparent)
// @Param:      minBright      - 0..1 	0.2  The minimum brightness for full opacity (fade from 0% to 100% between lowerBright and this)
// @Param:      maxBright      - 0..1 	0.8  The maximum brightness for full opacity (fade from 100% to 0% between this and upperBright)
// @Param:      upperBright    - 0..1 	0.9  The upper brightness threshold (above this becomes transparent)
// @Returns:    result         - -   		-   The image with only pixels in the specified brightness range visible (16-bit)
func colorSelectBrightness(img *image.NRGBA64,
	lowerBright, minBright, maxBright, upperBright float64,
) (*image.NRGBA64, error) {
	bounds := img.Bounds()
	result := IFromBounds(bounds)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, a := dsl.getColor(img, x, y)
			// Normalize to 0-1
			rf := float64(r) / 65535.0
			gf := float64(g) / 65535.0
			bf := float64(b) / 65535.0
			af := float64(a) / 65535.0

			// Calculate brightness using luminosity method
			brightness := rf*0.21 + gf*0.72 + bf*0.07

			// Calculate alpha based on brightness
			var alpha float64
			if brightness < lowerBright || brightness > upperBright {
				alpha = 0.0
			} else if brightness < minBright {
				alpha = (brightness - lowerBright) / (minBright - lowerBright)
			} else if brightness > maxBright {
				alpha = 1.0 - (brightness-maxBright)/(upperBright-maxBright)
			} else {
				alpha = 1.0
			}

			// Clamp alpha to [0,1]
			alpha = math.Clamp(alpha, 0.0, 1.0)

			// Apply the alpha to the original color, preserve original alpha channel
			finalAlpha := alpha * af
			dsl.setColor(img, x, y, uint32(rf*65535.0), uint32(gf*65535.0), uint32(bf*65535.0), uint32(finalAlpha*65535.0))
		}
	}

	return result, nil
}

// @Name: remove-brightness
// @Desc: Removes pixels based on their brightness, making pixels inside the specified range transparent
// @Param:      img            - -   		-   The image to process (16-bit)
// @Param:      lowerBright    - 0..1 	0.1  The lower brightness threshold (below this becomes transparent)
// @Param:      minBright      - 0..1 	0.2  The minimum brightness for full opacity (fade from 0% to 100% between lowerBright and this)
// @Param:      maxBright      - 0..1 	0.8  The maximum brightness for full opacity (fade from 100% to 0% between this and upperBright)
// @Param:      upperBright    - 0..1 	0.9  The upper brightness threshold (above this becomes transparent)
// @Returns:    result         - -   		-   The image with pixels in the specified brightness range removed (16-bit)
func colorRemoveBrightness(img *image.NRGBA64,
	lowerBright, minBright, maxBright, upperBright float64,
) (*image.NRGBA64, error) {
	bright, err := colorSelectBrightness(
		img,
		lowerBright, minBright, maxBright, upperBright,
	)
	if err != nil {
		return nil, err
	}
	return blendErase(img, bright)
}
