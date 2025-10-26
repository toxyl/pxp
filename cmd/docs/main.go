package main

import (
	"fmt"
	"os"
	"time"

	"github.com/toxyl/flo"
	"github.com/toxyl/pxp/language"
)

func printBlue(message string) {
	fmt.Printf("\033[1;34m%s\033[0m\n", message)
}

func printYellow(message string) {
	fmt.Printf("\033[1;33m%s\033[0m\n", message)
}

func printGreen(message string) {
	fmt.Printf("\033[1;32m%s\033[0m\n", message)
}

func printRed(message string) {
	fmt.Printf("\033[1;31m%s\033[0m\n", message)
}

func dieOnError(err error, msg string) {
	if err != nil {
		printRed(fmt.Sprintf("%s: %v", msg, err))
		os.Exit(1)
	}
}

func main() {
	/////////////////////////////////////////////////////////////////////////////////////////
	printYellow("PixelPipeline Documentation Generator")
	/////////////////////////////////////////////////////////////////////////////////////////
	var (
		dSrc     = flo.Dir(".")     // /src/pixelpipeline/
		dSrcBin  = dSrc.Dir("bin")  // /src/pixelpipeline/bin/
		dSrcDocs = dSrc.Dir("docs") // /src/pixelpipeline/docs/
	)
	/////////////////////////////////////////////////////////////////////////////////////////
	printBlue("Generating docs...")
	/////////////////////////////////////////////////////////////////////////////////////////

	dieOnError(dSrcDocs.File("README.md").StoreString(language.DocMarkdown()), "Failed to create Markdown doc")
	dieOnError(dSrcDocs.File("README.html").StoreString(language.DocHTML()), "Failed to create HTML doc")
	dieOnError(dSrcDocs.File("README.txt").StoreString(language.DocText()), "Failed to create text doc")

	/////////////////////////////////////////////////////////////////////////////////////////
	printBlue("Generating VSCode extension...")
	/////////////////////////////////////////////////////////////////////////////////////////

	time.Sleep(5 * time.Second)
	dieOnError(language.ExportToVSIX(dSrcBin.File("pxp.vsix").Path()), "Failed to export VSCode extension")

	/////////////////////////////////////////////////////////////////////////////////////////
	printGreen("Documentation generated successfully!")
	/////////////////////////////////////////////////////////////////////////////////////////
}
