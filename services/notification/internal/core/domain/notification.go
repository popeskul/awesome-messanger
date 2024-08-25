package domain

type SendNotificationRequest struct {
	RecipientId string `json:"recipient_id,omitempty"`
	Message     string `json:"message,omitempty"`
}

type SendNotificationResponse struct {
	Success bool `json:"success"`
}
