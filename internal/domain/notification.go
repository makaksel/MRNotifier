package domain

import "time"

type Notification struct {
	ID          int       `json:"id"`
	ProjectPath string    `json:"project_path"`
	MRIID       int       `json:"mr_iid"`
	EventType   string    `json:"event_type"`
	CreatedAt   time.Time `json:"created_at"`
}
