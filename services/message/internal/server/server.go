package server

import (
	"net/http"

	"github.com/popeskul/awesome-messanger/services/message/internal/config"
	"github.com/popeskul/awesome-messanger/services/message/internal/handlers"
)

// Server struct holds server configuration and routes
type Server struct {
	httpServer *http.Server
	handler    *handlers.Handler
}

// NewServer creates a new Server instance
func NewServer(cfg *config.Config) *Server {
	mux := http.NewServeMux()
	h := handlers.NewHandler()

	// Register handlers
	mux.HandleFunc("/message/send", h.SendMessageHandler)
	mux.HandleFunc("/live", h.LivenessHandler)
	mux.HandleFunc("/ready", h.ReadinessHandler)

	return &Server{
		httpServer: &http.Server{
			Addr:    cfg.ServerAddress,
			Handler: mux,
		},
		handler: h,
	}
}

// ListenAndServe starts the HTTP server
func (s *Server) ListenAndServe() error {
	return s.httpServer.ListenAndServe()
}
