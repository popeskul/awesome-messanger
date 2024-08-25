package grpc_gateway

import (
	"context"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rs/cors"

	"github.com/popeskul/awesome-messanger/services/auth/internal/delivery/grpc/grpc"
	"github.com/popeskul/awesome-messanger/services/auth/pkg/api/auth"
	"github.com/popeskul/awesome-messanger/services/auth/pkg/api/health"
)

type GatewayServer struct {
	handler http.Handler
	server  *http.Server
}

func NewGatewayServer(serviceServer *grpc.GrpcServer) *GatewayServer {
	httpMux := runtime.NewServeMux()

	err := auth.RegisterAuthServiceHandlerServer(context.Background(), httpMux, serviceServer.ServicesServer)
	if err != nil {
		log.Fatalf("Failed to register AuthService gateway: %v", err)
	}

	err = health.RegisterHealthServiceHandlerServer(context.Background(), httpMux, serviceServer.ServicesServer)
	if err != nil {
		log.Fatalf("Failed to register HealthService gateway: %v", err)
	}

	mux := http.NewServeMux()
	mux.Handle("/v1/", httpMux)
	fs := http.FileServer(http.Dir("swagger"))
	mux.Handle("/swagger/", http.StripPrefix("/swagger/", fs))

	corsWrapper := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:8001", "http://localhost:8000"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type"},
	})

	handler := corsWrapper.Handler(mux)

	return &GatewayServer{
		handler: handler,
	}
}

func (s *GatewayServer) ListenAndServe(address string) error {
	s.server = &http.Server{
		Addr: address,
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Printf("Received request: %s %s", r.Method, r.URL.Path)

			if r.Method == "OPTIONS" {
				w.WriteHeader(http.StatusOK)
				return
			}

			s.handler.ServeHTTP(w, r)
		}),
	}

	log.Printf("Starting HTTP server on %s", address)
	return s.server.ListenAndServe()
}

func (s *GatewayServer) Shutdown(ctx context.Context) error {
	log.Printf("Shutting down HTTP server")
	return s.server.Shutdown(ctx)
}
