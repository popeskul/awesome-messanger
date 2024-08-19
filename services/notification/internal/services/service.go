package services

import (
	"context"

	"github.com/popeskul/awesome-messanger/services/notification/internal/models"
)

type Logger interface {
	Printf(format string, v ...interface{})
}

type NotificationServiceI interface {
	SendNotification(ctx context.Context, req *models.SendNotificationRequest) error
}

type IServices interface {
	NotificationService() NotificationServiceI
}

type Service struct {
	notificator NotificationServiceI
}

func NewService(notificationService NotificationServiceI) IServices {
	return &Service{
		notificator: notificationService,
	}
}

func (s *Service) NotificationService() NotificationServiceI {
	return s.notificator
}
