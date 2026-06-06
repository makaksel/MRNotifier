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

