package usecases_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"github.com/popeskul/awesome-messanger/services/message/internal/core/ports"
	"github.com/popeskul/awesome-messanger/services/message/internal/usecases"
)

func TestNewUseCase(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockMessageUseCase := ports.NewMockMessageUseCase(ctrl)

	tests := []struct {
		name string
	}{
		{
			name: "Create new UseCase",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := usecases.NewUseCase(mockMessageUseCase)
			assert.NotNil(t, got)
			assert.Implements(t, (*ports.UserCase)(nil), got)
		})
	}
}

func TestUsecase_MessageUseCase(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockMessageUseCase := ports.NewMockMessageUseCase(ctrl)

	tests := []struct {
		name string
	}{
		{
			name: "Get MessageUseCase",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := usecases.NewUseCase(mockMessageUseCase)
			got := uc.MessageUseCase()
			assert.Equal(t, mockMessageUseCase, got)
		})
	}
}
