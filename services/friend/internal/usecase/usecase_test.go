package usecase_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"github.com/popeskul/awesome-messanger/services/friend/internal/core/ports"
	"github.com/popeskul/awesome-messanger/services/friend/internal/usecase"
)

func TestNewUseCase(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockFriendUseCase := ports.NewMockFriendUseCase(ctrl)

	uc := usecase.NewUseCase(mockFriendUseCase)

	assert.NotNil(t, uc)
	assert.Implements(t, (*ports.UseCase)(nil), uc)
}

func TestUseCase_FriendUseCase(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockFriendUseCase := ports.NewMockFriendUseCase(ctrl)

	uc := usecase.NewUseCase(mockFriendUseCase)

	retrievedFriendUseCase := uc.FriendUseCase()

	assert.NotNil(t, retrievedFriendUseCase)
	assert.Equal(t, mockFriendUseCase, retrievedFriendUseCase)
}
