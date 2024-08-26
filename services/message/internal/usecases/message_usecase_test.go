package usecases_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"github.com/popeskul/awesome-messanger/services/message/internal/core/domain"
	"github.com/popeskul/awesome-messanger/services/message/internal/core/ports"
	"github.com/popeskul/awesome-messanger/services/message/internal/usecases"
)

func TestMessageUsecase_GetMessages_Success(t *testing.T) {
	tests := []struct {
		name    string
		req     *domain.GetMessagesRequest
		wantLen int
	}{
		{
			name:    "Get messages from chat1",
			req:     &domain.GetMessagesRequest{ChatId: "chat1"},
			wantLen: 1,
		},
		{
			name:    "Get messages from chat2",
			req:     &domain.GetMessagesRequest{ChatId: "chat2"},
			wantLen: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockLogger := ports.NewMockLogger(ctrl)
			uc := usecases.NewMessageUseCase(mockLogger)

			mockLogger.EXPECT().Info("Getting messages for chat %s", tt.req.ChatId)

			resp, err := uc.GetMessages(context.Background(), tt.req)

			assert.NoError(t, err)
			assert.NotNil(t, resp)
			assert.Len(t, resp.Messages, tt.wantLen)
			assert.Equal(t, tt.req.ChatId, resp.Messages[0].ChatId)
		})
	}
}

func TestMessageUsecase_GetMessages_Fail(t *testing.T) {
	tests := []struct {
		name string
		req  *domain.GetMessagesRequest
	}{
		{
			name: "Empty ChatId",
			req:  &domain.GetMessagesRequest{ChatId: ""},
		},
		{
			name: "Whitespace ChatId",
			req:  &domain.GetMessagesRequest{ChatId: "   "},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockLogger := ports.NewMockLogger(ctrl)
			uc := usecases.NewMessageUseCase(mockLogger)

			resp, err := uc.GetMessages(context.Background(), tt.req)

			assert.Error(t, err)
			assert.Nil(t, resp)
		})
	}
}

func TestMessageUsecase_SendMessage_Success(t *testing.T) {
	tests := []struct {
		name string
		req  *domain.SendMessageRequest
	}{
		{
			name: "Send message to chat1",
			req:  &domain.SendMessageRequest{ChatId: "chat1", SenderId: "user1", Content: "Hello"},
		},
		{
			name: "Send message to chat2",
			req:  &domain.SendMessageRequest{ChatId: "chat2", SenderId: "user2", Content: "Hi there"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockLogger := ports.NewMockLogger(ctrl)
			uc := usecases.NewMessageUseCase(mockLogger)

			mockLogger.EXPECT().Info("Sending message from %s in chat %s", tt.req.SenderId, tt.req.ChatId)

			resp, err := uc.SendMessage(context.Background(), tt.req)

			assert.NoError(t, err)
			assert.NotNil(t, resp)
			assert.Equal(t, tt.req.ChatId, resp.Message.ChatId)
			assert.Equal(t, tt.req.SenderId, resp.Message.SenderId)
			assert.Equal(t, tt.req.Content, resp.Message.Content)
		})
	}
}

func TestMessageUsecase_SendMessage_Fail(t *testing.T) {
	tests := []struct {
		name string
		req  *domain.SendMessageRequest
	}{
		{
			name: "Empty ChatId",
			req:  &domain.SendMessageRequest{ChatId: "", SenderId: "user1", Content: "Hello"},
		},
		{
			name: "Empty SenderId",
			req:  &domain.SendMessageRequest{ChatId: "chat1", SenderId: "", Content: "Hello"},
		},
		{
			name: "Empty Content",
			req:  &domain.SendMessageRequest{ChatId: "chat1", SenderId: "user1", Content: ""},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockLogger := ports.NewMockLogger(ctrl)
			uc := usecases.NewMessageUseCase(mockLogger)

			resp, err := uc.SendMessage(context.Background(), tt.req)

			assert.Error(t, err)
			assert.Nil(t, resp)
		})
	}
}

func TestMessageUsecase_StreamMessages_Success(t *testing.T) {
	tests := []struct {
		name string
		req  *domain.StreamMessagesRequest
	}{
		{
			name: "Stream messages from chat1",
			req:  &domain.StreamMessagesRequest{ChatId: "chat1"},
		},
		{
			name: "Stream messages from chat2",
			req:  &domain.StreamMessagesRequest{ChatId: "chat2"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockLogger := ports.NewMockLogger(ctrl)
			uc := usecases.NewMessageUseCase(mockLogger)

			mockLogger.EXPECT().Info("Streaming messages for chat %s", tt.req.ChatId)

			messageChan, err := uc.StreamMessages(context.Background(), tt.req)

			assert.NoError(t, err)
			assert.NotNil(t, messageChan)

			select {
			case msg := <-messageChan:
				assert.NotNil(t, msg)
				assert.Equal(t, tt.req.ChatId, msg.ChatId)
			case <-time.After(6 * time.Second):
				t.Fatal("Timeout waiting for message")
			}
		})
	}
}

func TestMessageUsecase_StreamMessages_Fail(t *testing.T) {
	tests := []struct {
		name string
		req  *domain.StreamMessagesRequest
	}{
		{
			name: "Empty ChatId",
			req:  &domain.StreamMessagesRequest{ChatId: ""},
		},
		{
			name: "Whitespace ChatId",
			req:  &domain.StreamMessagesRequest{ChatId: "   "},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockLogger := ports.NewMockLogger(ctrl)
			uc := usecases.NewMessageUseCase(mockLogger)

			messageChan, err := uc.StreamMessages(context.Background(), tt.req)

			assert.Error(t, err)
			assert.Nil(t, messageChan)
		})
	}
}
