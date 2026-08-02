package cache

import (
	"github.com/makaksel/MRNotifier/internal/queue"
	"github.com/makaksel/MRNotifier/internal/repository"
	"github.com/makaksel/MRNotifier/internal/telegram"
)

type Client struct {
}

func New(q queue.NotificationQueue, r repository.NotificationRepository, tg telegram.Client) *Client {
	return &Client{}
}
