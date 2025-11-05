.PHONY: build run test clean migrate-up migrate-down migrate-create migrate-version migrate-force

# Build the application
build:
	go build -o bin/diro-be main.go

# Run the application
run:
	go run main.go

# Run tests
test:
	go test ./...

# Clean build artifacts
clean:
	rm -rf bin/

# Database migration commands
migrate-up:
	go run cmd/migrate/migrate.go -action=up

migrate-down:
	go run cmd/migrate/migrate.go -action=down

migrate-up-steps:
	go run cmd/migrate/migrate.go -action=up -steps=$(STEPS)

migrate-down-steps:
	go run cmd/migrate/migrate.go -action=down -steps=$(STEPS)

migrate-create:
	@if [ -z "$(NAME)" ]; then \
		echo "Error: NAME is required. Usage: make migrate-create NAME=migration_name"; \
		exit 1; \
	fi
	go run cmd/migrate/migrate.go -action=create -name=$(NAME)

migrate-version:
	go run cmd/migrate/migrate.go -action=version

migrate-force:
	@if [ -z "$(VERSION)" ]; then \
		echo "Error: VERSION is required. Usage: make migrate-force VERSION=version_number"; \
		exit 1; \
	fi
	go run cmd/migrate/migrate.go -action=force -version=$(VERSION)

# Development setup
dev-setup: migrate-up
	go mod tidy

# Docker commands (if needed)
docker-build:
	docker build -t diro-be .

docker-run:
	docker run -p 8080:8080 diro-be

# Help
help:
	@echo "Available commands:"
	@echo "  build              - Build the application"
	@echo "  run                - Run the application"
	@echo "  test               - Run tests"
	@echo "  clean              - Clean build artifacts"
	@echo "  migrate-up         - Run all pending migrations"
	@echo "  migrate-down       - Rollback all migrations"
	@echo "  migrate-up-steps   - Run N migrations up (use STEPS=N)"
	@echo "  migrate-down-steps - Rollback N migrations (use STEPS=N)"
	@echo "  migrate-create     - Create new migration files (use NAME=migration_name)"
	@echo "  migrate-version    - Show current migration version"
	@echo "  migrate-force      - Force migration to specific version (use VERSION=N)"
	@echo "  dev-setup          - Setup development environment (migrations + tidy)"
	@echo "  docker-build       - Build Docker image"
	@echo "  docker-run         - Run Docker container"
	@echo "  help               - Show this help"
