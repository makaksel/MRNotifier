package worker

import (
	"context"

	"github.com/makaksel/MRNotifier/internal/queue"
	"github.com/makaksel/MRNotifier/internal/repository/postgres"
	"github.com/makaksel/MRNotifier/internal/telegram"
)

type NotificationWorker struct {
	repo     postgres.MergeRequestRepository
	queue    queue.NotificationQueue
	telegram telegram.Client
}

func (w *NotificationWorker) Start(ctx context.Context) {
	ch, _ := w.queue.Consume(ctx)

	for id := range ch {
		mr, _ := w.repo.GetByID(ctx, id)
		w.telegram.SendMergeRequestNotification(mr)
	}
}
