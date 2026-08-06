package repository

import (
	"context"

	"github.com/makaksel/MRNotifier/internal/domain"
)

type MergeRequestRepository interface {
	UpsertMR(ctx context.Context, mr *domain.MergeRequestEvent) (bool, error)
}

type NotificationRepository interface {
	Get(ctx context.Context, id int) (*domain.Notification, error)
	GetByProject(ctx context.Context, project string, mrIID int) (*domain.Notification, error)
	Insert(ctx context.Context, projectPath string, mrIID int, n *domain.Notification) (bool, error)
	UpdateMessageID(ctx context.Context, id, msgID int, chatID int64) error
}
