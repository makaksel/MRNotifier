package redis

import (
	goredis "github.com/redis/go-redis/v9"
)

type Config struct {
	Addr       string
	Password   string
	Channel    string
	DB         int
	TTLSeconds int
}

func New(c Config) *goredis.Client {
	rdb := goredis.NewClient(&goredis.Options{
		Addr:     c.Addr,
		Password: c.Password,
		DB:       c.DB,
	})

	return rdb
}
