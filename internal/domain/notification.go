package domain

import "time"

type Notification struct {
	ID int `json:"id"`

	ProjectPath string `json:"project_path"`
	MRIID       int    `json:"mr_iid"`
	Status      string `json:"status"` // new | done
	Type        string `json:"type"`   // opened | merged

	Text string `json:"text"`

	MessageId int   `json:"message_id"`
	ChatId    int64 `json:"chat_id"`

	CreatedAt time.Time `json:"created_at"`

	IdForReply     int
	ChatIdForReply int64
}
