package main

import (
	"flag"
	"fmt"

	"github.com/toxyl/pxp/fonts"
)

func main() {
	var (
		name       = flag.String("font", "pixel-operator", "Font name")
		fOut       = flag.String("o", "image.png", "Output path")
		s          = flag.String("s", "░▒▓█▓▒░ Hi thËre ░▒▓█▓▒░\n ░▒▓█▓▒░ Hi there ░▒▓█▓▒░\n  ░▒▓█▓▒░ Hi there ░▒▓█▓▒░\n   ░▒▓█▓▒░ Hi there ░▒▓█▓▒░", "String to render")
		colText    = flag.String("text", "white", "Text color")
		colOutline = flag.String("outline", "black", "Outline color")
	)
	flag.Parse()

	var font *fonts.BitmapFont
	switch *name {
	case "pixel-operator":
		font = fonts.PixelOperator
	default:
		fmt.Println("Unknown font:", *name)
		return
	}

	if err := font.RenderToFile(*fOut, *s, fonts.Colors[*colText], fonts.Colors[*colOutline]); err != nil {
		fmt.Println("Rendering failed:", err)
		return
	}

}
