package usecase

import (
	"context"
	"log"

	"github.com/makaksel/MRNotifier/internal/domain"
	"github.com/makaksel/MRNotifier/internal/utils"
	"github.com/xanzy/go-gitlab"
)

func (uc *Client) HandleMr(ctx context.Context, input domain.CreateMRRequest) error {
	MRKey := utils.MakeKey(input.ProjectPath, input.MRIID)

	MR := new(gitlab.MergeRequest)

	// Проверяем МР в кеше
	err := uc.cache.Get(ctx, MRKey, MR)
	if err != nil {
		// Проверяем МР в бд если нет в кеше
		MR = uc.repo.GetByMRKey(ctx, MRKey)
	}
	log.Printf("GitLab data: %s; %d", input.ProjectPath, input.MRIID)

	// Запрашиваем МР из гитлаба
	MRData, resp, err := uc.gitlabClient.MergeRequests.GetMergeRequest(
		input.ProjectPath,
		input.MRIID,
		nil,
	)

	if err != nil {
		log.Println("GitLab error:", err)
		return err
	}

	log.Println("Status:", resp.StatusCode)

	// Обновляем ДБ
	uc.repo.Save(ctx, MRData)

	// Обновляем кеш
	uc.cache.Set(ctx, MRKey, MRData)

	// Пушим в очередь воркера

	//if err := uc.repo.Save(ctx, MRData); err != nil {
	//	return err
	//}

	//return uc.queue.Publish(ctx, mr.ID)
	return nil
}
