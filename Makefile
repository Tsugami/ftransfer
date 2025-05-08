.PHONY: build run test clean migrate-up migrate-down lint

# Variables
BINARY_NAME=./tmp/ftransfer
DB_NAME=ftransfer
DB_USER=user
DB_PASSWORD=pass
DB_HOST=localhost
DB_PORT=5432
MIGRATION_DIR=./migrations

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


dev:
	air -c .air.toml


pgsql:
	docker exec -it ftransfer_postgres psql -U $(DB_USER) -d $(DB_NAME)


test_dockerfile_build:
	docker compose -f docker-compose.prod.yaml build

test_dockerfile_up:
	docker compose -f docker-compose.prod.yaml up

test_dockerfile_down:
	docker compose -f docker-compose.prod.yaml down

clean_docker_volumes:
	docker compose -f docker-compose.prod.yaml down -v
