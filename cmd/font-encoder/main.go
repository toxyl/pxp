package main

import (
	"flag"
	"fmt"

	"github.com/toxyl/flo"
	"github.com/toxyl/pxp/fonts"
)

func main() {
	var (
		fIn  = flag.String("i", flo.Dir("pixel-operator").Path(), "Path to directory with font data")
		fOut = flag.String("o", flo.Dir("pixel-operator").File("font.yaml").Path(), "Path to directory where to store encoded font")
	)
	flag.Parse()

	if err := fonts.MakeBitmapFont(flo.File(*fOut), flo.Dir(*fIn).File("font.charset"), flo.Dir(*fIn).File("font.png")); err != nil {
		fmt.Println("Error creating font:", err)
		return
	}
}
