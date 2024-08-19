package handlers

import (
	"context"

	"github.com/popeskul/awesome-messanger/services/notification/internal/models"
	"github.com/popeskul/awesome-messanger/services/notification/internal/services"
	"github.com/popeskul/awesome-messanger/services/notification/pb/proto"
)

type Services interface {
	NotificationService() services.NotificationServiceI
}

type Validator interface {
	Struct(interface{}) error
}

type Handler struct {
	proto.UnimplementedNotificationServiceServer
	services  Services
	validator Validator
}

func NewHandler(services Services, validator Validator) *Handler {
	return &Handler{
		services:  services,
		validator: validator,
	}
}

func (h *Handler) SendNotification(ctx context.Context, req *proto.SendNotificationRequest) (*proto.SendNotificationResponse, error) {
	model := &models.SendNotificationRequest{
		RecipientId: req.GetRecipientId(),
		Message:     req.GetMessage(),
	}

	if err := h.validator.Struct(model); err != nil {
		return nil, err
	}

	err := h.services.NotificationService().SendNotification(ctx, model)
	if err != nil {
		return nil, err
	}

	return &proto.SendNotificationResponse{}, nil
}
