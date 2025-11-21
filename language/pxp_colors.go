package language

import (
	"fmt"
	"image/color"
	"sort"

	"github.com/toxyl/math"
)

// @Name: hsla
// @Desc: Creates a color from HSLA values
// @Param:      h      	"°" 0..360   	0.0   	The color's hue
// @Param:      s     	"%" 0.0..1.0   	0.5   	The color's saturation
// @Param:      l     	"%" 0.0..1.0   	0.5   	The color's luminosity
// @Param:      alpha  	"%" 0.0..1.0   	1.0   	The color's alpha
// @Returns:    result  - 	-   		-   	The color
func hsla(h float64, s float64, l float64, alpha float64) (color.RGBA64, error) {
	// Convert HSLA to RGB
	var r, g, b float64
	if s == 0 {
		r, g, b = l, l, l
	} else {
		var q float64
		if l < 0.5 {
			q = l * (1 + s)
		} else {
			q = l + s - l*s
		}
		p := 2*l - q

		r = hueToRGB(p, q, h/360+1.0/3.0)
		g = hueToRGB(p, q, h/360)
		b = hueToRGB(p, q, h/360-1.0/3.0)
	}

	// Convert to 16-bit color channels
	R := uint16(r * 65535)
	G := uint16(g * 65535)
	B := uint16(b * 65535)
	A := uint16(alpha * 65535)

	return color.RGBA64{
		R: R,
		G: G,
		B: B,
		A: A,
	}, nil
}

// @Name: rgba
// @Desc: Creates a color from RGBA values (8-bit per channel)
// @Param:      r      	"" 0..255   	0   	The red component
// @Param:      g     	"" 0..255   	0   	The green component
// @Param:      b     	"" 0..255   	0   	The blue component
// @Param:      alpha  	"%" 0.0..1.0   	1.0   	The alpha component
// @Returns:    result  - 	-   		-   	The color
func rgba(r float64, g float64, b float64, alpha float64) (color.RGBA64, error) {
	// Convert 8-bit values to 16-bit color channels
	R := uint16((r / 255.0) * 65535)
	G := uint16((g / 255.0) * 65535)
	B := uint16((b / 255.0) * 65535)
	A := uint16(alpha * 65535)

	return color.RGBA64{
		R: R,
		G: G,
		B: B,
		A: A,
	}, nil
}

// @Name: rgba64
// @Desc: Creates a color from RGBA values (16-bit per channel)
// @Param:      r      	"" 0..65535   	0   	The red component
// @Param:      g     	"" 0..65535   	0   	The green component
// @Param:      b     	"" 0..65535   	0   	The blue component
// @Param:      alpha  	"%" 0.0..1.0   	1.0   	The alpha component
// @Returns:    result  - 	-   		-   	The color
func rgba64(r float64, g float64, b float64, alpha float64) (color.RGBA64, error) {
	return color.RGBA64{
		R: uint16(r),
		G: uint16(g),
		B: uint16(b),
		A: uint16(alpha * 65535),
	}, nil
}

// @Name: cmyk
// @Desc: Creates a color from CMYK values
// @Param:      c      	"%" 0.0..1.0   	0.0   	The cyan component
// @Param:      m     	"%" 0.0..1.0   	0.0   	The magenta component
// @Param:      y     	"%" 0.0..1.0   	0.0   	The yellow component
// @Param:      k     	"%" 0.0..1.0   	0.0   	The key (black) component
// @Param:      alpha  	"%" 0.0..1.0   	1.0   	The alpha component
// @Returns:    result  - 	-   		-   	The color
func cmyk(c float64, m float64, y float64, k float64, alpha float64) (color.RGBA64, error) {
	// Convert CMYK to RGB
	r := (1 - c) * (1 - k)
	g := (1 - m) * (1 - k)
	b := (1 - y) * (1 - k)

	// Convert to 16-bit color channels
	R := uint16(r * 65535)
	G := uint16(g * 65535)
	B := uint16(b * 65535)
	A := uint16(alpha * 65535)

	return color.RGBA64{
		R: R,
		G: G,
		B: B,
		A: A,
	}, nil
}

// @Name: hsv
// @Desc: Creates a color from HSV values
// @Param:      h      	"°" 0..360   	0.0   	The color's hue
// @Param:      s     	"%" 0.0..1.0   	0.5   	The color's saturation
// @Param:      v     	"%" 0.0..1.0   	0.5   	The color's value (brightness)
// @Param:      alpha  	"%" 0.0..1.0   	1.0   	The color's alpha
// @Returns:    result  - 	-   		-   	The color
func hsv(h float64, s float64, v float64, alpha float64) (color.RGBA64, error) {
	// Convert HSV to RGB
	var r, g, b float64

	if s == 0 {
		r, g, b = v, v, v
	} else {
		sector := h / 60
		i := int(sector)
		f := sector - float64(i)
		p := v * (1 - s)
		q := v * (1 - s*f)
		t := v * (1 - s*(1-f))

		switch i {
		case 0:
			r, g, b = v, t, p
		case 1:
			r, g, b = q, v, p
		case 2:
			r, g, b = p, v, t
		case 3:
			r, g, b = p, q, v
		case 4:
			r, g, b = t, p, v
		default:
			r, g, b = v, p, q
		}
	}

	// Convert to 16-bit color channels
	R := uint16(r * 65535)
	G := uint16(g * 65535)
	B := uint16(b * 65535)
	A := uint16(alpha * 65535)

	return color.RGBA64{
		R: R,
		G: G,
		B: B,
		A: A,
	}, nil
}

// @Name: lab
// @Desc: Creates a color from CIELAB values
// @Param:      l      	"" 0..100   	50.0   	The lightness component
// @Param:      a     	"" -128..127   	0.0   	The green-red component
// @Param:      b     	"" -128..127   	0.0   	The blue-yellow component
// @Param:      alpha  	"%" 0.0..1.0   	1.0   	The color's alpha
// @Returns:    result  - 	-   		-   	The color
func lab(l float64, a float64, b float64, alpha float64) (color.RGBA64, error) {
	// Convert LAB to XYZ first
	// Using D65 illuminant
	fy := (l + 16) / 116
	fx := fy + (a / 500)
	fz := fy - (b / 200)

	// D65 white point
	xn, yn, zn := 0.95047, 1.0, 1.08883

	var x, y, z float64

	// Convert using piece-wise function
	if fx*fx*fx > 0.008856 {
		x = xn * fx * fx * fx
	} else {
		x = (116*fx - 16) / 903.3 * xn
	}

	if l > (903.3 * 0.008856) {
		y = yn * ((l + 16) / 116) * ((l + 16) / 116) * ((l + 16) / 116)
	} else {
		y = yn * l / 903.3
	}

	if fz*fz*fz > 0.008856 {
		z = zn * fz * fz * fz
	} else {
		z = (116*fz - 16) / 903.3 * zn
	}

	// Convert XYZ to RGB
	var r, g, b2 float64
	r = 3.2406*x - 1.5372*y - 0.4986*z
	g = -0.9689*x + 1.8758*y + 0.0415*z
	b2 = 0.0557*x - 0.2040*y + 1.0570*z

	// Clip values to 0-1 range
	r = math.Max(0.0, math.Min(1.0, r))
	g = math.Max(0.0, math.Min(1.0, g))
	b2 = math.Max(0.0, math.Min(1.0, b2))

	// Convert to 16-bit color channels
	R := uint16(r * 65535)
	G := uint16(g * 65535)
	B := uint16(b2 * 65535)
	A := uint16(alpha * 65535)

	return color.RGBA64{
		R: R,
		G: G,
		B: B,
		A: A,
	}, nil
}

// @Name: hwb
// @Desc: Creates a color from HWB (Hue, Whiteness, Blackness) values
// @Param:      h      	"°" 0..360   	0.0   	The color's hue
// @Param:      w     	"%" 0.0..1.0   	0.0   	The whiteness component
// @Param:      b     	"%" 0.0..1.0   	0.0   	The blackness component
// @Param:      alpha  	"%" 0.0..1.0   	1.0   	The color's alpha
// @Returns:    result  - 	-   		-   	The color
func hwb(h float64, w float64, b float64, alpha float64) (color.RGBA64, error) {
	// First convert HWB to HSV
	// Value is 1 - blackness
	v := 1 - b
	// Saturation is (1 - whiteness) / value
	var s float64
	if v != 0 {
		s = 1 - (w / v)
	}
	s = math.Max(0.0, math.Min(1.0, s))

	// Now we can use HSV conversion
	var r, g, black float64

	if s == 0 {
		r, g, black = v, v, v
	} else {
		sector := h / 60
		i := int(sector)
		f := sector - float64(i)
		p := v * (1 - s)
		q := v * (1 - s*f)
		t := v * (1 - s*(1-f))

		switch i {
		case 0:
			r, g, black = v, t, p
		case 1:
			r, g, black = q, v, p
		case 2:
			r, g, black = p, v, t
		case 3:
			r, g, black = p, q, v
		case 4:
			r, g, black = t, p, v
		default:
			r, g, black = v, p, q
		}
	}

	// Convert to 16-bit color channels
	R := uint16(r * 65535)
	G := uint16(g * 65535)
	B := uint16(black * 65535)
	A := uint16(alpha * 65535)

	return color.RGBA64{
		R: R,
		G: G,
		B: B,
		A: A,
	}, nil
}

// @Name: xyz
// @Desc: Creates a color from CIE XYZ values
// @Param:      x      	"" 0.0..0.95047   0.0   	The X component (red)
// @Param:      y     	"" 0.0..1.0   	0.0   	The Y component (green)
// @Param:      z     	"" 0.0..1.08883   0.0   	The Z component (blue)
// @Param:      alpha  	"%" 0.0..1.0   	1.0   	The color's alpha
// @Returns:    result  - 	-   		-   	The color
func xyz(x float64, y float64, z float64, alpha float64) (color.RGBA64, error) {
	// Convert XYZ to RGB using D65 matrix
	var r, g, b float64
	r = 3.2406*x - 1.5372*y - 0.4986*z
	g = -0.9689*x + 1.8758*y + 0.0415*z
	b = 0.0557*x - 0.2040*y + 1.0570*z

	// Clip values to 0-1 range
	r = math.Clamp(r, 0.0, 1.0)
	g = math.Clamp(g, 0.0, 1.0)
	b = math.Clamp(b, 0.0, 1.0)

	// Convert to 16-bit color channels
	R := uint16(r * 65535)
	G := uint16(g * 65535)
	B := uint16(b * 65535)
	A := uint16(alpha * 65535)

	return color.RGBA64{
		R: R,
		G: G,
		B: B,
		A: A,
	}, nil
}

// @Name: lch
// @Desc: Creates a color from LCH (Lightness, Chroma, Hue) values
// @Param:      l      	"" 0..100   	50.0   	The lightness component
// @Param:      c     	"" 0..128   	0.0   	The chroma component
// @Param:      h     	"°" 0..360   	0.0   	The hue component
// @Param:      alpha  	"%" 0.0..1.0   	1.0   	The color's alpha
// @Returns:    result  - 	-   		-   	The color
func lch(l float64, c float64, h float64, alpha float64) (color.RGBA64, error) {
	// Convert LCH to LAB
	// a = C * cos(h)
	// b = C * sin(h)
	hRad := h * (math.Pi / 180.0)
	a := c * math.Cos(hRad)
	b := c * math.Sin(hRad)

	// Now we can use our LAB conversion
	return lab(l, a, b, alpha)
}

// @Name: yuv
// @Desc: Creates a color from YUV values
// @Param:      y      	"" 0.0..1.0   	0.0   	The luminance component
// @Param:      u     	"" -0.436..0.436   0.0   	The U chrominance component
// @Param:      v     	"" -0.615..0.615   0.0   	The V chrominance component
// @Param:      alpha  	"%" 0.0..1.0   	1.0   	The color's alpha
// @Returns:    result  - 	-   		-   	The color
func yuv(y float64, u float64, v float64, alpha float64) (color.RGBA64, error) {
	// Convert YUV to RGB
	var r, g, b float64
	r = y + 1.13983*v
	g = y - 0.39465*u - 0.58060*v
	b = y + 2.03211*u

	// Clip values to 0-1 range
	r = math.Clamp(r, 0.0, 1.0)
	g = math.Clamp(g, 0.0, 1.0)
	b = math.Clamp(b, 0.0, 1.0)

	// Convert to 16-bit color channels
	R := uint16(r * 65535)
	G := uint16(g * 65535)
	B := uint16(b * 65535)
	A := uint16(alpha * 65535)

	return color.RGBA64{
		R: R,
		G: G,
		B: B,
		A: A,
	}, nil
}

// @Name: ycbcr
// @Desc: Creates a color from YCbCr values (digital video)
// @Param:      y      	"" 0..255   	0   	The luminance component
// @Param:      cb     	"" 0..255   	128   	The blue-difference chroma component
// @Param:      cr     	"" 0..255   	128   	The red-difference chroma component
// @Param:      alpha  	"%" 0.0..1.0   	1.0   	The color's alpha
// @Returns:    result  - 	-   		-   	The color
func ycbcr(y float64, cb float64, cr float64, alpha float64) (color.RGBA64, error) {
	// Normalize YCbCr values
	yn := y / 255.0
	cbn := (cb - 128) / 255.0
	crn := (cr - 128) / 255.0

	// Convert YCbCr to RGB
	var r, g, b float64
	r = yn + 1.402*crn
	g = yn - 0.344136*cbn - 0.714136*crn
	b = yn + 1.772*cbn

	// Clip values to 0-1 range
	r = math.Clamp(r, 0.0, 1.0)
	g = math.Clamp(g, 0.0, 1.0)
	b = math.Clamp(b, 0.0, 1.0)

	// Convert to 16-bit color channels
	R := uint16(r * 65535)
	G := uint16(g * 65535)
	B := uint16(b * 65535)
	A := uint16(alpha * 65535)

	return color.RGBA64{
		R: R,
		G: G,
		B: B,
		A: A,
	}, nil
}

// @Name: map-color
// @Desc: Maps a value to a color using color stops with HSLA interpolation
// @Param:      value      	"" -   	0.0   	The value to map to a color
// @Param:      min        	"" -   	0.0   	Minimum value of the range
// @Param:      max        	"" -   	1.0   	Maximum value of the range
// @Param:      stops      	"" -   	-   	Color stops as [][]any where each stop is [threshold, hue, saturation, lightness, alpha]; threshold is a raw value between min and max
// @Returns:    result      - -		-   	The interpolated color
func mapColor(value float64, min float64, max float64, stops [][]any) (color.RGBA64, error) {
	if len(stops) == 0 {
		return color.RGBA64{}, fmt.Errorf("no color stops provided")
	}

	type colorstop struct {
		threshold, thresholdNorm, hue, saturation, lightness, alpha float64
	}
	colorstops := []colorstop{}
	for _, s := range stops {
		t, h, sa, l, a := s[0].(float64), s[1].(float64), s[2].(float64), s[3].(float64), s[4].(float64)
		tn := 0.0
		if max != min {
			tn = (t - min) / (max - min)
		}
		tn = math.Max(0.0, math.Min(1.0, tn))
		colorstops = append(colorstops, colorstop{
			threshold:     t,
			thresholdNorm: tn,
			hue:           h,
			saturation:    sa,
			lightness:     l,
			alpha:         a,
		})
	}
	sort.Slice(colorstops, func(i, j int) bool {
		return colorstops[i].threshold < colorstops[j].threshold
	})

	var startStop, endStop *colorstop

	for i := 0; i < len(colorstops)-1; i++ {
		if value >= colorstops[i].threshold && value < colorstops[i+1].threshold {
			startStop = &colorstops[i]
			endStop = &colorstops[i+1]
			break
		}
	}

	// If value is before first stop or after last stop, use the nearest stop
	if startStop == nil {
		if value < colorstops[0].threshold {
			startStop = &colorstops[0]
			endStop = &colorstops[0]
		} else if value >= colorstops[len(colorstops)-1].threshold {
			startStop = &colorstops[len(colorstops)-1]
			endStop = &colorstops[len(colorstops)-1]
		} else {
			// This shouldn't happen if the loop above worked correctly, but handle edge case
			// where value equals the last threshold exactly
			startStop = &colorstops[len(colorstops)-1]
			endStop = &colorstops[len(colorstops)-1]
		}
	}

	// If start and end are the same, return that color directly
	if startStop.threshold == endStop.threshold {
		return hsla(startStop.hue, startStop.saturation, startStop.lightness, startStop.alpha)
	}

	// Calculate interpolation factor
	factor := (value - startStop.threshold) / (endStop.threshold - startStop.threshold)

	// Interpolate HSLA values
	// Handle hue wrapping (hue is circular 0-360)
	dh := endStop.hue - startStop.hue
	if math.Abs(dh) > 180 {
		if dh > 0 {
			dh -= 360
		} else {
			dh += 360
		}
	}
	h := startStop.hue + factor*dh
	if h < 0 {
		h += 360
	} else if h >= 360 {
		h -= 360
	}

	s := startStop.saturation + factor*(endStop.saturation-startStop.saturation)
	l := startStop.lightness + factor*(endStop.lightness-startStop.lightness)
	a := startStop.alpha + factor*(endStop.alpha-startStop.alpha)

	return hsla(h, s, l, a)
}

// @Name: set-alpha
// @Desc: Sets the alpha channel of a color
// @Param:      c      	"" -   	-   	The color to modify
// @Param:      alpha  	"%" 0.0..1.0   	1.0   	The new alpha value
// @Returns:    result  - 	-   		-   	The color with new alpha
func setAlpha(c color.RGBA64, alpha float64) (color.RGBA64, error) {
	A := uint16(alpha * 65535)
	return color.RGBA64{
		R: c.R,
		G: c.G,
		B: c.B,
		A: A,
	}, nil
}
