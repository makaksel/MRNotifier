package usecase

import (
	"context"
	"log"

	"github.com/makaksel/MRNotifier/internal/domain"
)

func (uc *Client) HandleMr(ctx context.Context, input *domain.MergeRequestEvent) error {
	log.Printf("input.MR.State: %s", input.MR.State)

	allowedStates := map[string]bool{"opened": true, "merged": true}
	if !allowedStates[input.MR.State] {
		log.Printf("MR state not opened or merged: %s", input.MR.State)
		return nil
	}

	// Обновляем MR в БД
	updated, err := uc.MRRepo.UpsertMR(ctx, input)
	if err != nil {
		return err
	}
	if !updated {
		return nil
	}

	// Обновляем уведомление в БД
	created, err := uc.NotificationRepo.InsertNotification(ctx, input.ProjectPath, input.MR.IID, input.MR.State)
	if err != nil {
		return err
	}
	if !created {
		return nil
	}

	// Пушим в очередь воркера
	return uc.queue.Publish(ctx, 1)
}
