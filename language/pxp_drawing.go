package language

import (
	"image"
	"image/color"

	"github.com/toxyl/pxp/fonts"
)

// @Name: fill
// @Desc: Fills the given image.
// @Param:      img        - - -   The image to fill
// @Param:      style      - - -   The color of the fill
// @Returns:    result     - - -   The filled image
func fill(img *image.NRGBA64, style FillStyle) (*image.NRGBA64, error) {
	r2, g2, b2, a2 := uint32(style.Color.R), uint32(style.Color.G), uint32(style.Color.B), uint32(style.Color.A)
	return dsl.parallelProcessNRGBA64(img, func(r1, g1, b1, a1 uint32) (r, g, b, a uint32) {
		return r2, g2, b2, a2
	}, NumColorConversionWorkers), nil
}

// @Name: border
// @Desc: Draws a border around the image.
// @Param:      img       - - -   The image to draw border around
// @Param:      style     - - -   The thickness and color of the border
// @Returns:    result    - - -	  The resulting image
func border(img *image.NRGBA64, style LineStyle) (*image.NRGBA64, error) {
	bounds := img.Bounds()
	return drawRectPx(img, R(0, 0, float64(bounds.Max.X), float64(bounds.Max.Y)), style)
}

// @Name: box
// @Desc: Fills the image with the given background color and then draws a border around it.
// @Param:      img           - - -   The image to draw border around
// @Param:      styleBorder   - - -   The thickness and color of the border
// @Param:      styleFill     - - -   The color of the fill
// @Returns:    result        - - -	  The resulting image
func box(img *image.NRGBA64, styleBorder LineStyle, styleFill FillStyle) (*image.NRGBA64, error) {
	img, _ = fill(img, styleFill)
	img, _ = border(img, styleBorder)
	return img, nil
}

// @Name: grid
// @Desc: Draws a grid from P(x1|y1) to P(x2|y2) with the given thickness and color.
// @Param:      img       - - -   The image to draw to
// @Param:      rows      - - -   The number of rows
// @Param:      cols      - - -   The number of cols
// @Param:      style     - - -   The thickness and color of the border
// @Returns:    result    - - -	  The resulting image
func grid(img *image.NRGBA64, rows, cols int, style LineStyle) (*image.NRGBA64, error) {
	bounds := img.Bounds()
	w, h := bounds.Max.X, bounds.Max.Y
	return drawGridPx(img, R(0, 0, float64(w), float64(h)), rows, cols, style)
}

// @Name: grid-v
// @Desc: Draws vertical grid lines on the image with the given thickness and color.
// @Param:      img       - - -   The image to draw to
// @Param:      cols      - - -   The number of columns
// @Param:      style     - - -   The thickness and color of the border
// @Returns:    result    - - -	  The resulting image
func gridV(img *image.NRGBA64, cols int, style LineStyle) (*image.NRGBA64, error) {
	bounds := img.Bounds()
	w, h := bounds.Max.X, bounds.Max.Y
	return drawGridVPx(img, R(0, 0, float64(w), float64(h)), cols, style)
}

// @Name: grid-h
// @Desc: Draws horizontal grid lines on the image with the given thickness and color.
// @Param:      img       - - -   The image to draw to
// @Param:      rows      - - -   The number of rows
// @Param:      style     - - -   The thickness and color of the border
// @Returns:    result    - - -	  The resulting image
func gridH(img *image.NRGBA64, rows int, style LineStyle) (*image.NRGBA64, error) {
	bounds := img.Bounds()
	w, h := bounds.Max.X, bounds.Max.Y
	return drawGridHPx(img, R(0, 0, float64(w), float64(h)), rows, style)
}

// @Name: group
// @Desc: Generates the given group with the given styles.
// @Param:      img       			- - -   	The image to wrap in a group
// @Param:      title     			- - -   	The title of the group
// @Param:      colTitle   			- - -   	The color of the title
// @Param:      colHeader  			- - -   	The color of the header
// @Param:      colBody    			- - -   	The color of the body
// @Param:      colBorder  			- - -   	The color of the border
// @Param:      padding 			0 - 3   	The padding for the image to wrap
// @Returns:    result    			- - -	  	The group wrapping the input image
func group(img *image.NRGBA64, title string, colTitle, colHeader, colBody, colBorder color.RGBA64, padding int) (*image.NRGBA64, error) {
	// Calculate dimensions
	bounds := img.Bounds()
	imgWidth := bounds.Max.X
	imgHeight := bounds.Max.Y

	// Add padding to all sides
	paddingPx := int(padding)
	totalWidth := imgWidth + 2*paddingPx
	totalHeight := int(float64(imgHeight + 2*paddingPx))
	headerHeight := fonts.PixelOperator.GlyphHeight + 4

	// Generate header text
	headerText, err := text(title, colTitle, colBorder)
	if err != nil {
		return nil, err
	}

	// Create header canvas
	headerCanvas := IC(totalWidth, headerHeight, color.RGBA64{0, 0, 0, 0})
	headerCanvas, _ = fill(headerCanvas, FS(colHeader))
	headerCanvas, _ = border(headerCanvas, LS(colBorder, 1))

	// Composite text onto header
	// Position text in header (centered horizontally, vertically aligned)
	textBounds := headerText.Bounds()
	textX := (totalWidth - textBounds.Max.X) / 2
	textY := (headerHeight - textBounds.Max.Y) / 2

	positionedText, _ := translateImage(headerText, *P(float64(textX), float64(textY)))
	headerCanvas, _ = blend(headerCanvas, positionedText, "normal")

	// Create body canvas
	bodyCanvas := IC(totalWidth, totalHeight, color.RGBA64{0, 0, 0, 0})
	bodyCanvas, _ = fill(bodyCanvas, FS(colBody))
	bodyCanvas, _ = border(bodyCanvas, LS(colBorder, 1))

	// Add padded input image to body
	positionedImg, _ := translateImage(img, *P(float64(paddingPx), float64(paddingPx)))
	bodyCanvas, _ = blend(bodyCanvas, positionedImg, "normal")

	// Position body below header
	positionedBody, _ := translateImage(bodyCanvas, *P(0, float64(headerHeight)-1))

	// Composite final result
	return blend(positionedBody, headerCanvas, "normal")
}
