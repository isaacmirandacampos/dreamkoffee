include .env

start:
	go run cmd/finkoffee/main.go

build:
	go build -o bin/finkoffee cmd/finkoffee/main.go

test:
	go test ./...

mockgen:
	mockgen -source=internal/storage/persistence/db.go -destination=internal/tests/mocks/expense.sql.mocks.go -package=mocks

graphql:
	go run github.com/99designs/gqlgen generate

queries:
	sqlc generate

create_migration:
	migrate create -ext=sql -dir=internal/infrastructure/database/migration -seq $(n)

migrate_up:
	migrate -path=internal/infrastructure/database/migration -database "postgresql://${DB_USERNAME}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable" -verbose up

migrate_down:
	migrate -path=internal/infrastructure/database/migration -database "postgresql://${DB_USERNAME}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable" -verbose down
