package ports

import (
	"context"

	"github.com/popeskul/awesome-messanger/services/message/internal/core/domain"
)

type MessageUseCase interface {
	GetMessages(ctx context.Context, req *domain.GetMessagesRequest) (*domain.GetMessagesResponse, error)
	SendMessage(ctx context.Context, req *domain.SendMessageRequest) (*domain.SendMessageResponse, error)
	StreamMessages(ctx context.Context, req *domain.StreamMessagesRequest) (<-chan *domain.Message, error)
}

type UserCase interface {
	MessageUseCase() MessageUseCase
}
