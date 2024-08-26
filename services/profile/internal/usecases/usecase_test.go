package usecases_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"github.com/popeskul/awesome-messanger/services/profile/internal/core/ports"
	"github.com/popeskul/awesome-messanger/services/profile/internal/usecases"
)

func TestNewUseCase(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "Create new UseCase",
		},
		{
			name: "Create another UseCase instance",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockProfileUseCase := ports.NewMockProfileUseCase(ctrl)

			uc := usecases.NewUseCase(mockProfileUseCase)

			assert.NotNil(t, uc)
			assert.Implements(t, (*ports.UserCase)(nil), uc)
		})
	}
}

func TestUseCase_ProfileUseCase(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "Get ProfileUseCase",
		},
		{
			name: "Get ProfileUseCase again",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockProfileUseCase := ports.NewMockProfileUseCase(ctrl)
			uc := usecases.NewUseCase(mockProfileUseCase)

			got := uc.ProfileUseCase()

			assert.NotNil(t, got)
			assert.Equal(t, mockProfileUseCase, got)
		})
	}
}
