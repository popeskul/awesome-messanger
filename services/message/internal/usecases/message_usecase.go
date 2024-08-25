package usecases

import (
	"context"
	"time"

	"github.com/popeskul/awesome-messanger/services/message/internal/core/domain"
	"github.com/popeskul/awesome-messanger/services/message/internal/core/ports"
)

type messageService struct {
	logger ports.Logger
}

func NewMessageUseCase(logger ports.Logger) ports.MessageUseCase {
	return &messageService{
		logger: logger,
	}
}

func (s *messageService) GetMessages(ctx context.Context, req *domain.GetMessagesRequest) ([]*domain.Message, error) {
	s.logger.Info("Getting messages for chat %s", req.ChatId)

	return []*domain.Message{
		{
			SenderId:  "1",
			Content:   "Hello",
			Timestamp: time.Now().Unix(),
		},
	}, nil
}

func (s *messageService) SendMessage(ctx context.Context, req *domain.SendMessageRequest) error {
	s.logger.Info("Sending message from %s to %s", req.SenderId, req.RecipientId)

	return nil
}
