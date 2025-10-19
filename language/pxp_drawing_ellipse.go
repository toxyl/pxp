package language

import (
	"image"
	"image/color"

	"github.com/toxyl/math"
)

// @Name: draw-circle
// @Desc: Draws a circle centered at (x,y) with the given radius.
// @Param:      img       - - -   The image to fill
// @Param:      p         - - -   The center of the circle (relative)
// @Param:      radius    - - -   The radius of the circle (relative)
// @Param:      thickness - - -   The thickness of the circle border (absolute)
// @Param:      cBorder	  - - -   The circle color
// @Returns:    result    - - -	  The resulting image
func drawCircle(img *image.NRGBA64, p Point, radius float64, thickness float64, cBorder color.RGBA64) (*image.NRGBA64, error) {
	bounds := img.Bounds()
	imgW, imgH := bounds.Max.X, bounds.Max.Y
	return drawCirclePx(img, P(p.X*float64(imgW), p.Y*float64(imgH)), int(radius*min(float64(imgW), float64(imgH))), thickness, cBorder)
}

// @Name: draw-circle-px
// @Desc: Draws a circle centered at (x,y) with the given radius.
// @Param:      img       - - -   The image to fill
// @Param:      p         - - -   The center of the circle
// @Param:      radius    - - -   The radius of the circle
// @Param:      thickness - - -   The thickness of the circle border
// @Param:      cBorder	  - - -   The circle color
// @Returns:    result    - - -	  The resulting image
func drawCirclePx(img *image.NRGBA64, p Point, radius int, thickness float64, cBorder color.RGBA64) (*image.NRGBA64, error) {
	return drawEllipsePx(img, p, radius, radius, thickness, cBorder)
}

// @Name: draw-ellipse
// @Desc: Draws an ellipse centered at (x,y) with the given radii.
// @Param:      img       - - -   The image to fill
// @Param:      p         - - -   The center of the ellipse (relative)
// @Param:      rx        - - -   The horizontal radius (relative)
// @Param:      ry        - - -   The vertical radius (relative)
// @Param:      thickness - - -   The thickness of the ellipse border (absolute)
// @Param:      cBorder	  - - -   The ellipse color
// @Returns:    result    - - -	  The resulting image
func drawEllipse(img *image.NRGBA64, p Point, rx, ry float64, thickness float64, cBorder color.RGBA64) (*image.NRGBA64, error) {
	bounds := img.Bounds()
	imgW, imgH := bounds.Max.X, bounds.Max.Y
	return drawEllipsePx(img, P(p.X*float64(imgW), p.Y*float64(imgH)), int(rx*float64(imgW)), int(ry*float64(imgH)), thickness, cBorder)
}

// @Name: draw-ellipse-px
// @Desc: Draws an ellipse centered at (x,y) with the given radii.
// @Param:      img       - - -   The image to fill
// @Param:      p         - - -   The center of the ellipse
// @Param:      rx        - - -   The horizontal radius
// @Param:      ry        - - -   The vertical radius
// @Param:      thickness - - -   The thickness of the ellipse border
// @Param:      cBorder	  - - -   The ellipse color
// @Returns:    result    - - -	  The resulting image
func drawEllipsePx(img *image.NRGBA64, p Point, rx, ry int, thickness float64, cBorder color.RGBA64) (*image.NRGBA64, error) {
	bounds := img.Bounds()
	result := IClone(img)

	// Calculate bounding box for efficiency
	minX := int(p.X - float64(rx) - thickness - 1)
	maxX := int(p.X + float64(rx) + thickness + 1)
	minY := int(p.Y - float64(ry) - thickness - 1)
	maxY := int(p.Y + float64(ry) + thickness + 1)

	// Draw ellipse border with thickness
	for py := minY; py <= maxY; py++ {
		for px := minX; px <= maxX; px++ {
			// Check bounds
			if px >= bounds.Min.X && px < bounds.Max.X && py >= bounds.Min.Y && py < bounds.Max.Y {
				// Calculate distance from ellipse center
				dx := float64(px - int(p.X))
				dy := float64(py - int(p.Y))

				// Calculate distance from ellipse boundary using absolute pixel thickness
				// Use parametric approach to find closest point on ellipse boundary
				// Then check if current point is within thickness pixels of that boundary

				// Normalize coordinates
				nx := dx / float64(rx)
				ny := dy / float64(ry)
				ellipseDist := nx*nx + ny*ny

				// Calculate distance from ellipse boundary
				var distToBoundary float64

				if ellipseDist < 1.0 {
					// Inside ellipse - distance is negative (going inward from boundary)
					distToBoundary = -(1.0 - ellipseDist) * float64(math.Min(rx, ry))
				} else {
					// Outside ellipse - find closest point on boundary
					norm := math.Sqrt(ellipseDist)
					closestX := nx / norm
					closestY := ny / norm

					closestPx := closestX * float64(rx)
					closestPy := closestY * float64(ry)

					distToBoundary = math.Sqrt((dx-closestPx)*(dx-closestPx) + (dy-closestPy)*(dy-closestPy))
				}

				// Get pixel coverage for anti-aliasing
				coverage := getCirclePixelCoverage(distToBoundary, thickness)

				if coverage > 0 {
					// Get existing pixel color
					existing := result.NRGBA64At(px, py)

					// Apply coverage to alpha
					alpha := uint32(float64(cBorder.A) * coverage)

					// Blend colors using normal blend mode
					r, g, b, a := blendWithAlpha(
						uint32(existing.R), uint32(existing.G), uint32(existing.B), uint32(existing.A),
						uint32(cBorder.R), uint32(cBorder.G), uint32(cBorder.B), alpha,
						func(r1, g1, b1, r2, g2, b2 uint32) (uint32, uint32, uint32) {
							return r2, g2, b2
						},
					)

					// Set blended color
					result.Set(px, py, color.NRGBA64{
						R: uint16(r),
						G: uint16(g),
						B: uint16(b),
						A: uint16(a),
					})
				}
			}
		}
	}

	return result, nil
}
