package worker

import (
	"context"
)

type NotificationWorker interface {
	Start(ctx context.Context) error
	Handle(ctx context.Context, id int) error
}
