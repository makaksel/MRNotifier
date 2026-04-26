package main

import (
	"context"
	"net/http"

	"github.com/makaksel/MRNotifier/internal/config"
	queue "github.com/makaksel/MRNotifier/internal/queue/redis"
	"github.com/makaksel/MRNotifier/internal/redis"
	"github.com/makaksel/MRNotifier/internal/repository/postgres"
	"github.com/makaksel/MRNotifier/internal/telegram"
	transportHttp "github.com/makaksel/MRNotifier/internal/transport/http"
	"github.com/makaksel/MRNotifier/internal/usecase"
	"github.com/makaksel/MRNotifier/internal/worker"
	"github.com/makaksel/MRNotifier/internal/worker/notification"
)

func main() {
	// 1. Загружаем конфиг (DB, Redis, GitLab token, Telegram и т.д.)
	cfg := config.Load()

	// 2. Подключение к PostgreSQL (источник правды)
	db := postgres.New(cfg.Postgres)

	// 3. Репозиторий для работы с MR
	repo := postgres.NewMergeRequestRepo(db)

	// 4. Редис - кеш
	redis := redis.New(cfg.Redis)

	// 5. Очередь
	queue := queue.NewQueue(redis, "mr_notify")

	// 7. Telegram клиент
	tg := telegram.New(cfg.Telegram)

	// 8. UseCase — основная бизнес-логика
	usecase := usecase.New(repo, queue)

	// 9. Worker — асинхронная обработка MR
	workerHandler := notification.New(repo, tg)
	worker := worker.NewWorker(queue, workerHandler)

	// 10. Запускаем воркер
	go worker.Start(context.Background())

	// 11. HTTP слой
	handler := transportHttp.NewHandler(usecase)
	router := transportHttp.NewRouter(handler)

	// 12. Старт сервера
	http.ListenAndServe(":8080", router)
}
