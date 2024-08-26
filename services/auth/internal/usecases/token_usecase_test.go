package usecases_test

import (
	"errors"
	"github.com/popeskul/awesome-messanger/services/auth/internal/usecases"
	"testing"

	"github.com/popeskul/awesome-messanger/services/auth/internal/core/ports"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestTokenUseCase_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLogger := ports.NewMockLogger(ctrl)
	mockTokenManager := ports.NewMockTokenManager(ctrl)

	uc := usecases.NewTokenUseCase(mockLogger, mockTokenManager)

	tests := []struct {
		name           string
		input          string
		expectedOutput string
		setup          func()
		operation      func(string) (string, error)
	}{
		{
			name:           "GenerateToken_Success",
			input:          "test@example.com",
			expectedOutput: "valid_token",
			setup: func() {
				mockTokenManager.EXPECT().GenerateToken("test@example.com").Return("valid_token", nil)
			},
			operation: uc.GenerateToken,
		},
		{
			name:           "ValidateToken_Success",
			input:          "valid_token",
			expectedOutput: "test@example.com",
			setup: func() {
				mockTokenManager.EXPECT().ValidateToken("valid_token").Return("test@example.com", nil)
			},
			operation: uc.ValidateToken,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			tt.setup()
			result, err := tt.operation(tt.input)
			assert.NoError(t, err)
			assert.Equal(t, tt.expectedOutput, result)
		})
	}
}

func TestTokenUseCase_Fail(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLogger := ports.NewMockLogger(ctrl)
	mockTokenManager := ports.NewMockTokenManager(ctrl)

	uc := usecases.NewTokenUseCase(mockLogger, mockTokenManager)

	tests := []struct {
		name          string
		input         string
		expectedError string
		setup         func()
		operation     func(string) (string, error)
	}{
		{
			name:          "GenerateToken_Failure",
			input:         "test@example.com",
			expectedError: "token generation failed",
			setup: func() {
				mockTokenManager.EXPECT().GenerateToken("test@example.com").Return("", errors.New("token generation failed"))
				mockLogger.EXPECT().Info("Error generating token: %v", gomock.Any())
			},
			operation: uc.GenerateToken,
		},
		{
			name:          "ValidateToken_Failure",
			input:         "invalid_token",
			expectedError: "token validation failed",
			setup: func() {
				mockTokenManager.EXPECT().ValidateToken("invalid_token").Return("", errors.New("token validation failed"))
				mockLogger.EXPECT().Info("Error validating token: %v", gomock.Any())
			},
			operation: uc.ValidateToken,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			tt.setup()
			result, err := tt.operation(tt.input)
			assert.Error(t, err)
			assert.Equal(t, "", result)
			assert.EqualError(t, err, tt.expectedError)
		})
	}
}
