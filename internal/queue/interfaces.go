package queue

import (
	"context"

	"github.com/makaksel/MRNotifier/internal/domain"
)

type NotificationQueue interface {
	Publish(ctx context.Context, n domain.Notification) error
	Consume(ctx context.Context) (<-chan domain.Notification, error)
}
