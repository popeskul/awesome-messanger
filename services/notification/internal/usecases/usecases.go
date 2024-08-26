package usecases

import "github.com/popeskul/awesome-messanger/services/notification/internal/core/ports"

type service struct {
	notificationUseCase ports.NotificationUseCase
}

func NewUseCase(useCase ports.NotificationUseCase) ports.UseCase {
	return &service{
		notificationUseCase: useCase,
	}
}

func (s *service) Notification() ports.NotificationUseCase {
	return s.notificationUseCase
}
