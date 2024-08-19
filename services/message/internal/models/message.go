package models

import (
	"github.com/popeskul/awesome-messanger/services/message/pb/proto"
)

type Message struct {
	SenderId  string
	Content   string
	Timestamp int64
}

func (m *Message) ConvertToProto() *proto.Message {
	return &proto.Message{
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
