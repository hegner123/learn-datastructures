.PHONY: help test build clean

# Default target
help:
	@echo "Available commands:"
	@echo "  make test    - Run all tests"
	@echo "  make build   - Build the application"
	@echo "  make clean   - Clean build artifacts"
	@echo "  make help    - Show this help message"

test:
	go test -v

build:
	go build -o balanced-parentheses

clean:
	go clean
	rm -f balanced-parentheses