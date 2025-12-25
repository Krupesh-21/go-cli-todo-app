# Makefile for go-cli-todo-app - Cross-platform Go CLI Todo App

BINARY_NAME=go-cli-todo-app
SRC_FILES=$(wildcard *.go)

.PHONY: all build run clean windows linux darwin help examples

all: build

build:
	@echo "Building for current OS..."
	go build -o $(BINARY_NAME) $(SRC_FILES)

run: build
	./$(BINARY_NAME) $(ARGS)

# Cross-compilation targets
windows:
	GOOS=windows GOARCH=amd64 go build -o $(BINARY_NAME).exe $(SRC_FILES)

linux:
	GOOS=linux GOARCH=amd64 go build -o $(BINARY_NAME)-linux $(SRC_FILES)

darwin:
	GOOS=darwin GOARCH=arm64 go build -o $(BINARY_NAME)-macos $(SRC_FILES)

# Build all platforms
all-platforms: windows linux darwin

clean:
	rm -f $(BINARY_NAME) $(BINARY_NAME).exe $(BINARY_NAME)-linux $(BINARY_NAME)-macos

help:
	@echo "Available targets:"
	@echo "  make build      - Build for current OS"
	@echo "  make run ARGS=\"add Buy milk\"  - Build and run"
	@echo "  make windows    - Build for Windows"
	@echo "  make linux      - Build for Linux"
	@echo "  make darwin     - Build for macOS"
	@echo "  make all-platforms - Build all platforms"
	@echo "  make clean      - Remove binaries"
	@echo "  make help       - Show this help"
	@echo "  make examples   - Show CLI command examples"

examples:
	@echo "=== CLI Todo App Examples ==="
	@echo ""
	@echo "make run ARGS=\"add Buy groceries\""
	@echo "make run ARGS=\"add Walk the dog\""
	@echo ""
	@echo "make run ARGS=\"list\""
	@echo ""
	@echo "make run ARGS=\"complete 1\""
	@echo ""
	@echo "make run ARGS=\"remove 2\""
	@echo ""
	@echo "make run ARGS=\"list\""
	@echo ""
	@echo "=== Or use the built binary directly ==="
	@echo "./$(BINARY_NAME) add \"Buy groceries\""
	@echo "./$(BINARY_NAME) list"
	@echo "./$(BINARY_NAME) complete 1"
	@echo "./$(BINARY_NAME) remove 2"