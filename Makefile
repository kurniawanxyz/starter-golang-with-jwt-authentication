# Load environment variables from .env file
ifneq (,$(wildcard ./.env))
	include .env
	export
endif

# Tools
GOLANGCI_LINT=golangci-lint
MIGRATE_CMD=migrate
MIGRATIONS_DIR=infrastructure/migrations

# Database URL (PostgreSQL example)
DB_URL=mysql://$(DB_USER):$(DB_PASS)@tcp($(DB_HOST):$(DB_PORT))/$(DB_NAME)

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

# Force migrate to a specific version (0 for clean state)
.PHONY: migrate-force
migrate-force:
	@read -p "Enter the migration version to force: " version; \
	echo "Forcing migration to version $$version..."; \
	$(MIGRATE_CMD) -path $(MIGRATIONS_DIR) -database "$(DB_URL)" force $$version

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

# Check migration version
.PHONY: migrate-version
migrate-version:
	@echo "Checking current migration version..."
	$(MIGRATE_CMD) -path $(MIGRATIONS_DIR) -database "$(DB_URL)" version

.PHONY: run
run:
	@echo "Running application..."
	go run cmd/app/main.go