package handlers

import "net/http"

// Handler holds dependencies for the handlers
type Handler struct {
	// Add any dependencies here
}

// NewHandler creates a new Handler instance
func NewHandler() *Handler {
	return &Handler{}
}

// SendMessageHandler handles sending messages
func (h *Handler) SendMessageHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Send message handler"))
}

// LivenessHandler checks if the service is alive
func (h *Handler) LivenessHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Liveness probe successful for message service"))
}

// ReadinessHandler checks if the service is ready
func (h *Handler) ReadinessHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Readiness probe successful for message service"))
}
