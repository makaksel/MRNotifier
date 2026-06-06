package notification

import (
	"context"
	"fmt"

	"github.com/go-telegram/bot"
)

func (c *Client) Handle(ctx context.Context, id int) error {
	// 1. достать notification
	n, err := c.repo.GetNotification(ctx, id)
	if err != nil {
		return err
	}

	msg := fmt.Sprintf(
		"📦 MR %s/%d: %s\n\nCreated: %s\n\nWeb URL here",
		n.ProjectPath, n.MRIID, n.EventType, n.CreatedAt,
	)

	// 2. отправить в Telegram
	_, err = c.tg.Bot.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: c.tg.Config.ChatID,
		Text:   msg,
	})
	return err
}
