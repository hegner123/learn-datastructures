.PHONY: help test build clean

help:
	@echo "Available commands:"
	@echo "  test  - Run all tests"
	@echo "  build - Build the project"
	@echo "  clean - Clean build artifacts"

test:
	go test -v

build:
	go build -o bst-intro

clean:
	rm -f bst-intro
	go clean