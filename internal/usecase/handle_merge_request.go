package usecase

import (
	"context"

	"github.com/makaksel/MRNotifier/internal/domain"
)

func (uc *Client) HandleMr(ctx context.Context, input domain.CreateMRRequest) error {
	mr := mapToDomain(input)

	if err := uc.repo.Save(ctx, mr); err != nil {
		return err
	}

	return uc.queue.Publish(ctx, mr.ID)
}
