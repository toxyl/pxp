package language

import (
	"image"
	"image/color"

	"github.com/toxyl/math"
)

// @Name: draw-line
// @Desc: Draws a line from P(x1|y1) to P(x2|y2) with the given thickness and color.
// @Param:      img       - - -   The image to fill
// @Param:      p1        - - -   The start position (relative)
// @Param:      p2        - - -   The end position (relative)
// @Param:      thickness - - -   The thickness of the line (absolute)
// @Param:      cBorder   - - -   The line color
// @Returns:    result    - - -	  The resulting image
func drawLine(img *image.NRGBA64, p1, p2 Point, thickness float64, cBorder color.RGBA64) (*image.NRGBA64, error) {
	bounds := img.Bounds()
	w, h := bounds.Max.X, bounds.Max.Y
	return drawLinePx(img, P(p1.X*float64(w), p1.Y*float64(h)), P(p2.X*float64(w), p2.Y*float64(h)), thickness, cBorder)
}

// @Name: draw-line-v
// @Desc: Draws a vertical line from P(x|y1) to P(x|y2) with the given thickness and color.
// @Param:      img       - - -   The image to fill
// @Param:      x         - - -   The position on the x-axis (relative)
// @Param:      y1        - - -   The start position on the y-axis (relative)
// @Param:      y2        - - -   The end position on the y-axis (relative)
// @Param:      thickness - - -   The thickness of the line (absolute)
// @Param:      cBorder	  - - -   The line color
// @Returns:    result    - - -	  The resulting image
func drawLineVertical(img *image.NRGBA64, x, y1, y2 float64, thickness float64, cBorder color.RGBA64) (*image.NRGBA64, error) {
	return drawLine(img, P(x, y1), P(x, y2), thickness, cBorder)
}

// @Name: draw-line-h
// @Desc: Draws a line from P(x1|y) to P(x2|y) with the given thickness and color.
// @Param:      img       - - -   The image to fill
// @Param:      y         - - -   The position on the y-axis (relative)
// @Param:      x1        - - -   The start position on the x-axis (relative)
// @Param:      x2        - - -   The end position on the x-axis (relative)
// @Param:      thickness - - -   The thickness of the line (absolute)
// @Param:      cBorder	  - - -   The line color
// @Returns:    result    - - -	  The resulting image
func drawLineHorizontal(img *image.NRGBA64, y, x1, x2 float64, thickness float64, cBorder color.RGBA64) (*image.NRGBA64, error) {
	return drawLine(img, P(x1, y), P(x2, y), thickness, cBorder)
}

// @Name: draw-line-px
// @Desc: Draws a line from P(x1|y1) to P(x2|y2) with the given thickness and color.
// @Param:      img       - - -   The image to fill
// @Param:      p1        - - -   The start position
// @Param:      p2        - - -   The end position
// @Param:      thickness - - -   The thickness of the line
// @Param:      cBorder	  - - -   The line color
// @Returns:    result    - - -	  The resulting image
func drawLinePx(img *image.NRGBA64, p1, p2 Point, thickness float64, cBorder color.RGBA64) (*image.NRGBA64, error) {
	bounds := img.Bounds()
	result := IClone(img)

	// Calculate line direction and perpendicular
	dx := float64(p2.X - p1.X)
	dy := float64(p2.Y - p1.Y)
	length := math.Sqrt(dx*dx + dy*dy)

	if length == 0 {
		// Single point case - draw a circle
		radius := thickness / 2.0
		minX := int(p1.X - radius - 1)
		maxX := int(p1.X + radius + 1)
		minY := int(p1.Y - radius - 1)
		maxY := int(p1.Y + radius + 1)

		for py := minY; py <= maxY; py++ {
			for px := minX; px <= maxX; px++ {
				if px >= bounds.Min.X && px < bounds.Max.X && py >= bounds.Min.Y && py < bounds.Max.Y {
					dist := math.Sqrt(float64((px-int(p1.X))*(px-int(p1.X)) + (py-int(p1.Y))*(py-int(p1.Y))))
					coverage := getCirclePixelCoverage(dist, radius)

					if coverage > 0 {
						existing := result.NRGBA64At(px, py)
						alpha := uint32(float64(cBorder.A) * coverage)
						r, g, b, a := blendWithAlpha(
							uint32(existing.R), uint32(existing.G), uint32(existing.B), uint32(existing.A),
							uint32(cBorder.R), uint32(cBorder.G), uint32(cBorder.B), alpha,
							func(r1, g1, b1, r2, g2, b2 uint32) (uint32, uint32, uint32) {
								return r2, g2, b2
							},
						)
						result.Set(px, py, color.NRGBA64{
							R: uint16(r), G: uint16(g), B: uint16(b), A: uint16(a),
						})
					}
				}
			}
		}
		return result, nil
	}

	// Normalize direction vector
	dx /= length
	dy /= length

	// Half thickness
	halfThickness := thickness / 2.0

	// Calculate bounding box for efficiency
	minX := int(math.Min(float64(p1.X), float64(p2.X)) - halfThickness - 1)
	maxX := int(math.Max(float64(p1.X), float64(p2.X)) + halfThickness + 1)
	minY := int(math.Min(float64(p1.Y), float64(p2.Y)) - halfThickness - 1)
	maxY := int(math.Max(float64(p1.Y), float64(p2.Y)) + halfThickness + 1)

	// Draw line with anti-aliasing
	for py := minY; py <= maxY; py++ {
		for px := minX; px <= maxX; px++ {
			// Check bounds
			if px >= bounds.Min.X && px < bounds.Max.X && py >= bounds.Min.Y && py < bounds.Max.Y {
				// Calculate distance from point to line
				// Line equation: (p2.Y-p1.Y)x - (p2.X-p1.X)y + (p2.X-p1.X)p1.Y - (y2-p1.Y)p1.X = 0
				// Distance = |ax + by + c| / sqrt(a² + b²)
				a := float64(p2.Y - p1.Y)
				b := float64(p1.X - p2.X)
				c := float64(p2.X*p1.Y - p1.X*p2.Y)
				dist := math.Abs(a*float64(px)+b*float64(py)+c) / length

				// Get pixel coverage for anti-aliasing
				coverage := getRectPixelCoverage(dist, halfThickness)

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

// @Name: draw-line-v-px
// @Desc: Draws a vertical line from P(x|y1) to P(x|y2) with the given thickness and color.
// @Param:      img       - - -   The image to fill
// @Param:      x         - - -   The position on the x-axis
// @Param:      y1        - - -   The start position on the y-axis
// @Param:      y2        - - -   The end position on the y-axis
// @Param:      thickness - - -   The thickness of the line
// @Param:      cBorder	  - - -   The line color
// @Returns:    result    - - -	  The resulting image
func drawLineVerticalPx(img *image.NRGBA64, x, y1, y2, thickness float64, cBorder color.RGBA64) (*image.NRGBA64, error) {
	return drawLinePx(img, P(x, y1), P(x, y2), thickness, cBorder)
}

// @Name: draw-line-h-px
// @Desc: Draws a line from P(x1|y) to P(x2|y) with the given thickness and color.
// @Param:      img       - - -   The image to fill
// @Param:      y         - - -   The position on the y-axis
// @Param:      x1        - - -   The start position on the x-axis
// @Param:      x2        - - -   The end position on the x-axis
// @Param:      thickness - - -   The thickness of the line
// @Param:      cBorder	  - - -   The line color
// @Returns:    result    - - -	  The resulting image
func drawLineHorizontalPx(img *image.NRGBA64, y, x1, x2, thickness float64, cBorder color.RGBA64) (*image.NRGBA64, error) {
	return drawLinePx(img, P(x1, y), P(x2, y), thickness, cBorder)
}
