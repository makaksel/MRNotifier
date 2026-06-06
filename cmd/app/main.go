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

	// 2. Подключение к PostgreSQL
	db := postgres.NewConnection(cfg.Postgres)

	// 3. Репозиторий для работы с MR
	MRRepo := postgres.NewMRRepo(db)
	NotificationRepo := postgres.NewNotificationRepo(db)

	// 4. Редис - для очереди
	redisS := redis.New(cfg.Redis)

	// 5. Очередь
	queueS := queue.NewQueue(redisS, cfg.Redis.Channel)

	// 7. Telegram клиент
	tg := telegram.New(cfg.Telegram)

	// 8. UseCase — основная бизнес-логика
	usecaseS := usecase.New(MRRepo, NotificationRepo, queueS)

	// 9. Worker — асинхронная обработка MR
	workerHandler := notification.New(NotificationRepo, tg)
	workerS := worker.NewWorker(queueS, workerHandler)

	// 10. Запускаем воркер
	go workerS.Start(context.Background())

	defer db.Close()
	defer redisS.Close()

	// 11. HTTP слой
	handler := transportHttp.NewHandler(usecaseS)
	router := transportHttp.NewRouter(handler)

	// 12. Старт сервера
	http.ListenAndServe(":8080", router)
}
