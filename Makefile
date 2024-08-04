# Makefile

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOFMT=$(GOCMD) fmt
GOVET=$(GOCMD) vet
BINARY_NAME=eventful

# Linter
GOLINT=golangci-lint

# Redis parameters
REDIS_PORT=6379

# Tailwind CSS
TAILWIND_CLI=npx tailwindcss

# Check if Redis is installed
REDIS_INSTALLED := $(shell command -v redis-server 2> /dev/null)

.PHONY: all build test clean run deps lint redis-start redis-stop tailwind-init tailwind-build

all: deps redis-start lint test tailwind-build build run

build:
	$(GOBUILD) -o $(BINARY_NAME) -v

test:
	$(GOTEST) -v ./...

clean:
	rm -f $(BINARY_NAME)
	rm -f coverage.out

run: build tailwind-build
	./$(BINARY_NAME)

deps:
	$(GOGET) github.com/gin-gonic/gin
	$(GOGET) github.com/go-redis/redis/v8
	$(GOGET) github.com/google/uuid
	$(GOGET) golang.org/x/crypto/bcrypt
	$(GOGET) github.com/stretchr/testify
	@if [ ! -f $$GOPATH/bin/golangci-lint ]; then \
		curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $$(go env GOPATH)/bin; \
	fi
	@if [ -z "$(REDIS_INSTALLED)" ]; then \
		echo "Redis is not installed. Installing Redis..."; \
		if [ $$(uname) = "Darwin" ]; then \
			brew install redis; \
		elif [ $$(uname) = "Linux" ]; then \
			sudo apt-get update && sudo apt-get install redis-server; \
		else \
			echo "Unsupported operating system. Please install Redis manually."; \
			exit 1; \
		fi; \
	fi
	@if [ ! -f package.json ]; then \
		npm init -y; \
	fi
	npm install tailwindcss

lint:
	$(GOFMT) ./...
	$(GOVET) ./...
	$(GOLINT) run

redis-start:
	@if [ -z "$$(pgrep redis-server)" ]; then \
		redis-server --daemonize yes --port $(REDIS_PORT); \
		echo "Redis server started on port $(REDIS_PORT)"; \
	else \
		echo "Redis server is already running"; \
	fi

redis-stop:
	@if [ -n "$$(pgrep redis-server)" ]; then \
		redis-cli shutdown; \
		echo "Redis server stopped"; \
	else \
		echo "Redis server is not running"; \
	fi

tailwind-init:
	@if [ ! -f tailwind.config.js ]; then \
		$(TAILWIND_CLI) init; \
	fi

tailwind-build: tailwind-init
	$(TAILWIND_CLI) -i ./static/css/input.css -o ./static/css/output.css --minify

coverage:
	$(GOTEST) -coverprofile=coverage.out ./...
	$(GOCMD) tool cover -html=coverage.out

# This ensures that we use the internal definition of the rules, not any external commands
.PHONY: all build test clean run deps lint redis-start redis-stop tailwind-init tailwind-build coverage