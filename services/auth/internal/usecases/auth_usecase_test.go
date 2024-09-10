package usecases_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"github.com/popeskul/awesome-messanger/services/auth/internal/core/ports"
	"github.com/popeskul/awesome-messanger/services/auth/internal/usecases"
)

func TestAuthUseCase(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLogger := ports.NewMockLogger(ctrl)

	tests := []struct {
		name     string
		setup    func(useCase ports.AuthUseCase)
		test     func(*testing.T, ports.AuthUseCase)
		wantErr  bool
		errorMsg string
	}{
		{
			name: "VerifyCredentials_Success",
			setup: func(uc ports.AuthUseCase) {
				mockLogger.EXPECT().Info(gomock.Any(), gomock.Any()).Times(1)
			},
			test: func(t *testing.T, uc ports.AuthUseCase) {
				err := uc.VerifyCredentials(context.Background(), "testuser", "password")
				assert.NoError(t, err)
			},
		},
		{
			name: "Logout_Success",
			setup: func(uc ports.AuthUseCase) {
				mockLogger.EXPECT().Info(gomock.Any(), gomock.Any()).Times(1)
			},
			test: func(t *testing.T, uc ports.AuthUseCase) {
				err := uc.Logout(context.Background(), "testtoken")
				assert.NoError(t, err)
			},
		},
		{
			name: "Register_Success",
			setup: func(uc ports.AuthUseCase) {
				mockLogger.EXPECT().Info(gomock.Any(), gomock.Any()).Times(1)
			},
			test: func(t *testing.T, uc ports.AuthUseCase) {
				req := ports.RegisterRequest{
					Username: "testuser",
					Email:    "test@example.com",
					Password: "password",
				}
				resp, err := uc.Register(context.Background(), req)
				assert.NoError(t, err)
				assert.Equal(t, "testuser", resp.User.Username)
				assert.Equal(t, "test@example.com", resp.User.Email)
				assert.NotEmpty(t, resp.Token)
			},
		},
		{
			name: "Refresh_Success",
			setup: func(uc ports.AuthUseCase) {
				mockLogger.EXPECT().Info(gomock.Any(), gomock.Any()).Times(1)
			},
			test: func(t *testing.T, uc ports.AuthUseCase) {
				token, err := uc.Refresh(context.Background(), "oldtoken")
				assert.NoError(t, err)
				assert.NotEmpty(t, token)
			},
		},
		{
			name: "Me_Success",
			setup: func(uc ports.AuthUseCase) {
				mockLogger.EXPECT().Info(gomock.Any(), gomock.Any()).Times(1)
			},
			test: func(t *testing.T, uc ports.AuthUseCase) {
				result, err := uc.Me(context.Background(), "testtoken")
				assert.NoError(t, err)
				assert.Equal(t, "Me", result)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			uc := usecases.NewAuthUseCase(mockLogger)
			tt.setup(uc)
			tt.test(t, uc)
		})
	}
}
