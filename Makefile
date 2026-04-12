.PHONY: build test test-verbose test-coverage bench clean install

# Build the CLI
# Creates bin/dan on Unix, bin/dan.exe on Windows
build:
	mkdir -p bin
	go build -o bin/dan cmd/dan/main.go

# Run tests (godan module — ../dango checkout in monorepo)
test:
	go -C ../dango test ./...

# Run tests with verbose output
test-verbose:
	go -C ../dango test ./... -v

# Run tests with coverage
test-coverage:
	go -C ../dango test ./... -cover

# Run benchmarks
bench:
	go -C ../dango test ./... -bench=. -benchmem

# Install the CLI
install:
	go install ./cmd/dan

# Clean build artifacts
clean:
	rm -rf bin
