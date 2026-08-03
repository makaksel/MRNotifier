package queue

import (
	"context"
	"encoding/json"
	"log"

	"github.com/makaksel/MRNotifier/internal/domain"
	"github.com/redis/go-redis/v9"
)

type RedisQueue struct {
	client  *redis.Client
	channel string
}

func New(client *redis.Client, channel string) *RedisQueue {
	return &RedisQueue{
		client:  client,
		channel: channel,
	}
}

func (q *RedisQueue) Publish(ctx context.Context, n domain.Notification) error {
	payload, err := json.Marshal(n)
	if err != nil {
		return err
	}
	return q.client.Publish(ctx, q.channel, payload).Err()
}

func (q *RedisQueue) Consume(ctx context.Context) (<-chan domain.Notification, error) {
	pubsub := q.client.Subscribe(ctx, q.channel)

	if _, err := pubsub.Receive(ctx); err != nil {
		return nil, err
	}

	out := make(chan domain.Notification)
	redisCh := pubsub.Channel()

	go func() {
		defer close(out)
		defer pubsub.Close()

		for {
			select {
			case msg, ok := <-redisCh:
				if !ok {
					return
				}

				var n domain.Notification
				err := json.Unmarshal([]byte(msg.Payload), &n)
				if err != nil {
					log.Printf("decode error: %v", err)
					continue
				}

				select {
				case out <- n:
				case <-ctx.Done():
					return
				}

			case <-ctx.Done():
				return
			}
		}
	}()

	return out, nil
}
