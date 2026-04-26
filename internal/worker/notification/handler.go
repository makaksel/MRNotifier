package notification

import (
	"context"

	"github.com/google/uuid"
)

func (c *Client) Handle(ctx context.Context, id uuid.UUID) error {
	// 1. достать notification
	n, err := c.repo.GetNotification(ctx, id)
	if err != nil {
		return err
	}

	// 2. отправить в Telegram
	return c.tg.SendMessage(ctx, n)
}
