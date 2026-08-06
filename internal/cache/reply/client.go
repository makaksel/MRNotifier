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
	data, err := c.rdb.Get(ctx, makeKey(project, mrIID)).Bytes()
	if err == redis.Nil {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	var info domain.ReplyInfo
	if err := json.Unmarshal(data, &info); err != nil {
		return nil, err
	}

	return &info, nil
}

func (c *Client) Delete(
	ctx context.Context,
	project string,
	mrIID int,
) error {
	return c.rdb.Del(ctx, makeKey(project, mrIID)).Err()
}
