package domain

import "time"

type Notification struct {
	ID int `json:"id"`

	ProjectPath string `json:"project_path"`
	MRIID       int    `json:"mr_iid"`
	Status      string `json:"status"` // new | done

	Text string `json:"text"`

	ReplyToMessageId string `json:"reply_to_message_id"`

	CreatedAt time.Time `json:"created_at"`
}
