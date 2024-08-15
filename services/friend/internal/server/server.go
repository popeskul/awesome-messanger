package server

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/popeskul/awesome-messanger/services/friend/internal/config"
	"github.com/popeskul/awesome-messanger/services/friend/internal/handlers"
	_ "github.com/popeskul/awesome-messanger/services/friend/swagger"
	"github.com/rs/cors"
)

type Server struct {
	httpServer *http.Server
	handler    http.Handler
}

func NewServer(cfg *config.Config, handler *handlers.Handler) *Server {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	mux := http.NewServeMux()
	mux.Handle("/v1/", r)
	fs := http.FileServer(http.Dir("swagger"))
	mux.Handle("/swagger/", http.StripPrefix("/swagger/", fs))

	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8091"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	}).Handler

	r.Use(corsHandler)

	r.Post("/add-friend", handler.PostAddFriend)
	r.Get("/friends", handler.GetFriends)
	r.Post("/respond-friend-request", handler.PostRespondFriendRequest)
	r.Get("/live", handler.GetLive)
	r.Get("/ready", handler.GetReady)

	return &Server{
		httpServer: &http.Server{
			Addr:    cfg.Server.HttpAddress,
			Handler: r,
		},
		handler: r,
	}
}

func (s *Server) ListenAndServe(address string) error {
	log.Printf("Starting HTTP server on %s", address)
	return http.ListenAndServe(address, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Received request: %s %s", r.Method, r.URL.Path)

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		s.handler.ServeHTTP(w, r)
	}))
}

func (s *Server) Shutdown(ctx context.Context) error {
	log.Println("Shutting down HTTP server")
	return s.httpServer.Shutdown(ctx)
}
