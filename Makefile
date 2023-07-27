dev:
	TZ=UTC air -c .air.toml

test:
	DB_NAME=test go test ./...

init-db:
	docker-compose exec db dropdb local -U root; \
	docker-compose exec db createdb local -U root && \
	make migrate

init-test-db:
	docker-compose exec db dropdb test -U root; \
	docker-compose exec db createdb test -U root && \
	DB_NAME=test make migrate

migrate:
	go run ./cmd/migrate/main.go

wire:
	wire ./di/wire.go
