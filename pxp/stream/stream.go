package stream

import (
	"crypto/sha256"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/toxyl/flo"
	"github.com/toxyl/pxp/pxp"
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
	mu      sync.RWMutex
	cfg     Config
	dir     *flo.DirObj
	latest  *flo.FileObj
	stopCh  chan struct{}
	running bool
	busy    bool
}

// New creates a new Stream instance. The stream writes into
// os.TempDir()/pxp/<sha256(Name)>/latest.png and overwrites that file on each
// successful render.
func New(cfg Config) *Stream {
	hash := fmt.Sprintf("%x", sha256.Sum256([]byte(cfg.Name)))
	base := flo.File(os.TempDir()).Dir("pxp").Dir(hash)
	return &Stream{cfg: cfg, dir: base}
}

// Start begins the rendering loop. It performs an immediate render and then
// continues at the configured interval. Calling Start when already running is a
// no-op and returns nil.
func (s *Stream) Start() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.running {
		return nil
	}

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
	s.stopCh = make(chan struct{})
	s.running = true

	go s.loop()
	return nil
}

// Stop signals the rendering loop to stop. It is safe to call Stop when the
// stream is not running.
func (s *Stream) Stop() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if !s.running {
		return nil
	}

	close(s.stopCh)
	s.running = false
	return nil
}

func (s *Stream) loop() {
	// Immediate first run
	s.render()

	ticker := time.NewTicker(s.cfg.Interval)
	defer ticker.Stop()

	for {
		select {
		case <-s.stopCh:
			return
		case <-ticker.C:
			s.render()
		}
	}
}

func (s *Stream) render() {
	// Prevent overlapping renders
	s.mu.Lock()
	if s.busy {
		s.mu.Unlock()
		return
	}
	s.busy = true
	s.mu.Unlock()
	defer func() {
		s.mu.Lock()
		s.busy = false
		s.mu.Unlock()
	}()

	now := time.Now()
	script := s.cfg.ScriptFn(now)
	if script = strings.TrimSpace(script); script == "" {
		// skip empty scripts
		return
	}
	script = ensureImgPrefix(script)

	// Render to latest.png
	if err := pxp.RenderToFile(script, s.latest.Path(), 0, 0); err != nil {
		// rendering failed, skip
		return
	}

	// Invoke hook if present
	if s.cfg.OnImage != nil {
		// Guard against panics in user hook to keep the loop alive
		func() {
			defer func() { _ = recover() }()
			s.cfg.OnImage(s.latest)
		}()
	}
}
