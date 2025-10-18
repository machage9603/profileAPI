.PHONY: run build test clean docker-build docker-run help

# Variables
BINARY_NAME=profile-api
DOCKER_IMAGE=profile-api:latest

# Default target
.DEFAULT_GOAL := help

## help: Display this help message
help:
	@echo "Available commands:"
	@echo "  make run          - Run the application"
	@echo "  make build        - Build the binary"
	@echo "  make test         - Run tests"
	@echo "  make test-verbose - Run tests with verbose output"
	@echo "  make test-cover   - Run tests with coverage"
	@echo "  make clean        - Remove binary and test cache"
	@echo "  make docker-build - Build Docker image"
	@echo "  make docker-run   - Run Docker container"
	@echo "  make lint         - Run golangci-lint"
	@echo "  make fmt          - Format code"

## run: Run the application locally
run:
	@echo "Starting application..."
	go run main.go

## build: Build the binary
build:
	@echo "Building binary..."
	go build -o $(BINARY_NAME) main.go
	@echo "Binary created: $(BINARY_NAME)"

## test: Run all tests
test:
	@echo "Running tests..."
	go test -v ./...

## test-verbose: Run tests with verbose output
test-verbose:
	@echo "Running tests with verbose output..."
	go test -v -race ./...

## test-cover: Run tests with coverage
test-cover:
	@echo "Running tests with coverage..."
	go test -v -cover -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report: coverage.html"

## clean: Remove binary and clean test cache
clean:
	@echo "Cleaning..."
	go clean
	rm -f $(BINARY_NAME)
	rm -f coverage.out coverage.html
	@echo "Clean complete"

## docker-build: Build Docker image
docker-build:
	@echo "Building Docker image..."
	docker build -t $(DOCKER_IMAGE) .
	@echo "Docker image built: $(DOCKER_IMAGE)"

## docker-run: Run Docker container
docker-run:
	@echo "Running Docker container..."
	docker run -p 8080:8080 \
		-e USER_EMAIL="your@email.com" \
		-e USER_NAME="Your Name" \
		-e USER_STACK="Go/Native HTTP" \
		$(DOCKER_IMAGE)

## fmt: Format Go code
fmt:
	@echo "Formatting code..."
	go fmt ./...

## lint: Run linter (requires golangci-lint)
lint:
	@echo "Running linter..."
	golangci-lint run

## dev: Run with auto-reload (requires air)
dev:
	@echo "Starting development server with hot reload..."
	air

## install-tools: Install development tools
install-tools:
	@echo "Installing development tools..."
	go install github.com/cosmtrek/air@latest
	@echo "Tools installed"