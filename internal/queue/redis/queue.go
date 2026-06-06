package queue

import (
	"context"
	"strconv"

	"github.com/redis/go-redis/v9"
)

type RedisQueue struct {
	client  *redis.Client
	channel string
}

func NewQueue(client *redis.Client, channel string) *RedisQueue {
	return &RedisQueue{
		client:  client,
		channel: channel,
	}
}

func (q *RedisQueue) Publish(ctx context.Context, id int) error {
	return q.client.Publish(ctx, q.channel, id).Err()
}

func (q *RedisQueue) Consume(ctx context.Context) (<-chan int, error) {
	pubsub := q.client.Subscribe(ctx, q.channel)

	// дождаться подписки
	if _, err := pubsub.Receive(ctx); err != nil {
		return nil, err
	}

	ch := make(chan int)

	go func() {
		defer close(ch)
		defer pubsub.Close()

		for {
			select {
			case msg := <-pubsub.Channel():
				id, err := strconv.Atoi(msg.Payload)
				if err != nil {
					continue
				}
				ch <- id

			case <-ctx.Done():
				return
			}
		}
	}()

	return ch, nil
}
