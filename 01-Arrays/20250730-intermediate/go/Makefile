.PHONY: help test build clean

help:
	@echo "Available commands:"
	@echo "  make test  - Run all tests"
	@echo "  make build - Build the solution"
	@echo "  make clean - Clean build artifacts"

test:
	go test -v

build:
	go build -o solution solution.go

clean:
	rm -f solution
	go clean