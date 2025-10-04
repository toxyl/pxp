package language

import (
	"fmt"
	"image"
	"image/color"

	"github.com/toxyl/math"
)

var (
	gaussianBlurKernels = map[float64][]float64{}
)

func makeGaussianBlurKernel(radius float64) (int, []float64) {
	if k, ok := gaussianBlurKernels[radius]; ok {
		return int(radius*2 + 1), k
	}
	kernelSize := int(radius*2 + 1)
	kernel := make([]float64, kernelSize*kernelSize)
	sum := 0.0

	// Create Gaussian kernel
	for y := range kernelSize {
		for x := range kernelSize {
			dx := float64(x) - radius
			dy := float64(y) - radius
			val := math.Exp(-(dx*dx + dy*dy) / (2 * radius * radius))
			kernel[y*kernelSize+x] = val
			sum += val
		}
	}

	// Normalize kernel
	for i := range kernel {
		kernel[i] /= sum
	}
	gaussianBlurKernels[radius] = kernel
	return kernelSize, kernel
}

// @Name: blur-gaussian
// @Desc: Applies a Gaussian blur to the image
// @Param:      img     - -   	-   The image to blur
// @Param:      radius  - 1..10 1   The blur radius (higher values create more blur)
// @Returns:    result  - -   	-   The blurred image
func blurGaussian(img *image.NRGBA64, radius float64) (*image.NRGBA64, error) {
	bounds := img.Bounds()
	result := image.NewNRGBA64(bounds)

	// Calculate kernel size and create kernel
	kernelSize, kernel := makeGaussianBlurKernel(radius)

	// Apply blur
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			var r, g, b float64
			var a uint16

			// Apply kernel
			for ky := range kernelSize {
				for kx := range kernelSize {
					px := x + kx - int(radius)
					py := y + ky - int(radius)

					// Handle boundaries
					if px < bounds.Min.X || px >= bounds.Max.X || py < bounds.Min.Y || py >= bounds.Max.Y {
						continue
					}

					c := img.NRGBA64At(px, py)
					weight := kernel[ky*kernelSize+kx]
					r += float64(c.R) * weight
					g += float64(c.G) * weight
					b += float64(c.B) * weight
				}
			}

			// Get original alpha
			a = img.NRGBA64At(x, y).A

			result.Set(x, y, color.NRGBA64{
				R: uint16(math.Clamp(r, 0, 65535)),
				G: uint16(math.Clamp(g, 0, 65535)),
				B: uint16(math.Clamp(b, 0, 65535)),
				A: a,
			})
		}
	}

	return result, nil
}

// @Name: blur-box
// @Desc: Applies a box blur to an image
// @Param:      img     - -   	-   The image to blur
// @Param:      radius  - 1..10 1   The blur radius (size of the box kernel)
// @Returns:    result  - -   	-   The blurred image
func blurBox(img *image.NRGBA64, radius int) (*image.NRGBA64, error) {
	bounds := img.Bounds()
	blurred := image.NewNRGBA64(bounds)
	size := radius*2 + 1
	kernelSize := size * size

	// Apply convolution using a box kernel
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			var r, g, b, count float64
			var a uint16 = 65535

			// Apply kernel
			for ky := -radius; ky <= radius; ky++ {
				for kx := -radius; kx <= radius; kx++ {
					ix := x + kx
					iy := y + ky

					if ix >= bounds.Min.X && ix < bounds.Max.X && iy >= bounds.Min.Y && iy < bounds.Max.Y {
						c := img.NRGBA64At(ix, iy)
						r += float64(c.R)
						g += float64(c.G)
						b += float64(c.B)
						a = c.A // Preserve original alpha from the center pixel for simplicity, or average alpha if needed
						count++ // Use actual count in case of edge pixels
					}
				}
			}

			// Use kernelSize for normalization if not handling edges specifically,
			// or use 'count' if handling edges by summing available pixels.
			// Using kernelSize assumes extension of edge pixels or wrapping.
			// Using 'count' averages only the pixels within bounds. For simplicity & common box blur, let's use kernelSize.
			normFactor := float64(kernelSize)
			if count > 0 { // Ensure we don't divide by zero if the kernel somehow falls outside the image entirely
				normFactor = count // Average based on actual pixels summed for edge handling
			}

			blurred.Set(x, y, color.NRGBA64{
				R: uint16(math.Clamp(r/normFactor, 0, 65535)),
				G: uint16(math.Clamp(g/normFactor, 0, 65535)),
				B: uint16(math.Clamp(b/normFactor, 0, 65535)),
				A: a, // Use the preserved alpha
			})
		}
	}
	return blurred, nil
}

// @Name: blur-motion
// @Desc: Applies a motion blur to an image along a specified angle.
// @Param:      img     - -       -   The image to blur
// @Param:      length  - 1..100  5   The length of the motion blur (in pixels)
// @Param:      angle   - 0..360  0   The angle of the motion blur (in degrees)
// @Returns:    result  - -       -   The blurred image
func blurMotion(img *image.NRGBA64, length int, angle float64) (*image.NRGBA64, error) {
	if angle < 0 || angle > 360 {
		// Allow 360 as it's equivalent to 0
		if angle != 360 {
			return nil, fmt.Errorf("motion blur angle must be between 0 and 360 degrees")
		}
		angle = 0 // Normalize 360 to 0
	}

	bounds := img.Bounds()
	blurred := image.NewNRGBA64(bounds)
	radAngle := angle * math.Pi / 180.0 // Convert angle to radians
	dx := math.Cos(radAngle)
	dy := math.Sin(radAngle)

	// Determine the number of steps based on length. More steps give smoother blur.
	// We can use 'length' directly as the number of samples along the line.
	numSamples := length

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			var r, g, b float64
			var count int = 0
			var finalA uint16 = img.NRGBA64At(x, y).A // Start with original alpha

			// Sample along the line defined by the angle and length
			for i := 0; i < numSamples; i++ {
				// Calculate sample position relative to the current pixel
				// We sample from -length/2 to +length/2 along the line
				offset := (float64(i) - float64(numSamples-1)/2.0)
				sx := float64(x) + offset*dx
				sy := float64(y) + offset*dy

				// Use bilinear interpolation for smoother results
				c := bilinearInterpolate(img, sx, sy)

				// Check if the interpolated color is valid (within bounds)
				// The interpolation function should handle boundary checks implicitly
				// If bilinearInterpolate returns a zero color.Color, it might indicate out of bounds
				// or we could add a check here based on sx, sy if needed.
				// For simplicity, let's assume bilinearInterpolate handles boundaries.

				nrgba, ok := c.(color.NRGBA64)
				if ok {
					r += float64(nrgba.R)
					g += float64(nrgba.G)
					b += float64(nrgba.B)
					// Alpha could be averaged, but let's keep the original center pixel's alpha
					count++
				}
			}

			if count > 0 {
				blurred.Set(x, y, color.NRGBA64{
					R: uint16(math.Clamp(r/float64(count), 0, 65535)),
					G: uint16(math.Clamp(g/float64(count), 0, 65535)),
					B: uint16(math.Clamp(b/float64(count), 0, 65535)),
					A: finalA,
				})
			} else {
				// If no samples were valid (e.g., extremely small image or large length), keep original pixel
				blurred.Set(x, y, img.NRGBA64At(x, y))
			}
		}
	}
	return blurred, nil
}

// @Name: blur-zoom
// @Desc: Applies a zoom blur effect to an image.
// @Param:      img     	- -       	-   	The image to blur
// @Param:      strength	- 0.0..1.0 	0.25 	The strength of the blur effect (higher means more blur)
// @Param:      centerX		- 0.0..1.0	0.5		X coordinate of the blur center (default: image center)
// @Param:      centerY		- 0.0..1.0	0.5		Y coordinate of the blur center (default: image center)
// @Returns:    result  	- -       	-   	The blurred image
func blurZoom(img *image.NRGBA64, strength float64, centerX float64, centerY float64) (*image.NRGBA64, error) {
	strength *= 0.085

	bounds := img.Bounds()
	blurred := image.NewNRGBA64(bounds)
	centerX *= float64(bounds.Dx())
	centerY *= float64(bounds.Dy())

	// Number of samples along the radial line. More samples = smoother but slower.
	numSamples := 10 // Adjust for quality/performance trade-off

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			var r, g, b float64
			var count int = 0
			finalA := img.NRGBA64At(x, y).A // Keep original alpha

			dx := float64(x) - centerX
			dy := float64(y) - centerY

			// Sample along the line from the center to the current pixel
			for i := 0; i < numSamples; i++ {
				// Scale factor determines how far along the line we sample
				// We sample points closer to the current pixel based on strength
				scale := 1.0 + (float64(i)/float64(numSamples-1)-0.5)*strength

				sx := centerX + dx*scale
				sy := centerY + dy*scale

				// Use bilinear interpolation for smoother results
				c := bilinearInterpolate(img, sx, sy)
				nrgba, ok := c.(color.NRGBA64)
				if ok {
					r += float64(nrgba.R)
					g += float64(nrgba.G)
					b += float64(nrgba.B)
					// Alpha could be averaged here too if desired
					count++
				}
			}

			if count > 0 {
				blurred.Set(x, y, color.NRGBA64{
					R: uint16(math.Clamp(r/float64(count), 0, 65535)),
					G: uint16(math.Clamp(g/float64(count), 0, 65535)),
					B: uint16(math.Clamp(b/float64(count), 0, 65535)),
					A: finalA,
				})
			} else {
				// Keep original if no samples were valid
				blurred.Set(x, y, img.NRGBA64At(x, y))
			}
		}
	}
	return blurred, nil
}
