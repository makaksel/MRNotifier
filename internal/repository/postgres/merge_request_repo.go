package postgres

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/xanzy/go-gitlab"
)

type MergeRequestRepo struct {
	db *pgxpool.Pool
}

func NewMergeRequestRepo(db *pgxpool.Pool) *MergeRequestRepo {
	return &MergeRequestRepo{db: db}
}

func (r *MergeRequestRepo) Save(ctx context.Context, mr *gitlab.MergeRequest) {

}

func (r *MergeRequestRepo) GetByMRKey(ctx context.Context, key string) gitlab.MergeRequest {

	return nil
}
