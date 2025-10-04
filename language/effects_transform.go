package language

import (
	"image"
	"image/color"

	"github.com/toxyl/math"
)

// @Name: translate
// @Desc: Translates (moves) an image by a specified amount
// @Param:      img     - -   	-   The image to translate
// @Param:      dx      - -   	0   The horizontal translation amount in % (positive = right)
// @Param:      dy      - -   	0   The vertical translation amount in % (positive = down)
// @Returns:    result  - -   	-   The translated image
func translate(img *image.NRGBA64, dx float64, dy float64) (*image.NRGBA64, error) {
	bounds := img.Bounds()
	result := image.NewNRGBA64(bounds)
	dx *= float64(bounds.Dx())
	dy *= float64(bounds.Dy())

	// Fill with transparent black initially
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			result.Set(x, y, color.NRGBA{0, 0, 0, 0})
		}
	}

	// Copy pixels with offset
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			// Calculate source position
			srcX := x - int(dx)
			srcY := y - int(dy)

			// Only copy if source pixel is within bounds
			if srcX >= bounds.Min.X && srcX < bounds.Max.X && srcY >= bounds.Min.Y && srcY < bounds.Max.Y {
				result.Set(x, y, img.NRGBA64At(srcX, srcY))
			}
		}
	}

	return result, nil
}

// @Name: rotate
// @Desc: Rotates an image around its center by a specified angle
// @Param:      img     - -   			-   The image to rotate
// @Param:      angle   - -360..360   	0   The rotation angle in degrees (positive = clockwise)
// @Returns:    result  - -   			-   The rotated image
func rotate(img *image.NRGBA64, angle float64) (*image.NRGBA64, error) {
	bounds := img.Bounds()
	width := bounds.Dx()
	height := bounds.Dy()

	// Convert angle to radians
	angleRad := angle * math.Pi / 180.0

	// Calculate sin and cos once
	sinA := math.Abs(math.Sin(angleRad)) // Positive angle for clockwise rotation
	cosA := math.Abs(math.Cos(angleRad))

	// Calculate new dimensions to fit the rotated image
	newWidth := int(math.Round(float64(width)*cosA + float64(height)*sinA))
	newHeight := int(math.Round(float64(width)*sinA + float64(height)*cosA))

	// Create a new image with the calculated dimensions
	newBounds := image.Rect(0, 0, newWidth, newHeight)
	result := image.NewNRGBA64(newBounds)

	// Calculate center points for both original and new image
	oldCenterX := float64(width) / 2.0
	oldCenterY := float64(height) / 2.0
	newCenterX := float64(newWidth) / 2.0
	newCenterY := float64(newHeight) / 2.0

	// Calculate actual sin and cos for rotation
	sinA = math.Sin(angleRad)
	cosA = math.Cos(angleRad)

	// Fill with transparent black initially
	for y := 0; y < newHeight; y++ {
		for x := 0; x < newWidth; x++ {
			result.Set(x, y, color.NRGBA{0, 0, 0, 0})
		}
	}

	// Rotate each pixel
	for y := range newHeight {
		for x := range newWidth {
			// Translate to origin of new image
			dx := float64(x) - newCenterX
			dy := float64(y) - newCenterY

			// Rotate backwards (use negative angle to get source position)
			srcX := int(math.Round(dx*cosA + dy*sinA + oldCenterX))
			srcY := int(math.Round(-dx*sinA + dy*cosA + oldCenterY))

			// Copy pixel if within bounds of original image
			if srcX >= bounds.Min.X && srcX < bounds.Max.X && srcY >= bounds.Min.Y && srcY < bounds.Max.Y {
				result.Set(x, y, img.NRGBA64At(srcX, srcY))
			}
		}
	}

	return result, nil
}

// @Name: scale
// @Desc: Scales an image by specified factors
// @Param:      img     - -   	-   The image to scale
// @Param:      sx      - -  	0   The horizontal scale factor
// @Param:      sy      - -  	0   The vertical scale factor
// @Returns:    result  - -   	-   The scaled image
func scale(img *image.NRGBA64, sx float64, sy float64) (*image.NRGBA64, error) {
	sx += 1
	sy += 1

	bounds := img.Bounds()
	width := bounds.Dx()
	height := bounds.Dy()

	// Calculate new dimensions
	newWidth := int(math.Round(float64(width) * sx))
	newHeight := int(math.Round(float64(height) * sy))

	// Create result image with new dimensions
	newBounds := image.Rect(0, 0, newWidth, newHeight)
	result := image.NewNRGBA64(newBounds)

	// Fill with transparent black initially
	for y := range newHeight {
		for x := range newWidth {
			result.Set(x, y, color.NRGBA{0, 0, 0, 0})
		}
	}

	// Scale each pixel
	for y := range newHeight {
		for x := range newWidth {
			// Calculate source pixel coordinates
			srcX := int(math.Round(float64(x) / sx))
			srcY := int(math.Round(float64(y) / sy))

			// Copy pixel if within bounds of source image
			if srcX >= bounds.Min.X && srcX < bounds.Max.X && srcY >= bounds.Min.Y && srcY < bounds.Max.Y {
				result.Set(x, y, img.NRGBA64At(srcX, srcY))
			}
		}
	}

	return result, nil
}

// @Name: transform
// @Desc: Applies translation, rotation, and scaling to an image in one operation
// @Param:      img     - -   	-   The image to transform
// @Param:      dx      - -   	0   The horizontal translation in pixels
// @Param:      dy      - -   	0   The vertical translation in pixels
// @Param:      angle   - -   	0   The rotation angle in degrees (clockwise)
// @Param:      sx      - -  	0   The horizontal scale factor
// @Param:      sy      - -  	0   The vertical scale factor
// @Returns:    result  - -   	-   The transformed image
func transform(img *image.NRGBA64, dx float64, dy float64, angle float64, sx float64, sy float64) (*image.NRGBA64, error) {
	sx += 1
	sy += 1

	bounds := img.Bounds()
	width := bounds.Dx()
	height := bounds.Dy()
	result := image.NewNRGBA64(bounds)

	dx *= float64(width)
	dy *= float64(height)

	// Convert angle to radians
	angleRad := angle * math.Pi / 180.0

	// Calculate center point
	centerX := float64(width) / 2.0
	centerY := float64(height) / 2.0

	// Calculate sin and cos once
	sinA := math.Sin(-angleRad) // Negative angle for clockwise rotation
	cosA := math.Cos(-angleRad)

	// Fill with transparent black initially
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			result.Set(x, y, color.NRGBA{0, 0, 0, 0})
		}
	}

	// Apply transformation to each pixel
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			// Translate to origin
			px := (float64(x) - dx) - centerX
			py := (float64(y) - dy) - centerY

			// Apply scale
			px /= sx
			py /= sy

			// Apply rotation
			srcX := int(math.Round(px*cosA - py*sinA + centerX))
			srcY := int(math.Round(px*sinA + py*cosA + centerY))

			// Copy pixel if within bounds
			if srcX >= bounds.Min.X && srcX < bounds.Max.X && srcY >= bounds.Min.Y && srcY < bounds.Max.Y {
				result.Set(x, y, img.NRGBA64At(srcX, srcY))
			}
		}
	}

	return result, nil
}

// @Name: flip-v
// @Desc: Flips an image vertically (top to bottom)
// @Param:      img     - -   	-   The image to flip vertically
// @Returns:    result  - -   	-   The vertically flipped image
func flipVertical(img *image.NRGBA64) (*image.NRGBA64, error) {
	bounds := img.Bounds()
	result := image.NewNRGBA64(bounds)
	height := bounds.Dy()

	// Copy pixels with vertical flip
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		// Calculate the corresponding y position in the flipped image
		flippedY := height - 1 - (y - bounds.Min.Y) + bounds.Min.Y

		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			result.Set(x, y, img.NRGBA64At(x, flippedY))
		}
	}

	return result, nil
}

// @Name: flip-h
// @Desc: Flips an image horizontally (left to right)
// @Param:      img     - -   	-   The image to flip horizontally
// @Returns:    result  - -   	-   The horizontally flipped image
func flipHorizontal(img *image.NRGBA64) (*image.NRGBA64, error) {
	bounds := img.Bounds()
	result := image.NewNRGBA64(bounds)
	width := bounds.Dx()

	// Copy pixels with horizontal flip
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			// Calculate the corresponding x position in the flipped image
			flippedX := width - 1 - (x - bounds.Min.X) + bounds.Min.X

			result.Set(x, y, img.NRGBA64At(flippedX, y))
		}
	}

	return result, nil
}

// @Name: crop
// @Desc: Crops an image by specified percentages from each side
// @Param:      img     - -   	-   The image to crop
// @Param:      left    - -   	0   The percentage to crop from the left side (0-1)
// @Param:      right   - -   	0   The percentage to crop from the right side (0-1)
// @Param:      top     - -   	0   The percentage to crop from the top side (0-1)
// @Param:      bottom  - -   	0   The percentage to crop from the bottom side (0-1)
// @Returns:    result  - -   	-   The cropped image
func crop(img *image.NRGBA64, left float64, right float64, top float64, bottom float64) (*image.NRGBA64, error) {
	bounds := img.Bounds()
	width := bounds.Dx()
	height := bounds.Dy()

	// Convert percentages to pixels
	leftPx := int(math.Round(float64(width) * left))
	rightPx := int(math.Round(float64(width) * right))
	topPx := int(math.Round(float64(height) * top))
	bottomPx := int(math.Round(float64(height) * bottom))

	return cropPx(img, leftPx, rightPx, topPx, bottomPx)
}

// @Name: crop-px
// @Desc: Crops an image by specified amounts of pixels from each side
// @Param:      img     - -   	-   The image to crop
// @Param:      left    - -   	0   The number of pixels to crop from the left side
// @Param:      right   - -   	0   The number of pixels to crop from the right side
// @Param:      top     - -   	0   The number of pixels to crop from the top side
// @Param:      bottom  - -   	0   The number of pixels to crop from the bottom side
// @Returns:    result  - -   	-   The cropped image
func cropPx(img *image.NRGBA64, left int, right int, top int, bottom int) (*image.NRGBA64, error) {
	bounds := img.Bounds()
	width := bounds.Dx()
	height := bounds.Dy()

	// Calculate new dimensions
	newWidth := width - left - right
	newHeight := height - top - bottom

	// Create result image with new dimensions
	newBounds := image.Rect(0, 0, newWidth, newHeight)
	result := image.NewNRGBA64(newBounds)

	// Copy the cropped region
	for y := range newHeight {
		for x := range newWidth {
			srcX := x + left
			srcY := y + top
			result.Set(x, y, img.NRGBA64At(srcX, srcY))
		}
	}

	return result, nil
}

// @Name: crop-circle
// @Desc: Crops an image using a circular mask. The circle is centered at (centerX+offsetX, centerY+offsetY) and the radius is a percentage (0-1) of half the minimum image dimension.
// @Param:      img      - -   	-   The image to crop
// @Param:      radius   - 0..1	1   Radius as a percentage of half the min(width, height)
// @Param:      offsetX  - -1..1	0   Horizontal offset from image center (percentage of width, -1..1)
// @Param:      offsetY  - -1..1	0   Vertical offset from image center (percentage of height, -1..1)
// @Returns:    result   - -   	-   The circularly cropped image (pixels outside the circle are transparent)
func cropCircle(img *image.NRGBA64, radius float64, offsetX float64, offsetY float64) (*image.NRGBA64, error) {
	bounds := img.Bounds()
	return cropCirclePx(img, radius, int(offsetX*(float64(bounds.Dx())/2.0)), int(offsetY*(float64(bounds.Dy())/2.0)))
}

// @Name: crop-circle-px
// @Desc: Crops an image using a circular mask. The circle is centered at (centerX+offsetX, centerY+offsetY) and the radius is a percentage (0-1) of half the minimum image dimension.
// @Param:      img      - -   	-   The image to crop
// @Param:      radius   - 0..1	1   Radius as a percentage of half the min(width, height)
// @Param:      offsetX  - -   	0   Horizontal offset from image center (pixels)
// @Param:      offsetY  - -   	0   Vertical offset from image center (pixels)
// @Returns:    result   - -   	-   The circularly cropped image (pixels outside the circle are transparent)
func cropCirclePx(img *image.NRGBA64, radius float64, offsetX int, offsetY int) (*image.NRGBA64, error) {
	bounds := img.Bounds()
	width := bounds.Dx()
	height := bounds.Dy()
	centerX := width/2 + offsetX
	centerY := height/2 + offsetY

	// Use the smaller of width or height for radius calculation
	minHalf := min(height, width)
	minHalf = minHalf / 2
	r := int(math.Round(float64(minHalf) * radius))
	r2 := r * r // squared radius for distance check

	// Calculate bounding box of the circle, clamp to image bounds
	cropMinX := centerX - r
	cropMaxX := centerX + r
	cropMinY := centerY - r
	cropMaxY := centerY + r

	if cropMinX < bounds.Min.X {
		cropMinX = bounds.Min.X
	}
	if cropMaxX > bounds.Max.X {
		cropMaxX = bounds.Max.X
	}
	if cropMinY < bounds.Min.Y {
		cropMinY = bounds.Min.Y
	}
	if cropMaxY > bounds.Max.Y {
		cropMaxY = bounds.Max.Y
	}

	newWidth := cropMaxX - cropMinX
	newHeight := cropMaxY - cropMinY
	newBounds := image.Rect(0, 0, newWidth, newHeight)
	result := image.NewNRGBA64(newBounds)
	transparent := color.NRGBA64{0, 0, 0, 0}

	for y := range newHeight {
		for x := range newWidth {
			srcX := cropMinX + x
			srcY := cropMinY + y
			dx := srcX - centerX
			dy := srcY - centerY
			if dx*dx+dy*dy <= r2 {
				result.Set(x, y, img.NRGBA64At(srcX, srcY))
			} else {
				result.Set(x, y, transparent)
			}
		}
	}

	return result, nil
}

// @Name: crop-square
// @Desc: Crops an image using a square mask. The square is centered at (centerX+offsetX, centerY+offsetY) and the size is a percentage (0-1) of the minimum image dimension.
// @Param:      img      - -   	-   The image to crop
// @Param:      size     - 0..1	1   Size as a percentage of the min(width, height)
// @Param:      offsetX  - -1..1	0   Horizontal offset from image center (percentage of width, -1..1)
// @Param:      offsetY  - -1..1	0   Vertical offset from image center (percentage of height, -1..1)
// @Returns:    result   - -   	-   The square-cropped image (pixels outside the square are transparent)
func cropSquare(img *image.NRGBA64, size float64, offsetX float64, offsetY float64) (*image.NRGBA64, error) {
	bounds := img.Bounds()
	return cropSquarePx(
		img,
		size,
		int(offsetX*(float64(bounds.Dx())/2.0)),
		int(offsetY*(float64(bounds.Dy())/2.0)),
	)
}

// @Name: crop-square-px
// @Desc: Crops an image using a square mask. The square is centered at (centerX+offsetX, centerY+offsetY) and the size is a percentage (0-1) of the minimum image dimension.
// @Param:      img      - -   	-   The image to crop
// @Param:      size     - 0..1	1   Size as a percentage of the min(width, height)
// @Param:      offsetX  - -   	0   Horizontal offset from image center (pixels)
// @Param:      offsetY  - -   	0   Vertical offset from image center (pixels)
// @Returns:    result   - -   	-   The square-cropped image (pixels outside the square are transparent)
func cropSquarePx(img *image.NRGBA64, size float64, offsetX int, offsetY int) (*image.NRGBA64, error) {
	bounds := img.Bounds()
	width := bounds.Dx()
	height := bounds.Dy()
	centerX := width/2 + offsetX
	centerY := height/2 + offsetY

	// Use the smaller of width or height for square size calculation
	minDim := min(height, width)
	halfSide := int(math.Round(float64(minDim) * size / 2.0))

	// Calculate the cropped square's bounds
	cropMinX := centerX - halfSide
	cropMaxX := centerX + halfSide
	cropMinY := centerY - halfSide
	cropMaxY := centerY + halfSide

	// Clamp crop bounds to image bounds
	if cropMinX < bounds.Min.X {
		cropMinX = bounds.Min.X
	}
	if cropMaxX > bounds.Max.X {
		cropMaxX = bounds.Max.X
	}
	if cropMinY < bounds.Min.Y {
		cropMinY = bounds.Min.Y
	}
	if cropMaxY > bounds.Max.Y {
		cropMaxY = bounds.Max.Y
	}

	cropWidth := cropMaxX - cropMinX
	cropHeight := cropMaxY - cropMinY

	resultBounds := image.Rect(0, 0, cropWidth, cropHeight)
	result := image.NewNRGBA64(resultBounds)

	for y := range cropHeight {
		for x := range cropWidth {
			srcX := cropMinX + x
			srcY := cropMinY + y
			if srcX >= bounds.Min.X && srcX < bounds.Max.X && srcY >= bounds.Min.Y && srcY < bounds.Max.Y {
				result.Set(x, y, img.NRGBA64At(srcX, srcY))
			} else {
				result.Set(x, y, color.NRGBA64{0, 0, 0, 0})
			}
		}
	}

	return result, nil
}

// @Name: expand
// @Desc: Expands an image by adding transparent borders with specified percentage widths
// @Param:      img     - -   	-   The image to expand
// @Param:      left    - -   	0   The percentage to add to the left side (relative to original width)
// @Param:      right   - -   	0   The percentage to add to the right side (relative to original width)
// @Param:      top     - -   	0   The percentage to add to the top side (relative to original height)
// @Param:      bottom  - -   	0   The percentage to add to the bottom side (relative to original height)
// @Returns:    result  - -   	-   The expanded image
func expand(img *image.NRGBA64, left float64, right float64, top float64, bottom float64) (*image.NRGBA64, error) {
	bounds := img.Bounds()
	width := bounds.Dx()
	height := bounds.Dy()

	// Convert percentages to pixels
	leftPx := int(math.Round(float64(width) * left))
	rightPx := int(math.Round(float64(width) * right))
	topPx := int(math.Round(float64(height) * top))
	bottomPx := int(math.Round(float64(height) * bottom))

	return expandPx(img, leftPx, rightPx, topPx, bottomPx)
}

// @Name: expand-px
// @Desc: Expands an image by adding transparent borders with specified pixel widths
// @Param:      img     - -   	-   The image to expand
// @Param:      left    - -   	0   The number of pixels to add to the left side
// @Param:      right   - -   	0   The number of pixels to add to the right side
// @Param:      top     - -   	0   The number of pixels to add to the top side
// @Param:      bottom  - -   	0   The number of pixels to add to the bottom side
// @Returns:    result  - -   	-   The expanded image
func expandPx(img *image.NRGBA64, left int, right int, top int, bottom int) (*image.NRGBA64, error) {
	bounds := img.Bounds()
	width := bounds.Dx()
	height := bounds.Dy()

	// Calculate new dimensions
	newWidth := width + left + right
	newHeight := height + top + bottom

	// Create result image with new dimensions
	newBounds := image.Rect(0, 0, newWidth, newHeight)
	result := image.NewNRGBA64(newBounds)

	// Fill with transparent black initially
	for y := range newHeight {
		for x := range newWidth {
			result.Set(x, y, color.NRGBA{0, 0, 0, 0})
		}
	}

	// Copy the original image to the expanded canvas
	for y := range height {
		for x := range width {
			result.Set(x+left, y+top, img.NRGBA64At(x, y))
		}
	}

	return result, nil
}

// @Name: resize-max-mp
// @Desc: Resize an image to stay within a maximum amount of megapixels
// @Param:      img     - -   	-   The image to resize
// @Param:      mpMax    - -   	0   The maximum amount of megapixels
// @Returns:    result  - -   	-   The resized image
func resizeToMaxMP(img *image.NRGBA64, mpMax int) (*image.NRGBA64, error) {
	bounds := img.Bounds()
	w := bounds.Dx()
	h := bounds.Dy()

	ow := w
	oh := h
	mp := w * h
	for mp > mpMax {
		w >>= 1
		h >>= 1
		mp = w * h
	}
	if ow != w || oh != h {
		// Create a new image with the target dimensions and resample using nearest-neighbor
		dst := image.NewNRGBA64(image.Rect(0, 0, w, h))
		for y := 0; y < h; y++ {
			srcY := y * oh / h
			for x := 0; x < w; x++ {
				srcX := x * ow / w
				dst.SetNRGBA64(x, y, img.NRGBA64At(srcX, srcY))
			}
		}
		img = dst
	}
	return img, nil
}
