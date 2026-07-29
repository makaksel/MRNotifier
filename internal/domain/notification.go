package domain

import "time"

type Notification struct {
	ID          int    `json:"id"`
	ProjectPath string `json:"project_path"`
	MRIID       int    `json:"mr_iid"`

	Text   string `json:"text"`
	Status string `json:"status"` // new | done

	ReplyToMessageId string `json:"reply_to_message_id"`

	CreatedAt time.Time `json:"created_at"`
}
