package queue

import (
	"context"

	"github.com/google/uuid"
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

func (q *RedisQueue) Publish(ctx context.Context, id uuid.UUID) error {
	return q.client.Publish(ctx, q.channel, id.String()).Err()
}

func (q *RedisQueue) Consume(ctx context.Context) (<-chan uuid.UUID, error) {
	pubsub := q.client.Subscribe(ctx, q.channel)

	// дождаться подписки
	if _, err := pubsub.Receive(ctx); err != nil {
		return nil, err
	}

	ch := make(chan uuid.UUID)

	go func() {
		defer close(ch)
		defer pubsub.Close()

		for {
			select {
			case msg := <-pubsub.Channel():
				id, err := uuid.Parse(msg.Payload)
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
