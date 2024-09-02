include .env

start:
	go run cmd/dreamkoffee/main.go

build:
	go build -o bin/dreamkoffee cmd/dreamkoffee/main.go

test:
	go test ./...

test_one:
	go test ./${p}

generate_mock:
	mockgen -source internal/domain/repository.go -destination internal/infrastructure/database/mocks/repository.go -package=mocks

generate_graph:
	go run github.com/99designs/gqlgen generate

generate_query:
	sqlc generate

create_migration:
	migrate create -ext=sql -dir=internal/infrastructure/database/postgres/migration -seq $(n)

migrate_up:
	migrate -path=internal/infrastructure/database/postgres/migration -database "postgresql://${DB_USERNAME}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable" -verbose up

migrate_down:
	migrate -path=internal/infrastructure/database/postgres/migration -database "postgresql://${DB_USERNAME}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable" -verbose down
