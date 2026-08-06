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

	// Обновляем MR в БД
	updated, err := uc.MRRepo.UpsertMR(ctx, data)
	if err != nil {
		return err
	}
	if !updated {
		log.Printf("MR already exist. Path: %s; IID: %d", data.ProjectPath, data.MR.IID)
	}

	// Создаем новое уведомление
	var n domain.Notification
	if data.MR.State == string(domain.TypeMerged) {
		n = uc.makeMergedNotification(ctx, data)
	} else {
		n = uc.makeOpenedNotification(data)
	}

	// Обновляем уведомление в БД
	created, err := uc.NotificationRepo.Insert(ctx, n.ProjectPath, n.MRIID, &n)
	if err != nil {
		return err
	}
	if !created {
		return fmt.Errorf("notification already exist. Path: %s; IID: %d; Type: %s", n.ProjectPath, n.MRIID, n.Status)
	}

	// Пушим в очередь воркера
	err = uc.queue.Publish(ctx, n)
	if err != nil {
		return fmt.Errorf("push to worker failed. Notification: %+v", n)
	}

	return nil
}

func (uc *Client) makeOpenedNotification(data *domain.MergeRequestEvent) domain.Notification {
	authorTg := uc.users[data.MR.Author.ID]

	if authorTg != "" {
		authorTg = fmt.Sprintf(" (tg: %s)", authorTg)
	}

	text := fmt.Sprintf(
		"🚀 Новый Merge Request создан\n\n📌 Название: %s\n\n🌿 Ветка: %s → %s\n\n👤 Автор: %s%s\n\n🔗 Ссылка: %s",
		data.MR.Title,
		data.MR.SourceBranch,
		data.MR.TargetBranch,
		data.MR.Author.Name,
		authorTg,
		data.MR.WebURL,
	)

	return domain.Notification{
		Text:        text,
		ProjectPath: data.ProjectPath,
		MRIID:       data.MR.IID,
		Type:        domain.TypeOpened,
	}
}

func (uc *Client) makeMergedNotification(ctx context.Context, data *domain.MergeRequestEvent) domain.Notification {
	authorTg := uc.users[data.MR.Author.ID]

	if authorTg != "" {
		authorTg = fmt.Sprintf("🧑 Обнови статус в Джире: %s\n\n", authorTg)
	}

	text := fmt.Sprintf(
		"✅ Merge Request принят!\n\n👤 Смержил: %s\n\n%s",
		data.MR.Author.Name,
		authorTg,
	)

	rId := uc.getReplyMsgId(ctx, data.ProjectPath, data.MR.IID)

	return domain.Notification{
		Text:        text,
		ProjectPath: data.ProjectPath,
		MRIID:       data.MR.IID,
		Type:        domain.TypeMerged,
		IdForReply:  rId,
	}
}

func (uc *Client) getReplyMsgId(ctx context.Context, project string, mrIID int) int {
	// Получить репли месседж из кеша
	r, err := uc.replyCache.Get(ctx, project, mrIID)
	if err != nil {
		log.Printf("err reply cache: %s", err)
	}
	if r != nil {
		return r.MessageID
	}

	// Получить репли месседж из БД
	n, err := uc.NotificationRepo.GetByProject(ctx, project, mrIID)
	if err != nil {
		log.Printf("err reply db: %s", err)
	}

	return n.MessageId
}
