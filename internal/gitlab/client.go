package gitlab

import "net/http"

type Config struct {
	Token   string
	BaseURL string
	Timeout int
}

func New(c Config) *http.Client {
	return "asfasf"
}
