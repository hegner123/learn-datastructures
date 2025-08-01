.PHONY: help test build clean

# Default target
help:
	@echo "Available commands:"
	@echo "  make test    - Run all tests"
	@echo "  make build   - Build the solution"
	@echo "  make clean   - Clean build artifacts"
	@echo "  make help    - Show this help message"

test:
	go test -v

build:
	go build -o solution

clean:
	go clean
	rm -f solution