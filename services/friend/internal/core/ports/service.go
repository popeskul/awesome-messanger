package ports

import (
	"context"
)

type OutboxProcessor interface {
	Start(ctx context.Context)
	Stop()
}

type Service interface {
	OutboxProcessor() OutboxProcessor
}
