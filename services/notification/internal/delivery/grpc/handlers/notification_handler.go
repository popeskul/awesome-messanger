package handlers

import (
	"context"
	"fmt"

	"github.com/popeskul/awesome-messanger/services/notification/internal/core/domain"
	"github.com/popeskul/awesome-messanger/services/notification/internal/core/ports"
	"github.com/popeskul/awesome-messanger/services/notification/pkg/api/notification"
	"github.com/popeskul/awesome-messanger/services/notification/proto/api/grpcutils"
)

type NotificationHandler struct {
	notification.UnimplementedNotificationServiceServer

	messageUseCase ports.NotificationUseCase
	validator      ports.Validator
	logger         ports.Logger
}

func NewMessageHandler(
	messageUseCase ports.NotificationUseCase,
	validator ports.Validator,
	logger ports.Logger,
) ports.NotificationHandler {
	return &NotificationHandler{
		messageUseCase: messageUseCase,
		validator:      validator,
		logger:         logger,
	}
}

func (h *NotificationHandler) SendNotification(ctx context.Context, req *notification.SendNotificationRequest) (*notification.SendNotificationResponse, error) {
	if err := h.validator.Validate(req); err != nil {
		return nil, grpcutils.RPCValidationError(err)
	}

	input := &domain.SendNotificationRequest{
		Message:     req.GetMessage(),
		RecipientId: req.GetRecipientId(),
	}

	res, err := h.messageUseCase.SendNotification(ctx, input)
	if err != nil {
		return nil, fmt.Errorf("failed to send notification: %w", err)
	}

	return &notification.SendNotificationResponse{
		Success: res.Success,
	}, nil
}
