package language

import (
	"bytes"
	"image"
	"strings"

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

// @Name: blend-aligned
// @Desc: Aligns two images using the given anchor (left-top, top, top-right, left, center, right, bottom-left, bottom, bottom-right) and blends them using the given blendmode (defaults to normal).
// @Param:      imgA    - -   	-   		The bottom image
// @Param:      imgB    - -   	-   		The top image
// @Param:      anchor  - -   	"C"         The anchor to align to (TL, T, TR, L, C, R, BL, B, BR)
// @Param:      mode    - -   	"normal"    The blendmode name
// @Returns:    result  - -   	-   		The aligned and blended image
func blendAligned(imgA, imgB *image.NRGBA64, anchor, mode string) (*image.NRGBA64, error) {
	wa, ha, wb, hb := imgA.Rect.Dx(), imgA.Rect.Dy(), imgB.Rect.Dx(), imgB.Rect.Dy()
	w, h := int(math.Max(float64(wa), float64(wb))), int(math.Max(float64(ha), float64(hb)))

	wda, hda := w-wa, h-ha
	wdb, hdb := w-wb, h-hb
	la, ra, ta, ba := 0, 0, 0, 0
	lb, rb, tb, bb := 0, 0, 0, 0

	switch strings.ToUpper(anchor) {
	case "TL":
		la, ra = 0, wda
		ta, ba = 0, hda
		lb, rb = 0, wdb
		tb, bb = 0, hdb
	case "T":
		la, ra = wda/2, wda/2
		ta, ba = 0, hda
		lb, rb = wdb/2, wdb/2
		tb, bb = 0, hdb
	case "TR":
		la, ra = wda, 0
		ta, ba = 0, hda
		lb, rb = wdb, 0
		tb, bb = 0, hdb
	case "L":
		la, ra = 0, wda
		ta, ba = hda/2, hda/2
		lb, rb = 0, wdb
		tb, bb = hdb/2, hdb/2
	case "C":
		la, ra = wda/2, wda/2
		ta, ba = hda/2, hda/2
		lb, rb = wdb/2, wdb/2
		tb, bb = hdb/2, hdb/2
	case "R":
		la, ra = wda, 0
		ta, ba = hda/2, hda/2
		lb, rb = wdb, 0
		tb, bb = hdb/2, hdb/2
	case "BL":
		la, ra = 0, wda
		ta, ba = hda, 0
		lb, rb = 0, wdb
		tb, bb = hdb, 0
	case "B":
		la, ra = wda/2, wda/2
		ta, ba = hda, 0
		lb, rb = wdb/2, wdb/2
		tb, bb = hdb, 0
	case "BR":
		la, ra = wda, 0
		ta, ba = hda, 0
		lb, rb = wdb, 0
		tb, bb = hdb, 0
	}

	imgA, _ = expandPx(imgA, la, ra, ta, ba)
	imgB, _ = expandPx(imgB, lb, rb, tb, bb)

	return blenders.BlendImages(mode, imgA, imgB), nil
}
