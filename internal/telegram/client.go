package telegram

import (
	"context"
	"log"
	"math"
	"time"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/makaksel/MRNotifier/internal/domain"
)

type Client interface {
	Send(ctx context.Context, n *domain.Notification) (*models.Message, error)
}

type Config struct {
	BotToken string
	ChatID   string
	Users    map[int]string
}

type Bot struct {
	*bot.Bot
	Config Config
}

func New(c Config) (*Bot, error) {
	var (
		b     *bot.Bot
		err   error
		retry = 3
	)

	for i := 0; i < retry; i++ {
		b, err = bot.New(c.BotToken)
		if err == nil {
			break
		}

		// Wait before retrying
		time.Sleep(time.Duration(math.Pow(2, float64(i))) * time.Second)
	}

	if err != nil {
		return nil, err
	}

	return &Bot{
		Bot:    b,
		Config: c,
	}, nil
}

func (b *Bot) Send(ctx context.Context, n *domain.Notification) (*models.Message, error) {
	log.Printf("n.IdForReply: %+v", n.IdForReply)
	res, err := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:          b.Config.ChatID, // из config
		Text:            n.Text,
		ReplyParameters: &models.ReplyParameters{MessageID: n.IdForReply},
	})

	return res, err
}
