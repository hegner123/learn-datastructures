.PHONY: help test build clean

# Default target - show available commands
help:
	@echo "Available commands:"
	@echo "  make test  - Run all tests"
	@echo "  make build - Build the solution"
	@echo "  make clean - Clean build artifacts"
	@echo "  make help  - Show this help message"

# Run tests with verbose output
test:
	go test -v

# Build the solution
build:
	go build -o solution .

# Clean build artifacts
clean:
	go clean
	rm -f solution