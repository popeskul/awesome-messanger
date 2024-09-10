package usecases

import (
	"context"
	"errors"
	"strings"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/popeskul/awesome-messanger/services/message/internal/core/domain"
	"github.com/popeskul/awesome-messanger/services/message/internal/core/ports"
)

type messageUsecase struct {
	logger ports.Logger
}

func NewMessageUseCase(logger ports.Logger) ports.MessageUseCase {
	return &messageUsecase{
		logger: logger,
	}
}

func (s *messageUsecase) GetMessages(ctx context.Context, req *domain.GetMessagesRequest) (*domain.GetMessagesResponse, error) {
	if strings.Trim(req.ChatId, " ") == "" {
		return nil, errors.New("invalid request: ChatId must be non-empty")
	}

	s.logger.Info("Getting messages for chat %s", req.ChatId)

	messages := []*domain.Message{
		{
			Id:        "1",
			ChatId:    req.ChatId,
			SenderId:  "user1",
			Content:   "Hello",
			Timestamp: timestamppb.Now(),
		},
	}

	return &domain.GetMessagesResponse{
		Messages: messages,
		HasMore:  false,
	}, nil
}

func (s *messageUsecase) SendMessage(ctx context.Context, req *domain.SendMessageRequest) (*domain.SendMessageResponse, error) {
	if req.ChatId == "" || req.SenderId == "" || req.Content == "" {
		return nil, errors.New("invalid request: all fields must be non-empty")
	}

	s.logger.Info("Sending message from %s in chat %s", req.SenderId, req.ChatId)

	message := &domain.Message{
		Id:        "new_message_id",
		ChatId:    req.ChatId,
		SenderId:  req.SenderId,
		Content:   req.Content,
		Timestamp: timestamppb.Now(),
	}

	return &domain.SendMessageResponse{
		Message: message,
	}, nil
}

func (s *messageUsecase) StreamMessages(ctx context.Context, req *domain.StreamMessagesRequest) (<-chan *domain.Message, error) {
	if strings.Trim(req.ChatId, " ") == "" {
		return nil, errors.New("invalid request: ChatId must be non-empty")
	}

	s.logger.Info("Streaming messages for chat %s", req.ChatId)

	messageChan := make(chan *domain.Message)

	go func() {
		defer close(messageChan)

		// Here should be a real logic of streaming messages
		// For example, we can send a message every 5 seconds
		for {
			select {
			case <-ctx.Done():
				return
			case <-time.After(5 * time.Second):
				messageChan <- &domain.Message{
					Id:        "streamed_message_id",
					ChatId:    req.ChatId,
					SenderId:  "user1",
					Content:   "Streamed message",
					Timestamp: timestamppb.Now(),
				}
			}
		}
	}()

	return messageChan, nil
}
