.PHONY: all deps build run clean lint package

# Default target
all: deps build

# Install or update dependencies
deps:
	@echo "Updating Go dependencies..."
	go mod tidy
	@echo "Updating npm dependencies..."
	cd frontend && npm install

# Build for the current platform
build: deps
	@echo "Building Eventful..."
	wails build

# Run the application in development mode
run:
	@echo "Running Eventful in development mode..."
	wails run

# Clean build artifacts
clean:
	@echo "Cleaning build artifacts..."
	rm -rf build
	cd frontend && npm run clean

# Lint the project
lint:
	@echo "Linting Go code..."
	golangci-lint run
	@echo "Linting Svelte code..."
	cd frontend && npm run lint

# Package the application for distribution
package:
	@echo "Packaging Eventful for distribution..."
	wails build -platform windows/amd64,darwin/amd64,linux/amd64

# Install necessary tools
install-tools:
	@echo "Installing necessary tools..."
	go install github.com/wailsapp/wails/v2/cmd/wails@latest
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	npm install -g svelte-check

# Help command to display available targets
help:
	@echo "Available targets:"
	@echo "  all       - Update dependencies and build the application (default)"
	@echo "  deps      - Install or update dependencies"
	@echo "  build     - Build the application"
	@echo "  run       - Run the application in development mode"
	@echo "  clean     - Clean build artifacts"
	@echo "  lint      - Lint the project"
	@echo "  package   - Package the application for distribution"
	@echo "  install-tools - Install necessary development tools"
	@echo "  help      - Display this help message"