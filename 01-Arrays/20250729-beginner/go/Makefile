.PHONY: help test build clean

# Default target shows available commands
help:
	@echo "Available commands:"
	@echo "  make test  - Run all tests"
	@echo "  make build - Build the solution"
	@echo "  make clean - Clean build artifacts"

# Run tests
test:
	go test -v

# Build the solution
build:
	go build -o solution

# Clean build artifacts
clean:
	rm -f solution
	go clean