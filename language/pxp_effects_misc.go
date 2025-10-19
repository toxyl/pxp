package language

import "image"

// @Name: enhance
// @Desc: Enhances colors and sharpness of an image
// @Param:      img     	- -   		-   	The image to enhance
// @Param:      brightness	- -1.0..1.0	0.0   	The brightness adjustment of the image
// @Param:      contrast	- -1.0..1.0	0.0   	The contrast adjustment of the image
// @Param:      sharpening	- 0.0..5.0  1.0 	The sharpening intensity in pixels (higher values detect larger edges)
// @Param:      rMin 		- 0.0..1.0  0.75   	The minimum intensity of the red channel
// @Param:      rMax 		- 0.0..1.0  1.00   	The maximum intensity of the red channel
// @Param:      gMin 		- 0.0..1.0  0.75   	The minimum intensity of the green channel
// @Param:      gMax 		- 0.0..1.0  1.00   	The maximum intensity of the green channel
// @Param:      bMin 		- 0.0..1.0  0.75   	The minimum intensity of the blue channel
// @Param:      bMax 		- 0.0..1.0  1.00   	The maximum intensity of the blue channel
// @Param:      rWeight		- 0.0..1.0  0.299   The weight of the red channel (used for sharpening)
// @Param:      gWeight		- 0.0..1.0  0.587   The weight of the green channel (used for sharpening)
// @Param:      bWeight		- 0.0..1.0  0.114   The weight of the blue channel (used for sharpening)
// @Returns:    result  	- -   		-   	The enhanceed image
func enhance(img *image.NRGBA64, brightness float64, contrast float64, sharpening float64, rMin float64, rMax float64, gMin float64, gMax float64, bMin float64, bMax float64, rWeight float64, gWeight float64, bWeight float64) (*image.NRGBA64, error) {
	var err error
	// prepare the base image
	img, err = colorBrightness(img, brightness+1)
	if err != nil {
		return nil, err
	}
	img, err = colorContrast(img, contrast+1)
	if err != nil {
		return nil, err
	}

	// prepare red channel
	r, err := colorBalance(img, rMax, gMin, bMin)
	if err != nil {
		return nil, err
	}
	r1, err := colorOpacity(r, 1)
	if err != nil {
		return nil, err
	}
	r2, err := blendMultiply(r, r1)
	if err != nil {
		return nil, err
	}

	// prepare green channel
	g, err := colorBalance(img, rMin, gMax, bMin)
	if err != nil {
		return nil, err
	}
	g1, err := colorOpacity(g, 1)
	if err != nil {
		return nil, err
	}
	g2, err := blendMultiply(g, g1)
	if err != nil {
		return nil, err
	}

	// prepare blue channel
	b, err := colorBalance(img, rMin, gMin, bMax)
	if err != nil {
		return nil, err
	}
	b1, err := colorOpacity(b, 1)
	if err != nil {
		return nil, err
	}
	b2, err := blendMultiply(b, b1)
	if err != nil {
		return nil, err
	}

	// merge channels
	gb, err := blendScreen(g2, b2)
	if err != nil {
		return nil, err
	}
	rgb, err := blendScreen(r2, gb)
	if err != nil {
		return nil, err
	}

	// sharpen and return result
	return sharpen(rgb, 1, sharpening, rWeight, gWeight, bWeight)
}
