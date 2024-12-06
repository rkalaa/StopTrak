# Variables
APP_NAME = gin-api
BUILD_DIR = build
DOCKER_IMAGE = gin-api:latest

# Default target
.PHONY: all
all: build

# Build the Go application
.PHONY: build
build:
	go build -o $(BUILD_DIR)/$(APP_NAME) .

# Run the Go application
.PHONY: run
run:
	go run main.go

# Run tests
.PHONY: test
test:
	go test ./... -v

# Clean build artifacts
.PHONY: clean
clean:
	rm -rf $(BUILD_DIR)

# Docker build
.PHONY: docker-build
docker-build:
	docker build -t $(DOCKER_IMAGE) .

# Docker run
.PHONY: docker-run
docker-run:
	docker run -p 8080:8080 $(DOCKER_IMAGE)

# Lint the code (requires golangci-lint installed)
.PHONY: lint
lint:
	golangci-lint run ./...

# Format the code
.PHONY: fmt
fmt:
	go fmt ./...

# Tidy dependencies
.PHONY: tidy
tidy:
	go mod tidy