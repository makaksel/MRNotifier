package postgres

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/makaksel/MRNotifier/internal/domain"
)

type MergeRequestRepo struct {
	db *pgxpool.Pool
}

func NewMRRepo(db *pgxpool.Pool) *MergeRequestRepo {

	return &MergeRequestRepo{db: db}
}

func (r *MergeRequestRepo) UpsertMR(ctx context.Context, e *domain.MergeRequestEvent) (bool, error) {
	res, err := r.db.Exec(ctx, `
		INSERT INTO merge_requests (
			project_path,
			mr_iid,
			title,
			description,
			state,
			web_url,
			source_branch,
			target_branch,
			changes_count,
			updated_at,
			merged_at,
			author_username
		)
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12)
		ON CONFLICT (project_path, mr_iid)
		DO UPDATE SET
			title = EXCLUDED.title,
			description = EXCLUDED.description,
			state = EXCLUDED.state,
			web_url = EXCLUDED.web_url,
			source_branch = EXCLUDED.source_branch,
			target_branch = EXCLUDED.target_branch,
			changes_count = EXCLUDED.changes_count,
			updated_at = EXCLUDED.updated_at,
			merged_at = EXCLUDED.merged_at,
			author_username = EXCLUDED.author_username
		WHERE (merge_requests.updated_at IS NULL
		       OR merge_requests.updated_at < EXCLUDED.updated_at)
	`,
		e.ProjectPath,
		e.MR.IID,
		e.MR.Title,
		e.MR.Description,
		e.MR.State,
		e.MR.WebURL,
		e.MR.SourceBranch,
		e.MR.TargetBranch,
		e.MR.ChangesCount,
		e.MR.UpdatedAt,
		e.MR.MergedAt,
		e.MR.Author.Username,
	)

	if err != nil {
		return false, err
	}

	return res.RowsAffected() > 0, nil
}
