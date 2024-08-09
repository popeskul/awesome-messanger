package services

import (
	"context"

	"github.com/popeskul/awesome-messanger/services/notification/internal/models"
)

type NotificationService struct {
	logger Logger
}

func NewNotificationService(logger Logger) *NotificationService {
	return &NotificationService{
		logger: logger,
	}
}

func (s *NotificationService) SendNotification(ctx context.Context, req *models.SendNotificationRequest) error {
	s.logger.Printf("Sending notification to %s", req.RecipientId)

	return nil
}
