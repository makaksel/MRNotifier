package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/makaksel/MRNotifier/internal/redis"
	"github.com/makaksel/MRNotifier/internal/repository/postgres"
	"github.com/makaksel/MRNotifier/internal/telegram"

	"github.com/joho/godotenv"
)

type Config struct {
	App      AppConfig
	Postgres postgres.Config
	Redis    redis.Config
	Telegram telegram.Config
}

type AppConfig struct {
	Name string
	Env  string
	Port string
}

func Load() *Config {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var u map[int]string
	err = json.Unmarshal([]byte(getEnv("USERS_MAP", "")), &u)
	if err != nil {
		log.Printf("USERS_MAP parsing error: %v", err)
	}

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
			Addr:     fmt.Sprintf("%s:%s", getEnv("REDIS_HOST", "localhost"), getEnv("REDIS_PORT", "6379")),
			Password: getEnv("REDIS_PASSWORD", ""),
			DB:       getEnvAsInt("REDIS_DB", 0),
			Channel:  getEnv("REDIS_PORT_CHANNEL", ""),
		},

		Telegram: telegram.Config{
			BotToken: getEnv("TELEGRAM_BOT_TOKEN", ""),
			ChatID:   getEnv("TELEGRAM_CHAT_ID", ""),
			Users:    u,
		},
	}
}

func getEnv(key, defaultVal string) string {
	val := os.Getenv(key)
	if val != "" {
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
	port := getEnv("$POSTGRES_PORT", "5432")
	user := getEnv("POSTGRES_USER", "postgress")
	password := getEnv("POSTGRES_PASSWORD", "")
	db := getEnv("POSTGRES_DB", "postgress")
	ssl := getEnv("POSTGRES_SSLMODE", "disable")
	log.Printf(
		"host=%s port=%s user=%s db=%s",
		host,
		port,
		user,
		db,
	)
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		user, password, host, port, db, ssl,
	)
}
