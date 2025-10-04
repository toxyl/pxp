package language

import (
	"fmt"
	"image"
	"image/color"

	"github.com/toxyl/math"
)

// @Name: rectangular-to-polar
// @Desc: Converts a rectangular coordinate image to polar coordinates
// @Param:      img     - -   	-   The image to transform
// @Returns:    result  - -   	-   The transformed image
func distortRectangularToPolar(img *image.NRGBA64) (*image.NRGBA64, error) {
	bounds := img.Bounds()
	width := bounds.Dx()
	height := bounds.Dy()
	result := image.NewNRGBA64(bounds)

	centerX := float64(width) / 2
	centerY := float64(height) / 2
	maxRadius := math.Min(centerX, centerY)

	for y := range height {
		for x := range width {
			// Convert rectangular coordinates to polar
			dx := float64(x) - centerX
			dy := float64(y) - centerY

			// Calculate radius (0 to 1) and angle (0 to 2π)
			radius := math.Sqrt(dx*dx+dy*dy) / maxRadius
			angle := math.Atan2(dy, dx)
			if angle < 0 {
				angle += 2 * math.Pi
			}

			// Map polar coordinates back to source image
			srcX := int(angle / (2 * math.Pi) * float64(width))
			srcY := int(radius * float64(height))

			// Get color from source coordinates
			if srcX >= 0 && srcX < width && srcY >= 0 && srcY < height {
				c := img.NRGBA64At(srcX+bounds.Min.X, srcY+bounds.Min.Y)
				result.Set(x+bounds.Min.X, y+bounds.Min.Y, c)
			}
		}
	}

	return result, nil
}

// @Name: polar-to-rectangular
// @Desc: Converts a polar coordinate image to rectangular coordinates
// @Param:      img     - -   	-   The image to transform
// @Returns:    result  - -   	-   The transformed image
func distortPolarToRectangular(img *image.NRGBA64) (*image.NRGBA64, error) {
	bounds := img.Bounds()
	width := bounds.Dx()
	height := bounds.Dy()
	result := image.NewNRGBA64(bounds)

	centerX := float64(width) / 2
	centerY := float64(height) / 2
	maxRadius := math.Min(centerX, centerY)

	for y := range height {
		for x := range width {
			// Convert output position to radius and angle
			radius := float64(y) / float64(height)
			angle := float64(x) / float64(width) * 2 * math.Pi

			// Convert polar to rectangular coordinates
			srcX := int(centerX + radius*maxRadius*math.Cos(angle))
			srcY := int(centerY + radius*maxRadius*math.Sin(angle))

			// Get color from source coordinates
			if srcX >= 0 && srcX < width && srcY >= 0 && srcY < height {
				c := img.NRGBA64At(srcX+bounds.Min.X, srcY+bounds.Min.Y)
				result.Set(x+bounds.Min.X, y+bounds.Min.Y, c)
			}
		}
	}

	return result, nil
}

// @Name: pixelate
// @Desc: Creates a pixelation effect by averaging colors in blocks
// @Param:      img      - -   	-   The image to pixelate
// @Param:      size     - 1..50 8   The size of the pixel blocks
// @Returns:    result   - -   	-   The pixelated image
func distortPixelate(img *image.NRGBA64, size int) (*image.NRGBA64, error) {
	if size < 1 || size > 50 {
		return nil, fmt.Errorf("pixel size must be between 1 and 50")
	}

	bounds := img.Bounds()
	width := bounds.Dx()
	height := bounds.Dy()
	result := image.NewNRGBA64(bounds)

	// Process each block
	for blockY := 0; blockY < height; blockY += size {
		for blockX := 0; blockX < width; blockX += size {
			var sumR, sumG, sumB, sumA, count float64

			// Calculate average color for block
			for y := 0; y < size && blockY+y < height; y++ {
				for x := 0; x < size && blockX+x < width; x++ {
					c := img.NRGBA64At(blockX+x+bounds.Min.X, blockY+y+bounds.Min.Y)
					sumR += float64(c.R)
					sumG += float64(c.G)
					sumB += float64(c.B)
					sumA += float64(c.A)
					count++
				}
			}

			// Calculate average
			avgR := uint16(sumR / count)
			avgG := uint16(sumG / count)
			avgB := uint16(sumB / count)
			avgA := uint16(sumA / count)

			// Fill block with average color
			for y := 0; y < size && blockY+y < height; y++ {
				for x := 0; x < size && blockX+x < width; x++ {
					result.Set(blockX+x+bounds.Min.X, blockY+y+bounds.Min.Y,
						color.NRGBA64{R: avgR, G: avgG, B: avgB, A: avgA})
				}
			}
		}
	}

	return result, nil
}

// @Name: displace
// @Desc: Displaces pixels based on the brightness of a displacement map
// @Param:      img      - -   	-   The image to displace
// @Param:      dMap      - -   	-   The displacement map image
// @Param:      amount   - 0..50 10  The amount of displacement
// @Returns:    result   - -   	-   The displaced image
func distortDisplace(img *image.NRGBA64, dMap *image.NRGBA64, amount float64) (*image.NRGBA64, error) {
	if amount < 0 || amount > 50 {
		return nil, fmt.Errorf("displacement amount must be between 0 and 50")
	}

	bounds := img.Bounds()
	width := bounds.Dx()
	height := bounds.Dy()
	result := image.NewNRGBA64(bounds)

	// Verify displacement map size matches input image
	if dMap.Bounds().Dx() != width || dMap.Bounds().Dy() != height {
		return nil, fmt.Errorf("displacement map dimensions must match input image")
	}

	for y := range height {
		for x := range width {
			// Get displacement from map
			mapColor := dMap.NRGBA64At(x+bounds.Min.X, y+bounds.Min.Y)

			// Use red channel for X displacement and green for Y displacement
			// Normalize from [0,65535] to [-1,1]
			dispX := (float64(mapColor.R) - 32767.5) / 32767.5
			dispY := (float64(mapColor.G) - 32767.5) / 32767.5

			// Calculate source position
			srcX := int(float64(x) + dispX*amount)
			srcY := int(float64(y) + dispY*amount)

			// Get color from source position with bounds checking
			if srcX >= 0 && srcX < width && srcY >= 0 && srcY < height {
				c := img.NRGBA64At(srcX+bounds.Min.X, srcY+bounds.Min.Y)
				result.Set(x+bounds.Min.X, y+bounds.Min.Y, c)
			} else {
				// Use edge color instead of black for out-of-bounds
				edgeX := math.Clamp(srcX, 0, width-1)
				edgeY := math.Clamp(srcY, 0, height-1)
				c := img.NRGBA64At(edgeX+bounds.Min.X, edgeY+bounds.Min.Y)
				result.Set(x+bounds.Min.X, y+bounds.Min.Y, c)
			}
		}
	}

	return result, nil
}

// @Name: defisheye
// @Desc: Corrects fisheye lens distortion in an image
// @Param:      img      - -   	-   The image to correct
// @Param:      strength - 0..2  1   The strength of the correction
// @Returns:    result   - -   	-   The corrected image
func distortDefisheye(img *image.NRGBA64, strength float64) (*image.NRGBA64, error) {
	bounds := img.Bounds()
	width := bounds.Dx()
	height := bounds.Dy()
	result := image.NewNRGBA64(bounds)

	centerX := float64(width) / 2
	centerY := float64(height) / 2
	maxRadius := math.Min(centerX, centerY)

	for y := range height {
		for x := range width {
			// Calculate normalized coordinates from center (-1 to 1)
			dx := (float64(x) - centerX) / maxRadius
			dy := (float64(y) - centerY) / maxRadius

			// Calculate radius from center
			radius := math.Sqrt(dx*dx + dy*dy)

			// Apply defisheye correction using arctangent function
			// This creates a more natural lens correction
			scale := math.Atan(radius*strength) / (radius * strength)
			if radius == 0 {
				scale = 1.0
			}

			// Scale the coordinates
			srcX := int(centerX + dx*maxRadius*scale)
			srcY := int(centerY + dy*maxRadius*scale)

			// Get color from source coordinates with bounds checking
			if srcX >= 0 && srcX < width && srcY >= 0 && srcY < height {
				c := img.NRGBA64At(srcX+bounds.Min.X, srcY+bounds.Min.Y)
				result.Set(x+bounds.Min.X, y+bounds.Min.Y, c)
			} else {
				// Use edge color for out-of-bounds pixels
				edgeX := math.Clamp(srcX, 0, width-1)
				edgeY := math.Clamp(srcY, 0, height-1)
				c := img.NRGBA64At(edgeX+bounds.Min.X, edgeY+bounds.Min.Y)
				result.Set(x+bounds.Min.X, y+bounds.Min.Y, c)
			}
		}
	}

	return result, nil
}

// @Name: fisheye
// @Desc: Applies a fisheye lens distortion effect to the image
// @Param:      img      - -   	-   The image to distort
// @Param:      strength - 0..2  1   The strength of the fisheye effect
// @Returns:    result   - -   	-   The distorted image
func distortFisheye(img *image.NRGBA64, strength float64) (*image.NRGBA64, error) {
	bounds := img.Bounds()
	width := bounds.Dx()
	height := bounds.Dy()
	result := image.NewNRGBA64(bounds)

	centerX := float64(width) / 2
	centerY := float64(height) / 2
	maxRadius := math.Min(centerX, centerY)

	for y := range height {
		for x := range width {
			// Calculate normalized coordinates from center (-1 to 1)
			dx := (float64(x) - centerX) / maxRadius
			dy := (float64(y) - centerY) / maxRadius

			// Calculate radius from center
			radius := math.Sqrt(dx*dx + dy*dy)

			// Apply fisheye distortion using tangent function
			// This creates the barrel distortion characteristic of fisheye lenses
			scale := math.Tan(radius*strength) / (radius * strength)
			if radius == 0 {
				scale = 1.0
			}

			// Scale the coordinates
			srcX := int(centerX + dx*maxRadius*scale)
			srcY := int(centerY + dy*maxRadius*scale)

			// Get color from source coordinates with bounds checking
			if srcX >= 0 && srcX < width && srcY >= 0 && srcY < height {
				c := img.NRGBA64At(srcX+bounds.Min.X, srcY+bounds.Min.Y)
				result.Set(x+bounds.Min.X, y+bounds.Min.Y, c)
			} else {
				// Use edge color for out-of-bounds pixels
				edgeX := math.Clamp(srcX, 0, width-1)
				edgeY := math.Clamp(srcY, 0, height-1)
				c := img.NRGBA64At(edgeX+bounds.Min.X, edgeY+bounds.Min.Y)
				result.Set(x+bounds.Min.X, y+bounds.Min.Y, c)
			}
		}
	}

	return result, nil
}
