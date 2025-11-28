package language

import (
	_ "embed"
	"image"
	"image/color"

	"github.com/toxyl/pxp/fonts"
)

// @Name: draw-text
// @Desc: Draws a text at position (x,y).
// @Param:      img        - - -        The image to draw to
// @Param:      p          - - -        The upper-left coordinate of the text
// @Param:      t          - - -        The text to draw
// @Param:      colText    - - -        The text color
// @Param:      colOutline - - -        The outline color
// @Param:      blendMode  - - "normal" The blend mode to use
// @Returns:    result     - - -        The resulting image
func drawText(img *image.NRGBA64, p Point, t string, colText, colOutline color.RGBA64, blendMode string) (*image.NRGBA64, error) {
	bounds := img.Bounds()
	return drawTextPx(img, *p.Denorm(float64(bounds.Max.X), float64(bounds.Max.Y)), t, colText, colOutline, blendMode)
}

// @Name: draw-text-px
// @Desc: Draws text at position (x,y).
// @Param:      img        - - -        The image to draw to
// @Param:      p          - - -        The upper-left coordinate of the text
// @Param:      t          - - -        The text to draw
// @Param:      colText    - - -   		The text color
// @Param:      colOutline - - -   		The outline color
// @Param:      blendMode  - - "normal" The blend mode to use
// @Returns:    result     - - -	    The resulting image
func drawTextPx(img *image.NRGBA64, p Point, t string, colText, colOutline color.RGBA64, blendMode string) (*image.NRGBA64, error) {
	result := IClone(img)
	text, _ := text(t, colText, colOutline)
	res, _ := translateImage(text, p)
	return blend(result, res, blendMode)
}

// @Name: text
// @Desc: Generates the given text.
// @Param:      t         	- - -   		The text to generate
// @Param:      colText   	- - -   		The text color
// @Param:      colOutline  - - -   		The outline color
// @Returns:    result    	- - -			The resulting image
func text(t string, colText, colOutline color.RGBA64) (*image.NRGBA64, error) {
	// Convert non-premultiplied colors to premultiplied
	colText = color.RGBA64{
		R: uint16(uint32(colText.R) * uint32(colText.A) / 0xFFFF),
		G: uint16(uint32(colText.G) * uint32(colText.A) / 0xFFFF),
		B: uint16(uint32(colText.B) * uint32(colText.A) / 0xFFFF),
		A: colText.A,
	}
	colOutline = color.RGBA64{
		R: uint16(uint32(colOutline.R) * uint32(colOutline.A) / 0xFFFF),
		G: uint16(uint32(colOutline.G) * uint32(colOutline.A) / 0xFFFF),
		B: uint16(uint32(colOutline.B) * uint32(colOutline.A) / 0xFFFF),
		A: colOutline.A,
	}

	return fonts.PixelOperator.Render(t, colText, colOutline), nil
}
