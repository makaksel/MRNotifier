package postgres

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/makaksel/MRNotifier/internal/domain"
)

type MergeRequestRepo struct {
	db *pgxpool.Pool
}

func NewMergeRequestRepo(db *pgxpool.Pool) *MergeRequestRepo {
	return &MergeRequestRepo{db: db}
}

func (r *MergeRequestRepo) Save(ctx context.Context, mr *domain.MergeRequest) {

}
func (r *MergeRequestRepo) GetByID(ctx context.Context, id uuid.UUID) {

}
