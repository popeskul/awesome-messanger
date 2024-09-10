package server

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5/middleware"

	"github.com/popeskul/awesome-messanger/services/friend/internal/core/ports"
)

func LoggingMiddleware(logger ports.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)

			defer func() {
				logger.Info("HTTP Request",
					"method", r.Method,
					"remote_addr", r.RemoteAddr,
					"url", r.URL.Path,
					"status", ww.Status(),
					"status_text", http.StatusText(ww.Status()),
					"bytes", ww.BytesWritten(),
					"duration", time.Since(start),
				)
			}()

			next.ServeHTTP(ww, r)
		})
	}
}
