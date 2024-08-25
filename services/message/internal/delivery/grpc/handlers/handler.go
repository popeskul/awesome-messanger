package handlers

import (
	"github.com/popeskul/awesome-messanger/services/message/internal/core/ports"
	"github.com/popeskul/awesome-messanger/services/message/pkg/api/health"
	"github.com/popeskul/awesome-messanger/services/message/pkg/api/message"
)

type Handler struct {
	message.MessageServiceServer
	health.HealthServiceServer
}

func NewHandler(
	messageHandler message.MessageServiceServer,
	healthHandler health.HealthServiceServer,
) ports.ServicesServer {
	return &Handler{
		MessageServiceServer: messageHandler,
		HealthServiceServer:  healthHandler,
	}
}
