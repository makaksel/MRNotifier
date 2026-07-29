include .env
export

MIGRATIONS_PATH=./migrations

.PHONY: up down logs migrate-up migrate-down migrate-create restart

up:
	docker compose up --build

down:
	docker compose down

build:
	docker compose build

logs:
	docker compose logs -f

restart:
	docker compose down && docker compose up --build

migrate-up:
	docker compose run --rm migrate up

migrate-down:
	docker run --rm -v $(PWD)/migrations:/migrations migrate/migrate \
	-database postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@postgres:5432/$(POSTGRES_DB)?sslmode=disable \
	-path /migrations down 1

migrate-create:
	@read -p "name: " name; \
	migrate create -ext sql -dir $(MIGRATIONS_PATH) -seq $$name