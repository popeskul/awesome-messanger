package handlers

import (
	"context"
	"fmt"

	"github.com/bufbuild/protovalidate-go"
	"github.com/popeskul/awesome-messanger/services/notification/internal/models"
	"github.com/popeskul/awesome-messanger/services/notification/internal/services"
	"github.com/popeskul/awesome-messanger/services/notification/pkg/api/health"
	"github.com/popeskul/awesome-messanger/services/notification/pkg/api/notification"
	"github.com/popeskul/awesome-messanger/services/notification/proto/api/grpcutils"
)

type Services interface {
	NotificationService() services.NotificationServiceI
}

type Handler struct {
	notification.UnimplementedNotificationServiceServer
	health.UnimplementedHealthServiceServer
	services  Services
	validator *protovalidate.Validator
}

func NewHandler(services Services) (*Handler, error) {
	validator, err := protovalidate.New(
		protovalidate.WithDisableLazy(true),
		protovalidate.WithMessages(
			&notification.SendNotificationRequest{},
			&health.HealthCheckRequest{},
		),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create validator: %w", err)
	}

	return &Handler{
		services:  services,
		validator: validator,
	}, nil
}

func (h *Handler) SendNotification(
	ctx context.Context,
	req *notification.SendNotificationRequest,
) (*notification.SendNotificationResponse, error) {
	if err := h.validator.Validate(req); err != nil {
		return nil, grpcutils.RPCValidationError(err)
	}

	model := &models.SendNotificationRequest{
		RecipientId: req.GetRecipientId(),
		Message:     req.GetMessage(),
	}

	err := h.services.NotificationService().SendNotification(ctx, model)
	if err != nil {
		return nil, err
	}

	return &notification.SendNotificationResponse{}, nil
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
