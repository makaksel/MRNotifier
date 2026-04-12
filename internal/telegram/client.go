package telegram

import (
	"github.com/go-telegram/bot"
)

type Config struct {
	BotToken string
	ChatID   string
}

func New(c Config) *bot.Bot {
	b, err := bot.New(c.BotToken)
	if err != nil {
		panic(err)
	}

	return b
}
