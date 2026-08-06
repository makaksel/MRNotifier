package notification

import (
	"context"
	"log"

	"github.com/makaksel/MRNotifier/internal/cache"
	"github.com/makaksel/MRNotifier/internal/domain"
	queue "github.com/makaksel/MRNotifier/internal/queue"
	"github.com/makaksel/MRNotifier/internal/repository"
	"github.com/makaksel/MRNotifier/internal/telegram"
)

type Worker struct {
	queue      queue.NotificationQueue
	repo       repository.NotificationRepository
	tg         telegram.Client
	replyCache cache.ReplyCache
}

func New(q queue.NotificationQueue, r repository.NotificationRepository, tg telegram.Client, rc cache.ReplyCache) *Worker {
	return &Worker{
		queue:      q,
		repo:       r,
		tg:         tg,
		replyCache: rc,
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

			go func(n domain.Notification) {
				defer func() { <-sem }()

				if err := w.Handle(ctx, n); err != nil {
					log.Printf("handle error: %v, notification: %+v", err, n)
				}
			}(n)

		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

func (w *Worker) Handle(ctx context.Context, n domain.Notification) error {
	// Send to tg
	res, err := w.tg.Send(ctx, &n)

	// Update notification in db
	err = w.repo.UpdateMessageID(ctx, n.ID, res.ID, res.Chat.ID)

	// Update reply in cache
	if n.Type == domain.TypeMerged {
		err = w.replyCache.Delete(ctx, n.ProjectPath, n.MRIID)
	} else {
		err = w.replyCache.Set(ctx, n.ProjectPath, n.MRIID, res.Chat.ID, res.ID)
	}

	return err
}
