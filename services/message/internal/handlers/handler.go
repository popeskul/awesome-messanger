package handlers

import (
	"context"

	"github.com/popeskul/awesome-messanger/services/message/internal/models"
	"github.com/popeskul/awesome-messanger/services/message/internal/services"
	"github.com/popeskul/awesome-messanger/services/message/pb/proto"
)

type Service interface {
	MessageService() services.MessageServiceI
}

type Handler struct {
	proto.UnimplementedMessageServiceServer
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) GetMessages(ctx context.Context, req *proto.GetMessagesRequest) (*proto.GetMessagesResponse, error) {
	input := &models.GetMessagesRequest{
		ChatId: req.GetChatId(),
	}

	messages, err := h.service.MessageService().GetMessages(ctx, input)
	if err != nil {
		return nil, err
	}

	var pbMessages []*proto.Message
	for _, message := range messages {
		pbMessages = append(pbMessages, message.ConvertToProto())
	}

	return &proto.GetMessagesResponse{Messages: pbMessages}, nil
}

func (h *Handler) SendMessage(ctx context.Context, req *proto.SendMessageRequest) (*proto.SendMessageResponse, error) {
	input := &models.SendMessageRequest{
		SenderId:    req.GetSenderId(),
		RecipientId: req.GetRecipientId(),
		Content:     req.GetContent(),
	}

	err := h.service.MessageService().SendMessage(ctx, input)
	if err != nil {
		return nil, err
	}

	return &proto.SendMessageResponse{
		Success: true,
	}, nil
}
