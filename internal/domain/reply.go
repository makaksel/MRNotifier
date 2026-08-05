package domain

type ReplyInfo struct {
	MessageID int   `json:"message_id"`
	ChatID    int64 `json:"chat_id"`
}
