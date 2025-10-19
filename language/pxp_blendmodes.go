package language

import (
	"image"
)

// @Name: blend
// @Desc: Blends the two images using the given blendmode (defaults to normal)
// @Param:      imgA    - -   	-   		The bottom image
// @Param:      imgB    - -   	-   		The top image
// @Param:      mode    - -   	"normal"    The blendmode name
// @Returns:    result  - -   	-   		The blended image
func blend(imgA, imgB *image.NRGBA64, mode string) (*image.NRGBA64, error) {
	return blenders.BlendImages(mode, imgA, imgB), nil
}

// @Name: blend-normal
// @Desc: Blends the two images using the normal blend mode (alpha compositing)
// @Param:      imgA     - -   	-   The bottom image
// @Param:      imgB     - -   	-   The top image
// @Returns:    result  - -   	-   The blended image
func blendNormal(imgA, imgB *image.NRGBA64) (*image.NRGBA64, error) {
	return blend(imgA, imgB, NORMAL)
}

// @Name: blend-erase
// @Desc: Erases the bottom image wherever the top image is present (destination out)
// @Param:      imgA     - -   	-   The bottom image
// @Param:      imgB     - -   	-   The top image
// @Returns:    result  - -   	-   The blended image
func blendErase(imgA, imgB *image.NRGBA64) (*image.NRGBA64, error) {
	return blend(imgA, imgB, ERASE)
}

// @Name: blend-multiply
// @Desc: Blends the two images using the multiply blend mode
// @Param:      imgA     - -   	-   The bottom image
// @Param:      imgB     - -   	-   The top image
// @Returns:    result  - -   	-   The blended image
func blendMultiply(imgA, imgB *image.NRGBA64) (*image.NRGBA64, error) {
	return blend(imgA, imgB, MULTIPLY)
}

// @Name: blend-screen
// @Desc: Blends the two images using the screen blend mode
// @Param:      imgA     - -   	-   The bottom image
// @Param:      imgB     - -   	-   The top image
// @Returns:    result  - -   	-   The blended image
func blendScreen(imgA, imgB *image.NRGBA64) (*image.NRGBA64, error) {
	return blend(imgA, imgB, SCREEN)
}

// @Name: blend-exclusion
// @Desc: Blends the two images using the exclusion blend mode
// @Param:      imgA     - -   	-   The bottom image
// @Param:      imgB     - -   	-   The top image
// @Returns:    result  - -   	-   The blended image
func blendExclusion(imgA, imgB *image.NRGBA64) (*image.NRGBA64, error) {
	return blend(imgA, imgB, EXCLUSION)
}

// @Name: blend-overlay
// @Desc: Blends the two images using the overlay blend mode
// @Param:      imgA     - -   	-   The bottom image
// @Param:      imgB     - -   	-   The top image
// @Returns:    result  - -   	-   The blended image
func blendOverlay(imgA, imgB *image.NRGBA64) (*image.NRGBA64, error) {
	return blend(imgA, imgB, OVERLAY)
}

// @Name: blend-color-burn
// @Desc: Blends the two images using the color burn blend mode
// @Param:      imgA     - -   	-   The bottom image
// @Param:      imgB     - -   	-   The top image
// @Returns:    result  - -   	-   The blended image
func blendColorBurn(imgA, imgB *image.NRGBA64) (*image.NRGBA64, error) {
	return blend(imgA, imgB, COLOR_BURN)
}

// @Name: blend-color-dodge
// @Desc: Blends the two images using the color dodge blend mode
// @Param:      imgA     - -   	-   The bottom image
// @Param:      imgB     - -   	-   The top image
// @Returns:    result  - -   	-   The blended image
func blendColorDodge(imgA, imgB *image.NRGBA64) (*image.NRGBA64, error) {
	return blend(imgA, imgB, COLOR_DODGE)
}

// @Name: blend-soft-light
// @Desc: Blends the two images using the soft light blend mode
// @Param:      imgA     - -   	-   The bottom image
// @Param:      imgB     - -   	-   The top image
// @Returns:    result  - -   	-   The blended image
func blendSoftLight(imgA, imgB *image.NRGBA64) (*image.NRGBA64, error) {
	return blend(imgA, imgB, SOFT_LIGHT)
}

// @Name: blend-hard-light
// @Desc: Blends the two images using the hard light blend mode
// @Param:      imgA     - -   	-   The bottom image
// @Param:      imgB     - -   	-   The top image
// @Returns:    result  - -   	-   The blended image
func blendHardLight(imgA, imgB *image.NRGBA64) (*image.NRGBA64, error) {
	return blend(imgA, imgB, HARD_LIGHT)
}

// @Name: blend-difference
// @Desc: Blends the two images using the difference blend mode
// @Param:      imgA     - -   	-   The bottom image
// @Param:      imgB     - -   	-   The top image
// @Returns:    result  - -   	-   The blended image
func blendDifference(imgA, imgB *image.NRGBA64) (*image.NRGBA64, error) {
	return blend(imgA, imgB, DIFFERENCE)
}

// @Name: blend-subtract
// @Desc: Blends the two images using the subtract blend mode
// @Param:      imgA     - -   	-   The bottom image
// @Param:      imgB     - -   	-   The top image
// @Returns:    result  - -   	-   The blended image
func blendSubtract(imgA, imgB *image.NRGBA64) (*image.NRGBA64, error) {
	return blend(imgA, imgB, SUBTRACT)
}

// @Name: blend-divide
// @Desc: Blends the two images using the divide blend mode
// @Param:      imgA     - -   	-   The bottom image
// @Param:      imgB     - -   	-   The top image
// @Returns:    result  - -   	-   The blended image
func blendDivide(imgA, imgB *image.NRGBA64) (*image.NRGBA64, error) {
	return blend(imgA, imgB, DIVIDE)
}

// @Name: blend-hue
// @Desc: Blends the two images using the hue blend mode
// @Param:      imgA     - -   	-   The bottom image
// @Param:      imgB     - -   	-   The top image
// @Returns:    result  - -   	-   The blended image
func blendHue(imgA, imgB *image.NRGBA64) (*image.NRGBA64, error) {
	return blend(imgA, imgB, HUE)
}

// @Name: blend-saturation
// @Desc: Blends the two images using the saturation blend mode
// @Param:      imgA     - -   	-   The bottom image
// @Param:      imgB     - -   	-   The top image
// @Returns:    result  - -   	-   The blended image
func blendSaturation(imgA, imgB *image.NRGBA64) (*image.NRGBA64, error) {
	return blend(imgA, imgB, SATURATION)
}

// @Name: blend-color
// @Desc: Blends the two images using the color blend mode
// @Param:      imgA     - -   	-   The bottom image
// @Param:      imgB     - -   	-   The top image
// @Returns:    result  - -   	-   The blended image
func blendColor(imgA, imgB *image.NRGBA64) (*image.NRGBA64, error) {
	return blend(imgA, imgB, COLOR)
}

// @Name: blend-luminosity
// @Desc: Blends the two images using the luminosity blend mode
// @Param:      imgA     - -   	-   The bottom image
// @Param:      imgB     - -   	-   The top image
// @Returns:    result  - -   	-   The blended image
func blendLuminosity(imgA, imgB *image.NRGBA64) (*image.NRGBA64, error) {
	return blend(imgA, imgB, LUMINOSITY)
}

// @Name: blend-average
// @Desc: Blends the two images using the average blend mode
// @Param:      imgA     - -   	-   The bottom image
// @Param:      imgB     - -   	-   The top image
// @Returns:    result  - -   	-   The blended image
func blendAverage(imgA, imgB *image.NRGBA64) (*image.NRGBA64, error) {
	return blend(imgA, imgB, AVERAGE)
}

// @Name: blend-negation
// @Desc: Blends the two images using the negation blend mode
// @Param:      imgA     - -   	-   The bottom image
// @Param:      imgB     - -   	-   The top image
// @Returns:    result  - -   	-   The blended image
func blendNegation(imgA, imgB *image.NRGBA64) (*image.NRGBA64, error) {
	return blend(imgA, imgB, NEGATION)
}

// @Name: blend-reflect
// @Desc: Blends the two images using the reflect blend mode
// @Param:      imgA     - -   	-   The bottom image
// @Param:      imgB     - -   	-   The top image
// @Returns:    result  - -   	-   The blended image
func blendReflect(imgA, imgB *image.NRGBA64) (*image.NRGBA64, error) {
	return blend(imgA, imgB, REFLECT)
}

// @Name: blend-glow
// @Desc: Blends the two images using the glow blend mode
// @Param:      imgA     - -   	-   The bottom image
// @Param:      imgB     - -   	-   The top image
// @Returns:    result  - -   	-   The blended image
func blendGlow(imgA, imgB *image.NRGBA64) (*image.NRGBA64, error) {
	return blend(imgA, imgB, GLOW)
}

// @Name: blend-contrast-negate
// @Desc: Blends the two images using the contrast negate blend mode
// @Param:      imgA     - -   	-   The bottom image
// @Param:      imgB     - -   	-   The top image
// @Returns:    result  - -   	-   The blended image
func blendContrastNegate(imgA, imgB *image.NRGBA64) (*image.NRGBA64, error) {
	return blend(imgA, imgB, CONTRAST_NEGATE)
}

// @Name: blend-vivid-light
// @Desc: Blends the two images using the vivid light blend mode
// @Param:      imgA     - -   	-   The bottom image
// @Param:      imgB     - -   	-   The top image
// @Returns:    result  - -   	-   The blended image
func blendVividLight(imgA, imgB *image.NRGBA64) (*image.NRGBA64, error) {
	return blend(imgA, imgB, VIVID_LIGHT)
}

// @Name: blend-linear-light
// @Desc: Blends the two images using the linear light blend mode
// @Param:      imgA     - -   	-   The bottom image
// @Param:      imgB     - -   	-   The top image
// @Returns:    result  - -   	-   The blended image
func blendLinearLight(imgA, imgB *image.NRGBA64) (*image.NRGBA64, error) {
	return blend(imgA, imgB, LINEAR_LIGHT)
}

// @Name: blend-pin-light
// @Desc: Blends the two images using the pin light blend mode
// @Param:      imgA     - -   	-   The bottom image
// @Param:      imgB     - -   	-   The top image
// @Returns:    result  - -   	-   The blended image
func blendPinLight(imgA, imgB *image.NRGBA64) (*image.NRGBA64, error) {
	return blend(imgA, imgB, PIN_LIGHT)
}

// @Name: blend-darken
// @Desc: Blends the two images using the darken blend mode
// @Param:      imgA     - -   	-   The bottom image
// @Param:      imgB     - -   	-   The top image
// @Returns:    result  - -   	-   The blended image
func blendDarken(imgA, imgB *image.NRGBA64) (*image.NRGBA64, error) {
	return blend(imgA, imgB, DARKEN)
}

// @Name: blend-darker-color
// @Desc: Blends the two images using the darker color blend mode
// @Param:      imgA     - -   	-   The bottom image
// @Param:      imgB     - -   	-   The top image
// @Returns:    result  - -   	-   The blended image
func blendDarkerColor(imgA, imgB *image.NRGBA64) (*image.NRGBA64, error) {
	return blend(imgA, imgB, DARKER_COLOR)
}

// @Name: blend-lighten
// @Desc: Blends the two images using the lighten blend mode
// @Param:      imgA     - -   	-   The bottom image
// @Param:      imgB     - -   	-   The top image
// @Returns:    result  - -   	-   The blended image
func blendLighten(imgA, imgB *image.NRGBA64) (*image.NRGBA64, error) {
	return blend(imgA, imgB, LIGHTEN)
}

// @Name: blend-lighter-color
// @Desc: Blends the two images using the lighter color blend mode
// @Param:      imgA     - -   	-   The bottom image
// @Param:      imgB     - -   	-   The top image
// @Returns:    result  - -   	-   The blended image
func blendLighterColor(imgA, imgB *image.NRGBA64) (*image.NRGBA64, error) {
	return blend(imgA, imgB, LIGHTER_COLOR)
}

// @Name: blend-hard-mix
// @Desc: Blends the two images using the hard mix blend mode
// @Param:      imgA     - -   	-   The bottom image
// @Param:      imgB     - -   	-   The top image
// @Returns:    result  - -   	-   The blended image
func blendHardMix(imgA, imgB *image.NRGBA64) (*image.NRGBA64, error) {
	return blend(imgA, imgB, HARD_MIX)
}
