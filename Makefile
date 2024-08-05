.PHONY: all build run clean windows linux darwin

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
	rm -f eventful eventful.exe

# Build for Windows
windows:
	fyne-cross windows

# Build for Linux
linux:
	fyne-cross linux

# Build for macOS
darwin:
	fyne-cross darwin

# Run the Windows executable using Wine (make sure Wine is installed)
run-windows: windows
	wine fyne-cross/bin/windows-amd64/eventful.exe

# For running Linux build on WSL, just use the normal run target

# For macOS, you can't directly run it on WSL, but you can build it