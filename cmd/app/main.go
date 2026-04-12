package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/makaksel/MRNotifier/internal/cache/redis"
	"github.com/makaksel/MRNotifier/internal/config"
	"github.com/makaksel/MRNotifier/internal/gitlab"
	queue "github.com/makaksel/MRNotifier/internal/queue/memory"
	"github.com/makaksel/MRNotifier/internal/repository/postgres"
	"github.com/makaksel/MRNotifier/internal/telegram"
	"github.com/makaksel/MRNotifier/internal/usecase"
	"github.com/makaksel/MRNotifier/internal/worker"
)

func main() {

	// 1. Загружаем конфиг (DB, Redis, GitLab token, Telegram и т.д.)
	cfg := config.Load()

	// 2. Подключение к PostgreSQL (источник правды)
	db, err := postgres.New(cfg.Postgres)
	if err != nil {
		log.Fatal(err)
	}

	// 3. Репозиторий для работы с MR
	repo := postgres.NewMergeRequestRepo(db)

	// 4. Redis — используется как кэш + быстрый доступ
	cache := redis.New(cfg.Redis.Addr, cfg.Redis.Password, cfg.Redis.DB, cfg.Redis.TTLSeconds)

	// 5. Очередь (может быть Redis/Rabbit/etc)
	queue := queue.New(cfg.Queue)

	// 6. GitLab клиент (для получения актуальных данных MR)
	gitlabClient := gitlab.New(cfg.GitLab)

	// 7. Telegram клиент
	tg := telegram.New(cfg.Telegram)

	// 8. UseCase — основная бизнес-логика
	usecase := usecase.New(repo, cache, queue, gitlabClient)

	// 9. Worker — асинхронная обработка MR
	worker := worker.New(repo, cache, queue, gitlab, tg)

	// 10. Запускаем воркер
	go worker.Start(context.Background())

	// 11. HTTP слой
	router := http.NewRouter(usecase)

	// 12. Старт сервера
	http.ListenAndServe(":8080", router)

	fmt.Println("Hello world")
}
