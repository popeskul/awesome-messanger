package models

type SendNotificationRequest struct {
	RecipientId string `json:"recipient_id,omitempty"`
	Message     string `json:"message,omitempty"`
}
