package usecases

import (
	"github.com/popeskul/awesome-messanger/services/message/internal/core/ports"
)

type service struct {
	messageUseCase ports.MessageUseCase
}

func NewUseCase(useCase ports.MessageUseCase) ports.UserCase {
	return &service{
		messageUseCase: useCase,
	}
}

func (s *service) MessageUseCase() ports.MessageUseCase {
	return s.MessageUseCase()
}
