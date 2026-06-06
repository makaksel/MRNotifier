package postgres

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/makaksel/MRNotifier/internal/domain"
)

type NotificationRepo struct {
	db *pgxpool.Pool
}

func NewNotificationRepo(db *pgxpool.Pool) *NotificationRepo {

	return &NotificationRepo{db: db}
}

func (r *NotificationRepo) InsertNotification(ctx context.Context, projectPath string, mrIID int, eventType string) (bool, error) {
	res, err := r.db.Exec(ctx, `
		INSERT INTO notifications (project_path, mr_iid, event_type)
		VALUES ($1, $2, $3)
		ON CONFLICT DO NOTHING
	`,
		projectPath,
		mrIID,
		eventType,
	)

	if err != nil {
		return false, err
	}

	return res.RowsAffected() > 0, nil
}

func (r *NotificationRepo) GetNotification(ctx context.Context, id int) (*domain.Notification, error) {
	var n domain.Notification

	err := r.db.QueryRow(ctx, `
		SELECT
			id,
			project_path,
			mr_iid,
			event_type,
			created_at
		FROM notifications
		WHERE id = $1
	`, id).Scan(
		&n.ID,
		&n.ProjectPath,
		&n.MRIID,
		&n.EventType,
		&n.CreatedAt,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &n, nil
}
