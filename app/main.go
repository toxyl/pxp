package main

import (
	"embed"
	"os"

	"github.com/toxyl/pxp/language"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	if len(os.Args) == 2 && os.Args[1] == "shell" {
		language.Shell()
		return
	}
	// Create an instance of the app structure
	app := NewApp()
	lang := &language.Language{}

	// Create application with options
	err := wails.Run(&options.App{
		Title:            "PixelPipeline Studio",
		Width:            1024,
		Height:           768,
		Assets:           assets,
		BackgroundColour: &options.RGBA{R: 27, G: 27, B: 27, A: 1},
		OnStartup:        app.startup,
		Bind: []any{
			app,
			lang,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
