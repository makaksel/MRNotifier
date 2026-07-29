package usecase

import (
	"context"
	"fmt"
	"log"

	"github.com/makaksel/MRNotifier/internal/domain"
)

func (uc *Client) HandleMr(ctx context.Context, data *domain.MergeRequestEvent) error {
	allowedStates := map[string]bool{"opened": true, "merged": true}
	if !allowedStates[data.MR.State] {
		return fmt.Errorf("MR state not opened or merged: %s", data.MR.State)
	}

	log.Printf("input.MR.State: %s", data.MR.State)
	// Обновляем MR в БД
	updated, err := uc.MRRepo.UpsertMR(ctx, data)

	if err != nil {
		return err
	}
	if !updated {
		return fmt.Errorf("MR already exist. Path: %s; IID: %d", data.ProjectPath, data.MR.IID)
	}

	// Создаем новое уведомление
	n := makeNotification(data)

	// Обновляем уведомление в БД
	created, err := uc.NotificationRepo.InsertNotification(ctx, n.ProjectPath, n.MRIID, n)
	if err != nil {
		return err
	}
	if !created {
		return fmt.Errorf("Notification already exist. Path: %s; IID: %d", data.ProjectPath, data.MR.IID)
	}

	// Кладем уведомление в кеш

	// Пушим в очередь воркера
	err = uc.queue.Publish(ctx, data.MR.ID)

	if err != nil {
		return fmt.Errorf("Push to worker failed. Path: %s; and IID: %d", data.ProjectPath, data.MR.IID)
	}

	return nil
}

func makeNotification(data *domain.MergeRequestEvent) domain.Notification {
	var text string

	if data.MR.State == "merged" {
		// TODO from there I can get TG for !!data.MR.Author.Username ?
		text = fmt.Sprintf(
			"✅ *Merge Request принят!*\n\n👤 Смержил: %s\n\n🧑 Обнови статус в Джире: %s",
			data.MR.Author.Username,
			data.MR.Author.Username,
		)

	}
	// TODO from there I can get TG for !!data.MR.Author.Username ?
	text = fmt.Sprintf(
		"🚀 Новый Merge Request создан\n\n📌 Название: %s\n\n🌿 Ветка: %s → %s\n\n👤 Автор: %s (TG: %s)\n\n🔗 Ссылка: %s",
		data.MR.Title,
		data.MR.SourceBranch,
		data.MR.TargetBranch,
		data.MR.Author.Name,
		data.MR.Author.Username,
		data.MR.WebURL,
	)

	return domain.Notification{
		Text:             text,
		ProjectPath:      data.ProjectPath,
		MRIID:            data.MR.IID,
		Status:           "new",
		ReplyToMessageId: "1", // from there?
	}

}
