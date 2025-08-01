.PHONY: help test build clean

help:
	@echo "Available commands:"
	@echo "  test    - Run all tests"
	@echo "  build   - Build the solution"
	@echo "  clean   - Clean build artifacts"

test:
	go test -v

build:
	go build -o solution .

clean:
	rm -f solution
	go clean