.PHONY: help test build clean

help:
	@echo "Available commands:"
	@echo "  make test    - Run all tests"
	@echo "  make build   - Build the application"
	@echo "  make clean   - Clean build artifacts"
	@echo "  make help    - Show this help message"

test:
	go test -v

build:
	go build -o bst-intro

clean:
	rm -f bst-intro
	go clean