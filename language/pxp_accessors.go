package language

import "image"

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
	return max(w, h) / min(w, h), nil
}
