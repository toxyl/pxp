package language

import (
	"image"
	"image/color"

	"github.com/toxyl/math"
)

// @Name: draw-grid
// @Desc: Draws a grid from P(x1|y1) to P(x2|y2) with the given thickness and color.
// @Param:      img       - - -   The image to draw to
// @Param:      r         - - -   The area to draw the grid on (relative)
// @Param:      rows      - - -   The number of rows
// @Param:      cols      - - -   The number of cols
// @Param:      style     - - -   The thickness and color of the line
// @Returns:    result    - - -	  The resulting image
func drawGrid(img *image.NRGBA64, r Rect, rows, cols int, style LineStyle) (*image.NRGBA64, error) {
	bounds := img.Bounds()
	wi, hi := float64(bounds.Dx()), float64(bounds.Dy())
	x, y, w, h := r.X1()*wi, r.Y1()*hi, r.X2()*wi, r.Y2()*hi
	return drawGridPx(img, R(x, y, w, h), rows, cols, style)
}

// @Name: draw-grid-h
// @Desc: Draws a grid from P(x1|y1) to P(x2|y2) with the given thickness and color.
// @Param:      img       - - -   The image to draw to
// @Param:      r         - - -   The area to draw the grid on (relative)
// @Param:      rows      - - -   The number of rows
// @Param:      style     - - -   The thickness and color of the line
// @Returns:    result    - - -	  The resulting image
func drawGridH(img *image.NRGBA64, r Rect, rows int, style LineStyle) (*image.NRGBA64, error) {
	bounds := img.Bounds()
	wi, hi := float64(bounds.Dx()), float64(bounds.Dy())
	x, y, w, h := r.X1()*wi, r.Y1()*hi, r.X2()*wi, r.Y2()*hi
	return drawGridHPx(img, R(x, y, w, h), rows, style)
}

// @Name: draw-grid-v
// @Desc: Draws a grid from P(x1|y1) to P(x2|y2) with the given thickness and color.
// @Param:      img       - - -   The image to draw to
// @Param:      r         - - -   The area to draw the grid on (relative)
// @Param:      cols      - - -   The number of cols
// @Param:      style     - - -   The thickness and color of the line
// @Returns:    result    - - -	  The resulting image
func drawGridV(img *image.NRGBA64, r Rect, cols int, style LineStyle) (*image.NRGBA64, error) {
	bounds := img.Bounds()
	wi, hi := float64(bounds.Dx()), float64(bounds.Dy())
	x, y, w, h := r.X1()*wi, r.Y1()*hi, r.X2()*wi, r.Y2()*hi
	return drawGridVPx(img, R(x, y, w, h), cols, style)
}

// @Name: draw-grid-px
// @Desc: Draws a grid from P(x1|y1) to P(x2|y2) with the given thickness and color.
// @Param:      img       - - -   The image to draw to
// @Param:      r         - - -   The area to draw the grid on
// @Param:      rows      - - -   The number of rows
// @Param:      cols      - - -   The number of cols
// @Param:      style     - - -   The thickness and color of the line
// @Returns:    result    - - -	  The resulting image
func drawGridPx(img *image.NRGBA64, r Rect, rows, cols int, style LineStyle) (*image.NRGBA64, error) {
	x, y, w, h := r.X1(), r.Y1(), r.W()-1, r.H()-1
	wB, hB := w/float64(cols), h/float64(rows)

	for c := range cols {
		img, _ = drawLineVerticalPx(img, x+float64(c)*wB, 0, h, style)
	}
	img, _ = drawLineVerticalPx(img, x+w, 0, h, style)

	for r := range rows {
		img, _ = drawLineHorizontalPx(img, y+float64(r)*hB, 0, w, style)
	}
	img, _ = drawLineHorizontalPx(img, y+h, 0, w, style)

	return img, nil
}

// @Name: draw-grid-h-px
// @Desc: Draws a grid from P(x1|y1) to P(x2|y2) with the given thickness and color.
// @Param:      img       - - -   The image to draw to
// @Param:      r         - - -   The area to draw the grid on
// @Param:      rows      - - -   The number of rows
// @Param:      style     - - -   The thickness and color of the line
// @Returns:    result    - - -	  The resulting image
func drawGridHPx(img *image.NRGBA64, r Rect, rows int, style LineStyle) (*image.NRGBA64, error) {
	if rows <= 1 {
		return img, nil
	}
	y, w, h := r.Y1(), r.W()-1, r.H()-1
	cellH := h / float64(rows)

	for row := range rows - 1 {
		img, _ = drawLineHorizontalPx(img, y+cellH+float64(row)*cellH, 0, w, style)
	}

	return img, nil
}

// @Name: draw-grid-v-px
// @Desc: Draws a grid from P(x1|y1) to P(x2|y2) with the given thickness and color.
// @Param:      img       - - -   The image to draw to
// @Param:      r         - - -   The area to draw the grid on
// @Param:      cols      - - -   The number of cols
// @Param:      style     - - -   The thickness and color of the line
// @Returns:    result    - - -	  The resulting image
func drawGridVPx(img *image.NRGBA64, r Rect, cols int, style LineStyle) (*image.NRGBA64, error) {
	if cols <= 2 {
		return img, nil
	}
	x, w, h := r.X1(), r.W()-1, r.H()-1
	cellW := w / float64(cols)

	for col := range cols - 1 {
		img, _ = drawLineVerticalPx(img, x+cellW+float64(col)*cellW, 0, h, style)
	}

	return img, nil
}

// @Name: draw-line
// @Desc: Draws a line from P(x1|y1) to P(x2|y2) with the given thickness and color.
// @Param:      img       - - -   The image to draw to
// @Param:      p1        - - -   The start position (relative)
// @Param:      p2        - - -   The end position (relative)
// @Param:      style     - - -   The thickness and color of the line
// @Returns:    result    - - -	  The resulting image
func drawLine(img *image.NRGBA64, p1, p2 Point, style LineStyle) (*image.NRGBA64, error) {
	bounds := img.Bounds()
	w, h := bounds.Max.X, bounds.Max.Y
	return drawLinePx(img, *P(p1.X*float64(w), p1.Y*float64(h)), *P(p2.X*float64(w), p2.Y*float64(h)), style)
}

// @Name: draw-line-v
// @Desc: Draws a line from P(x|y1) to P(x|y2) with the given thickness and color.
// @Param:      img       - - -   The image to draw to
// @Param:      x         - - -   The position on the x-axis (relative)
// @Param:      y1        - - -   The start position on the y-axis (relative)
// @Param:      y2        - - -   The end position on the y-axis (relative)
// @Param:      style     - - -   The thickness and color of the line
// @Returns:    result    - - -	  The resulting image
func drawLineVertical(img *image.NRGBA64, x, y1, y2 float64, style LineStyle) (*image.NRGBA64, error) {
	return drawLine(img, *P(x, y1), *P(x, y2), style)
}

// @Name: draw-line-h
// @Desc: Draws a line from P(x1|y) to P(x2|y) with the given thickness and color.
// @Param:      img       - - -   The image to draw to
// @Param:      y         - - -   The position on the y-axis (relative)
// @Param:      x1        - - -   The start position on the x-axis (relative)
// @Param:      x2        - - -   The end position on the x-axis (relative)
// @Param:      style     - - -   The thickness and color of the line
// @Returns:    result    - - -	  The resulting image
func drawLineHorizontal(img *image.NRGBA64, y, x1, x2 float64, style LineStyle) (*image.NRGBA64, error) {
	return drawLine(img, *P(x1, y), *P(x2, y), style)
}

// @Name: draw-line-px
// @Desc: Draws a line from P(x1|y1) to P(x2|y2) with the given thickness and color.
// @Param:      img       - - -   The image to draw to
// @Param:      p1        - - -   The start position
// @Param:      p2        - - -   The end position
// @Param:      style     - - -   The thickness and color of the line
// @Returns:    result    - - -	  The resulting image
func drawLinePx(img *image.NRGBA64, p1, p2 Point, style LineStyle) (*image.NRGBA64, error) {
	bounds := img.Bounds()
	result := IClone(img)

	// Calculate line direction and perpendicular
	dx := float64(p2.X - p1.X)
	dy := float64(p2.Y - p1.Y)
	length := math.Sqrt(dx*dx + dy*dy)

	if length == 0 {
		// Single point case - draw a circle
		radius := style.Thickness
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
						alpha := uint32(float64(style.Color.A) * coverage)
						r, g, b, a := blendWithAlpha(
							uint32(existing.R), uint32(existing.G), uint32(existing.B), uint32(existing.A),
							uint32(style.Color.R), uint32(style.Color.G), uint32(style.Color.B), alpha,
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
	halfThickness := style.Thickness

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

				// Project pixel onto line segment to check if within bounds
				pxVec := float64(px) - float64(p1.X)
				pyVec := float64(py) - float64(p1.Y)
				projection := pxVec*dx + pyVec*dy

				// Check if projection is within line segment bounds
				if projection < -halfThickness || projection > length+halfThickness {
					continue
				}

				// Get pixel coverage for anti-aliasing
				coverage := getRectPixelCoverage(dist, halfThickness)

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

// @Name: draw-line-v-px
// @Desc: Draws a line from P(x|y1) to P(x|y2) with the given thickness and color.
// @Param:      img       - - -   The image to draw to
// @Param:      x         - - -   The position on the x-axis
// @Param:      y1        - - -   The start position on the y-axis
// @Param:      y2        - - -   The end position on the y-axis
// @Param:      style     - - -   The thickness and color of the line
// @Returns:    result    - - -	  The resulting image
func drawLineVerticalPx(img *image.NRGBA64, x, y1, y2 float64, style LineStyle) (*image.NRGBA64, error) {
	return drawLinePx(img, *P(x, y1), *P(x, y2), style)
}

// @Name: draw-line-h-px
// @Desc: Draws a line from P(x1|y) to P(x2|y) with the given thickness and color.
// @Param:      img       - - -   The image to draw to
// @Param:      y         - - -   The position on the y-axis
// @Param:      x1        - - -   The start position on the x-axis
// @Param:      x2        - - -   The end position on the x-axis
// @Param:      style     - - -   The thickness and color of the line
// @Returns:    result    - - -	  The resulting image
func drawLineHorizontalPx(img *image.NRGBA64, y, x1, x2 float64, style LineStyle) (*image.NRGBA64, error) {
	return drawLinePx(img, *P(x1, y), *P(x2, y), style)
}

// @Name: draw-line-polar
// @Desc: Draws a line from polar point P(r1|θ1) to P(r2|θ2) with the given thickness and color.
// @Param:      img       - - -   The image to draw to
// @Param:      origin    - - -   The origin point for the polar coordinate system (relative)
// @Param:      r1        - - -   The start radius
// @Param:      theta1    - - -   The start angle in radians
// @Param:      r2        - - -   The end radius
// @Param:      theta2    - - -   The end angle in radians
// @Param:      style     - - -   The thickness and color of the line
// @Returns:    result    - - -	  The resulting image
func drawLinePolar(img *image.NRGBA64, origin Point, r1, theta1, r2, theta2 float64, style LineStyle) (*image.NRGBA64, error) {
	bounds := img.Bounds()
	w, h := float64(bounds.Dx()), float64(bounds.Dy())
	ox, oy := origin.X*float64(w), origin.Y*float64(h)
	return drawLinePolarPx(img, *P(ox, oy), r1*math.Min(w, h), theta1, r2*math.Min(w, h), theta2, style)
}

// @Name: draw-line-polar-px
// @Desc: Draws a line from polar point P(r1|θ1) to P(r2|θ2) with the given thickness and color.
// @Param:      img       - - -   The image to draw to
// @Param:      origin    - - -   The origin point for the polar coordinate system
// @Param:      r1        - - -   The start radius
// @Param:      theta1    - - -   The start angle in radians
// @Param:      r2        - - -   The end radius
// @Param:      theta2    - - -   The end angle in radians
// @Param:      style     - - -   The thickness and color of the line
// @Returns:    result    - - -	  The resulting image
func drawLinePolarPx(img *image.NRGBA64, origin Point, r1, theta1, r2, theta2 float64, style LineStyle) (*image.NRGBA64, error) {
	x1 := origin.X + r1*math.Cos(theta1)
	y1 := origin.Y + r1*math.Sin(theta1)
	x2 := origin.X + r2*math.Cos(theta2)
	y2 := origin.Y + r2*math.Sin(theta2)
	return drawLinePx(img, *P(x1, y1), *P(x2, y2), style)
}

// @Name: draw-line-radial
// @Desc: Draws a radial line from radius r1 to r2 at angle θ with the given thickness and color.
// @Param:      img       - - -   The image to draw to
// @Param:      origin    - - -   The origin point for the polar coordinate system (relative)
// @Param:      theta     - - -   The angle in radians
// @Param:      r1        - - -   The start radius
// @Param:      r2        - - -   The end radius
// @Param:      style     - - -   The thickness and color of the line
// @Returns:    result    - - -	  The resulting image
func drawLineRadial(img *image.NRGBA64, origin Point, theta, r1, r2 float64, style LineStyle) (*image.NRGBA64, error) {
	bounds := img.Bounds()
	w, h := float64(bounds.Dx()), float64(bounds.Dy())
	ox, oy := origin.X*float64(w), origin.Y*float64(h)
	return drawLineRadialPx(img, *P(ox, oy), theta, r1*math.Min(w, h), r2*math.Min(w, h), style)
}

// @Name: draw-line-radial-px
// @Desc: Draws a radial line from radius r1 to r2 at angle θ with the given thickness and color.
// @Param:      img       - - -   The image to draw to
// @Param:      origin    - - -   The origin point for the polar coordinate system
// @Param:      theta     - - -   The angle in radians
// @Param:      r1        - - -   The start radius
// @Param:      r2        - - -   The end radius
// @Param:      style     - - -   The thickness and color of the line
// @Returns:    result    - - -	  The resulting image
func drawLineRadialPx(img *image.NRGBA64, origin Point, theta, r1, r2 float64, style LineStyle) (*image.NRGBA64, error) {
	x1 := origin.X + r1*math.Cos(theta)
	y1 := origin.Y + r1*math.Sin(theta)
	x2 := origin.X + r2*math.Cos(theta)
	y2 := origin.Y + r2*math.Sin(theta)
	return drawLinePx(img, *P(x1, y1), *P(x2, y2), style)
}

// @Name: draw-arc
// @Desc: Draws an arc at radius r from angle θ1 to θ2 with the given thickness and color.
// @Param:      img       - - -   The image to draw to
// @Param:      origin    - - -   The origin point for the polar coordinate system (relative)
// @Param:      r         - - -   The radius
// @Param:      theta1    - - -   The start angle in radians
// @Param:      theta2    - - -   The end angle in radians
// @Param:      style     - - -   The thickness and color of the line
// @Returns:    result    - - -	  The resulting image
func drawArc(img *image.NRGBA64, origin Point, r, theta1, theta2 float64, style LineStyle) (*image.NRGBA64, error) {
	bounds := img.Bounds()
	w, h := float64(bounds.Dx()), float64(bounds.Dy())
	ox, oy := origin.X*float64(w), origin.Y*float64(h)
	return drawArcPx(img, *P(ox, oy), r*math.Min(w, h), theta1, theta2, style)
}

// @Name: draw-arc-px
// @Desc: Draws an arc at radius r from angle θ1 to θ2 with the given thickness and color.
// @Param:      img       - - -   The image to draw to
// @Param:      origin    - - -   The origin point for the polar coordinate system
// @Param:      r         - - -   The radius
// @Param:      theta1    - - -   The start angle in radians
// @Param:      theta2    - - -   The end angle in radians
// @Param:      style     - - -   The thickness and color of the line
// @Returns:    result    - - -	  The resulting image
func drawArcPx(img *image.NRGBA64, origin Point, r, theta1, theta2 float64, style LineStyle) (*image.NRGBA64, error) {
	result := IClone(img)

	halfThickness := style.Thickness
	deltaTheta := theta2 - theta1
	steps := int(math.Max(1, math.Ceil(math.Abs(deltaTheta)*r/math.Max(1, halfThickness*2))))

	if steps < 2 {
		steps = 2
	}

	prevX := origin.X + r*math.Cos(theta1)
	prevY := origin.Y + r*math.Sin(theta1)

	for i := 1; i <= steps; i++ {
		theta := theta1 + float64(i)*deltaTheta/float64(steps)
		currX := origin.X + r*math.Cos(theta)
		currY := origin.Y + r*math.Sin(theta)

		segment, err := drawLinePx(result, *P(prevX, prevY), *P(currX, currY), style)
		if err != nil {
			return result, err
		}
		result = segment

		prevX = currX
		prevY = currY
	}

	return result, nil
}

// @Name: draw-grid-polar
// @Desc: Draws a polar grid with concentric circles and radial lines.
// @Param:      img       - - -   The image to draw to
// @Param:      origin    - - -   The origin point for the polar coordinate system (relative)
// @Param:      maxRadius - - -   The maximum radius (relative)
// @Param:      circles   - - -   The number of concentric circles
// @Param:      radials   - - -   The number of radial lines
// @Param:      style     - - -   The thickness and color of the line
// @Returns:    result    - - -	  The resulting image
func drawGridPolar(img *image.NRGBA64, origin Point, maxRadius float64, circles, radials int, style LineStyle) (*image.NRGBA64, error) {
	bounds := img.Bounds()
	w, h := float64(bounds.Dx()), float64(bounds.Dy())
	ox, oy := origin.X*float64(w), origin.Y*float64(h)
	maxR := maxRadius * math.Min(w, h)
	return drawGridPolarPx(img, *P(ox, oy), maxR, circles, radials, style)
}

// @Name: draw-grid-polar-px
// @Desc: Draws a polar grid with concentric circles and radial lines.
// @Param:      img       - - -   The image to draw to
// @Param:      origin    - - -   The origin point for the polar coordinate system
// @Param:      maxRadius - - -   The maximum radius
// @Param:      circles   - - -   The number of concentric circles
// @Param:      radials   - - -   The number of radial lines
// @Param:      style     - - -   The thickness and color of the line
// @Returns:    result    - - -	  The resulting image
func drawGridPolarPx(img *image.NRGBA64, origin Point, maxRadius float64, circles, radials int, style LineStyle) (*image.NRGBA64, error) {
	result := IClone(img)

	if circles > 0 {
		radiusStep := maxRadius / float64(circles+1)
		for i := 1; i <= circles; i++ {
			r := float64(i) * radiusStep
			arc, err := drawArcPx(result, origin, r, 0, 2*math.Pi, style)
			if err != nil {
				return result, err
			}
			result = arc
		}
	}

	if radials > 0 {
		angleStep := 2 * math.Pi / float64(radials)
		for i := 0; i < radials; i++ {
			theta := float64(i) * angleStep
			radial, err := drawLineRadialPx(result, origin, theta, 0, maxRadius, style)
			if err != nil {
				return result, err
			}
			result = radial
		}
	}

	return result, nil
}

// @Name: draw-grid-radial
// @Desc: Draws radial lines from the origin at evenly spaced angles.
// @Param:      img       - - -   The image to draw to
// @Param:      origin    - - -   The origin point for the polar coordinate system (relative)
// @Param:      maxRadius - - -   The maximum radius (relative)
// @Param:      radials   - - -   The number of radial lines
// @Param:      style     - - -   The thickness and color of the line
// @Returns:    result    - - -	  The resulting image
func drawGridRadial(img *image.NRGBA64, origin Point, maxRadius float64, radials int, style LineStyle) (*image.NRGBA64, error) {
	bounds := img.Bounds()
	w, h := float64(bounds.Dx()), float64(bounds.Dy())
	ox, oy := origin.X*float64(w), origin.Y*float64(h)
	maxR := maxRadius * math.Min(w, h)
	return drawGridRadialPx(img, *P(ox, oy), maxR, radials, style)
}

// @Name: draw-grid-radial-px
// @Desc: Draws radial lines from the origin at evenly spaced angles.
// @Param:      img       - - -   The image to draw to
// @Param:      origin    - - -   The origin point for the polar coordinate system
// @Param:      maxRadius - - -   The maximum radius
// @Param:      radials   - - -   The number of radial lines
// @Param:      style     - - -   The thickness and color of the line
// @Returns:    result    - - -	  The resulting image
func drawGridRadialPx(img *image.NRGBA64, origin Point, maxRadius float64, radials int, style LineStyle) (*image.NRGBA64, error) {
	if radials <= 0 {
		return img, nil
	}
	result := IClone(img)
	angleStep := 2 * math.Pi / float64(radials)
	for i := 0; i < radials; i++ {
		theta := float64(i) * angleStep
		radial, err := drawLineRadialPx(result, origin, theta, 0, maxRadius, style)
		if err != nil {
			return result, err
		}
		result = radial
	}
	return result, nil
}

// @Name: draw-grid-concentric
// @Desc: Draws concentric circles centered at the origin.
// @Param:      img       - - -   The image to draw to
// @Param:      origin    - - -   The origin point for the polar coordinate system (relative)
// @Param:      maxRadius - - -   The maximum radius (relative)
// @Param:      circles   - - -   The number of concentric circles
// @Param:      style     - - -   The thickness and color of the line
// @Returns:    result    - - -	  The resulting image
func drawGridConcentric(img *image.NRGBA64, origin Point, maxRadius float64, circles int, style LineStyle) (*image.NRGBA64, error) {
	bounds := img.Bounds()
	w, h := float64(bounds.Dx()), float64(bounds.Dy())
	ox, oy := origin.X*float64(w), origin.Y*float64(h)
	maxR := maxRadius * math.Min(w, h)
	return drawGridConcentricPx(img, *P(ox, oy), maxR, circles, style)
}

// @Name: draw-grid-concentric-px
// @Desc: Draws concentric circles centered at the origin.
// @Param:      img       - - -   The image to draw to
// @Param:      origin    - - -   The origin point for the polar coordinate system
// @Param:      maxRadius - - -   The maximum radius
// @Param:      circles   - - -   The number of concentric circles
// @Param:      style     - - -   The thickness and color of the line
// @Returns:    result    - - -	  The resulting image
func drawGridConcentricPx(img *image.NRGBA64, origin Point, maxRadius float64, circles int, style LineStyle) (*image.NRGBA64, error) {
	if circles <= 0 {
		return img, nil
	}
	result := IClone(img)
	radiusStep := maxRadius / float64(circles+1)
	for i := 1; i <= circles; i++ {
		r := float64(i) * radiusStep
		arc, err := drawArcPx(result, origin, r, 0, 2*math.Pi, style)
		if err != nil {
			return result, err
		}
		result = arc
	}
	return result, nil
}
