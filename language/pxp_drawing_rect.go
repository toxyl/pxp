package language

import (
	"image"
	"image/color"

	"github.com/toxyl/math"
)

// @Name: draw-rect
// @Desc: Draws a rectangle at position (x,y) with the given width and height.
// @Param:      img       - - -   The image to draw to
// @Param:      r         - - -   The rectangle to draw (relative)
// @Param:      style     - - -   The thickness and color of the line
// @Returns:    result    - - -	  The resulting image
func drawRect(img *image.NRGBA64, r Rect, style LineStyle) (*image.NRGBA64, error) {
	bounds := img.Bounds()
	return drawRectPx(img, *r.Denorm(float64(bounds.Max.X), float64(bounds.Max.Y)), style)
}

// @Name: draw-square
// @Desc: Draws a square at position (x,y) with the given size.
// @Param:      img       - - -   The image to draw to
// @Param:      s         - - -   The square to draw (relative)
// @Param:      style     - - -   The thickness and color of the line
// @Returns:    result    - - -	  The resulting image
func drawSquare(img *image.NRGBA64, s Rect, style LineStyle) (*image.NRGBA64, error) {
	bounds := img.Bounds()
	return drawSquarePx(img, *s.Denorm(float64(bounds.Max.X), float64(bounds.Max.Y)), style)
}

// @Name: draw-rect-px
// @Desc: Draws a rectangle at position (x,y) with the given width and height.
// @Param:      img       - - -   The image to draw to
// @Param:      r         - - -   The rectangle to draw (absolute)
// @Param:      style     - - -   The thickness and color of the line
// @Returns:    result    - - -	  The resulting image
func drawRectPx(img *image.NRGBA64, r Rect, style LineStyle) (*image.NRGBA64, error) {
	bounds := img.Bounds()
	result := IClone(img)
	r.P2.X -= 1
	r.P2.Y -= 1

	// Draw rectangle border with proper thickness and anti-aliasing
	// Calculate bounding box for efficiency
	minX := int(r.X1() - style.Thickness - 1)
	maxX := int(r.X2() + style.Thickness + 1)
	minY := int(r.Y1() - style.Thickness - 1)
	maxY := int(r.Y2() + style.Thickness + 1)

	for py := minY; py <= maxY; py++ {
		for px := minX; px <= maxX; px++ {
			// Check bounds
			if px >= bounds.Min.X && px < bounds.Max.X && py >= bounds.Min.Y && py < bounds.Max.Y {
				// Calculate distance to rectangle boundary
				var dist float64

				// Distance to left edge
				distLeft := float64(px) - r.X1()
				// Distance to right edge
				distRight := r.X2() - float64(px)
				// Distance to top edge
				distTop := float64(py) - r.Y1()
				// Distance to bottom edge
				distBottom := r.Y2() - float64(py)

				// Check if point is inside rectangle
				if distLeft < 0 && distRight < 0 && distTop < 0 && distBottom < 0 {
					// Inside rectangle - don't draw (borders only)
					continue
				}

				// Find minimum distance to any edge (only for outside points)
				dist = math.Min(distLeft, distRight)
				dist = math.Min(dist, distTop)
				dist = math.Min(dist, distBottom)

				// Get pixel coverage for anti-aliasing
				coverage := getRectPixelCoverage(dist, style.Thickness)

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
						R: uint16(r), G: uint16(g), B: uint16(b), A: uint16(a),
					})
				}
			}
		}
	}

	return result, nil
}

// @Name: draw-square-px
// @Desc: Draws a square at position (x,y) with the given size.
// @Param:      img       - - -   The image to draw to
// @Param:      s         - - -   The square to draw (absolute)
// @Param:      style     - - -   The thickness and color of the line
// @Returns:    result    - - -	  The resulting image
func drawSquarePx(img *image.NRGBA64, s Rect, style LineStyle) (*image.NRGBA64, error) {
	if s.W() != s.H() {
		m := math.Min(s.W(), s.H())
		s = Rect{
			P1: s.P1,
			P2: P(s.P1.X+m, s.P1.Y+m),
		}
	}
	return drawRectPx(img, s, style)
}
