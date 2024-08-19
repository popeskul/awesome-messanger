package http

import (
	"context"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/popeskul/awesome-messanger/services/profile/pkg/api/health"
	"github.com/popeskul/awesome-messanger/services/profile/pkg/api/profile"
)

type GatewayServer struct {
	httpMux *runtime.ServeMux
}

type ServicesServer interface {
	profile.ProfileServiceServer
	health.HealthServiceServer
}

func NewGatewayServer(serviceServer ServicesServer) *GatewayServer {
	httpMux := runtime.NewServeMux()

	err := profile.RegisterProfileServiceHandlerServer(context.Background(), httpMux, serviceServer)
	if err != nil {
		log.Fatalf("Failed to register ProfileService gateway: %v", err)
	}

	err = health.RegisterHealthServiceHandlerServer(context.Background(), httpMux, serviceServer)
	if err != nil {
		log.Fatalf("Failed to register HealthService gateway: %v", err)
	}

	return &GatewayServer{
		httpMux: httpMux,
	}
}

func (s *GatewayServer) ListenAndServe(address string) error {
	log.Printf("Starting HTTP server on %s", address)
	return http.ListenAndServe(address, s.httpMux)
}
