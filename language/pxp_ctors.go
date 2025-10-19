package language

import (
	"image"
	"image/color"
)

// @Name: P
// @Desc: Creates a new point at P(x|y).
// @Param:      x        - - 0   The start position on the x-axis
// @Param:      y        - - 0   The start position on the y-axis
// @Returns:    result   - - -	 A new point
func makePoint(x, y float64) (Point, error) {
	return P(x, y), nil
}

// @Name: R
// @Desc: Creates a new rectangle with the given dimensions at P(x|y).
// @Param:      x        - - -   The upper-left corner of the rectangle
// @Param:      y        - - -   The upper-left corner of the rectangle
// @Param:      w        - - -   The width of the rectangle
// @Param:      h        - - -   The width of the rectangle
// @Returns:    result   - - -	 A new rectangle
func makeRect(x, y, w, h float64) (Rect, error) {
	return R(x, y, w, h), nil
}

// @Name: IC
// @Desc: Creates a new image with the given color.
// @Param:      w       - - -   The width of the image
// @Param:      h       - - -   The height of the image
// @Param:      cFill  	- - -   The fill color
// @Returns:    result  - - -	The new image
func makeImage(w, h int, cFill color.RGBA64) (*image.NRGBA64, error) {
	return IC(w, h, cFill), nil
}

// @Name: I
// @Desc: Creates a new transparent image.
// @Param:      w       - - -   The width of the image
// @Param:      h       - - -   The height of the image
// @Returns:    result  - - -	The new image
func makeImageTransparent(w, h int) (*image.NRGBA64, error) {
	return I(w, h), nil
}

// @Name: SI
// @Desc: Copies an area from a source image and returns it as a new image.
// @Param:      img     - - -   The source image
// @Param:      r       - - -   The selection to copy
// @Returns:    result  - - -	The new image
func extractSubImage(img *image.NRGBA64, r Rect) (*image.NRGBA64, error) {
	bounds := img.Bounds()
	iw, ih := bounds.Dx(), bounds.Dy()
	return cropPx(img, int(r.X1()), iw-int(r.X2()), int(r.Y1()), ih-int(r.Y2()))
}

func P(x, y float64) Point {
	return Point{
		X: x,
		Y: y,
	}
}

func R(x, y, w, h float64) Rect {
	return Rect{
		P1: P(x, y),
		P2: P(x+w, y+h),
	}
}

func IC(w, h int, cFill color.RGBA64) *image.NRGBA64 {
	img := image.NewNRGBA64(image.Rect(0, 0, w, h))
	if cFill.A == 0 {
		return img // no need to fill the image
	}
	r2, g2, b2, a2 := uint32(cFill.R), uint32(cFill.G), uint32(cFill.B), uint32(cFill.A)
	return dsl.parallelProcessNRGBA64(img, func(r1, g1, b1, a1 uint32) (r, g, b, a uint32) {
		return r2, g2, b2, a2
	}, numWorkers)
}

func I(w, h int) *image.NRGBA64 {
	return IC(w, h, color.RGBA64{0, 0, 0, 0})
}

func IFromBounds(bounds image.Rectangle) *image.NRGBA64 {
	return IC(bounds.Dx(), bounds.Dy(), color.RGBA64{0, 0, 0, 0})
}

func IClone(img *image.NRGBA64) *image.NRGBA64 {
	bounds := img.Bounds()
	result := IFromBounds(bounds)

	// Copy original image to result
	for py := bounds.Min.Y; py < bounds.Max.Y; py++ {
		for px := bounds.Min.X; px < bounds.Max.X; px++ {
			result.Set(px, py, img.NRGBA64At(px, py))
		}
	}
	return result
}
