package usecases_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"github.com/popeskul/awesome-messanger/services/profile/internal/core/domain"
	"github.com/popeskul/awesome-messanger/services/profile/internal/core/ports"
	"github.com/popeskul/awesome-messanger/services/profile/internal/usecases"
)

func TestProfileUseCase_CreateProfile_Success(t *testing.T) {
	tests := []struct {
		name  string
		input *ports.CreateProfileRequest
		want  *domain.Profile
	}{
		{
			name: "Create profile with all fields",
			input: &ports.CreateProfileRequest{
				UserId:    "user1",
				Nickname:  "nick1",
				Bio:       "bio1",
				AvatarUrl: "avatar1",
			},
			want: &domain.Profile{
				UserId:    "user1",
				Nickname:  "nick1",
				Bio:       "bio1",
				AvatarUrl: "avatar1",
			},
		},
		{
			name: "Create profile with minimal fields",
			input: &ports.CreateProfileRequest{
				UserId:   "user2",
				Nickname: "nick2",
			},
			want: &domain.Profile{
				UserId:   "user2",
				Nickname: "nick2",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockLogger := ports.NewMockLogger(ctrl)
			uc := usecases.NewProfileUseCase(mockLogger)

			mockLogger.EXPECT().Info("CreateProfile called")

			got, err := uc.CreateProfile(context.Background(), tt.input)

			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestProfileUseCase_CreateProfile_Fail(t *testing.T) {
	tests := []struct {
		name    string
		input   *ports.CreateProfileRequest
		wantErr string
	}{
		{
			name: "Create profile with empty user id",
			input: &ports.CreateProfileRequest{
				UserId:   "",
				Nickname: "nick",
			},
			wantErr: "user_id is required",
		},
		{
			name: "Create profile with empty nickname",
			input: &ports.CreateProfileRequest{
				UserId:   "user1",
				Nickname: "",
			},
			wantErr: "nickname is required",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockLogger := ports.NewMockLogger(ctrl)
			uc := usecases.NewProfileUseCase(mockLogger)

			mockLogger.EXPECT().Info("CreateProfile called")

			_, err := uc.CreateProfile(context.Background(), tt.input)

			assert.Error(t, err)
			assert.Contains(t, err.Error(), tt.wantErr)
		})
	}
}

func TestProfileUseCase_UpdateProfile_Success(t *testing.T) {
	tests := []struct {
		name  string
		input *ports.UpdateProfileRequest
		want  *domain.Profile
	}{
		{
			name: "Update all profile fields",
			input: &ports.UpdateProfileRequest{
				UserId:    "user1",
				Nickname:  "nick1_updated",
				Bio:       "bio1_updated",
				AvatarUrl: "avatar1_updated",
			},
			want: &domain.Profile{
				UserId:    "user1",
				Nickname:  "nick1_updated",
				Bio:       "bio1_updated",
				AvatarUrl: "avatar1_updated",
			},
		},
		{
			name: "Update only nickname",
			input: &ports.UpdateProfileRequest{
				UserId:   "user2",
				Nickname: "nick2_updated",
			},
			want: &domain.Profile{
				UserId:   "user2",
				Nickname: "nick2_updated",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockLogger := ports.NewMockLogger(ctrl)
			uc := usecases.NewProfileUseCase(mockLogger)

			mockLogger.EXPECT().Info("UpdateProfile called")

			got, err := uc.UpdateProfile(context.Background(), tt.input)

			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestProfileUseCase_UpdateProfile_Fail(t *testing.T) {
	tests := []struct {
		name    string
		input   *ports.UpdateProfileRequest
		wantErr string
	}{
		{
			name: "Update profile with empty user id",
			input: &ports.UpdateProfileRequest{
				UserId:   "",
				Nickname: "nick",
			},
			wantErr: "invalid user id",
		},
		{
			name: "Update profile with empty fields",
			input: &ports.UpdateProfileRequest{
				UserId: "user1",
			},
			wantErr: "no fields to update",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockLogger := ports.NewMockLogger(ctrl)
			uc := usecases.NewProfileUseCase(mockLogger)

			mockLogger.EXPECT().Info("UpdateProfile called")

			_, err := uc.UpdateProfile(context.Background(), tt.input)

			assert.Error(t, err)
			assert.Contains(t, err.Error(), tt.wantErr)
		})
	}
}

func TestProfileUseCase_GetProfile_Success(t *testing.T) {
	tests := []struct {
		name  string
		input *ports.GetProfileRequest
		want  *domain.Profile
	}{
		{
			name: "Get existing profile",
			input: &ports.GetProfileRequest{
				UserId: "user1",
			},
			want: &domain.Profile{
				UserId:    "user1",
				Nickname:  "nickname",
				Bio:       "bio",
				AvatarUrl: "avatar_url",
			},
		},
		{
			name: "Get another existing profile",
			input: &ports.GetProfileRequest{
				UserId: "user2",
			},
			want: &domain.Profile{
				UserId:    "user2",
				Nickname:  "nickname",
				Bio:       "bio",
				AvatarUrl: "avatar_url",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockLogger := ports.NewMockLogger(ctrl)
			uc := usecases.NewProfileUseCase(mockLogger)

			mockLogger.EXPECT().Info("GetProfile called")

			got, err := uc.GetProfile(context.Background(), tt.input)

			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestProfileUseCase_GetProfile_Fail(t *testing.T) {
	tests := []struct {
		name    string
		input   *ports.GetProfileRequest
		wantErr string
	}{
		{
			name: "Get profile with empty user id",
			input: &ports.GetProfileRequest{
				UserId: "",
			},
			wantErr: "invalid user id",
		},
		{
			name: "Get non-existent profile",
			input: &ports.GetProfileRequest{
				UserId: "nonexistent",
			},
			wantErr: "profile not found",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockLogger := ports.NewMockLogger(ctrl)
			uc := usecases.NewProfileUseCase(mockLogger)

			mockLogger.EXPECT().Info("GetProfile called")

			_, err := uc.GetProfile(context.Background(), tt.input)

			assert.Error(t, err)
			assert.Contains(t, err.Error(), tt.wantErr)
		})
	}
}
