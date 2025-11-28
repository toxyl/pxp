package fonts

import (
	_ "embed"
)

//go:embed pixel-operator.yaml
var bytesPixelOperator []byte

var (
	PixelOperator, _ = LoadBitmapFontFromBytes(bytesPixelOperator)
)
