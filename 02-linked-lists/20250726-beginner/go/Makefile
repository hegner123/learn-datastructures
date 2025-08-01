.PHONY: help test build clean

# Default target - show available commands
help:
	@echo "Available commands:"
	@echo "  make test    - Run all tests"
	@echo "  make build   - Build the project"
	@echo "  make clean   - Remove build artifacts"
	@echo "  make help    - Show this help message"

# Run tests
test:
	go test -v

# Build the project
build:
	go build -o solution

# Clean build artifacts
clean:
	rm -f solution
	go clean

# Default target
.DEFAULT_GOAL := help