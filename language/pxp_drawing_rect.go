package language

import (
	"image"
	"image/color"

	"github.com/toxyl/math"
)

// @Name: draw-rect
// @Desc: Draws a rectangle at position (x,y) with the given width and height.
// @Param:      img       - - -   The image to fill
// @Param:      p         - - -   The center of the rectangle (relative)
// @Param:      w         - - -   The width of the rectangle (relative)
// @Param:      h         - - -   The height of the rectangle (relative)
// @Param:      thickness - - -   The thickness of the rectangle border (absolute)
// @Param:      cBorder	  - - -   The rectangle border color
// @Returns:    result    - - -	  The resulting image
func drawRect(img *image.NRGBA64, p Point, w, h, thickness float64, cBorder color.RGBA64) (*image.NRGBA64, error) {
	bounds := img.Bounds()
	imgW, imgH := bounds.Max.X, bounds.Max.Y
	return drawRectPx(img, P(p.X*float64(imgW), p.Y*float64(imgH)), int(w*float64(imgW)), int(h*float64(imgH)), thickness, cBorder)
}

// @Name: draw-square
// @Desc: Draws a square at position (x,y) with the given size.
// @Param:      img       - - -   The image to fill
// @Param:      p         - - -   The center of the square (relative)
// @Param:      size      - - -   The size of the square (relative)
// @Param:      thickness - - -   The thickness of the square border (absolute)
// @Param:      cBorder	  - - -   The square color
// @Returns:    result    - - -	  The resulting image
func drawSquare(img *image.NRGBA64, p Point, size float64, thickness float64, cBorder color.RGBA64) (*image.NRGBA64, error) {
	bounds := img.Bounds()
	imgW, imgH := bounds.Max.X, bounds.Max.Y
	x1, y1, w1, h1 := int(p.X*float64(imgW)), int(p.Y*float64(imgH)), int(size*float64(imgW)), int(size*float64(imgH))
	w1 = min(w1, h1)
	return drawRectPx(img, P(float64(x1), float64(y1)), w1, w1, thickness, cBorder)
}

// @Name: draw-rect-px
// @Desc: Draws a rectangle at position (x,y) with the given width and height.
// @Param:      img       - - -   The image to fill
// @Param:      p         - - -   The center of the rectangle
// @Param:      w         - - -   The width of the rectangle
// @Param:      h         - - -   The height of the rectangle
// @Param:      thickness - - -   The thickness of the rectangle border
// @Param:      cBorder	  - - -   The rectangle color
// @Returns:    result    - - -	  The resulting image
func drawRectPx(img *image.NRGBA64, p Point, w, h int, thickness float64, cBorder color.RGBA64) (*image.NRGBA64, error) {
	bounds := img.Bounds()
	result := IClone(img)

	// Draw rectangle border with proper thickness and anti-aliasing
	// Calculate bounding box for efficiency
	halfW := float64(w) / 2.0
	halfH := float64(h) / 2.0
	minX := int(p.X - halfW - thickness - 1)
	maxX := int(p.X + halfW + thickness + 1)
	minY := int(p.Y - halfH - thickness - 1)
	maxY := int(p.Y + halfH + thickness + 1)

	for py := minY; py <= maxY; py++ {
		for px := minX; px <= maxX; px++ {
			// Check bounds
			if px >= bounds.Min.X && px < bounds.Max.X && py >= bounds.Min.Y && py < bounds.Max.Y {
				// Calculate distance to rectangle boundary
				var dist float64

				// Distance to left edge
				distLeft := float64(px) - (p.X - halfW)
				// Distance to right edge
				distRight := (p.X + halfW) - float64(px)
				// Distance to top edge
				distTop := float64(py) - (p.Y - halfH)
				// Distance to bottom edge
				distBottom := (p.Y + halfH) - float64(py)

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
				coverage := getRectPixelCoverage(dist, thickness)

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
// @Param:      img       - - -   The image to fill
// @Param:      p         - - -   The center of the square
// @Param:      size      - - -   The size of the square
// @Param:      thickness - - -   The thickness of the square border
// @Param:      cBorder	  - - -   The square color
// @Returns:    result    - - -	  The resulting image
func drawSquarePx(img *image.NRGBA64, p Point, size int, thickness float64, cBorder color.RGBA64) (*image.NRGBA64, error) {
	return drawRectPx(img, p, size, size, thickness, cBorder)
}
