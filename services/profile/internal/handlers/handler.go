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

// ProfileHandler handles profile requests
func (h *Handler) ProfileHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Profile handler for profile service"))
}

// LivenessHandler checks if the service is alive
func (h *Handler) LivenessHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Liveness probe successful for profile service"))
}

// ReadinessHandler checks if the service is ready
func (h *Handler) ReadinessHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Readiness probe successful for profile service"))
}
