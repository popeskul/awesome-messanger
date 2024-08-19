package handlers

import (
	"context"
	"fmt"

	"github.com/bufbuild/protovalidate-go"
	"github.com/popeskul/awesome-messanger/services/message/internal/models"
	"github.com/popeskul/awesome-messanger/services/message/internal/services"
	"github.com/popeskul/awesome-messanger/services/message/pkg/api/health"
	"github.com/popeskul/awesome-messanger/services/message/pkg/api/message"
	"github.com/popeskul/awesome-messanger/services/message/proto/api/grpcutils"
)

type Service interface {
	MessageService() services.MessageServiceI
}

type Handler struct {
	message.UnimplementedMessageServiceServer
	health.UnimplementedHealthServiceServer
	service   Service
	validator *protovalidate.Validator
}

func NewHandler(service Service) (*Handler, error) {
	validator, err := protovalidate.New(
		protovalidate.WithDisableLazy(true),
		protovalidate.WithMessages(
			&message.GetMessagesRequest{},
			&message.SendMessageRequest{},
			&health.HealthCheckRequest{},
		),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create validator: %w", err)
	}

	return &Handler{
		service:   service,
		validator: validator,
	}, nil
}

func (h *Handler) GetMessages(ctx context.Context, req *message.GetMessagesRequest) (*message.GetMessagesResponse, error) {
	if err := h.validator.Validate(req); err != nil {
		return nil, grpcutils.RPCValidationError(err)
	}

	input := &models.GetMessagesRequest{
		ChatId: req.GetChatId(),
	}

	messages, err := h.service.MessageService().GetMessages(ctx, input)
	if err != nil {
		return nil, err
	}

	var pbMessages []*message.Message
	for _, m := range messages {
		pbMessages = append(pbMessages, m.ConvertToProto())
	}

	return &message.GetMessagesResponse{Messages: pbMessages}, nil
}

func (h *Handler) SendMessage(ctx context.Context, req *message.SendMessageRequest) (*message.SendMessageResponse, error) {
	input := &models.SendMessageRequest{
		SenderId:    req.GetSenderId(),
		RecipientId: req.GetRecipientId(),
		Content:     req.GetContent(),
	}

	err := h.service.MessageService().SendMessage(ctx, input)
	if err != nil {
		return nil, err
	}

	return &message.SendMessageResponse{
		Success: true,
	}, nil
}

func (h *Handler) Check(ctx context.Context, req *health.HealthCheckRequest) (*health.HealthCheckResponse, error) {
	return &health.HealthCheckResponse{Status: "healthy"}, nil
}

func (h *Handler) Liveness(ctx context.Context, req *health.HealthCheckRequest) (*health.HealthCheckResponse, error) {
	return &health.HealthCheckResponse{Status: "alive"}, nil
}

func (h *Handler) Readiness(ctx context.Context, req *health.HealthCheckRequest) (*health.HealthCheckResponse, error) {
	return &health.HealthCheckResponse{Status: "ready"}, nil
}

func (h *Handler) Healthz(ctx context.Context, req *health.HealthCheckRequest) (*health.HealthCheckResponse, error) {
	return &health.HealthCheckResponse{Status: "healthy"}, nil
}
