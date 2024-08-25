package ports

import (
	"context"

	"github.com/popeskul/awesome-messanger/services/notification/internal/core/domain"
)

type NotificationUseCase interface {
	SendNotification(ctx context.Context, req *domain.SendNotificationRequest) (*domain.SendNotificationResponse, error)
}

type UserCase interface {
	MessageUseCase() NotificationUseCase
}
