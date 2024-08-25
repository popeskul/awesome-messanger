package handlers

import (
	"context"

	"github.com/popeskul/awesome-messanger/services/message/internal/core/domain"
	"github.com/popeskul/awesome-messanger/services/message/internal/core/ports"
	"github.com/popeskul/awesome-messanger/services/message/pkg/api/message"
	"github.com/popeskul/awesome-messanger/services/message/proto/api/grpcutils"
)

type MessageHandler struct {
	message.UnimplementedMessageServiceServer

	messageUseCase ports.MessageUseCase
	validator      ports.Validator
	logger         ports.Logger
}

func NewMessageHandler(
	messageUseCase ports.MessageUseCase,
	validator ports.Validator,
	logger ports.Logger,
) ports.MessageHandler {
	return &MessageHandler{
		messageUseCase: messageUseCase,
		validator:      validator,
		logger:         logger,
	}
}

func (h *MessageHandler) GetMessages(ctx context.Context, req *message.GetMessagesRequest) (*message.GetMessagesResponse, error) {
	if err := h.validator.Validate(req); err != nil {
		return nil, grpcutils.RPCValidationError(err)
	}

	input := &domain.GetMessagesRequest{
		ChatId: req.GetChatId(),
	}

	messages, err := h.messageUseCase.GetMessages(ctx, input)
	if err != nil {
		return nil, err
	}

	var pbMessages []*message.Message
	for _, m := range messages {
		pbMessages = append(pbMessages, m.ConvertToProto())
	}

	return &message.GetMessagesResponse{Messages: pbMessages}, nil
}

func (h *MessageHandler) SendMessage(ctx context.Context, req *message.SendMessageRequest) (*message.SendMessageResponse, error) {
	input := &domain.SendMessageRequest{
		SenderId:    req.GetSenderId(),
		RecipientId: req.GetRecipientId(),
		Content:     req.GetContent(),
	}

	err := h.messageUseCase.SendMessage(ctx, input)
	if err != nil {
		return nil, err
	}

	return &message.SendMessageResponse{
		Success: true,
	}, nil
}
