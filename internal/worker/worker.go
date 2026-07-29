package worker

import (
	"context"
	"log"

	"github.com/makaksel/MRNotifier/internal/domain"
	queue "github.com/makaksel/MRNotifier/internal/queue/redis"
)

type NotificationHandler interface {
	Handle(ctx context.Context, n *domain.Notification) error
}

type Worker struct {
	queue   *queue.RedisQueue
	handler NotificationHandler
}

func NewWorker(q *queue.RedisQueue, h NotificationHandler) *Worker {
	return &Worker{
		queue:   q,
		handler: h,
	}
}

func (w *Worker) Start(ctx context.Context) error {
	ch, err := w.queue.Consume(ctx)
	if err != nil {
		return err
	}

	const workers = 5
	sem := make(chan struct{}, workers)

	for {
		select {
		case n, ok := <-ch:
			if !ok {
				return nil
			}

			sem <- struct{}{}

			go func(n *domain.Notification) {
				defer func() { <-sem }()

				if err := w.handler.Handle(ctx, n); err != nil {
					log.Printf("handle error: %v, notification: %+v", err, n)
				}
			}(n)

		case <-ctx.Done():
			return ctx.Err()
		}
	}
}
