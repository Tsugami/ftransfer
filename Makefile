.PHONY: build run test clean migrate-up migrate-down lint

# Variables
BINARY_NAME=./tmp/ftransfer
DB_NAME=ftransfer
DB_USER=user
DB_PASSWORD=pass
DB_HOST=localhost
DB_PORT=5432
MIGRATION_DIR=./migrations
DATABASE_URL=postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable

# Build the application
build:
	go build -o $(BINARY_NAME) ./cmd/server

# Run the application
run:
	go run ./cmd/server

# Run tests
test:
	go test -v ./...

# Clean build artifacts
clean:
	rm -f $(BINARY_NAME)

# Run database migrations up
migrate-up:
	migrate -path $(MIGRATION_DIR) -database "postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable" up

# Run database migrations down
migrate-down:
	migrate -path $(MIGRATION_DIR) -database "postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable" down

migrate-create:
	migrate create -ext sql -dir $(MIGRATION_DIR) -seq $(name)

# Run linter
lint:
	golangci-lint run

# Install dependencies
deps:
	go mod download

# Install development tools
dev-tools:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
	go install github.com/air-verse/air@latest


docker_dev:
	docker compose -f docker-compose.dev.yml up -d

docker_dev_down:
	docker compose -f docker-compose.dev.yml down -v

docker_dev_logs:
	docker compose -f docker-compose.dev.yml logs -f

psql_dev:
	docker exec -it ftransfer_postgres_dev psql -U $(DB_USER) -d $(DB_NAME)

dev:
	DATABASE_URL=$(DATABASE_URL) air -c .air.toml


pgsql:
	docker exec -it ftransfer_postgres psql -U $(DB_USER) -d $(DB_NAME)

test_dockerfile_build:
	docker compose -f docker-compose.prod.yaml build

test_dockerfile_up:
	docker compose -f docker-compose.prod.yaml up

test_dockerfile_down:
	docker compose -f docker-compose.prod.yaml down

test_dockerfile_psql:
	docker exec -it ftransfer-postgres-prod psql -U ftransfer -d ftransfer

clean_docker_volumes: test_dockerfile_down
	docker compose -f docker-compose.prod.yaml down -v
