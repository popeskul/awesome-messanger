package services

import (
	"context"
	"time"

	"github.com/popeskul/awesome-messanger/services/message/internal/models"
)

type MessageService struct {
	logger Logger
}

func NewMessageService(logger Logger) *MessageService {
	return &MessageService{
		logger: logger,
	}
}

func (s *MessageService) GetMessages(ctx context.Context, req *models.GetMessagesRequest) ([]*models.Message, error) {
	s.logger.Printf("Getting messages for chat %s", req.ChatId)

	return []*models.Message{
		{
			SenderId:  "1",
			Content:   "Hello",
			Timestamp: time.Now().Unix(),
		},
	}, nil
}

func (s *MessageService) SendMessage(ctx context.Context, req *models.SendMessageRequest) error {
	s.logger.Printf("Sending message from %s to %s", req.SenderId, req.RecipientId)

	return nil
}
