package usecase

import (
	"context"
	"log"

	"github.com/makaksel/MRNotifier/internal/domain"
)

func (uc *Client) HandleMr(ctx context.Context, input *domain.MergeRequestEvent) error {
	log.Printf("input.MR.State: %s", input.MR.State)

	if input.MR.State != "opened" || input.MR.State != "merged" {
		log.Printf("MR state not opened or merged: %s", input.MR.State)
		return nil
	}

	// Обновляем MR в БД
	updated, err := uc.repo.UpsertMR(ctx, input)
	if err != nil {
		return err
	}
	if !updated {
		return nil
	}

	// Обновляем уведомление в БД
	created, err := uc.repo.InsertNotification(ctx, input.ProjectPath, input.MR.IID, input.MR.State)
	if err != nil {
		return err
	}
	if !created {
		return nil
	}

	// Пушим в очередь воркера

	return nil
}
