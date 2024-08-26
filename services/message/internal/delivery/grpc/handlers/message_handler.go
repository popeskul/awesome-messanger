package handlers

import (
	"context"

	"github.com/popeskul/awesome-messanger/services/message/internal/core/ports"
	"github.com/popeskul/awesome-messanger/services/message/internal/delivery/grpc/converters"
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

	input := converters.GRPCGetMessagesRequestToDomainGetMessagesRequest(req)

	messages, err := h.messageUseCase.GetMessages(ctx, input)
	if err != nil {
		return nil, err
	}

	return converters.DomainGetMessagesResponseToGRPCGetMessagesResponse(messages), nil
}

func (h *MessageHandler) SendMessage(ctx context.Context, req *message.SendMessageRequest) (*message.SendMessageResponse, error) {
	if err := h.validator.Validate(req); err != nil {
		return nil, grpcutils.RPCValidationError(err)
	}

	input := converters.GRPCSendMessageRequestToDomainSendMessageRequest(req)

	resp, err := h.messageUseCase.SendMessage(ctx, input)
	if err != nil {
		return nil, err
	}

	return converters.DomainSendMessageResponseToGRPCSendMessageResponse(resp), nil
}

func (h *MessageHandler) StreamMessages(req *message.StreamMessagesRequest, stream message.MessageService_StreamMessagesServer) error {
	if err := h.validator.Validate(req); err != nil {
		return grpcutils.RPCValidationError(err)
	}

	input := converters.GRPCStreamMessagesRequestToDomainStreamMessagesRequest(req)

	messageChan, err := h.messageUseCase.StreamMessages(stream.Context(), input)
	if err != nil {
		return err
	}

	for msg := range messageChan {
		grpcMsg := converters.DomainMessageToGRPCMessage(msg)
		if err := stream.Send(grpcMsg); err != nil {
			return err
		}
	}

	return nil
}
