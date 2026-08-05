package cache

import (
	"context"

	"github.com/makaksel/MRNotifier/internal/domain"
)

type ReplyCache interface {
	Set(
		ctx context.Context,
		project string,
		mrIID int,
		chatID int64,
		messageID int,
	) error

	Get(
		ctx context.Context,
		project string,
		mrIID int,
	) (*domain.ReplyInfo, error)

	Delete(
		ctx context.Context,
		project string,
		mrIID int,
	) error
}
