package language

import (
	"bytes"
	_ "embed"
	"image"
	"image/color"
	"image/png"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"golang.org/x/image/math/fixed"

	"github.com/toxyl/flo"
	"github.com/toxyl/math"
)

//go:embed font.png
var fontPNGData []byte

// @Name: draw-text
// @Desc: Draws a text at position (x,y).
// @Param:      img       - - -   The image to draw to
// @Param:      p         - - -   The upper-left coordinate of the text
// @Param:      t         - - -   The text to draw
// @Returns:    result    - - -	  The resulting image
func drawText(img *image.NRGBA64, p Point, t Text) (*image.NRGBA64, error) {
	bounds := img.Bounds()
	return drawTextPx(img, *p.Denorm(float64(bounds.Max.X), float64(bounds.Max.Y)), t)
}

// @Name: draw-text-px
// @Desc: Draws text at position (x,y) with the given style using TrueType fonts.
// @Param:      img       - - -   The image to draw to
// @Param:      p         - - -   The upper-left coordinate of the text
// @Param:      t         - - -   The text to draw
// @Returns:    result    - - -	  The resulting image
func drawTextPx(img *image.NRGBA64, p Point, t Text) (*image.NRGBA64, error) {
	result := IClone(img)

	if t.Style.Family == "mono" {
		return drawBitmapText(result, p, t)
	}

	// Try TrueType font first
	face, err := getFontFace(t.Style.Family, t.Style.Size)
	if err != nil {
		// Fallback to bitmap font
		return drawBitmapText(result, p, t)
	}

	metrics := face.Metrics()
	ascent := metrics.Ascent
	lineHeight := metrics.Height

	// Split text into lines
	lines := splitLines(t.Text)

	// Draw each line
	for i, line := range lines {
		if line == "" {
			continue
		}

		y := p.Y + float64(i)*float64(lineHeight>>6)
		baselineY := fixed.Int26_6(y*64) + ascent

		drawer := &font.Drawer{
			Dst:  result,
			Src:  image.NewUniform(color.RGBA64{R: t.Style.Color.R, G: t.Style.Color.G, B: t.Style.Color.B, A: t.Style.Color.A}),
			Face: face,
			Dot:  fixed.Point26_6{X: fixed.Int26_6(float64(p.X) * 64), Y: baselineY},
		}

		drawer.DrawString(line)
	}

	return result, nil
}

// @Name: draw-text-outline
// @Desc: Draws only the outline of text at position (x,y).
// @Param:      img       - - -   The image to draw to
// @Param:      p         - - -   The upper-left coordinate of the text
// @Param:      t         - - -   The text to outline
// @Param:      outline   - - -   The outline style (thickness and color)
// @Returns:    result    - - -   The resulting image
func drawTextOutline(img *image.NRGBA64, p Point, t Text, outline LineStyle) (*image.NRGBA64, error) {
	bounds := img.Bounds()
	return drawTextOutlinePx(img, *p.Denorm(float64(bounds.Max.X), float64(bounds.Max.Y)), t, outline)
}

// @Name: draw-text-outline-px
// @Desc: Draws only the outline of text at position (x,y).
// @Param:      img       - - -   The image to draw to
// @Param:      p         - - -   The upper-left coordinate of the text
// @Param:      t         - - -   The text to outline
// @Param:      outline   - - -   The outline style (thickness and color)
// @Returns:    result    - - -   The resulting image
func drawTextOutlinePx(img *image.NRGBA64, p Point, t Text, outline LineStyle) (*image.NRGBA64, error) {
	result := IClone(img)
	bounds := img.Bounds()

	// Create temporary transparent image to render text
	tempImg := I(bounds.Dx(), bounds.Dy())

	// Draw text to temporary image
	tempImg, err := drawTextPx(tempImg, p, t)
	if err != nil {
		return result, err
	}

	// Find all solid text pixels (alpha > 0.8) and calculate outline
	thickness := outline.Thickness
	searchRadius := int(thickness) + 1

	// For each pixel in the result image
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			// Check if this pixel is already text (solid text, not anti-aliased)
			tempPixel := tempImg.NRGBA64At(x, y)
			if tempPixel.A > 52428 { // Only consider solid text pixels (alpha > 0.8)
				// This is a solid text pixel, skip outline drawing
				continue
			}

			// Find minimum distance to any solid text pixel
			minDist := thickness + 1.0
			for dy := -searchRadius; dy <= searchRadius; dy++ {
				for dx := -searchRadius; dx <= searchRadius; dx++ {
					checkX := x + dx
					checkY := y + dy

					// Check bounds
					if checkX < bounds.Min.X || checkX >= bounds.Max.X ||
						checkY < bounds.Min.Y || checkY >= bounds.Max.Y {
						continue
					}

					// Check if this is a solid text pixel (not anti-aliased)
					checkPixel := tempImg.NRGBA64At(checkX, checkY)
					if checkPixel.A > 52428 { // Only consider solid text pixels (alpha > 0.8)
						// Calculate distance
						dist := math.Sqrt(float64(dx*dx + dy*dy))
						if dist < minDist {
							minDist = dist
						}
					}
				}
			}

			// If within outline thickness, draw outline pixel
			if minDist <= thickness {
				coverage := getRectPixelCoverage(minDist, thickness)
				if coverage > 0 {
					existing := result.NRGBA64At(x, y)
					alpha := uint32(float64(outline.Color.A) * coverage)

					r, g, b, a := blendWithAlpha(
						uint32(existing.R), uint32(existing.G), uint32(existing.B), uint32(existing.A),
						uint32(outline.Color.R), uint32(outline.Color.G), uint32(outline.Color.B), alpha,
						func(r1, g1, b1, r2, g2, b2 uint32) (uint32, uint32, uint32) {
							return r2, g2, b2
						},
					)

					result.Set(x, y, color.NRGBA64{
						R: uint16(r), G: uint16(g), B: uint16(b), A: uint16(a),
					})
				}
			}
		}
	}

	// Return only the outline, no text rendering
	return result, nil
}

// splitLines splits text by \n characters and removes \r
func splitLines(text string) []string {
	var lines []string
	var current strings.Builder

	for _, r := range text {
		if r == '\n' {
			lines = append(lines, current.String())
			current.Reset()
		} else if r != '\r' {
			current.WriteRune(r)
		}
	}

	// Add the last line if it's not empty or if there was no trailing newline
	if current.Len() > 0 || len(lines) == 0 {
		lines = append(lines, current.String())
	}

	return lines
}

var (
	fontCache     = make(map[string]*opentype.Font)
	fontCacheLock sync.RWMutex
	faceCache     = make(map[string]font.Face)
	faceCacheLock sync.RWMutex
)

// getFontFace returns a font face for the given family and size
func getFontFace(family string, size float64) (font.Face, error) {

	cacheKey := family + ":" + string(rune(int(size)))

	faceCacheLock.RLock()
	if face, ok := faceCache[cacheKey]; ok {
		faceCacheLock.RUnlock()
		return face, nil
	}
	faceCacheLock.RUnlock()

	ttf, err := loadFont(family)
	if err != nil {
		return nil, err
	}

	face, err := opentype.NewFace(ttf, &opentype.FaceOptions{
		Size:    size,
		DPI:     96,                   // Higher DPI for crisper rendering
		Hinting: font.HintingVertical, // Better hinting for crisp text
	})
	if err != nil {
		return nil, err
	}

	faceCacheLock.Lock()
	faceCache[cacheKey] = face
	faceCacheLock.Unlock()

	return face, nil
}

// loadFont loads a TrueType font from embedded data or system
func loadFont(family string) (*opentype.Font, error) {
	fontCacheLock.RLock()
	if ttf, ok := fontCache[family]; ok {
		fontCacheLock.RUnlock()
		return ttf, nil
	}
	fontCacheLock.RUnlock()

	var fontData []byte
	var err error

	switch family {
	case "sans", "sans-serif", "":
		fontData, err = loadSystemFont("arial.ttf", "DejaVuSans.ttf", "liberation-sans.ttf")
	case "serif":
		fontData, err = loadSystemFont("times.ttf", "DejaVuSerif.ttf", "liberation-serif.ttf")
	case "bold":
		fontData, err = loadSystemFont("arialbd.ttf", "DejaVuSans-Bold.ttf", "liberation-sans-bold.ttf")
	default:
		// Try exact filename first, then fallback
		fontData, err = loadSystemFont(family+".ttf", family+".TTF", family+".otf", family+".OTF")
		if err != nil {
			return nil, err
		}
	}

	if err != nil {
		return nil, err
	}

	ttf, err := opentype.Parse(fontData)
	if err != nil {
		return nil, err
	}

	fontCacheLock.Lock()
	fontCache[family] = ttf
	fontCacheLock.Unlock()

	return ttf, nil
}

// loadSystemFont attempts to load a font from common system locations
func loadSystemFont(names ...string) ([]byte, error) {
	fontPaths := []string{
		"C:/Windows/Fonts/",
		"/usr/share/fonts/truetype/",
		"/usr/share/fonts/TTF/",
		"/System/Library/Fonts/",
		"/Library/Fonts/",
	}

	// Add user-specific font directories
	if homeDir, err := os.UserHomeDir(); err == nil {
		fontPaths = append(fontPaths,
			filepath.Join(homeDir, "AppData", "Local", "Microsoft", "Windows", "Fonts")+"/",
			filepath.Join(homeDir, ".fonts")+"/",
			filepath.Join(homeDir, "Library", "Fonts")+"/",
		)
	}

	for _, name := range names {
		for _, basePath := range fontPaths {
			path := basePath + name
			f := flo.File(path)
			if f.Exists() {
				return f.AsBytes(), nil
			}
		}
	}

	return nil, &FontLoadError{Message: "no suitable font found"}
}

type FontLoadError struct {
	Message string
}

func (e *FontLoadError) Error() string {
	return e.Message
}

var bitmapFontCache *image.NRGBA64

// drawBitmapText renders text using the bitmap font from font.png
func drawBitmapText(img *image.NRGBA64, p Point, t Text) (*image.NRGBA64, error) {
	result := IClone(img)

	// Load bitmap font if not cached
	if bitmapFontCache == nil {
		fontImg, err := png.Decode(bytes.NewReader(fontPNGData))
		if err != nil {
			return result, err
		}

		// Convert to NRGBA64
		bitmapFontCache = image.NewNRGBA64(fontImg.Bounds())
		for y := fontImg.Bounds().Min.Y; y < fontImg.Bounds().Max.Y; y++ {
			for x := fontImg.Bounds().Min.X; x < fontImg.Bounds().Max.X; x++ {
				bitmapFontCache.Set(x, y, fontImg.At(x, y))
			}
		}
	}

	// Simple bitmap font rendering (9x10 characters)
	const (
		charsPerRow = 16
		charWidth   = 9
		charHeight  = 10
	)
	scale := int(t.Style.Size) / charHeight
	if scale <= 0 {
		scale = 1
	}

	// Calculate grid layout from sprite sheet dimensions
	spriteWidth := bitmapFontCache.Bounds().Max.X
	spriteHeight := bitmapFontCache.Bounds().Max.Y

	lines := splitLines(t.Text)

	for lineIdx, line := range lines {
		y := int(p.Y) + lineIdx*(charHeight*scale)

		for charIdx, char := range line {
			charIndex := int(char) - 33
			if charIndex < 0 || charIndex >= 16*6 {
				charIndex = 95 // Default to underscore for unsupported characters
			}

			x := int(p.X) + charIdx*(charWidth*scale)

			// Calculate source position in font atlas using grid layout
			srcX := (charIndex % charsPerRow) * charWidth
			srcY := (charIndex / charsPerRow) * charHeight

			// Draw scaled character
			for dy := 0; dy < charHeight*scale; dy++ {
				for dx := 0; dx < charWidth*scale; dx++ {
					srcPixelX := srcX + dx/scale
					srcPixelY := srcY + dy/scale

					if srcPixelX >= spriteWidth || srcPixelY >= spriteHeight {
						continue
					}

					fontPixel := bitmapFontCache.NRGBA64At(srcPixelX, srcPixelY)
					if fontPixel.A > 0 {
						// Apply text color
						finalColor := color.NRGBA64{
							R: t.Style.Color.R,
							G: t.Style.Color.G,
							B: t.Style.Color.B,
							A: t.Style.Color.A,
						}

						dstX := x + dx
						dstY := y + dy
						if dstX >= 0 && dstY >= 0 && dstX < result.Bounds().Max.X && dstY < result.Bounds().Max.Y {
							result.Set(dstX, dstY, finalColor)
						}
					}
				}
			}
		}
	}

	return result, nil
}
