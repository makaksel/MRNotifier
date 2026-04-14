package usecase

import (
	"github.com/makaksel/MRNotifier/internal/cache/redis"
	"github.com/makaksel/MRNotifier/internal/queue/memory"
	"github.com/makaksel/MRNotifier/internal/repository/postgres"
	"github.com/xanzy/go-gitlab"
)

type Client struct {
	repo         *postgres.MergeRequestRepo
	cache        *redis.Cache
	queue        *memory.Queue
	gitlabClient *gitlab.Client
}

func New(repo *postgres.MergeRequestRepo, cache *redis.Cache, queue *memory.Queue, gitlabClient *gitlab.Client) *Client {

	return &Client{
		repo,
		cache,
		queue,
		gitlabClient,
	}
}
