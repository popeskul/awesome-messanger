package repository

import (
	"context"
	"encoding/json"
	"time"

	"github.com/popeskul/awesome-messanger/services/friend/internal/core/models"
	"github.com/popeskul/awesome-messanger/services/friend/internal/core/ports"
	platformPorts "github.com/popeskul/awesome-messanger/services/platform/database/postgres/ports"
)

type outboxRepository struct {
	conn platformPorts.Connection
}

func NewOutboxRepository(conn platformPorts.Connection) ports.OutboxRepository {
	return &outboxRepository{conn: conn}
}

func (r *outboxRepository) Add(ctx context.Context, event models.OutboxEvent) error {
	query := `INSERT INTO outbox (event_type, payload, status) VALUES ($1, $2, $3)`
	payload, err := json.Marshal(event.Payload)
	if err != nil {
		return err
	}
	_, err = r.conn.Exec(ctx, query, event.EventType, payload, "pending")
	return err
}

func (r *outboxRepository) GetPendingEvents(ctx context.Context, limit int) ([]models.OutboxEvent, error) {
	query := `SELECT id, event_type, payload, status, created_at FROM outbox WHERE status = 'pending' ORDER BY created_at LIMIT $1`
	rows, err := r.conn.Query(ctx, query, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []models.OutboxEvent
	for rows.Next() {
		var event models.OutboxEvent
		var payload []byte
		err := rows.Scan(&event.ID, &event.EventType, &payload, &event.Status, &event.CreatedAt)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(payload, &event.Payload)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	return events, nil
}

func (r *outboxRepository) MarkAsProcessed(ctx context.Context, id int) error {
	query := `UPDATE outbox SET status = 'processed', processed_at = $1 WHERE id = $2`
	_, err := r.conn.Exec(ctx, query, time.Now(), id)
	return err
}
