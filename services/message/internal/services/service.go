package services

import (
	"context"

	"github.com/popeskul/awesome-messanger/services/message/internal/models"
)

type Logger interface {
	Printf(format string, v ...interface{})
}

type MessageServiceI interface {
	GetMessages(ctx context.Context, req *models.GetMessagesRequest) ([]*models.Message, error)
	SendMessage(ctx context.Context, req *models.SendMessageRequest) error
}

type IServices interface {
	MessageService() MessageServiceI
}

type Service struct {
	messageService MessageServiceI
}

func NewService(messageService MessageServiceI) IServices {
	return &Service{
		messageService: messageService,
	}
}

func (s *Service) MessageService() MessageServiceI {
	return s.messageService
}
