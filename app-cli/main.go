package main

import (
	"bytes"
	"flag"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"

	"github.com/toxyl/flo"
	"github.com/toxyl/pxp/language"
)

func main() {
	var scriptPath = flag.String("i", "", "Path to the PXP script file")
	var outputPath = flag.String("o", "", "Path to the output image file")
	flag.Parse()

	if *scriptPath == "" {
		fmt.Fprintf(os.Stderr, "ERROR: script path is required\n")
		flag.Usage()
		os.Exit(1)
	}

	if *outputPath == "" {
		fmt.Fprintf(os.Stderr, "ERROR: output path is required\n")
		flag.Usage()
		os.Exit(1)
	}

	scriptFile := flo.File(*scriptPath)
	if !scriptFile.Exists() {
		fmt.Fprintf(os.Stderr, "ERROR: script file does not exist: %s\n", *scriptPath)
		os.Exit(1)
	}

	script := scriptFile.AsString()
	baseDir := filepath.Dir(*scriptPath)

	res, err := language.New().Run(string(script), baseDir, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %s\n", err.Error())
		os.Exit(1)
	}

	if res == nil {
		fmt.Fprintf(os.Stderr, "ERROR: No result returned\n")
		os.Exit(1)
	}

	// Try to get the image from the result
	var img image.Image
	switch val := res.Value().(type) {
	case *image.RGBA:
		img = val
	case *image.NRGBA:
		img = val
	case *image.RGBA64:
		img = val
	case *image.NRGBA64:
		img = val
	default:
		fmt.Fprintf(os.Stderr, "ERROR: Unexpected result type: %T\n", res.Value())
		os.Exit(1)
	}

	var buf bytes.Buffer
	ext := strings.ToLower(filepath.Ext(*outputPath))
	switch ext {
	case ".png":
		if err := png.Encode(&buf, img); err != nil {
			fmt.Fprintf(os.Stderr, "ERROR: Failed to encode PNG: %v\n", err)
			os.Exit(1)
		}
	case ".jpg", ".jpeg":
		if err := jpeg.Encode(&buf, img, &jpeg.Options{Quality: 100}); err != nil {
			fmt.Fprintf(os.Stderr, "ERROR: Failed to encode JPEG: %v\n", err)
			os.Exit(1)
		}
	default:
		fmt.Fprintf(os.Stderr, "ERROR: Unsupported file format: %s (supported: .png, .jpg, .jpeg)\n", ext)
		os.Exit(1)
	}

	outputFile := flo.File(*outputPath)
	if err := outputFile.Mkparent(0755); err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: Failed to create output directory: %v\n", err)
		os.Exit(1)
	}

	if err := outputFile.StoreBytes(buf.Bytes()); err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: Failed to save image: %v\n", err)
		os.Exit(1)
	}
}
