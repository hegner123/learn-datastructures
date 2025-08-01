.PHONY: help test build clean

# Default target
help:
	@echo "Available commands:"
	@echo "  make test   - Run tests"
	@echo "  make build  - Build the project"
	@echo "  make clean  - Clean build artifacts"
	@echo "  make help   - Show this help message"

test:
	go test -v

build:
	go build -o calculator

clean:
	rm -f calculator
	go clean