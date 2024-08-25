package usecases

import (
	"context"

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
	s.logger.Info("Sending notification to %s", req.RecipientId)

	return &domain.SendNotificationResponse{
		Success: true,
	}, nil
}
