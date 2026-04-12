package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/makaksel/MRNotifier/internal/domain"
)

type MergeRequestRepository interface {
	Save(ctx context.Context, mr *domain.MergeRequest) error
	GetByID(ctx context.Context, id uuid.UUID) (*domain.MergeRequest, error)
}
