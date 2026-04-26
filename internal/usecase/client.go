package usecase

import (
	queue "github.com/makaksel/MRNotifier/internal/queue/redis"
	"github.com/makaksel/MRNotifier/internal/repository/postgres"
)

type Client struct {
	repo  *postgres.MergeRequestRepo
	queue *queue.RedisQueue
}

func New(repo *postgres.MergeRequestRepo, queue *queue.RedisQueue) *Client {

	return &Client{
		repo,
		queue,
	}
}
