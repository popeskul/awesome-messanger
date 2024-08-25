package ports

import (
	"context"
	"github.com/popeskul/awesome-messanger/services/message/internal/core/domain"
)

type MessageUseCase interface {
	GetMessages(ctx context.Context, req *domain.GetMessagesRequest) ([]*domain.Message, error)
	SendMessage(ctx context.Context, req *domain.SendMessageRequest) error
}

type UserCase interface {
	MessageUseCase() MessageUseCase
}
