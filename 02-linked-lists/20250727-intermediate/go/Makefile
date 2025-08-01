.PHONY: help test build clean

help:
	@echo "Available commands:"
	@echo "  make test    - Run all tests"
	@echo "  make build   - Build the project"
	@echo "  make clean   - Clean build artifacts"

test:
	go test -v

build:
	go build -o solution .

clean:
	go clean
	rm -f solution