package language

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"strings"

	"github.com/gen2brain/heic"
	"github.com/rwcarlsen/goexif/exif"
	"github.com/tidwall/gjson"
	"github.com/toxyl/math"
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

// @Name: load
// @Desc: Loads an image
// @Param:      path    - -   -   Path to the image
// @Returns:    result  - -   -   The loaded image
func load(path string) (any, error) {
	path = strings.TrimSpace(path)
	var nrgba *image.NRGBA64

	// Check if the image is in cache
	if cachedImg, found := ImagesCache.Get(path); found {
		ImagesCache.UpdateTimestamp(path)
		return cachedImg, nil
	}
	localPath, data, err := loadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to load image: %v", err)
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
	case "gif":
		img, err = gif.Decode(reader)
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
	nrgba = IFromBounds(bounds)
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			nrgba.Set(x, y, img.At(x, y))
		}
	}

	// Get EXIF
	f, err := os.Open(localPath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	defer func() {
		if nrgba != nil {
			// Store the result in cache before returning
			ImagesCache.Put(localPath, nrgba)
		}
	}()

	// Apple is a kid with special needs, as usual ...
	metaData, err := exif.Decode(f)
	if err == nil {
		jsonByte, err := metaData.MarshalJSON()
		if err == nil {
			jsonString := string(jsonByte)
			or := gjson.Get(jsonString, "Orientation").Array()
			if len(or) > 0 {
				orientation := or[0]
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

// @Name: plot-data
// @Desc: Renders a chart from CSV data by plotting selected columns with specified colors
// @Param:      width   - - -   Chart width in pixels
// @Param:      height  - - -   Chart height in pixels
// @Param:      data    - - -   2D array of data from CSV
// @Param:      columns - - -   Array of column indices to plot
// @Param:      colors  - - -   Array of colors for each column
// @Returns:    result  - - -   The chart image
func plotData(width, height int, data [][]float64, columns []any, colors []any) (*image.NRGBA64, error) {
	if width <= 0 || height <= 0 {
		return nil, fmt.Errorf("width and height must be greater than 0")
	}

	if len(data) == 0 {
		return IC(width, height, color.RGBA64{0, 0, 0, 0}), nil
	}

	if len(columns) != len(colors) {
		return nil, fmt.Errorf("columns and colors arrays must have the same length")
	}

	if len(columns) == 0 {
		return IC(width, height, color.RGBA64{0, 0, 0, 0}), nil
	}

	columnsInt := make([]int, len(columns))
	for i, v := range columns {
		if intVal, ok := v.(int); ok {
			columnsInt[i] = intVal
		} else if floatVal, ok := v.(float64); ok {
			columnsInt[i] = int(floatVal)
		} else {
			return nil, fmt.Errorf("invalid column index type at position %d", i)
		}
	}

	colorsRGBA64 := make([]color.RGBA64, len(colors))
	for i, v := range colors {
		if col, ok := v.(color.RGBA64); ok {
			colorsRGBA64[i] = col
		} else {
			return nil, fmt.Errorf("invalid color type at position %d", i)
		}
	}

	numDatapoints := len(data)
	if numDatapoints == 0 {
		return IC(width, height, color.RGBA64{0, 0, 0, 0}), nil
	}

	img := IC(width, height, color.RGBA64{0, 0, 0, 0})
	xScale := float64(width) / float64(numDatapoints)

	for colIdx, colIndex := range columnsInt {
		if colIndex < 0 {
			continue
		}

		colColor := colorsRGBA64[colIdx]
		nrgbaColor := color.NRGBA64{R: colColor.R, G: colColor.G, B: colColor.B, A: colColor.A}

		for i := 0; i < numDatapoints; i++ {
			if colIndex >= len(data[i]) {
				continue
			}

			yValue := data[i][colIndex]
			if math.IsNaN(yValue) {
				continue
			}

			yNormalized := math.Clamp(yValue, 0.0, 1.0)
			yPixel := height - 1 - int(yNormalized*float64(height))
			xPixel := int(float64(i) * xScale)

			if xPixel < 0 {
				xPixel = 0
			}
			if xPixel >= width {
				xPixel = width - 1
			}
			if yPixel < 0 {
				yPixel = 0
			}
			if yPixel >= height {
				yPixel = height - 1
			}

			img.Set(xPixel, yPixel, nrgbaColor)
		}
	}

	return img, nil
}
