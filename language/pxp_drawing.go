package language

import (
	"image"
	"image/color"

	"github.com/toxyl/math"
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
	}, numWorkers), nil
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

// @Name: text
// @Desc: Generates the given text with the given styles.
// @Param:      t         - - -   The text to generate
// @Param:      style     - - -   The text style (font, size, color)
// @Param:      outline   - - -   The thickness and color of the outline
// @Returns:    result    - - -	  The resulting image
func text(t string, style TextStyle, outline LineStyle) (*image.NRGBA64, error) {
	// Create a Text object with the provided style
	textObj := Text{
		Text:  t,
		Style: &style,
	}

	// Calculate generous initial canvas size
	lines := splitLines(t)
	maxLineLength := 0
	for _, line := range lines {
		if len(line) > maxLineLength {
			maxLineLength = len(line)
		}
	}
	outline.Thickness++

	// Estimate dimensions with generous padding
	estimatedWidth := maxLineLength*int(style.Size) + int(outline.Thickness)*2 + 20
	estimatedHeight := len(lines)*int(style.Size) + int(outline.Thickness)*2 + 20

	// Create oversized canvas
	canvas := I(estimatedWidth, estimatedHeight)

	// Render text fill
	canvas, err := drawTextPx(canvas, *P(10, 10), textObj)
	if err != nil {
		return nil, err
	}

	// Render text outline
	canvas, err = drawTextOutlinePx(canvas, *P(10, 10), textObj, outline)
	if err != nil {
		return nil, err
	}

	// Detect content bounds
	bounds := canvas.Bounds()
	minX, minY := bounds.Max.X, bounds.Max.Y
	maxX, maxY := bounds.Min.X, bounds.Min.Y

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			pixel := canvas.NRGBA64At(x, y)
			if pixel.A > 0 {
				if x < minX {
					minX = x
				}
				if x > maxX {
					maxX = x
				}
				if y < minY {
					minY = y
				}
				if y > maxY {
					maxY = y
				}
			}
		}
	}

	// Handle edge case where no content is found
	if minX > maxX || minY > maxY {
		return I(1, 1), nil
	}

	// Crop to content bounds
	left := minX
	right := bounds.Max.X - maxX - 1
	top := minY
	bottom := bounds.Max.Y - maxY - 1

	return cropPx(canvas, left, right, top, bottom)
}

// @Name: group
// @Desc: Generates the given group with the given styles.
// @Param:      img       			- - -   	The image to wrap in a group
// @Param:      title     			- - -   	The title of the group
// @Param:      colText   			- - -   	The color of the title
// @Param:      colFill    			- - -   	The color of the group
// @Param:      colBorder  			- - -   	The color of the border
// @Param:      padding 			0 - 3   	The padding for the image to wrap
// @Param:      fillAlphaHeader 	0 1 0.9   	The alpha of header fill
// @Param:      fillAlphaBody   	0 1 0.8   	The alpha of body fill
// @Param:      borderThickness   	0 - 1   	The thickness of the border
// @Param:      borderAlphaHeader 	0 1 0.95   	The alpha of header border
// @Param:      borderAlphaBody   	0 1 0.95   	The alpha of body border
// @Returns:    result    			- - -	  	The group wrapping the input image
func group(img *image.NRGBA64, title string, colText, colFill, colBorder color.RGBA64, padding, fillAlphaHeader, fillAlphaBody, borderThickness, borderAlphaHeader, borderAlphaBody float64) (*image.NRGBA64, error) {
	// Calculate dimensions
	bounds := img.Bounds()
	imgWidth := bounds.Max.X
	imgHeight := bounds.Max.Y

	// Add padding to all sides
	paddingPx := int(padding)
	totalWidth := imgWidth + 2*paddingPx
	totalHeight := int(float64(imgHeight+2*paddingPx) - borderThickness)

	// Estimate header height based on title text size (fallback to 20px)
	headerHeight := 20
	if title != "" {
		// Estimate based on text length and typical font size
		estimatedTextHeight := int(math.Max(12, math.Min(24, float64(len(title))/3)))
		headerHeight = estimatedTextHeight + 8 // Add padding
	}

	// Calculate final dimensions
	finalWidth := totalWidth

	// Create header styles
	headerFillColor := colFill
	headerFillColor.A = uint16(float64(headerFillColor.A) * fillAlphaHeader)
	headerFillStyle := FS(headerFillColor)

	headerBorderColor := colBorder
	headerBorderColor.A = uint16(float64(headerBorderColor.A) * borderAlphaHeader)
	headerBorderStyle := LS(headerBorderColor, borderThickness)

	// Create body styles
	bodyFillColor := colFill
	bodyFillColor.A = uint16(float64(bodyFillColor.A) * fillAlphaBody)
	bodyFillStyle := FS(bodyFillColor)

	bodyBorderColor := colBorder
	bodyBorderColor.A = uint16(float64(bodyBorderColor.A) * borderAlphaBody)
	bodyBorderStyle := LS(bodyBorderColor, borderThickness)

	// Generate header text
	var headerText *image.NRGBA64
	if title != "" {
		textColorVal := colText
		textStyle := TS(textColorVal, 12, "mono")
		outlineStyle := LS(headerBorderColor, 1)

		var err error
		headerText, err = text(title, textStyle, outlineStyle)
		if err != nil {
			return nil, err
		}
	}

	// Create header canvas
	headerCanvas := IC(finalWidth, headerHeight, color.RGBA64{0, 0, 0, 0})
	headerCanvas, _ = fill(headerCanvas, headerFillStyle)
	headerCanvas, _ = border(headerCanvas, headerBorderStyle)

	// Composite text onto header
	if headerText != nil {
		// Position text in header (centered horizontally, vertically aligned)
		textBounds := headerText.Bounds()
		textX := (finalWidth - textBounds.Max.X) / 2
		textY := (headerHeight - textBounds.Max.Y) / 2

		textPos := P(float64(textX), float64(textY))
		positionedText, _ := translateImage(headerText, *textPos)
		headerCanvas, _ = blend(headerCanvas, positionedText, "normal")
	}

	// Create body canvas
	bodyCanvas := IC(finalWidth, totalHeight, color.RGBA64{0, 0, 0, 0})
	bodyCanvas, _ = fill(bodyCanvas, bodyFillStyle)
	bodyCanvas, _ = border(bodyCanvas, bodyBorderStyle)

	// Add padded input image to body
	paddedImgPos := P(float64(paddingPx), float64(paddingPx))
	positionedImg, _ := translateImage(img, *paddedImgPos)
	bodyCanvas, _ = blend(bodyCanvas, positionedImg, "normal")

	// Position body below header
	bodyPos := P(0, float64(headerHeight)-borderThickness)
	positionedBody, _ := translateImage(bodyCanvas, *bodyPos)

	// Composite final result
	result, _ := blend(positionedBody, headerCanvas, "normal")

	return result, nil
}
