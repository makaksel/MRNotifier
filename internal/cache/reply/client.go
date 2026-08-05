package replyCache

import (
	"context"
	"encoding/json"
	"time"

	"github.com/makaksel/MRNotifier/internal/domain"
	"github.com/redis/go-redis/v9"
)

type Client struct {
	rdb *redis.Client
	ttl time.Duration
}

const ttl = 30 * 24 * time.Hour

func New(rdb *redis.Client) *Client {
	return &Client{
		rdb: rdb,
		ttl: ttl,
	}
}

func (c *Client) Set(
	ctx context.Context,
	project string,
	mrIID int,
	chatID int64,
	messageID int,
) error {
	value, err := json.Marshal(domain.ReplyInfo{
		ChatID:    chatID,
		MessageID: messageID,
	})
	if err != nil {
		return err
	}

	return c.rdb.Set(
		ctx,
		makeKey(project, mrIID),
		value,
		c.ttl,
	).Err()
}

func (c *Client) Get(
	ctx context.Context,
	project string,
	mrIID int,
) (*domain.ReplyInfo, error) {
	makeKey(project, mrIID)

	return nil, nil
}

func (c *Client) Delete(
	ctx context.Context,
	project string,
	mrIID int,
) error {
	makeKey(project, mrIID)

	return nil
}
