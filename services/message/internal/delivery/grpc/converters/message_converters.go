package converters

import (
	"github.com/popeskul/awesome-messanger/services/message/internal/core/domain"
	"github.com/popeskul/awesome-messanger/services/message/pkg/api/message"
)

// GRPCGetMessagesRequestToDomainGetMessagesRequest converts gRPC GetMessagesRequest to domain GetMessagesRequest
func GRPCGetMessagesRequestToDomainGetMessagesRequest(req *message.GetMessagesRequest) *domain.GetMessagesRequest {
	return &domain.GetMessagesRequest{
		ChatId:          req.GetChatId(),
		Limit:           req.GetLimit(),
		BeforeTimestamp: req.GetBeforeTimestamp(),
	}
}

// DomainGetMessagesResponseToGRPCGetMessagesResponse converts domain GetMessagesResponse to gRPC GetMessagesResponse
func DomainGetMessagesResponseToGRPCGetMessagesResponse(resp *domain.GetMessagesResponse) *message.GetMessagesResponse {
	messages := make([]*message.Message, len(resp.Messages))
	for i, msg := range resp.Messages {
		messages[i] = DomainMessageToGRPCMessage(msg)
	}
	return &message.GetMessagesResponse{
		Messages: messages,
		HasMore:  resp.HasMore,
	}
}

// GRPCSendMessageRequestToDomainSendMessageRequest converts gRPC SendMessageRequest to domain SendMessageRequest
func GRPCSendMessageRequestToDomainSendMessageRequest(req *message.SendMessageRequest) *domain.SendMessageRequest {
	return &domain.SendMessageRequest{
		ChatId:   req.GetChatId(),
		SenderId: req.GetSenderId(),
		Content:  req.GetContent(),
	}
}

// DomainSendMessageResponseToGRPCSendMessageResponse converts domain SendMessageResponse to gRPC SendMessageResponse
func DomainSendMessageResponseToGRPCSendMessageResponse(resp *domain.SendMessageResponse) *message.SendMessageResponse {
	return &message.SendMessageResponse{
		Message: DomainMessageToGRPCMessage(resp.Message),
	}
}

// GRPCStreamMessagesRequestToDomainStreamMessagesRequest converts gRPC StreamMessagesRequest to domain StreamMessagesRequest
func GRPCStreamMessagesRequestToDomainStreamMessagesRequest(req *message.StreamMessagesRequest) *domain.StreamMessagesRequest {
	return &domain.StreamMessagesRequest{
		ChatId: req.GetChatId(),
	}
}

// DomainMessageToGRPCMessage converts domain Message to gRPC Message
func DomainMessageToGRPCMessage(msg *domain.Message) *message.Message {
	return &message.Message{
		Id:        msg.Id,
		ChatId:    msg.ChatId,
		SenderId:  msg.SenderId,
		Content:   msg.Content,
		Timestamp: msg.Timestamp,
	}
}

// GRPCMessageToDomainMessage converts gRPC Message to domain Message
func GRPCMessageToDomainMessage(msg *message.Message) *domain.Message {
	return &domain.Message{
		Id:        msg.GetId(),
		ChatId:    msg.GetChatId(),
		SenderId:  msg.GetSenderId(),
		Content:   msg.GetContent(),
		Timestamp: msg.GetTimestamp(),
	}
}
