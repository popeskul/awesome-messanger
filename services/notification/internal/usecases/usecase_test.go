package usecases_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"github.com/popeskul/awesome-messanger/services/notification/internal/core/ports"
	"github.com/popeskul/awesome-messanger/services/notification/internal/usecases"
)

func TestNewUseCase(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockNotificationUseCase := ports.NewMockNotificationUseCase(ctrl)

	tests := []struct {
		name string
	}{
		{
			name: "Create new UseCase",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := usecases.NewUseCase(mockNotificationUseCase)
			assert.NotNil(t, got)
			assert.Implements(t, (*ports.UseCase)(nil), got)
		})
	}
}

func TestUseCase_NotificationUseCase(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockNotificationUseCase := ports.NewMockNotificationUseCase(ctrl)

	tests := []struct {
		name string
	}{
		{
			name: "Get NotificationUseCase",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := usecases.NewUseCase(mockNotificationUseCase)
			got := uc.Notification()
			assert.NotNil(t, got)
			assert.Implements(t, (*ports.NotificationUseCase)(nil), got)
			assert.Equal(t, mockNotificationUseCase, got)
		})
	}
}
