package language

import (
	"image"
	"image/color"

	"github.com/toxyl/math"
)

// @Name: draw-circle
// @Desc: Draws a circle.
// @Param:      img       - - -   The image to draw to
// @Param:      c         - - -   The circle to draw
// @Param:      style     - - -   The thickness and color of the line
// @Returns:    result    - - -	  The resulting image
func drawCircle(img *image.NRGBA64, c Ellipse, style LineStyle) (*image.NRGBA64, error) {
	bounds := img.Bounds()
	return drawCirclePx(img, *c.Denorm(float64(bounds.Max.X), float64(bounds.Max.Y)), style)
}

// @Name: draw-circle-px
// @Desc: Draws a circle.
// @Param:      img       - - -   The image to draw to
// @Param:      c         - - -   The circle to draw
// @Param:      style     - - -   The thickness and color of the line
// @Returns:    result    - - -	  The resulting image
func drawCirclePx(img *image.NRGBA64, c Ellipse, style LineStyle) (*image.NRGBA64, error) {
	if c.Radius.X != c.Radius.Y {
		r := math.Min(c.Radius.X, c.Radius.Y)
		c.Radius = &Point{
			X: r,
			Y: r,
		}
	}
	return drawEllipsePx(img, c, style)
}

// @Name: draw-ellipse
// @Desc: Draws an ellipse.
// @Param:      img       - - -   The image to draw to
// @Param:      e         - - -   The ellipse to draw
// @Param:      style     - - -   The thickness and color of the line
// @Returns:    result    - - -	  The resulting image
func drawEllipse(img *image.NRGBA64, e Ellipse, style LineStyle) (*image.NRGBA64, error) {
	bounds := img.Bounds()
	return drawEllipsePx(img, *e.Denorm(float64(bounds.Max.X), float64(bounds.Max.Y)), style)
}

// @Name: draw-ellipse-px
// @Desc: Draws an ellipse.
// @Param:      img       - - -   The image to draw to
// @Param:      e         - - -   The ellipse to draw
// @Param:      style     - - -   The thickness and color of the line
// @Returns:    result    - - -	  The resulting image
func drawEllipsePx(img *image.NRGBA64, e Ellipse, style LineStyle) (*image.NRGBA64, error) {
	bounds := img.Bounds()
	result := IClone(img)

	centerX := e.Center.X
	centerY := e.Center.Y

	// Calculate bounding box for efficiency
	minX := int(centerX - e.Radius.X - style.Thickness - 1)
	maxX := int(centerX + e.Radius.X + style.Thickness + 1)
	minY := int(centerY - e.Radius.Y - style.Thickness - 1)
	maxY := int(centerY + e.Radius.Y + style.Thickness + 1)

	// Draw ellipse border with thickness
	for py := minY; py <= maxY; py++ {
		for px := minX; px <= maxX; px++ {
			// Check bounds
			if px >= bounds.Min.X && px < bounds.Max.X && py >= bounds.Min.Y && py < bounds.Max.Y {
				// Calculate distance from ellipse center
				dx := float64(px - int(centerX))
				dy := float64(py - int(centerY))

				// Calculate distance from ellipse boundary using absolute pixel thickness
				// Use parametric approach to find closest point on ellipse boundary
				// Then check if current point is within thickness pixels of that boundary

				// Normalize coordinates
				nx := dx / e.Radius.X
				ny := dy / e.Radius.Y
				ellipseDist := nx*nx + ny*ny

				// Calculate distance from ellipse boundary
				var distToBoundary float64

				if ellipseDist < 1.0 {
					// Inside ellipse - distance is negative (going inward from boundary)
					distToBoundary = -(1.0 - ellipseDist) * float64(math.Min(e.Radius.X, e.Radius.Y))
				} else {
					// Outside ellipse - find closest point on boundary
					norm := math.Sqrt(ellipseDist)
					closestX := nx / norm
					closestY := ny / norm

					closestPx := closestX * e.Radius.X
					closestPy := closestY * e.Radius.Y

					distToBoundary = math.Sqrt((dx-closestPx)*(dx-closestPx) + (dy-closestPy)*(dy-closestPy))
				}

				// Get pixel coverage for anti-aliasing
				coverage := getCirclePixelCoverage(distToBoundary, style.Thickness)

				if coverage > 0 {
					// Get existing pixel color
					existing := result.NRGBA64At(px, py)

					// Apply coverage to alpha
					alpha := uint32(float64(style.Color.A) * coverage)

					// Blend colors using normal blend mode
					r, g, b, a := blendWithAlpha(
						uint32(existing.R), uint32(existing.G), uint32(existing.B), uint32(existing.A),
						uint32(style.Color.R), uint32(style.Color.G), uint32(style.Color.B), alpha,
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
