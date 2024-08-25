package handlers

import (
	"github.com/popeskul/awesome-messanger/services/auth/internal/delivery/grpc/grpc"
	"github.com/popeskul/awesome-messanger/services/auth/pkg/api/auth"
	"github.com/popeskul/awesome-messanger/services/auth/pkg/api/health"
)

func NewHandler(
	authHandler auth.AuthServiceServer,
	healthHandler health.HealthServiceServer,
) grpc.ServicesServer {
	return &struct {
		auth.AuthServiceServer
		health.HealthServiceServer
	}{
		AuthServiceServer:   authHandler,
		HealthServiceServer: healthHandler,
	}
}
