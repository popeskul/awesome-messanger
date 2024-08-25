package handlers

import (
	"context"

	"github.com/popeskul/awesome-messanger/services/notification/pkg/api/health"
)

type HealthHandler struct {
	health.UnimplementedHealthServiceServer
}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

func (h *HealthHandler) Check(context.Context, *health.HealthCheckRequest) (*health.HealthCheckResponse, error) {
	return &health.HealthCheckResponse{
		Status: health.HealthStatus_HEALTHY,
	}, nil
}

func (h *HealthHandler) Liveness(context.Context, *health.HealthCheckRequest) (*health.HealthCheckResponse, error) {
	return &health.HealthCheckResponse{
		Status: health.HealthStatus_HEALTHY,
	}, nil
}

func (h *HealthHandler) Readiness(context.Context, *health.HealthCheckRequest) (*health.HealthCheckResponse, error) {
	return &health.HealthCheckResponse{
		Status: health.HealthStatus_HEALTHY,
	}, nil
}

func (h *HealthHandler) Healthz(context.Context, *health.HealthCheckRequest) (*health.HealthCheckResponse, error) {
	return &health.HealthCheckResponse{
		Status: health.HealthStatus_HEALTHY,
	}, nil
}
