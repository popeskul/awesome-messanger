package usecases

import "github.com/popeskul/awesome-messanger/services/notification/internal/core/ports"

type service struct {
	notificationUseCase ports.NotificationUseCase
}

func NewUseCase(useCase ports.NotificationUseCase) ports.UserCase {
	return &service{
		notificationUseCase: useCase,
	}
}

func (s *service) MessageUseCase() ports.NotificationUseCase {
	return s.MessageUseCase()
}
