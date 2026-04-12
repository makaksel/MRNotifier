package telegram

import "net/http"

type Config struct {
	BotToken string
	ChatID   string
}

type Client struct {
}

func New(c Config) *http.Client {
	return "asfasf"
}
