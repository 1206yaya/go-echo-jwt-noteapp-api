
.PHONY: run run-db stop-db migrate


run:
	GO_ENV=dev go run main.go

run-db:
	docker compose up -d

stop-db:
	docker compose down

migrate:
	GO_ENV=dev go run migrate/migrate.go
