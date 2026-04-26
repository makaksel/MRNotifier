package notification

import (
	"github.com/go-telegram/bot"
	"github.com/makaksel/MRNotifier/internal/repository/postgres"
)

type Client struct {
	repo *postgres.MergeRequestRepo
	tg   *bot.Bot
}

func New(repo *postgres.MergeRequestRepo, tg *bot.Bot) *Client {
	return &Client{
		repo: repo,
		tg:   tg,
	}
}
