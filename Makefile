# Load environment variables from .env file
ifneq (,$(wildcard ./.env))
	include .env
	export
endif
# Tools
GOLANGCI_LINT=golangci-lint
MIGRATE_CMD=migrate
MIGRATIONS_DIR=./infrastructure/migrations

# Database URL (PostgreSQL example)
DB_URL=postgres://$(DB_USER):$(DB_PASS)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable

# Run migrations
.PHONY: migrate-up
migrate-up:
	@echo "Running database migrations..."
	$(MIGRATE_CMD) -path $(MIGRATIONS_DIR) -database "$(DB_URL)" up

# Rollback migrations (1 step)
.PHONY: migrate-down
migrate-down:
	@echo "Rolling back last database migration..."
	$(MIGRATE_CMD) -path $(MIGRATIONS_DIR) -database "$(DB_URL)" down 1

# Rollback all migrations
.PHONY: migrate-reset
migrate-reset:
	@echo "Resetting database migrations..."
	$(MIGRATE_CMD) -path $(MIGRATIONS_DIR) -database "$(DB_URL)" down
	$(MIGRATE_CMD) -path $(MIGRATIONS_DIR) -database "$(DB_URL)" up

# Create new migration
.PHONY: migrate-create
migrate-create:
	@echo "Creating new migration..."
	$(MIGRATE_CMD) create -dir $(MIGRATIONS_DIR) -ext sql $(name)

# Seed database
.PHONY: seed
seed:
	@echo "Seeding database..."
	go run main.go seed

# Linting
.PHONY: lint
lint:
	@echo "Running golangci-lint..."
	$(GOLANGCI_LINT) run ./...

# Clean generated binaries (optional)
.PHONY: clean
clean:
	@echo "Cleaning up..."
	go clean
