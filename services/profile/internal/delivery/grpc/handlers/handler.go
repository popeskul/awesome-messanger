package handlers

import (
	"github.com/popeskul/awesome-messanger/services/profile/internal/core/ports"
	"github.com/popeskul/awesome-messanger/services/profile/pkg/api/health"
	"github.com/popeskul/awesome-messanger/services/profile/pkg/api/profile"
)

type Handler struct {
	profile.ProfileServiceServer
	health.HealthServiceServer
}

func NewHandler(
	profileHandler profile.ProfileServiceServer,
	healthHandler health.HealthServiceServer,
) ports.ServicesServer {
	return &Handler{
		ProfileServiceServer: profileHandler,
		HealthServiceServer:  healthHandler,
	}
}
