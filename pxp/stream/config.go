package stream

import (
	"time"

	"github.com/toxyl/flo"
)

// Config defines the stream configuration.
//
// ScriptFn must return a complete PXP script. The last non-empty line will be
// ensured to start with "img:" (if it doesn't already) before rendering.
// OnImage, when provided, is invoked synchronously after each successful render
// with the file object pointing to the latest image on disk.
type Config struct {
	Name     string
	Interval time.Duration
	ScriptFn func(time.Time) string
	OnImage  func(*flo.FileObj)
}
