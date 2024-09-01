package service

import (
	"context"
	"time"

	"github.com/popeskul/awesome-messanger/services/friend/internal/core/ports"
)

var (
	interval      = 5 * time.Second
	processAtTime = 10
)

type OutboxProcessor struct {
	repo   ports.Repository
	logger ports.Logger
}

func NewOutboxProcessor(repo ports.Repository, logger ports.Logger) *OutboxProcessor {
	return &OutboxProcessor{
		repo:   repo,
		logger: logger,
	}
}

func (p *OutboxProcessor) Start(ctx context.Context) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			if err := p.processEvents(ctx); err != nil {
				p.logger.Error("Error processing outbox events: %v", err)
			}
		}
	}
}

func (p *OutboxProcessor) processEvents(ctx context.Context) error {
	events, err := p.repo.Outbox().GetPendingEvents(ctx, processAtTime)
	if err != nil {
		return err
	}

	for _, event := range events {
		// Here you would publish the event to your message broker
		// For example:
		// err := p.kafkaClient.Publish(event.EventType, event.Payload)
		// if err != nil {
		//     p.logger.Error("Failed to publish event: %v", err)
		//     continue
		// }

		// For now, we'll just log the event
		p.logger.Info("Processing event: %s", event.EventType)

		if err := p.repo.Outbox().MarkAsProcessed(ctx, event.ID); err != nil {
			p.logger.Error("Failed to mark event as processed: %v", err)
		}
	}

	return nil
}

func (p *OutboxProcessor) Stop() {
	p.logger.Info("Stopping outbox processor")
}
