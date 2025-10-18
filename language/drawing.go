package language

import (
	"image"
	"image/color"

	"github.com/toxyl/math"
)

// Helper function for anti-aliased pixel coverage (for rectangles)
func getRectPixelCoverage(distance float64, thickness float64) float64 {
	// For rectangles: only draw pixels outside the shape
	if distance < 0 {
		// Inside shape - don't draw
		return 0.0
	}
	if distance >= thickness {
		// Outside border - don't draw
		return 0.0
	}
	if distance >= thickness-1.0 {
		// Anti-aliasing zone (last pixel of border)
		return thickness - distance
	}
	// Solid border area
	return 1.0
}

// Helper function for anti-aliased pixel coverage (for circles/ellipses)
func getCirclePixelCoverage(distance float64, thickness float64) float64 {
	// For circles: draw pixels within thickness distance from boundary
	// Use absolute distance to handle both inner and outer edges

	absDist := math.Abs(distance)

	if absDist >= thickness {
		// Outside border range - don't draw
		return 0.0
	}

	if absDist >= thickness-1.0 {
		// Anti-aliasing zone (last pixel of border)
		return thickness - absDist
	}

	// Solid border area
	return 1.0
}

func P(x, y float64) Point {
	return Point{
		X: x,
		Y: y,
	}
}

// @Name: P
// @Desc: Creates the point P(x|y).
// @Param:      x        - - 0   The start position on the x-axis (relative)
// @Param:      y        - - 0   The start position on the y-axis (relative)
// @Returns:    result    - - -	 A new point
func makePoint(x, y float64) (Point, error) {
	return P(x, y), nil
}

// @Name: Px
// @Desc: Returns the x-coordinate of a point.
// @Param:      p        - - -   The point to return the x-coordinate of
// @Returns:    result    - - -	 The x-coordinate of p
func pointX(p Point) (any, error) {
	return p.X, nil
}

// @Name: Py
// @Desc: Returns the y-coordinate of a point.
// @Param:      p        - - -   The point to return the y-coordinate of
// @Returns:    result    - - -	 The y-coordinate of p
func pointY(p Point) (any, error) {
	return p.Y, nil
}

// @Name: draw-line
// @Desc: Draws a line from P(x1|y1) to P(x2|y2) with the given thickness and color.
// @Param:      img       - - -   The image to fill
// @Param:      p1        - - -   The start position (relative)
// @Param:      p2        - - -   The end position (relative)
// @Param:      thickness - - -   The thickness of the line (absolute)
// @Param:      col  	  - - -   The line color
// @Returns:    result    - - -	  The resulting image
func drawLine(img *image.NRGBA64, p1, p2 Point, thickness float64, col color.RGBA64) (*image.NRGBA64, error) {
	bounds := img.Bounds()
	w, h := bounds.Max.X, bounds.Max.Y
	return drawLinePx(img, P(p1.X*float64(w), p1.Y*float64(h)), P(p2.X*float64(w), p2.Y*float64(h)), thickness, col)
}

// @Name: draw-line-v
// @Desc: Draws a vertical line from P(x|y1) to P(x|y2) with the given thickness and color.
// @Param:      img       - - -   The image to fill
// @Param:      x         - - -   The position on the x-axis (relative)
// @Param:      y1        - - -   The start position on the y-axis (relative)
// @Param:      y2        - - -   The end position on the y-axis (relative)
// @Param:      thickness - - -   The thickness of the line (absolute)
// @Param:      col  	  - - -   The line color
// @Returns:    result    - - -	  The resulting image
func drawLineVertical(img *image.NRGBA64, x, y1, y2 float64, thickness float64, col color.RGBA64) (*image.NRGBA64, error) {
	return drawLine(img, P(x, y1), P(x, y2), thickness, col)
}

// @Name: draw-line-h
// @Desc: Draws a line from P(x1|y) to P(x2|y) with the given thickness and color.
// @Param:      img       - - -   The image to fill
// @Param:      y         - - -   The position on the y-axis (relative)
// @Param:      x1        - - -   The start position on the x-axis (relative)
// @Param:      x2        - - -   The end position on the x-axis (relative)
// @Param:      thickness - - -   The thickness of the line (absolute)
// @Param:      col  	  - - -   The line color
// @Returns:    result    - - -	  The resulting image
func drawLineHorizontal(img *image.NRGBA64, y, x1, x2 float64, thickness float64, col color.RGBA64) (*image.NRGBA64, error) {
	return drawLine(img, P(x1, y), P(x2, y), thickness, col)
}

// @Name: draw-line-px
// @Desc: Draws a line from P(x1|y1) to P(x2|y2) with the given thickness and color.
// @Param:      img       - - -   The image to fill
// @Param:      p1        - - -   The start position
// @Param:      p2        - - -   The end position
// @Param:      thickness - - -   The thickness of the line
// @Param:      col  	  - - -   The line color
// @Returns:    result    - - -	  The resulting image
func drawLinePx(img *image.NRGBA64, p1, p2 Point, thickness float64, col color.RGBA64) (*image.NRGBA64, error) {
	bounds := img.Bounds()
	result := image.NewNRGBA64(bounds)

	// Copy original image to result
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			result.Set(x, y, img.NRGBA64At(x, y))
		}
	}

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
						alpha := uint32(float64(col.A) * coverage)
						r, g, b, a := blendWithAlpha(
							uint32(existing.R), uint32(existing.G), uint32(existing.B), uint32(existing.A),
							uint32(col.R), uint32(col.G), uint32(col.B), alpha,
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
					alpha := uint32(float64(col.A) * coverage)

					// Blend colors using normal blend mode
					r, g, b, a := blendWithAlpha(
						uint32(existing.R), uint32(existing.G), uint32(existing.B), uint32(existing.A),
						uint32(col.R), uint32(col.G), uint32(col.B), alpha,
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
// @Param:      col  	  - - -   The line color
// @Returns:    result    - - -	  The resulting image
func drawLineVerticalPx(img *image.NRGBA64, x, y1, y2, thickness float64, col color.RGBA64) (*image.NRGBA64, error) {
	return drawLinePx(img, P(x, y1), P(x, y2), thickness, col)
}

// @Name: draw-line-h-px
// @Desc: Draws a line from P(x1|y) to P(x2|y) with the given thickness and color.
// @Param:      img       - - -   The image to fill
// @Param:      y         - - -   The position on the y-axis
// @Param:      x1        - - -   The start position on the x-axis
// @Param:      x2        - - -   The end position on the x-axis
// @Param:      thickness - - -   The thickness of the line
// @Param:      col  	  - - -   The line color
// @Returns:    result    - - -	  The resulting image
func drawLineHorizontalPx(img *image.NRGBA64, y, x1, x2, thickness float64, col color.RGBA64) (*image.NRGBA64, error) {
	return drawLinePx(img, P(x1, y), P(x2, y), thickness, col)
}

// @Name: draw-rect
// @Desc: Draws a rectangle at position (x,y) with the given width and height.
// @Param:      img       - - -   The image to fill
// @Param:      p         - - -   The center of the rectangle (relative)
// @Param:      w         - - -   The width of the rectangle (relative)
// @Param:      h         - - -   The height of the rectangle (relative)
// @Param:      thickness - - -   The thickness of the rectangle border (absolute)
// @Param:      col  	  - - -   The rectangle border color
// @Returns:    result    - - -	  The resulting image
func drawRect(img *image.NRGBA64, p Point, w, h, thickness float64, col color.RGBA64) (*image.NRGBA64, error) {
	bounds := img.Bounds()
	imgW, imgH := bounds.Max.X, bounds.Max.Y
	return drawRectPx(img, P(p.X*float64(imgW), p.Y*float64(imgH)), int(w*float64(imgW)), int(h*float64(imgH)), thickness, col)
}

// @Name: draw-square
// @Desc: Draws a square at position (x,y) with the given size.
// @Param:      img       - - -   The image to fill
// @Param:      p         - - -   The center of the square (relative)
// @Param:      size      - - -   The size of the square (relative)
// @Param:      thickness - - -   The thickness of the square border (absolute)
// @Param:      col  	  - - -   The square color
// @Returns:    result    - - -	  The resulting image
func drawSquare(img *image.NRGBA64, p Point, size float64, thickness float64, col color.RGBA64) (*image.NRGBA64, error) {
	bounds := img.Bounds()
	imgW, imgH := bounds.Max.X, bounds.Max.Y
	x1, y1, w1, h1 := int(p.X*float64(imgW)), int(p.Y*float64(imgH)), int(size*float64(imgW)), int(size*float64(imgH))
	w1 = min(w1, h1)
	return drawRectPx(img, P(float64(x1), float64(y1)), w1, w1, thickness, col)
}

// @Name: draw-rect-px
// @Desc: Draws a rectangle at position (x,y) with the given width and height.
// @Param:      img       - - -   The image to fill
// @Param:      p         - - -   The center of the rectangle
// @Param:      w         - - -   The width of the rectangle
// @Param:      h         - - -   The height of the rectangle
// @Param:      thickness - - -   The thickness of the rectangle border
// @Param:      col  	  - - -   The rectangle color
// @Returns:    result    - - -	  The resulting image
func drawRectPx(img *image.NRGBA64, p Point, w, h int, thickness float64, col color.RGBA64) (*image.NRGBA64, error) {
	bounds := img.Bounds()
	result := image.NewNRGBA64(bounds)

	// Copy original image to result
	for py := bounds.Min.Y; py < bounds.Max.Y; py++ {
		for px := bounds.Min.X; px < bounds.Max.X; px++ {
			result.Set(px, py, img.NRGBA64At(px, py))
		}
	}

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
					alpha := uint32(float64(col.A) * coverage)

					// Blend colors using normal blend mode
					r, g, b, a := blendWithAlpha(
						uint32(existing.R), uint32(existing.G), uint32(existing.B), uint32(existing.A),
						uint32(col.R), uint32(col.G), uint32(col.B), alpha,
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
// @Param:      col  	  - - -   The square color
// @Returns:    result    - - -	  The resulting image
func drawSquarePx(img *image.NRGBA64, p Point, size int, thickness float64, col color.RGBA64) (*image.NRGBA64, error) {
	return drawRectPx(img, p, size, size, thickness, col)
}

// @Name: draw-circle
// @Desc: Draws a circle centered at (x,y) with the given radius.
// @Param:      img       - - -   The image to fill
// @Param:      p         - - -   The center of the circle (relative)
// @Param:      radius    - - -   The radius of the circle (relative)
// @Param:      thickness - - -   The thickness of the circle border (absolute)
// @Param:      col  	  - - -   The circle color
// @Returns:    result    - - -	  The resulting image
func drawCircle(img *image.NRGBA64, p Point, radius float64, thickness float64, col color.RGBA64) (*image.NRGBA64, error) {
	bounds := img.Bounds()
	imgW, imgH := bounds.Max.X, bounds.Max.Y
	return drawCirclePx(img, P(p.X*float64(imgW), p.Y*float64(imgH)), int(radius*min(float64(imgW), float64(imgH))), thickness, col)
}

// @Name: draw-circle-px
// @Desc: Draws a circle centered at (x,y) with the given radius.
// @Param:      img       - - -   The image to fill
// @Param:      p         - - -   The center of the circle
// @Param:      radius    - - -   The radius of the circle
// @Param:      thickness - - -   The thickness of the circle border
// @Param:      col  	  - - -   The circle color
// @Returns:    result    - - -	  The resulting image
func drawCirclePx(img *image.NRGBA64, p Point, radius int, thickness float64, col color.RGBA64) (*image.NRGBA64, error) {
	return drawEllipsePx(img, p, radius, radius, thickness, col)
}

// @Name: draw-ellipse
// @Desc: Draws an ellipse centered at (x,y) with the given radii.
// @Param:      img       - - -   The image to fill
// @Param:      p         - - -   The center of the ellipse (relative)
// @Param:      rx        - - -   The horizontal radius (relative)
// @Param:      ry        - - -   The vertical radius (relative)
// @Param:      thickness - - -   The thickness of the ellipse border (absolute)
// @Param:      col  	  - - -   The ellipse color
// @Returns:    result    - - -	  The resulting image
func drawEllipse(img *image.NRGBA64, p Point, rx, ry float64, thickness float64, col color.RGBA64) (*image.NRGBA64, error) {
	bounds := img.Bounds()
	imgW, imgH := bounds.Max.X, bounds.Max.Y
	return drawEllipsePx(img, P(p.X*float64(imgW), p.Y*float64(imgH)), int(rx*float64(imgW)), int(ry*float64(imgH)), thickness, col)
}

// @Name: draw-ellipse-px
// @Desc: Draws an ellipse centered at (x,y) with the given radii.
// @Param:      img       - - -   The image to fill
// @Param:      p         - - -   The center of the ellipse
// @Param:      rx        - - -   The horizontal radius
// @Param:      ry        - - -   The vertical radius
// @Param:      thickness - - -   The thickness of the ellipse border
// @Param:      col  	  - - -   The ellipse color
// @Returns:    result    - - -	  The resulting image
func drawEllipsePx(img *image.NRGBA64, p Point, rx, ry int, thickness float64, col color.RGBA64) (*image.NRGBA64, error) {
	bounds := img.Bounds()
	result := image.NewNRGBA64(bounds)

	// Copy original image to result
	for py := bounds.Min.Y; py < bounds.Max.Y; py++ {
		for px := bounds.Min.X; px < bounds.Max.X; px++ {
			result.Set(px, py, img.NRGBA64At(px, py))
		}
	}

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
					alpha := uint32(float64(col.A) * coverage)

					// Blend colors using normal blend mode
					r, g, b, a := blendWithAlpha(
						uint32(existing.R), uint32(existing.G), uint32(existing.B), uint32(existing.A),
						uint32(col.R), uint32(col.G), uint32(col.B), alpha,
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
