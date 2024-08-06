.PHONY: all build run clean

# Default target
all: build

# Build for the current platform
build:
	go build -o eventful

# Run the application
run: build
	./eventful

# Clean build artifacts
clean:
	rm -f eventful