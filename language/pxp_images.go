package language

import (
	"bytes"
	"image"

	"github.com/toxyl/math"
)

// DetectImageType determines the image type from magic bytes
func DetectImageType(data []byte) string {
	if len(data) < 8 {
		return ""
	}

	// Check for PNG
	if bytes.Equal(data[:8], []byte{0x89, 0x50, 0x4E, 0x47, 0x0D, 0x0A, 0x1A, 0x0A}) {
		return "png"
	}

	// Check for JPEG
	if bytes.Equal(data[:2], []byte{0xFF, 0xD8}) && bytes.Equal(data[len(data)-2:], []byte{0xFF, 0xD9}) {
		return "jpeg"
	}

	// Check for GIF
	if len(data) >= 6 && (bytes.Equal(data[:6], []byte("GIF87a")) || bytes.Equal(data[:6], []byte("GIF89a"))) {
		return "gif"
	}

	// Check for HEIC
	// HEIC files start with 'ftyp' box which contains 'heic' or 'mif1'
	if len(data) >= 12 && bytes.Equal(data[4:8], []byte("ftyp")) {
		if bytes.Equal(data[8:12], []byte("heic")) || bytes.Equal(data[8:12], []byte("mif1")) {
			return "heic"
		}
	}

	return ""
}

// @Name: It
// @Desc: Translates the given image by expanding/cropping the left + top borders.
// @Param:      img     - - -   The image to translate
// @Param:      dt      - - -   The translation to apply
// @Returns:    result  - - -	The new image
func translateImage(img *image.NRGBA64, dt Point) (*image.NRGBA64, error) {
	left, top := 0, 0
	if dt.X != 0 {
		if dt.X < 0 {
			left = int(dt.X)
		} else {
			left = math.Abs(int(dt.X))
		}
	}
	if dt.Y != 0 {
		if dt.Y < 0 {
			top = int(dt.Y)
		} else {
			top = math.Abs(int(dt.Y))
		}
	}
	return expandPx(img, left, 0, top, 0)
}
