package language

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gen2brain/heic"
	"github.com/rwcarlsen/goexif/exif"
	"github.com/tidwall/gjson"
	"github.com/toxyl/flo"
)

func loadFile(filePath string) (string, []byte, error) {
	filePath = strings.TrimSpace(filePath)
	isRemote := strings.HasPrefix(filePath, "http://") || strings.HasPrefix(filePath, "https://")

	var data []byte
	if isRemote {
		resp, err := http.Get(filePath)
		if err != nil {
			return "", nil, fmt.Errorf("failed to download file '%s': %s", filePath, err.Error())
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			return "", nil, fmt.Errorf("failed to download file: status %s", resp.Status)
		}

		data, err = io.ReadAll(resp.Body)
		if err != nil {
			return "", nil, fmt.Errorf("failed to read body: %s", err.Error())
		}
		hash := fmt.Sprintf("%x", sha256.Sum256([]byte(filePath)))
		filePath = flo.File(os.TempDir()).Dir("pxp").Dir(hash).File(filepath.Base(filePath)).Path()
		f := flo.File(filePath)
		_ = f.Mkparent(0755)
		if err = f.StoreBytes(data); err != nil {
			return "", nil, fmt.Errorf("could not store downloaded file: %s", err.Error())
		}
	}

	return filePath, flo.File(filePath).AsBytes(), nil
}

// @Name: load
// @Desc: Loads an image
// @Param:      path    - -   -   Path to the image
// @Returns:    result  - -   -   The loaded image
func load(path string) (any, error) {
	path = strings.TrimSpace(path)
	var nrgba *image.NRGBA64

	// // Check if the image is in cache
	// if cachedImg, found := ImagesCache.Get(path); found {
	// 	ImagesCache.UpdateTimestamp(path)
	// 	return cachedImg, nil
	// }
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

	// defer func() {
	// 	if nrgba != nil {
	// 		// Store the result in cache before returning
	// 		ImagesCache.Put(localPath, nrgba)
	// 	}
	// }()

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
// @Param:      path    - -   -    Path where to save
func save(img *image.NRGBA64, path string) (any, error) {
	tmp := flo.File(path + ".part")
	w := bytes.Buffer{}
	png.Encode(&w, img) // note that we use NRGBA because storing PNGs is much faster that way
	if err := tmp.StoreBytes(w.Bytes()); err != nil {
		return img, err
	}
	return img, os.Rename(tmp.Path(), path)
}
