package gitlab

import (
	"log"

	"github.com/xanzy/go-gitlab"
)

type Config struct {
	Token   string
	BaseURL string
}

func New(c Config) *gitlab.Client {
	git, err := gitlab.NewClient(c.Token, gitlab.WithBaseURL(c.BaseURL))
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	return git
}
