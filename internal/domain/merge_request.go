package domain

import "time"

type MergeRequestEvent struct {
	ProjectPath string           `json:"project_path"`
	MR          MergeRequestData `json:"mr"`
}

type MergeRequestData struct {
	ID          int    `json:"id"`
	IID         int    `json:"iid"`
	Title       string `json:"title"`
	Description string `json:"description"`
	State       string `json:"state"`
	WebURL      string `json:"web_url"`

	SourceBranch string `json:"source_branch"`
	TargetBranch string `json:"target_branch"`

	Author GitlabUser `json:"author"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	MergedAt  time.Time `json:"merged_at"`

	ChangesCount string `json:"changes_count"`
	Draft        bool   `json:"draft"`
}

type GitlabUser struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Name     string `json:"name"`
}
