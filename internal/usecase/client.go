package usecase

import (
	"github.com/makaksel/MRNotifier/internal/cache"
	queue "github.com/makaksel/MRNotifier/internal/queue"
	"github.com/makaksel/MRNotifier/internal/repository"
)

type Client struct {
	MRRepo           repository.MergeRequestRepository
	NotificationRepo repository.NotificationRepository
	queue            queue.NotificationQueue
	replyCache       cache.ReplyCache
	users            map[int]string
}

func New(
	mrRepo repository.MergeRequestRepository,
	nRepo repository.NotificationRepository,
	queue queue.NotificationQueue,
	replyCache cache.ReplyCache,
	users map[int]string) *Client {

	return &Client{
		mrRepo,
		nRepo,
		queue,
		replyCache,
		users,
	}
}
