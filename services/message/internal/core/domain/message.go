package domain

import (
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Message struct {
	Id        string                 `json:"id,omitempty"`
	ChatId    string                 `json:"chat_id,omitempty"`
	SenderId  string                 `json:"sender_id,omitempty"`
	Content   string                 `json:"content,omitempty"`
	Timestamp *timestamppb.Timestamp `json:"timestamp,omitempty"`
}

type GetMessagesRequest struct {
	ChatId          string                 `json:"chat_id,omitempty"`
	Limit           int32                  `json:"limit,omitempty"`
	BeforeTimestamp *timestamppb.Timestamp `json:"before_timestamp,omitempty"`
}

type GetMessagesResponse struct {
	Messages []*Message `json:"messages,omitempty"`
	HasMore  bool       `json:"has_more,omitempty"`
}

type SendMessageRequest struct {
	ChatId   string `json:"chat_id,omitempty"`
	SenderId string `json:"sender_id,omitempty"`
	Content  string `json:"content,omitempty"`
}

type SendMessageResponse struct {
	Message *Message `json:"message,omitempty"`
}

type StreamMessagesRequest struct {
	ChatId string `json:"chat_id,omitempty"`
}
