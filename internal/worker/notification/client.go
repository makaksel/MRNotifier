package notification

import (
	"github.com/makaksel/MRNotifier/internal/repository"
	"github.com/makaksel/MRNotifier/internal/telegram"
)

type Client struct {
	repo repository.NotificationRepository
	tg   *telegram.Bot
}

func New(repo repository.NotificationRepository, tg *telegram.Bot) *Client {
	return &Client{
		repo: repo,
		tg:   tg,
	}
}
