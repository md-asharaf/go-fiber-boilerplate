# Go Fiber Boilerplate Makefile

# Variables
APP_NAME := go-fiber-boilerplate
BINARY_NAME := $(APP_NAME)
DOCKER_IMAGE := $(APP_NAME):latest
GO_VERSION := 1.21

# Build info
BUILD_TIME := $(shell date -u +%Y%m%d.%H%M%S)
GIT_COMMIT := $(shell git rev-parse --short HEAD 2>/dev/null || echo "unknown")
VERSION := $(shell git describe --tags --exact-match 2>/dev/null || echo "dev")

# Paths
CMD_PATH := ./cmd/server
BUILD_PATH := ./build
MAIN_FILE := $(CMD_PATH)/main.go

# Go build flags
LDFLAGS := -ldflags "-X main.version=$(VERSION) -X main.buildTime=$(BUILD_TIME) -X main.gitCommit=$(GIT_COMMIT)"

.PHONY: help build run test test-coverage clean fmt lint vet deps security docker-build docker-run docker-compose-up docker-compose-down generate-mocks install-tools

# Default target
help: ## Show this help message
	@echo "$(APP_NAME) - Available commands:"
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  \033[36m%-20s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

# Development
run: ## Run the application
	@echo "Running $(APP_NAME)..."
	@go run $(MAIN_FILE)

dev: ## Run the application with air for hot reloading
	@echo "Running $(APP_NAME) with hot reload..."
	@air

build: ## Build the application
	@echo "Building $(APP_NAME)..."
	@mkdir -p $(BUILD_PATH)
	@go build $(LDFLAGS) -o $(BUILD_PATH)/$(BINARY_NAME) $(MAIN_FILE)
	@echo "Binary built: $(BUILD_PATH)/$(BINARY_NAME)"

clean: ## Clean build artifacts
	@echo "Cleaning build artifacts..."
	@rm -rf $(BUILD_PATH)
	@go clean

# Dependencies
deps: ## Download dependencies
	@echo "Downloading dependencies..."
	@go mod download

tidy: ## Tidy up dependencies
	@echo "Tidying up dependencies..."
	@go mod tidy

vendor: ## Vendor dependencies
	@echo "Vendoring dependencies..."
	@go mod vendor

# Code quality
fmt: ## Format Go code
	@echo "Formatting code..."
	@go fmt ./...

lint: ## Run golangci-lint
	@echo "Running linter..."
	@golangci-lint run

vet: ## Run go vet
	@echo "Running go vet..."
	@go vet ./...

security: ## Run gosec security scanner
	@echo "Running security scanner..."
	@gosec ./...

# Testing
test: ## Run tests
	@echo "Running tests..."
	@go test -v ./...

test-coverage: ## Run tests with coverage
	@echo "Running tests with coverage..."
	@go test -v -coverprofile=coverage.out ./...
	@go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

test-integration: ## Run integration tests
	@echo "Running integration tests..."
	@go test -v -tags=integration ./tests/...

benchmark: ## Run benchmarks
	@echo "Running benchmarks..."
	@go test -bench=. -benchmem ./...

# Database
db-migrate-up: ## Run database migrations up
	@echo "Running database migrations up..."
	@migrate -path migrations -database "$(DB_URL)" up

db-migrate-down: ## Run database migrations down
	@echo "Running database migrations down..."
	@migrate -path migrations -database "$(DB_URL)" down

db-migrate-create: ## Create a new migration file (usage: make db-migrate-create name=create_users_table)
	@echo "Creating migration: $(name)"
	@migrate create -ext sql -dir migrations $(name)

# Docker
docker-build: ## Build Docker image
	@echo "Building Docker image..."
	@docker build -t $(DOCKER_IMAGE) .

docker-run: ## Run Docker container
	@echo "Running Docker container..."
	@docker run -p 8000:8000 --env-file .env $(DOCKER_IMAGE)

docker-compose-up: ## Start services with docker-compose
	@echo "Starting services with docker-compose..."
	@docker-compose up --build

docker-compose-down: ## Stop services with docker-compose
	@echo "Stopping services with docker-compose..."
	@docker-compose down

# Tools installation
install-tools: ## Install development tools
	@echo "Installing development tools..."
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	@go install github.com/securecodewarrior/gosec/v2/cmd/gosec@latest
	@go install github.com/cosmtrek/air@latest
	@go install github.com/golang-migrate/migrate/v4/cmd/migrate@latest
	@go install go.uber.org/mock/mockgen@latest

# Code generation
generate: ## Generate code (mocks, docs, etc.)
	@echo "Generating code..."
	@go generate ./...

generate-mocks: ## Generate mocks for testing
	@echo "Generating mocks..."
	@mockgen -source=internal/services/interfaces.go -destination=tests/mocks/services.go

generate-docs: ## Generate API documentation
	@echo "Generating API documentation..."
	@swag init -g cmd/server/main.go -o docs/swagger

# Release
release: test lint build ## Run tests, lint, and build for release
	@echo "Release build completed successfully!"

# Git hooks
install-hooks: ## Install git hooks
	@echo "Installing git hooks..."
	@cp scripts/hooks/* .git/hooks/
	@chmod +x .git/hooks/*

# Health check
health: ## Check if the application is healthy
	@echo "Checking application health..."
	@curl -f http://localhost:8000/api/v1/health || exit 1
