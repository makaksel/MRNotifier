package redis

import (
	"context"
	"encoding/json"
	"time"

	goredis "github.com/redis/go-redis/v9"
)

type Config struct {
	Addr       string
	Password   string
	DB         int
	TTLSeconds int
}

type Cache struct {
	client *goredis.Client
	ttl    time.Duration
}

func New(c Config) *Cache {
	rdb := goredis.NewClient(&goredis.Options{
		Addr:     c.Addr,
		Password: c.Password,
		DB:       c.DB,
	})

	return &Cache{client: rdb, ttl: time.Duration(c.TTLSeconds) * time.Second}
}

func (c *Cache) Get(ctx context.Context, key string, dest any) error {
	val, err := c.client.Get(ctx, key).Result()
	if err != nil {
		return err
	}

	return json.Unmarshal([]byte(val), dest)
}

func (c *Cache) Set(ctx context.Context, key string, value any) error {
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}

	return c.client.Set(ctx, key, data, c.ttl).Err()
}

func (c *Cache) Delete(ctx context.Context, key string) error {
	return c.client.Del(ctx, key).Err()
}
