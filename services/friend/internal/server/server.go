package server

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/popeskul/awesome-messanger/services/friend/internal/config"
	"github.com/popeskul/awesome-messanger/services/friend/internal/handlers"
)

//go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen -config ../../api/openapi/cfg.yaml ../../api/openapi/api.yaml

type Server struct {
	httpServer *http.Server
	handler    *handlers.Handler
}

func NewServer(cfg *config.Config, handler *handlers.Handler) *Server {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	r.Post("/add-friend", handler.PostAddFriend)
	r.Get("/friends", handler.GetFriends)
	r.Post("/respond-friend-request", handler.PostRespondFriendRequest)
	r.Get("/live", handler.GetLive)
	r.Get("/ready", handler.GetReady)

	return &Server{
		httpServer: &http.Server{
			Addr:    cfg.ServerAddress,
			Handler: r,
		},
		handler: handler,
	}
}

func (s *Server) ListenAndServe() error {
	return s.httpServer.ListenAndServe()
}
