package language

import (
	"image"
	"image/color"

	"github.com/toxyl/math"
)

// //////// TEXT ///////////////////////////////////////////////////
func T(style TextStyle, text string) *Text {
	return &Text{
		Text:  text,
		Style: &style,
	}
}

// @Name: T
// @Desc: Creates a new text.
// @Param:      style    - - -   The text style to use
// @Param:      text     - - -   The text to print
// @Returns:    result   - - -	 A new text
func makeText(style TextStyle, text string) (Text, error) {
	return *T(style, text), nil
}

// //////// POINT ///////////////////////////////////////////////////
func P(x, y float64) *Point {
	return &Point{
		X: x,
		Y: y,
	}
}

// @Name: P
// @Desc: Creates a new point at P(x|y).
// @Param:      x        - - 0   The start position on the x-axis
// @Param:      y        - - 0   The start position on the y-axis
// @Returns:    result   - - -	 A new point
func makePoint(x, y float64) (Point, error) {
	return *P(x, y), nil
}

// @Name: Px
// @Desc: Returns the x-coordinate of a point.
// @Param:      p        - - -   The point to return the x-coordinate of
// @Returns:    result    - - -	 The x-coordinate of p
func pointX(p Point) (any, error) {
	return p.X, nil
}

// @Name: Py
// @Desc: Returns the y-coordinate of a point.
// @Param:      p        - - -   The point to return the y-coordinate of
// @Returns:    result    - - -	 The y-coordinate of p
func pointY(p Point) (any, error) {
	return p.Y, nil
}

// //////// VECTOR ///////////////////////////////////////////////////
func V(x, y, z float64) Vector {
	return Vector{
		X: x,
		Y: y,
	}
}

// @Name: V
// @Desc: Creates a new Vector from x, y and z.
// @Param:      x        - - 0   The x-component
// @Param:      y        - - 0   The y-component
// @Param:      z        - - 0   The z-component
// @Returns:    result   - - -	 A new vector
func makeVector(x, y, z float64) (Vector, error) {
	return V(x, y, z), nil
}

// @Name: Vx
// @Desc: Returns the x-component of a vector.
// @Param:      v        - - -   The vector to return the x-component of
// @Returns:    result    - - -	 The x-component of v
func vectorX(v Vector) (any, error) {
	return v.X, nil
}

// @Name: Vy
// @Desc: Returns the y-component of a vector.
// @Param:      v        - - -   The vector to return the y-component of
// @Returns:    result    - - -	 The y-component of vp
func vectorY(v Vector) (any, error) {
	return v.Y, nil
}

// @Name: Vz
// @Desc: Returns the z-component of a vector.
// @Param:      v        - - -   The vector to return the z-component of
// @Returns:    result    - - -	 The z-component of vp
func vectorZ(v Vector) (any, error) {
	return v.Z, nil
}

// //////// RECT ///////////////////////////////////////////////////
func R(x, y, w, h float64) Rect {
	return Rect{
		P1: P(x, y),
		P2: P(x+w, y+h),
	}
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

// @Name: Rx
// @Desc: Returns the x-coordinate of a rect.
// @Param:      r        - - -   The rect to return the x-coordinate of
// @Returns:    result    - - -	 The x-coordinate of r
func rectX(r Rect) (any, error) {
	return r.X1(), nil
}

// @Name: Ry
// @Desc: Returns the y-coordinate of a rect.
// @Param:      r        - - -   The rect to return the y-coordinate of
// @Returns:    result    - - -	 The y-coordinate of r
func rectY(r Rect) (any, error) {
	return r.Y1(), nil
}

// @Name: Rw
// @Desc: Returns the width of a rect.
// @Param:      r        - - -   The rect to return the width of
// @Returns:    result    - - -	 The width of r
func rectW(r Rect) (any, error) {
	return r.W(), nil
}

// @Name: Rh
// @Desc: Returns the height of a rect.
// @Param:      r        - - -   The rect to return the height of
// @Returns:    result    - - -	 The height of r
func rectH(r Rect) (any, error) {
	return r.H(), nil
}

// //////// ELLIPSE ///////////////////////////////////////////////////
func E(centerX, centerY, radiusX, radiusY float64) Ellipse {
	return Ellipse{
		Center: &Point{
			X: centerX,
			Y: centerY,
		},
		Radius: &Point{
			X: centerX,
			Y: centerY,
		},
	}
}

// @Name: E
// @Desc: Creates a new ellipse with the given radius at P(x|y).
// @Param:      centerX   - - -  The center of the ellipse on the x-axis
// @Param:      centerY   - - -  The center of the ellipse on the y-axis
// @Param:      radiusX   - - -  The radius of the ellipse on the x-axis
// @Param:      radiusY   - - -  The radius of the ellipse on the y-axis
// @Returns:    result   - - -	 A new ellipse
func makeEllipse(centerX, centerY, radiusX, radiusY float64) (Ellipse, error) {
	return E(centerX, centerY, radiusX, radiusY), nil
}

// @Name: Ex
// @Desc: Returns the center x-coordinate of an ellipse.
// @Param:      e        - - -   The ellipse to return the center x-coordinate of
// @Returns:    result    - - -	 the center x-coordinate of e
func ellipseX(e Ellipse) (any, error) {
	return e.Center.X, nil
}

// @Name: Ey
// @Desc: Returns the center y-coordinate of an ellipse.
// @Param:      e        - - -   The ellipse to return the center y-coordinate of
// @Returns:    result    - - -	 the center y-coordinate of e
func ellipseY(e Ellipse) (any, error) {
	return e.Center.Y, nil
}

// @Name: Erx
// @Desc: Returns the x-component of the radius of an ellipse.
// @Param:      e        - - -   The ellipse to return the x-component of the radius of
// @Returns:    result    - - -	 the x-component of the radius of e
func ellipseRadiusX(e Ellipse) (any, error) {
	return e.Radius.X, nil
}

// @Name: Ery
// @Desc: Returns the y-component of the radius of an ellipse.
// @Param:      e        - - -   The ellipse to return the y-component of the radius of
// @Returns:    result    - - -	 the y-component of the radius of e
func ellipseRadiusY(e Ellipse) (any, error) {
	return e.Radius.Y, nil
}

// @Name: C
// @Desc: Creates a new circle with the given radius at P(x|y).
// @Param:      centerX   - - -  The center of the circle on the x-axis
// @Param:      centerY   - - -  The center of the circle on the y-axis
// @Param:      radius    - - -  The radius of the circle
// @Returns:    result   - - -	 A new circle
func makeCircle(centerX, centerY, radius float64) (Ellipse, error) {
	return E(centerX, centerY, radius, radius), nil
}

// //////// IMAGES ///////////////////////////////////////////////////

// @Name: Iw
// @Desc: Returns the width of an image.
// @Param:      img       - - -  The image to return the width of
// @Returns:    result    - - -	 The width of img
func imageW(img *image.NRGBA64) (any, error) {
	return img.Bounds().Max.X, nil
}

// @Name: Ih
// @Desc: Returns the height of an image.
// @Param:      img       - - -  The image to return the height of
// @Returns:    result    - - -	 The height of img
func imageH(img *image.NRGBA64) (any, error) {
	return img.Bounds().Max.Y, nil
}

// @Name: Ir
// @Desc: Returns the aspect ratio of the given image
// @Param:      img     - -   - The image to return the aspect ratio of
// @Returns:    result  - -   - Aspect ratio of the image
func imageAspectRatio(img *image.NRGBA64) (float64, error) {
	w, h := float64(img.Bounds().Max.X), float64(img.Bounds().Max.Y)
	return math.Max(w, h) / math.Min(w, h), nil
}

// //////// IMAGE: SOLID ///////////////////////////////////////////////////
func IC(w, h int, cFill color.RGBA64) *image.NRGBA64 {
	img := image.NewNRGBA64(image.Rect(0, 0, w, h))
	if cFill.A == 0 {
		return img // no need to fill the image
	}
	r2, g2, b2, a2 := uint32(cFill.R), uint32(cFill.G), uint32(cFill.B), uint32(cFill.A)
	return dsl.parallelProcessNRGBA64(img, func(r1, g1, b1, a1 uint32) (r, g, b, a uint32) {
		return r2, g2, b2, a2
	}, NumColorConversionWorkers)
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

// //////// IMAGE: TRANSPARENT ///////////////////////////////////////////////////
func I(w, h int) *image.NRGBA64 {
	return IC(w, h, color.RGBA64{0, 0, 0, 0})
}

// @Name: I
// @Desc: Creates a new transparent image.
// @Param:      w       - - -   The width of the image
// @Param:      h       - - -   The height of the image
// @Returns:    result  - - -	The new image
func makeImageTransparent(w, h int) (*image.NRGBA64, error) {
	return I(w, h), nil
}

// //////// SUB-IMAGE ///////////////////////////////////////////////////
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

// //////// LINE STYLE ///////////////////////////////////////////////////
func LS(color color.RGBA64, thickness float64) LineStyle {
	return LineStyle{
		Thickness: thickness,
		Color:     &color,
	}
}

// @Name: LS
// @Desc: Creates a new line style.
// @Param:      color     - - -  The line color
// @Param:      thickness 1 - 1  The line thickness
// @Returns:    result    - - -	 A new line style
func makeLineStyle(color color.RGBA64, thickness float64) (LineStyle, error) {
	return LS(color, thickness), nil
}

// //////// FILL STYLE ///////////////////////////////////////////////////
func FS(color color.RGBA64) FillStyle {
	return FillStyle{
		Color: &color,
	}
}

// @Name: FS
// @Desc: Creates a new fill style.
// @Param:      color     - - -  The fill color
// @Returns:    result    - - -	 A new fill style
func makeFillStyle(color color.RGBA64) (FillStyle, error) {
	return FS(color), nil
}

// //////// TEXT STYLE ///////////////////////////////////////////////////
func TS(color color.RGBA64, size float64, family string) TextStyle {
	return TextStyle{
		Family: family,
		Size:   size,
		Color:  &color,
	}
}

// @Name: TS
// @Desc: Creates a new font style.
// @Param:      color     - - -  		The font color
// @Param:      size      1 - 10  		The font size
// @Param:      family    - - "mono"  	The font family
// @Returns:    result    - - -	 		A new font style
func makeTextStyle(color color.RGBA64, size float64, family string) (TextStyle, error) {
	return TS(color, size, family), nil
}

// //////// MISC (not exposed) ///////////////////////////////////////////////////
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
