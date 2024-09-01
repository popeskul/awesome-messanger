package models

import "time"

type OutboxEvent struct {
	ID          int
	EventType   string
	Payload     interface{}
	Status      string
	CreatedAt   time.Time
	ProcessedAt *time.Time
}
