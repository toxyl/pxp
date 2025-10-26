package stream_test

import (
	"fmt"
	"io"
	"net/http"
	"testing"
	"time"

	"github.com/toxyl/pxp/pxp/stream"
)

func fetch(addr, route string) error {
	resp1, err := http.Get("http://localhost" + addr + route)
	if err != nil {
		return fmt.Errorf("GET failed: %v", err)
	}
	defer resp1.Body.Close()
	if resp1.StatusCode != http.StatusOK {
		b, _ := io.ReadAll(resp1.Body)
		return fmt.Errorf("GET status=%d body=%s", resp1.StatusCode, string(b))
	}
	return nil
}

func TestStartServer(t *testing.T) {
	sc1 := stream.Config{
		Name:     "test cam 1",
		Interval: 20 * time.Second,
		ScriptFn: func(time.Time) string {
			return `auto-white-balance(crop-square(polar-to-rectangular(load("https://www.irf.se/alis/allsky/krn/latest_medium.jpeg"))))`
		},
	}
	sc2 := stream.Config{
		Name:     "test cam 2",
		Interval: 20 * time.Second,
		ScriptFn: func(time.Time) string {
			return `auto-white-balance(crop-square(polar-to-rectangular(load("https://www.irf.se/alis/allsky/krn/latest_medium.jpeg"))))`
		},
	}
	addr := ":9123"
	route1, s1 := "/img1.png", stream.New(sc1)
	route2, s2 := "/img2.png", stream.New(sc2)
	closer := stream.StartServer(addr, map[string]*stream.Stream{route1: s1, route2: s2}, func(err error) { t.Errorf("server error: %v", err) })

	// first fetch
	time.Sleep(sc1.Interval) // wait one interval, rendering must have completed after that
	if err := fetch(addr, route1); err != nil {
		t.Fatalf("[0] first GET failed: %v", err)
	}
	if err := fetch(addr, route2); err != nil {
		t.Fatalf("[1] first GET failed: %v", err)
	}
	// second fetch
	time.Sleep(sc1.Interval) // wait another interval, the next render must have completed after that
	if err := fetch(addr, route1); err != nil {
		t.Fatalf("[0] second GET failed: %v", err)
	}
	time.Sleep(sc2.Interval) // wait another interval, the next render must have completed after that
	if err := fetch(addr, route2); err != nil {
		t.Fatalf("[1] second GET failed: %v", err)
	}

	closer() // graceful shutdown
}
