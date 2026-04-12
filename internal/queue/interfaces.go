package queue

import (
	"context"

	"github.com/google/uuid"
)

type NotificationQueue interface {
	Publish(ctx context.Context, id uuid.UUID) error
	Consume(ctx context.Context) (<-chan uuid.UUID, error)
}
