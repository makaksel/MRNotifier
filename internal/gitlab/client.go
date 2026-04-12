package gitlab

import (
	"log"

	"github.com/xanzy/go-gitlab"
)

func New(token, baseURL string) *gitlab.Client {
	git, err := gitlab.NewClient(token, gitlab.WithBaseURL(baseURL))
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	return git
}
