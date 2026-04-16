package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/makaksel/MRNotifier/internal/cache/redis"
	"github.com/makaksel/MRNotifier/internal/gitlab"
	"github.com/makaksel/MRNotifier/internal/repository/postgres"
	"github.com/makaksel/MRNotifier/internal/telegram"
)

type Config struct {
	App      AppConfig
	Postgres postgres.Config
	Redis    redis.Config
	GitLab   gitlab.Config
	Telegram telegram.Config
}

type AppConfig struct {
	Name string
	Env  string
	Port string
}

func Load() *Config {
	return &Config{
		App: AppConfig{
			Name: getEnv("APP_NAME", "mr-notifier"),
			Env:  getEnv("APP_ENV", "local"),
			Port: getEnv("APP_PORT", "8080"),
		},

		Postgres: postgres.Config{
			DSN: buildPostgresDSN(),
		},

		Redis: redis.Config{
			Addr:       fmt.Sprintf("%s:%s", getEnv("REDIS_HOST", "localhost"), getEnv("REDIS_PORT", "6379")),
			Password:   getEnv("REDIS_PASSWORD", ""),
			DB:         getEnvAsInt("REDIS_DB", 0),
			TTLSeconds: getEnvAsInt("CACHE_TTL_SECONDS", 3600),
		},

		GitLab: gitlab.Config{
			Token:   getEnv("GITLAB_TOKEN", ""),
			BaseURL: getEnv("GITLAB_BASE_URL", "https://gitlab-edu.mos.ru/api/v4"),
		},

		Telegram: telegram.Config{
			BotToken: getEnv("TELEGRAM_BOT_TOKEN", ""),
			ChatID:   getEnv("TELEGRAM_CHAT_ID", ""),
		},
	}
}

func getEnv(key, defaultVal string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return defaultVal
}

func getEnvAsInt(key string, defaultVal int) int {
	valStr := getEnv(key, "")
	if val, err := strconv.Atoi(valStr); err == nil {
		return val
	}
	return defaultVal
}

func buildPostgresDSN() string {
	host := getEnv("POSTGRES_HOST", "localhost")
	port := getEnv("POSTGRES_PORT", "5432")
	user := getEnv("POSTGRES_USER", "postgress")
	password := getEnv("POSTGRES_PASSWORD", "")
	db := getEnv("POSTGRES_DB", "postgress")
	ssl := getEnv("POSTGRES_SSLMODE", "disable")

	return fmt.Sprintf(
		"postgress://%s:%s@%s:%s/%s?sslmode=%s",
		user, password, host, port, db, ssl,
	)
}
