package queue

import (
	"context"
	"log"
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
	return q.client.Publish(ctx, q.channel, strconv.Itoa(id)).Err()
}

func (q *RedisQueue) Consume(ctx context.Context) (<-chan int, error) {
	pubsub := q.client.Subscribe(ctx, q.channel)

	if _, err := pubsub.Receive(ctx); err != nil {
		return nil, err
	}

	out := make(chan int)
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

				notification, err := strconv.Atoi(msg.Payload)
				if err != nil {
					log.Printf("decode error: %v", err)
					continue
				}

				select {
				case out <- notification:
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
