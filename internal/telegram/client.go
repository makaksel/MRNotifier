package telegram

import (
	"context"
	"fmt"

	"github.com/go-telegram/bot"
	"github.com/makaksel/MRNotifier/internal/domain"
)

type Config struct {
	BotToken string
	ChatID   string
}

type Bot struct {
	*bot.Bot
	Config Config
}

func New(c Config) *Bot {
	b, err := bot.New(c.BotToken)
	if err != nil {
		panic(err)
	}

	return &Bot{
		Bot:    b,
		Config: c,
	}
}

func (b *Bot) SendNotification(ctx context.Context, n *domain.Notification) error {
	msg := fmt.Sprintf(
		"📦 MR %s/%d: %s\n%s\n%s",
		n.ProjectPath, n.MRIID, n.EventType, n.CreatedAt, "Web URL here",
	)
	_, err := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: b.Config.ChatID, // из config
		Text:   msg,
	})
	return err
}
