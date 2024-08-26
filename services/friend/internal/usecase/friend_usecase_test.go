package usecase_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"github.com/popeskul/awesome-messanger/services/friend/internal/core/models"
	"github.com/popeskul/awesome-messanger/services/friend/internal/core/ports"
	"github.com/popeskul/awesome-messanger/services/friend/internal/usecase"
)

func TestFriendUseCase_Success(t *testing.T) {
	tests := []struct {
		name      string
		mockSetup func(*gomock.Controller) ports.Logger
		method    func(ports.FriendUseCase) (interface{}, error)
		want      interface{}
	}{
		{
			name: "AddFriend_Success",
			mockSetup: func(ctrl *gomock.Controller) ports.Logger {
				mockLogger := ports.NewMockLogger(ctrl)
				mockLogger.EXPECT().Info("Add friend request received for user %s", "1")
				return mockLogger
			},
			method: func(uc ports.FriendUseCase) (interface{}, error) {
				return uc.AddFriend(context.Background(), &models.Friend{UserId: "1", FriendId: "2"})
			},
			want: &models.Friend{UserId: "1", FriendId: "2"},
		},
		{
			name: "GetFriends_Success",
			mockSetup: func(ctrl *gomock.Controller) ports.Logger {
				mockLogger := ports.NewMockLogger(ctrl)
				mockLogger.EXPECT().Info("Get friends request received for user %s", "1")
				return mockLogger
			},
			method: func(uc ports.FriendUseCase) (interface{}, error) {
				return uc.GetFriends(context.Background())
			},
			want: []*models.Friend{{UserId: "1", FriendId: "2"}},
		},
		{
			name: "RespondToFriendRequest_Success",
			mockSetup: func(ctrl *gomock.Controller) ports.Logger {
				mockLogger := ports.NewMockLogger(ctrl)
				mockLogger.EXPECT().Info("Respond friend request received for user %s", "1")
				return mockLogger
			},
			method: func(uc ports.FriendUseCase) (interface{}, error) {
				return uc.RespondToFriendRequest(context.Background(), &models.Friend{UserId: "1", FriendId: "2"})
			},
			want: &models.Friend{UserId: "1", FriendId: "2"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockLogger := tt.mockSetup(ctrl)
			uc := usecase.NewFriendUseCase(mockLogger)
			got, err := tt.method(uc)

			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestFriendUseCase_Fail(t *testing.T) {
	tests := []struct {
		name    string
		method  func(ports.FriendUseCase) (interface{}, error)
		wantErr error
	}{
		{
			name: "AddFriend_Fail",
			method: func(uc ports.FriendUseCase) (interface{}, error) {
				return uc.AddFriend(context.Background(), &models.Friend{UserId: "1", FriendId: "2"})
			},
			wantErr: errors.New("friend already added"),
		},
		{
			name: "GetFriends_Fail",
			method: func(uc ports.FriendUseCase) (interface{}, error) {
				return uc.GetFriends(context.Background())
			},
			wantErr: errors.New("failed to retrieve friends"),
		},
		{
			name: "RespondToFriendRequest_Fail",
			method: func(uc ports.FriendUseCase) (interface{}, error) {
				return uc.RespondToFriendRequest(context.Background(), &models.Friend{UserId: "1", FriendId: "2"})
			},
			wantErr: errors.New("friend request not found"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockUC := ports.NewMockFriendUseCase(ctrl)
			switch tt.name {
			case "AddFriend_Fail":
				mockUC.EXPECT().AddFriend(gomock.Any(), gomock.Any()).Return(nil, tt.wantErr)
			case "GetFriends_Fail":
				mockUC.EXPECT().GetFriends(gomock.Any()).Return(nil, tt.wantErr)
			case "RespondToFriendRequest_Fail":
				mockUC.EXPECT().RespondToFriendRequest(gomock.Any(), gomock.Any()).Return(nil, tt.wantErr)
			}

			got, err := tt.method(mockUC)

			assert.Error(t, err)
			assert.Equal(t, tt.wantErr.Error(), err.Error())
			assert.Nil(t, got)
		})
	}
}

func TestNewFriendUseCase(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLogger := ports.NewMockLogger(ctrl)
	uc := usecase.NewFriendUseCase(mockLogger)

	assert.NotNil(t, uc)
	assert.Implements(t, (*ports.FriendUseCase)(nil), uc)
}
