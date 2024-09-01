package server

import (
	"net/http"
	"time"

	"github.com/popeskul/awesome-messanger/services/platform/app/ports"
)

func NewHTTPServer(addr string, logger ports.Logger) (ports.HTTPServer, error) {
	mux := http.NewServeMux()

	// Add your HTTP routes here
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	// Add logging middleware
	loggingMiddleware := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			next.ServeHTTP(w, r)
			logger.Info("HTTP request",
				"method", r.Method,
				"path", r.URL.Path,
				"duration", time.Since(start),
			)
		})
	}

	return &http.Server{
		Addr:    addr,
		Handler: loggingMiddleware(mux),
	}, nil
}
