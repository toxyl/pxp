package language

import (
	"image"

	"github.com/toxyl/math"
)

// @Name: sharpen
// @Desc: Sharpens an image using a highpass combined with vivid light blending
// @Param:      img     	- -   		-   	The image to sharpen
// @Param:      intensity  	- 0.0..1.0  1   	The intensity of the sharpening effect
// @Param:      radius  	- 0.1..2.0  1   	The radius of the filter in pixels (higher values detect larger edges)
// @Param:      rWeight 	- 0.0..1.0  0.299   The weight of the red channel
// @Param:      gWeight 	- 0.0..1.0  0.587   The weight of the green channel
// @Param:      bWeight 	- 0.0..1.0  0.114   The weight of the blue channel
// @Returns:    result  	- -   		-   	The sharpened image
func sharpen(img *image.NRGBA64, intensity float64, radius float64, rWeight float64, gWeight float64, bWeight float64) (*image.NRGBA64, error) {
	if radius == 0 {
		return img, nil
	}

	hp, err := sharpenHighpass(img, radius, rWeight, gWeight, bWeight)
	if err != nil {
		return nil, err
	}
	hp, err = colorOpacity(hp, intensity)
	if err != nil {
		return nil, err
	}
	res, err := blendVividLight(img, hp)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// @Name: highpass
// @Desc: Creates a high pass filter effect, resulting in a gray image with embossed edges
// @Param:      img     - -   	    -   	The image to apply the high-pass filter to
// @Param:      radius  - 0.1..2.0  1   	The radius of the filter in pixels (higher values detect larger edges)
// @Param:      rWeight - 0.0..1.0  0.299   The weight of the red channel
// @Param:      gWeight - 0.0..1.0  0.587   The weight of the green channel
// @Param:      bWeight - 0.0..1.0  0.114   The weight of the blue channel
// @Returns:    result  - -   	    -   	The filtered image
func sharpenHighpass(img *image.NRGBA64, radius float64, rWeight float64, gWeight float64, bWeight float64) (*image.NRGBA64, error) {
	bounds := img.Bounds()
	width := bounds.Dx()
	height := bounds.Dy()
	result := IFromBounds(bounds)

	// Create 1D Gaussian kernel (separable convolution)
	sigma, kernelSize, halfSize := calcKernelSize(radius)
	kernel := makeKernelGaussian1D(kernelSize, sigma, halfSize)

	// Convert image to grayscale and store in tempBuf
	tempBuf := parallelPixelBuffer(img, func(r, g, b, a uint32) (res float64) {
		return float64(r)*rWeight + float64(g)*gWeight + float64(b)*bWeight
	})

	// Create temporary buffer for separable convolution and the result
	blurBuf := make([]float64, width*height)
	resBuf := make([]uint32, width*height)

	// Horizontal pass
	for y := range height {
		for x := range width {
			var sumH float64
			var totalWeightH float64

			for i := -halfSize; i <= halfSize; i++ {
				srcX := x + i
				if srcX >= 0 && srcX < width {
					weight := kernel[i+halfSize]
					sumH += tempBuf[y*width+srcX] * weight
					totalWeightH += weight
				}
			}

			if totalWeightH > 0 {
				blurBuf[y*width+x] = sumH / totalWeightH
			}
		}
	}

	// Vertical pass
	for y := range height {
		for x := range width {
			var sumV float64
			var totalWeightV float64

			for i := -halfSize; i <= halfSize; i++ {
				srcY := y + i
				if srcY >= 0 && srcY < height {
					weight := kernel[i+halfSize]
					sumV += blurBuf[srcY*width+x] * weight
					totalWeightV += weight
				}
			}

			if totalWeightV > 0 {
				// Calculate high-pass by subtracting blurred from original and add 50% gray to make the result visible
				resBuf[y*width+x] = uint32(math.Clamp(0x7FFF+(tempBuf[y*width+x]-(sumV/totalWeightV)), 0, 0xFFFF))
			}
		}
	}

	// Apply the result
	for y := range height {
		for x := range width {
			gray := resBuf[y*width+x]
			dsl.setColor(result, x, y, gray, gray, gray, dsl.getColorAlphaChannel(img, x, y))
		}
	}
	return result, nil
}

// @Name: clarity
// @Desc: Enhances local contrast while preserving overall image structure
// @Param:      img     	- -   		-   	The image to enhance
// @Param:      intensity  	- 0.0..1.0  1   	The intensity of the clarity effect
// @Param:      radius  	- 0.1..2.0  1   	The radius of the filter in pixels (higher values affect larger areas)
// @Param:      rWeight 	- 0.0..1.0  0.299   The weight of the red channel
// @Param:      gWeight 	- 0.0..1.0  0.587   The weight of the green channel
// @Param:      bWeight 	- 0.0..1.0  0.114   The weight of the blue channel
// @Returns:    result  	- -   		-   	The enhanced image
func clarity(img *image.NRGBA64, intensity float64, radius float64, rWeight float64, gWeight float64, bWeight float64) (*image.NRGBA64, error) {
	if radius == 0 || intensity == 0 {
		return img, nil
	}

	// Create high-pass with medium radius for strong local contrast
	hp, err := sharpenHighpass(img, radius*10.0, rWeight, gWeight, bWeight)
	if err != nil {
		return nil, err
	}

	// Boost the high-pass signal for stronger effect
	hp, err = colorOpacity(hp, intensity*2.0)
	if err != nil {
		return nil, err
	}

	// Use soft light blend for more natural contrast enhancement
	res, err := blendSoftLight(img, hp)
	if err != nil {
		return nil, err
	}

	// Apply a second pass with soft light for extra punch
	res, err = blendSoftLight(res, hp)
	if err != nil {
		return nil, err
	}

	return res, nil
}
