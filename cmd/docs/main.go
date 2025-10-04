package main

import (
	"fmt"

	"github.com/toxyl/flo"
	"github.com/toxyl/pxp/language"
)

func main() {
	if err := language.ExportToVSIX("pxp.vsix"); err != nil {
		fmt.Printf("Failed to export VSCode extension: %v", err)
	}
	if err := flo.File("docs/pxp.md").StoreString(language.DocMarkdown()); err != nil {
		fmt.Printf("Failed to create Markdown doc: %v", err)
	}
	if err := flo.File("docs/pxp.html").StoreString(language.DocHTML()); err != nil {
		fmt.Printf("Failed to create HTML doc: %v", err)
	}
	if err := flo.File("docs/pxp.txt").StoreString(language.DocText()); err != nil {
		fmt.Printf("Failed to create text doc: %v", err)
	}
}
