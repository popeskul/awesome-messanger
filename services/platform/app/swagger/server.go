package swagger

import (
	"context"
	"net/http"

	"github.com/popeskul/awesome-messanger/services/platform/app/ports"
)

type Server struct {
	address string
	handler http.HandlerFunc
	server  *http.Server
}

func NewSwaggerServer(address string, handler http.HandlerFunc) ports.SwaggerServer {
	return &Server{
		address: address,
		handler: handler,
	}
}

func (s *Server) Run() error {
	s.server = &http.Server{
		Addr:    s.address,
		Handler: s.handler,
	}
	return s.server.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	if s.server != nil {
		return s.server.Shutdown(ctx)
	}
	return nil
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.handler.ServeHTTP(w, r)
}
