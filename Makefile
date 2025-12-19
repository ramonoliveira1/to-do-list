.PHONY: help build run test clean docker-build docker-up docker-down install lint fmt

# Variables
APP_NAME=api
MAIN_PATH=cmd/api/main.go
BUILD_DIR=bin

help: ## Show this help message
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Available targets:'
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  %-15s %s\n", $$1, $$2}' $(MAKEFILE_LIST)

install: ## Install dependencies
	go mod download
	go mod tidy

build: ## Build the application
	@echo "Building..."
	@mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(APP_NAME) $(MAIN_PATH)
	@echo "Build complete: $(BUILD_DIR)/$(APP_NAME)"

build-prod: ## Build for production (optimized)
	@echo "Building for production..."
	@mkdir -p $(BUILD_DIR)
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o $(BUILD_DIR)/$(APP_NAME) $(MAIN_PATH)
	@echo "Production build complete: $(BUILD_DIR)/$(APP_NAME)"

run: ## Run the application
	go run $(MAIN_PATH)

test: ## Run tests
	go test -v ./...

test-coverage: ## Run tests with coverage
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

lint: ## Run linter
	@which golangci-lint > /dev/null || (echo "golangci-lint not installed. Install it from https://golangci-lint.run/usage/install/" && exit 1)
	golangci-lint run ./...

fmt: ## Format code
	go fmt ./...
	gofmt -s -w .

vet: ## Run go vet
	go vet ./...

clean: ## Clean build artifacts
	@echo "Cleaning..."
	rm -rf $(BUILD_DIR)
	rm -f coverage.out coverage.html
	@echo "Clean complete"

docker-build: ## Build Docker image
	docker build -t todo-api:latest .

docker-up: ## Start Docker containers
	docker-compose up -d
	@echo "Containers started. API available at http://localhost:8080"

docker-down: ## Stop Docker containers
	docker-compose down

docker-logs: ## Show Docker logs
	docker-compose logs -f api

db-migrate: ## Run database migrations
	mysql -u$(DB_USER) -p$(DB_PASSWORD) -h$(DB_HOST) $(DB_NAME) < schema.sql

dev: ## Run in development mode with hot reload (requires air)
	@which air > /dev/null || (echo "air not installed. Install it: go install github.com/cosmtrek/air@latest" && exit 1)
	air

all: clean fmt vet lint test build ## Run all checks and build

