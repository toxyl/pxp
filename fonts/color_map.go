package fonts

import (
	"image/color"
	"math"
)

func makeColRGBA64(r, g, b, a float64) color.RGBA64 {
	return color.RGBA64{R: uint16(math.Round(r * float64(0xFFFF))), G: uint16(math.Round(g * float64(0xFFFF))), B: uint16(math.Round(b * float64(0xFFFF))), A: uint16(math.Round(a * float64(0xFFFF)))}
}

var (
	opaque           = 1.000
	TRANSPARENT      = makeColRGBA64(0.000, 0.000, 0.000, 0.000)
	SHADOW           = makeColRGBA64(0.000, 0.000, 0.000, 0.600)
	BLACK            = makeColRGBA64(0.000, 0.000, 0.000, opaque)
	DARK_GRAY        = makeColRGBA64(0.250, 0.250, 0.250, opaque)
	GRAY             = makeColRGBA64(0.500, 0.500, 0.500, opaque)
	LIGHT_GRAY       = makeColRGBA64(0.750, 0.750, 0.750, opaque)
	WHITE            = makeColRGBA64(1.000, 1.000, 1.000, opaque)
	RED              = makeColRGBA64(1.000, 0.000, 0.000, opaque)
	ORANGE           = makeColRGBA64(1.000, 0.500, 0.000, opaque)
	YELLOW           = makeColRGBA64(1.000, 1.000, 0.000, opaque)
	LIGHT_GREEN      = makeColRGBA64(0.500, 1.000, 0.000, opaque)
	GREEN            = makeColRGBA64(0.000, 1.000, 0.000, opaque)
	TURQUOISE        = makeColRGBA64(0.000, 1.000, 0.500, opaque)
	CYAN             = makeColRGBA64(0.000, 1.000, 1.000, opaque)
	OCEAN            = makeColRGBA64(0.000, 0.500, 1.000, opaque)
	BLUE             = makeColRGBA64(0.000, 0.000, 1.000, opaque)
	VIOLET           = makeColRGBA64(0.500, 0.000, 1.000, opaque)
	MAGENTA          = makeColRGBA64(1.000, 0.000, 1.000, opaque)
	RASPBERRY        = makeColRGBA64(1.000, 0.000, 0.500, opaque)
	DARK_RED         = makeColRGBA64(0.500, 0.000, 0.000, opaque)
	DARK_ORANGE      = makeColRGBA64(0.500, 0.250, 0.000, opaque)
	DARK_YELLOW      = makeColRGBA64(0.500, 0.500, 0.000, opaque)
	DARK_LIGHT_GREEN = makeColRGBA64(0.250, 0.500, 0.000, opaque)
	DARK_GREEN       = makeColRGBA64(0.000, 0.500, 0.000, opaque)
	DARK_TURQUOISE   = makeColRGBA64(0.000, 0.500, 0.250, opaque)
	DARK_CYAN        = makeColRGBA64(0.000, 0.500, 0.500, opaque)
	DARK_OCEAN       = makeColRGBA64(0.000, 0.250, 0.500, opaque)
	DARK_BLUE        = makeColRGBA64(0.000, 0.000, 0.500, opaque)
	DARK_VIOLET      = makeColRGBA64(0.250, 0.000, 0.500, opaque)
	DARK_MAGENTA     = makeColRGBA64(0.500, 0.000, 0.500, opaque)
	DARK_RASPBERRY   = makeColRGBA64(0.500, 0.000, 0.250, opaque)
	Colors           = map[string]color.RGBA64{
		"transparent":      TRANSPARENT,
		"shadow":           SHADOW,
		"black":            BLACK,
		"dark-gray":        DARK_GRAY,
		"gray":             GRAY,
		"light-gray":       LIGHT_GRAY,
		"white":            WHITE,
		"red":              RED,
		"orange":           ORANGE,
		"yellow":           YELLOW,
		"light-green":      LIGHT_GREEN,
		"green":            GREEN,
		"turquoise":        TURQUOISE,
		"cyan":             CYAN,
		"ocean":            OCEAN,
		"blue":             BLUE,
		"violet":           VIOLET,
		"magenta":          MAGENTA,
		"raspberry":        RASPBERRY,
		"dark-red":         DARK_RED,
		"dark-orange":      DARK_ORANGE,
		"dark-yellow":      DARK_YELLOW,
		"dark-light-green": DARK_LIGHT_GREEN,
		"dark-green":       DARK_GREEN,
		"dark-turquoise":   DARK_TURQUOISE,
		"dark-cyan":        DARK_CYAN,
		"dark-ocean":       DARK_OCEAN,
		"dark-blue":        DARK_BLUE,
		"dark-violet":      DARK_VIOLET,
		"dark-magenta":     DARK_MAGENTA,
		"dark-raspberry":   DARK_RASPBERRY,
	}
)
