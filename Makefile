.PHONY: build test test-verbose test-coverage bench clean install

# Build the CLI
# Creates bin/dan on Unix, bin/dan.exe on Windows
build:
	mkdir -p bin
	go build -o bin/dan cmd/dan/main.go

# Run tests
test:
	go test ./tests

# Run tests with verbose output
test-verbose:
	go test ./tests -v

# Run tests with coverage
test-coverage:
	go test ./tests -cover

# Run benchmarks
bench:
	go test ./tests -bench=. -benchmem

# Install the CLI
install:
	go install ./cmd/dan

# Clean build artifacts
clean:
	rm -rf bin
