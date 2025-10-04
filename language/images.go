package language

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"

	"github.com/gen2brain/heic"
	"github.com/rwcarlsen/goexif/exif"
	"github.com/tidwall/gjson"
)

var ImagesCache = NewImageCache()

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

	// Check for HEIC
	// HEIC files start with 'ftyp' box which contains 'heic' or 'mif1'
	if len(data) >= 12 && bytes.Equal(data[4:8], []byte("ftyp")) {
		if bytes.Equal(data[8:12], []byte("heic")) || bytes.Equal(data[8:12], []byte("mif1")) {
			return "heic"
		}
	}

	return ""
}

// @Name: load
// @Desc: Loads an image
// @Param:      path    - -   -   Path to the image
// @Returns:    result  - -   -   The loaded image
func load(path string) (any, error) {
	var nrgba *image.NRGBA64

	// Check if the image is in cache
	if cachedImg, found := ImagesCache.Get(path); found {
		ImagesCache.UpdateTimestamp(path)
		return cachedImg, nil
	}

	// Read the entire file
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	// Create a bytes reader for image decoding
	reader := bytes.NewReader(data)

	// Detect image type from magic bytes
	imgType := DetectImageType(data)
	if imgType == "" {
		return nil, fmt.Errorf("unsupported image format or corrupted file")
	}

	// Decode the image based on its type
	var img image.Image
	switch imgType {
	case "png":
		img, err = png.Decode(reader)
	case "jpeg":
		img, err = jpeg.Decode(reader)
	case "heic":
		img, err = heic.Decode(reader)
	default:
		return nil, fmt.Errorf("unsupported image format: %s", imgType)
	}

	if err != nil {
		return nil, fmt.Errorf("failed to decode image: %v", err)
	}

	// Convert to NRGBA format which our functions expect
	bounds := img.Bounds()
	nrgba = image.NewNRGBA64(bounds)
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			nrgba.Set(x, y, img.At(x, y))
		}
	}

	// Get EXIF
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	defer func() {
		if nrgba != nil {
			// Store the result in cache before returning
			ImagesCache.Put(path, nrgba)
		}
	}()

	// Apple is a kid with special needs, as usual ...
	metaData, err := exif.Decode(f)
	if err == nil {
		jsonByte, err := metaData.MarshalJSON()
		if err == nil {
			jsonString := string(jsonByte)
			orientation := gjson.Get(jsonString, "Orientation").Array()[0]
			switch orientation.String() {
			case "1": // 1 = Horizontal (normal) = no change
			case "2": // 2 = Mirror horizontal
				nrgba, err = flipHorizontal(nrgba)
				return nrgba, err
			case "3": // 3 = Rotate 180
				nrgba, err = rotate(nrgba, 180)
				return nrgba, err
			case "4": // 4 = Mirror vertical
				nrgba, err = flipVertical(nrgba)
				return nrgba, err
			case "5": // 5 = Mirror horizontal and rotate 270 CW
				res, err := flipHorizontal(nrgba)
				if err != nil {
					return nil, err
				}
				nrgba, err = rotate(res, 270)
				return nrgba, err
			case "6": // 6 = Rotate 90 CW
				nrgba, err = rotate(nrgba, 90)
				return nrgba, err
			case "7": // 7 = Mirror horizontal and rotate 90 CW
				res, err := flipHorizontal(nrgba)
				if err != nil {
					return nil, err
				}
				nrgba, err = rotate(res, 90)
				return nrgba, err
			case "8": // 8 = Rotate 270 CW
				nrgba, err = rotate(nrgba, 270)
				return nrgba, err
			}
		}
	}

	return nrgba, nil
}

// @Name: save
// @Desc: Saves an image
// @Param:      img     - -   -    The image to save
// @Param:      path     - -   -   Path where to save
func save(img *image.NRGBA64, path string) (any, error) {
	file, err := os.Create(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	png.Encode(file, img) // note that we use NRGBA because storing PNGs is much faster that way
	return img, nil
}
