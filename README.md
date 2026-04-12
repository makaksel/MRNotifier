📦 Структура проекта

```
mr-notifier/
│
├── cmd/
│   └── app/
│       └── main.go
│
├── internal/
│   │
│   ├── config/
│   │   └── config.go
│   │
│   ├── domain/
│   │   ├── merge_request.go
│   │   └── notification.go
│   │
│   ├── usecase/
│   │   └── handle_merge_request.go
│   │
│   ├── repository/
│   │   ├── postgres/
│   │   │   ├── connection.go
│   │   │   └── merge_request_repo.go
│   │   └── interfaces.go
│   │
│   ├── queue/
│   │   ├── interfaces.go
│   │   ├── redis/
│   │   │   └── queue.go
│   │   └── memory/
│   │       └── queue.go
│   │
│   ├── telegram/
│   │   └── client.go
│   │
│   ├── worker/
│   │   └── notification_worker.go
│   │
│   └── transport/
│       └── http/
│           ├── handler.go
│           └── router.go
│
├── migrations/
│   └── 001_init.sql
│
├── .air.toml
├── docker-compose.yml
├── go.mod
└── README.md
```
🔥 Поток обработки запроса

```
POST /merge-request
↓
HTTP handler
↓
UseCase (бизнес логика)
↓
1. Сохраняем в Postgres
2. Кладём задачу в очередь
   ↓
   Worker читает очередь
   ↓
   Telegram client отправляет сообщение
```

🧠 По слоям

1️⃣ domain/

Чистые модели без зависимостей.
```
type MergeRequest struct {
ID          uuid.UUID
ProjectID   uuid.UUID
Title       string
Author      string
URL         string
CreatedAt   time.Time
}
```
2️⃣ usecase/

Тут логика:
```
type MergeRequestUseCase struct {
repo  repository.MergeRequestRepository
queue queue.NotificationQueue
}

func (uc *MergeRequestUseCase) Handle(ctx context.Context, input CreateMRInput) error {
mr := mapToDomain(input)

    if err := uc.repo.Save(ctx, mr); err != nil {
        return err
    }

    return uc.queue.Publish(ctx, mr.ID)
}
```
UseCase ничего не знает:

ни про Postgres

ни про Redis

ни про Telegram

3️⃣ repository/
```
interfaces.go
type MergeRequestRepository interface {
Save(ctx context.Context, mr *domain.MergeRequest) error
GetByID(ctx context.Context, id uuid.UUID) (*domain.MergeRequest, error)
}
```
postgres/

Реализация через pgx или sqlx.

4️⃣ queue/
```
interfaces.go
type NotificationQueue interface {
Publish(ctx context.Context, id uuid.UUID) error
Consume(ctx context.Context) (<-chan uuid.UUID, error)
}
```

Можно начать с memory-очереди, потом заменить на Redis/Rabbit без изменения usecase.

5️⃣ worker/
```
type NotificationWorker struct {
repo     repository.MergeRequestRepository
queue    queue.NotificationQueue
telegram telegram.Client
}

func (w *NotificationWorker) Start(ctx context.Context) {
ch, _ := w.queue.Consume(ctx)

    for id := range ch {
        mr, _ := w.repo.GetByID(ctx, id)
        w.telegram.SendMergeRequestNotification(mr)
    }
}
```
6️⃣ telegram/

Клиент-обёртка над Telegram Bot API.
```
type Client interface {
SendMergeRequestNotification(mr *domain.MergeRequest) error
}
```
7️⃣ transport/http/

Handler только валидирует и передаёт в usecase:
```
func (h *Handler) CreateMR(w http.ResponseWriter, r *http.Request) {
var req CreateMRRequest
json.NewDecoder(r.Body).Decode(&req)

    err := h.usecase.Handle(r.Context(), req)
    ...
}
```