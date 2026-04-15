package usecase

import (
	"context"

	"github.com/makaksel/MRNotifier/internal/domain"
	"github.com/makaksel/MRNotifier/internal/utils"
	"github.com/xanzy/go-gitlab"
)

func (uc *Client) HandleMr(ctx context.Context, input domain.CreateMRRequest) error {
	MRKey := utils.MakeKey(input.ProjectPath, input.MRID)

	var MR gitlab.MergeRequest

	// Проверяем МР в кеше
	err := uc.cache.Get(ctx, MRKey, MR)
	if err != nil {
		// Проверяем МР в бд если нет в кеше
		MR = uc.repo.GetByMRKey(ctx, MRKey)
	}

	// Запрашиваем МР из гитлаба
	MRData, _, _ := uc.gitlabClient.MergeRequests.GetMergeRequest(input.ProjectPath, input.MRID, &gitlab.GetMergeRequestsOptions{})

	// Обновляем ДБ
	uc.repo.Save(ctx, MRData)

	// Обновляем кеш
	uc.cache.Set(ctx, MRKey, MRData)

	// Пушим в очередь воркера

	if err := uc.repo.Save(ctx, mr); err != nil {
		return err
	}

	return uc.queue.Publish(ctx, mr.ID)
}
