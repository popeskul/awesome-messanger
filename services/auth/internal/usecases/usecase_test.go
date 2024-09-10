package usecases_test

import (
	"github.com/popeskul/awesome-messanger/services/auth/internal/usecases"
	"testing"

	"github.com/popeskul/awesome-messanger/services/auth/internal/core/ports"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestUseCases(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTokenUseCase := ports.NewMockTokenUseCase(ctrl)
	mockAuthUseCase := ports.NewMockAuthUseCase(ctrl)

	useCases := usecases.NewUseCase(mockAuthUseCase, mockTokenUseCase)

	tests := []struct {
		name       string
		getUseCase func(cases ports.UseCase) interface{}
		expected   interface{}
	}{
		{
			name: "AuthUseCase",
			getUseCase: func(u ports.UseCase) interface{} {
				return u.AuthUseCase()
			},
			expected: mockAuthUseCase,
		},
		{
			name: "TokenUseCase",
			getUseCase: func(u ports.UseCase) interface{} {
				return u.TokenUseCase()
			},
			expected: mockTokenUseCase,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := tt.getUseCase(useCases)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestNewUseCase(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAuthUseCase := ports.NewMockAuthUseCase(ctrl)
	mockTokenUseCase := ports.NewMockTokenUseCase(ctrl)

	useCases := usecases.NewUseCase(mockAuthUseCase, mockTokenUseCase)

	assert.NotNil(t, useCases)
	assert.Equal(t, mockAuthUseCase, useCases.AuthUseCase())
	assert.Equal(t, mockTokenUseCase, useCases.TokenUseCase())
}
