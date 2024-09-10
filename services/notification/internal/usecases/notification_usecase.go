package usecases

import (
	"context"
	"errors"
	"strings"

	"github.com/popeskul/awesome-messanger/services/notification/internal/core/domain"
	"github.com/popeskul/awesome-messanger/services/notification/internal/core/ports"
)

type notificationUseCase struct {
	logger ports.Logger
}

func NewNotificationUseCase(logger ports.Logger) ports.NotificationUseCase {
	return &notificationUseCase{
		logger: logger,
	}
}

func (s *notificationUseCase) SendNotification(
	ctx context.Context,
	req *domain.SendNotificationRequest,
) (*domain.SendNotificationResponse, error) {
	if strings.Trim(req.RecipientId, " ") == "" || strings.Trim(req.Message, " ") == "" {
		return nil, errors.New("recipient id is required")
	}

	s.logger.Info("Sending notification to %s", req.RecipientId)

	return &domain.SendNotificationResponse{
		Success: true,
	}, nil
}
