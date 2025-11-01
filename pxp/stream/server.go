package stream

import (
	"context"
	"net/http"
	"time"
)

// StartServer starts a single HTTP server that serves multiple streams on
// different routes. The map key is the route (e.g. "/cam1.png"), the value is
// the stream to serve at that route.
func StartServer(addr string, routeToStream map[string]*Stream, onError func(err error)) (closer func()) {
	srv, err := newServer(addr, routeToStream)
	if err != nil {
		if onError != nil {
			onError(err)
		}
		return func() {}
	}

	shutdownCh := make(chan struct{})
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			if onError != nil {
				onError(err)
			}
		}
		close(shutdownCh)
	}()

	return func() {
		// graceful shutdown with timeout
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		_ = srv.Shutdown(ctx)
		<-shutdownCh
	}
}

// newServer constructs an *http.Server that serves multiple streams under
// different routes on a single HTTP server.
func newServer(addr string, routeToStream map[string]*Stream) (*http.Server, error) {
	mux := http.NewServeMux()

	for route, s := range routeToStream {
		r := route
		st := s
		if r == "" {
			r = "/"
		}

		if err := st.Start(); err != nil {
			return nil, err
		}

		mux.HandleFunc(r, func(w http.ResponseWriter, r *http.Request) {
			latest := st.dir.File("latest.png")
			var data []byte
			if err := latest.LoadBytes(&data); err != nil || len(data) == 0 {
				http.Error(w, "no image available", http.StatusNotFound)
				return
			}
			w.Header().Set("Content-Type", "image/png")
			w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
			w.Header().Set("Pragma", "no-cache")
			w.Header().Set("Expires", "0")
			_, _ = w.Write(data)
		})
	}

	return &http.Server{Addr: addr, Handler: mux}, nil
}
