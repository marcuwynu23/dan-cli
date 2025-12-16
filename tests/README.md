# DAN Parser Test Suite

This directory contains comprehensive test cases for the DAN (Data Advanced Notation) parser and encoder.

## Test Files

### `dan_test.go`

Core unit tests covering:

- **Basic decoding**: Key-value pairs, numbers, booleans, strings
- **Nested blocks**: Simple and deeply nested object structures
- **Arrays**: Empty, string, number, and mixed arrays
- **Tables**: Simple tables, multi-column tables, empty tables, mixed types
- **Comments**: Hash comments (`#`), double-slash comments (`//`), inline and line comments
- **Input types**: String and `[]byte` input validation
- **Encoding**: Basic encoding, nested blocks, tables, arrays
- **Round-trip**: Encoding and decoding to ensure data integrity

### `integration_test.go`

Integration tests covering:

- **Complex documents**: Real-world scenarios with multiple features
- **Whitespace handling**: Empty lines and spacing
- **Edge cases**: Unclosed blocks/tables, empty strings, zero/negative/decimal numbers
- **Table column order**: Ensuring column order is preserved
- **Nested arrays in blocks**: Arrays within nested objects

### `benchmark_test.go`

Performance benchmarks:

- `BenchmarkDecode`: Decoding performance
- `BenchmarkEncode`: Encoding performance
- `BenchmarkRoundTrip`: Full encode/decode cycle
- `BenchmarkDecodeLargeTable`: Performance with large tables (1000 rows)

## Running Tests

Run all tests from the project root:

```bash
go test ./tests
```

Or with verbose output:

```bash
go test ./tests -v
```

Run specific test:

```bash
go test ./tests -v -run TestDecodeBasic
```

Run benchmarks:

```bash
go test ./tests -bench=. -benchmem
```

Run with coverage:

```bash
go test ./tests -cover
```

**Note**: Use `./tests` (with `./`) to reference the local directory. Using just `tests` will try to find a package in the standard library.

## Test Coverage

The test suite covers:

- ✅ Basic data types (string, number, boolean)
- ✅ Nested objects/blocks
- ✅ Arrays (empty, simple, mixed types)
- ✅ Tables (empty, simple, multi-column, mixed types)
- ✅ Comments (hash and double-slash)
- ✅ Input validation (string, []byte, invalid types)
- ✅ Encoding/decoding round-trips
- ✅ Edge cases (empty values, zero, negative numbers, decimals)
- ✅ Complex real-world documents

## Known Limitations

Some features are not fully supported and tests are commented out:

- Nested arrays (e.g., `[[1, 2], [3, 4]]`) - parser treats as strings
- Extra whitespace around colons - parser expects strict `key: value` format
- Mixed line endings (`\r\n`) - may not be handled correctly

These limitations are documented in the test files where applicable.
