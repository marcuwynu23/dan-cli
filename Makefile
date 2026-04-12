.PHONY: build test test-verbose test-coverage bench clean install

# Build the CLI
# Creates bin/dan on Unix, bin/dan.exe on Windows
build:
	mkdir -p bin
	go build -o bin/dan cmd/dan/main.go

# Run tests (library lives in ../dan-go)
test:
	go -C ../dan-go test ./...

# Run tests with verbose output
test-verbose:
	go -C ../dan-go test ./... -v

# Run tests with coverage
test-coverage:
	go -C ../dan-go test ./... -cover

# Run benchmarks
bench:
	go -C ../dan-go test ./... -bench=. -benchmem

# Install the CLI
install:
	go install ./cmd/dan

# Clean build artifacts
clean:
	rm -rf bin
