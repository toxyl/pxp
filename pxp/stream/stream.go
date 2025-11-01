package stream

import (
	"crypto/sha256"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/toxyl/flo"
	"github.com/toxyl/pxp/pxp"
	"github.com/toxyl/scheduler"
)

func ensureImgPrefix(script string) string {
	lines := strings.Split(strings.TrimSpace(script), "\n")
	i := len(lines) - 1
	for i >= 0 && strings.TrimSpace(lines[i]) == "" {
		i--
	}
	if i >= 0 {
		last := strings.TrimSpace(lines[i])
		if !strings.HasPrefix(last, "img:") {
			lines[i] = "img:" + last
		}
	}
	return strings.Join(lines, "\n")
}

// Stream repeatedly renders scripts produced by ScriptFn to a single file in
// the OS temp directory and invokes a hook after each successful render.
type Stream struct {
	cfg    Config
	dir    *flo.DirObj
	latest *flo.FileObj
}

// New creates a new Stream instance. The stream writes into
// os.TempDir()/pxp/<sha256(Name)>/latest.png and overwrites that file on each
// successful render.
func New(cfg Config) *Stream {
	hash := fmt.Sprintf("%x", sha256.Sum256([]byte(cfg.Name)))
	base := flo.File(os.TempDir()).Dir("pxp").Dir(hash)
	_ = base.Mkdir(0755)
	return &Stream{cfg: cfg, dir: base}
}

// Start begins the rendering loop. It performs an immediate render and then
// continues at the configured interval. Calling Start when already running is a
// no-op and returns nil.
func (s *Stream) Start() error {
	if s.cfg.ScriptFn == nil {
		return fmt.Errorf("scriptFn must be provided")
	}
	if s.cfg.Interval <= 0 {
		return fmt.Errorf("interval must be greater than zero")
	}

	if err := s.dir.Mkdir(0755); err != nil {
		return err
	}

	s.latest = s.dir.File("latest.png")

	scheduler.Run(s.cfg.Interval, 0, func() (stop bool) {
		script := ensureImgPrefix(strings.TrimSpace(s.cfg.ScriptFn(time.Now())))
		if script == "" {
			// skip empty scripts
			return false
		}

		// Render to latest.png
		if err := pxp.RenderToFile(script, s.latest.Path(), 0, 0); err != nil {
			// rendering failed, skip
			return false
		}

		// Invoke hook if present
		if s.cfg.OnImage != nil {
			// Guard against panics in user hook to keep the loop alive
			func() {
				defer func() { _ = recover() }()
				s.cfg.OnImage(s.latest)
			}()
		}
		return false
	}, nil)
	return nil
}
