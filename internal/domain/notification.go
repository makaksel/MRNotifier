package domain

import "time"

type Notification struct {
	ID int `json:"id"`

	ProjectPath string `json:"project_path"`
	MRIID       int    `json:"mr_iid"`
	Status      Status `json:"status"`
	Type        Type   `json:"type"`

	Text string `json:"text"`

	MessageId int   `json:"message_id"`
	ChatId    int64 `json:"chat_id"`

	CreatedAt time.Time `json:"created_at"`

	IdForReply     int
	ChatIdForReply int64
}

type Status string

const (
	StatusNew    Status = "new"
	StatusSended Status = "sended"
)

type Type string

const (
	TypeOpened Type = "opened"
	TypeMerged Type = "merged"
)
