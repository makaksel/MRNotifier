package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	replyCache "github.com/makaksel/MRNotifier/internal/cache/reply"
	"github.com/makaksel/MRNotifier/internal/config"
	queue "github.com/makaksel/MRNotifier/internal/queue/redis"
	"github.com/makaksel/MRNotifier/internal/redis"
	"github.com/makaksel/MRNotifier/internal/repository/postgres"
	"github.com/makaksel/MRNotifier/internal/telegram"
	transportHttp "github.com/makaksel/MRNotifier/internal/transport/http"
	"github.com/makaksel/MRNotifier/internal/usecase"
	"github.com/makaksel/MRNotifier/internal/worker/notification"
)

func main() {
	// 1. Загружаем конфиг (DB, Redis, GitLab token, Telegram и т.д.)
	cfg := config.Load()

	// 2. Подключение к PostgreSQL
	db := postgres.NewConnection(cfg.Postgres)

	// 3. Репозиторий для работы с MR
	MRRepo := postgres.NewMRRepo(db)
	NRepo := postgres.NewNotificationRepo(db)

	// 4. Редис - для очереди
	r := redis.New(cfg.Redis)

	// 5. Очередь
	q := queue.New(r, cfg.Redis.Channel)

	// 6. Кеш ответов
	rc := replyCache.New(r)

	// 7. Telegram клиент
	tg, err := telegram.New(cfg.Telegram)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to telegram: %v\n", err)
	}

	// 8. UseCase — основная бизнес-логика
	uc := usecase.New(MRRepo, NRepo, q, rc, cfg.Telegram.Users)

	// 9. Worker — асинхронная обработка MR
	worker := notification.New(q, NRepo, tg, rc)

	// 10. Запускаем воркер
	go worker.Start(context.Background())

	defer db.Close()
	defer r.Close()

	// 11. HTTP слой
	handler := transportHttp.NewHandler(uc)
	router := transportHttp.NewRouter(handler)

	// 12. Старт сервера
	http.ListenAndServe(":8080", router)
}
