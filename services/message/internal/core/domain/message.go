package domain

import "github.com/popeskul/awesome-messanger/services/message/pkg/api/message"

type Message struct {
	SenderId  string
	Content   string
	Timestamp int64
}

func (m *Message) ConvertToProto() *message.Message {
	return &message.Message{
		SenderId:  m.SenderId,
		Content:   m.Content,
		Timestamp: m.Timestamp,
	}
}

type GetMessagesRequest struct {
	ChatId string `json:"chat_id,omitempty"`
}

type SendMessageRequest struct {
	SenderId    string `json:"sender_id,omitempty"`
	RecipientId string `json:"recipient_id,omitempty"`
	Content     string `json:"content,omitempty"`
}
