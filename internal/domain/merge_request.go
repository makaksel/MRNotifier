package domain

import (
	"time"

	"github.com/google/uuid"
)

type MergeRequest struct {
	ID        uuid.UUID
	ProjectID uuid.UUID
	Title     string
	Author    string
	URL       string
	CreatedAt time.Time
}
