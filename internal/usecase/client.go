package usecase

import (
	queue "github.com/makaksel/MRNotifier/internal/queue/redis"
	"github.com/makaksel/MRNotifier/internal/repository"
)

type Client struct {
	MRRepo           repository.MergeRequestRepository
	NotificationRepo repository.NotificationRepository
	queue            *queue.RedisQueue
}

func New(mrRepo repository.MergeRequestRepository, nRepo repository.NotificationRepository, queue *queue.RedisQueue) *Client {

	return &Client{
		mrRepo,
		nRepo,
		queue,
	}
}
