package handlers

import (
	"github.com/popeskul/awesome-messanger/services/notification/internal/core/ports"
	"github.com/popeskul/awesome-messanger/services/notification/pkg/api/health"
	"github.com/popeskul/awesome-messanger/services/notification/pkg/api/notification"
)

type Handler struct {
	notification.NotificationServiceServer
	health.HealthServiceServer
}

func NewHandler(
	messageHandler notification.NotificationServiceServer,
	healthHandler health.HealthServiceServer,
) ports.ServicesServer {
	return &Handler{
		NotificationServiceServer: messageHandler,
		HealthServiceServer:       healthHandler,
	}
}
