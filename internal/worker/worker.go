package worker

import (
	"context"
	"log"

	"github.com/google/uuid"
	queue "github.com/makaksel/MRNotifier/internal/queue/redis"
)

type NotificationHandler interface {
	Handle(ctx context.Context, id uuid.UUID) error
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
		case id, ok := <-ch:
			if !ok {
				return nil
			}

			sem <- struct{}{}

			go func(id uuid.UUID) {
				defer func() { <-sem }()

				if err := w.handler.Handle(ctx, id); err != nil {
					log.Printf("handle error: %v, id=%s", err, id)
				}
			}(id)

		case <-ctx.Done():
			return ctx.Err()
		}
	}
}
