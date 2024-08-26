package usecases

import (
	"github.com/popeskul/awesome-messanger/services/message/internal/core/ports"
)

type usecase struct {
	messageUseCase ports.MessageUseCase
}

func NewUseCase(messageUseCase ports.MessageUseCase) ports.UserCase {
	return &usecase{
		messageUseCase: messageUseCase,
	}
}

func (s *usecase) MessageUseCase() ports.MessageUseCase {
	return s.messageUseCase
}
