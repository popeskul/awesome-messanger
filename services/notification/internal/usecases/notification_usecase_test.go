package usecases_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"github.com/popeskul/awesome-messanger/services/notification/internal/core/domain"
	"github.com/popeskul/awesome-messanger/services/notification/internal/core/ports"
	"github.com/popeskul/awesome-messanger/services/notification/internal/usecases"
)

func TestNotificationUseCase_SendNotification_Success(t *testing.T) {
	tests := []struct {
		name    string
		req     *domain.SendNotificationRequest
		wantRes *domain.SendNotificationResponse
	}{
		{
			name: "Send notification successfully",
			req: &domain.SendNotificationRequest{
				RecipientId: "user1",
				Message:     "Test notification",
			},
			wantRes: &domain.SendNotificationResponse{
				Success: true,
			},
		},
		{
			name: "Send notification to another user",
			req: &domain.SendNotificationRequest{
				RecipientId: "user2",
				Message:     "Another test notification",
			},
			wantRes: &domain.SendNotificationResponse{
				Success: true,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockLogger := ports.NewMockLogger(ctrl)

			mockLogger.EXPECT().Info("Sending notification to %s", tt.req.RecipientId)

			uc := usecases.NewNotificationUseCase(mockLogger)
			res, err := uc.SendNotification(context.Background(), tt.req)

			assert.NoError(t, err)
			assert.Equal(t, tt.wantRes, res)
		})
	}
}

func TestNotificationUseCase_SendNotification_Failure(t *testing.T) {
	tests := []struct {
		name      string
		req       *domain.SendNotificationRequest
		setupMock func(*ports.MockLogger)
		wantErr   bool
	}{
		{
			name: "Send notification with empty recipient",
			req: &domain.SendNotificationRequest{
				RecipientId: "",
				Message:     "Test notification",
			},
			setupMock: func(ml *ports.MockLogger) {
				// No logging expected for invalid input
			},
			wantErr: true,
		},
		{
			name: "Send notification with empty message",
			req: &domain.SendNotificationRequest{
				RecipientId: "user1",
				Message:     "",
			},
			setupMock: func(ml *ports.MockLogger) {
				// No logging expected for invalid input
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockLogger := ports.NewMockLogger(ctrl)

			tt.setupMock(mockLogger)

			uc := usecases.NewNotificationUseCase(mockLogger)
			res, err := uc.SendNotification(context.Background(), tt.req)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, res)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, res)
			}
		})
	}
}

func TestNewNotificationUseCase(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLogger := ports.NewMockLogger(ctrl)
	uc := usecases.NewNotificationUseCase(mockLogger)

	assert.NotNil(t, uc)
	assert.Implements(t, (*ports.NotificationUseCase)(nil), uc)
}
