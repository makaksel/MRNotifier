package notification

import (
	"context"
	"fmt"
	"log"

	"github.com/go-telegram/bot"
)

func (c *Client) Handle(ctx context.Context, id int) error {
	// Get from cache by ID

	// Get from DB by ID
	log.Printf("Handle MRIID: %d", id)

	msg := fmt.Sprintf(`
		🚀 Новый Merge Request создан
		📌 Название: 
		🌿 Ветка: → 
		👤 Автор:  (TG:@chingaevaes)
		🔗 Ссылка: 
		`,
		id,
	)

	// Send to tg
	_, err := c.tg.Bot.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: c.tg.Config.ChatID,
		Text:   msg,
	})

	// Update notification in cache
	// Update notification in db

	// If it new, update reply message

	return err
}
