package repository

import (
	"context"

	"github.com/makaksel/MRNotifier/internal/domain"
)

type MergeRequestRepository interface {
	UpsertMR(ctx context.Context, mr *domain.MergeRequestEvent) (bool, error)
}

type NotificationRepository interface {
	GetNotification(ctx context.Context, id int) (*domain.Notification, error)
	InsertNotification(ctx context.Context, projectPath string, mrIID int, eventType string) (bool, error)
}
