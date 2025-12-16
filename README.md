# DAN CLI - Data Advanced Notation Command Line Tool

A Go implementation of the DAN (Data Advanced Notation) parser and encoder with a command-line interface.

## Features

- **Parse DAN format** - Convert DAN text to structured data
- **Encode to DAN** - Convert structured data (JSON or DAN) to DAN format
- **Pretty-print** - Format DAN files for better readability
- **JSON support** - Convert between DAN and JSON formats

## Installation

```bash
go build -o dan cmd/dan/main.go
```

Or install directly:

```bash
go install ./cmd/dan
```

## Usage

### Decode DAN

Convert DAN format to JSON or pretty-printed DAN:

```bash
# Decode from file
dan decode file.dan

# Decode from stdin
cat file.dan | dan decode

# Output as JSON
dan decode -json file.dan
```

### Encode to DAN

Convert JSON or DAN to DAN format:

```bash
# Encode JSON file to DAN
dan encode -json file.json

# Encode DAN file (pretty-print)
dan encode file.dan
```

### Pretty-print

Format DAN files:

```bash
dan pretty file.dan
```

## DAN Format

DAN (Data Advanced Notation) is a human-readable data format that supports:

- **Key-value pairs**: `key: value`
- **Nested objects**: `key { ... }`
- **Tables**: `key: table(col1, col2) [ ... ]`
- **Arrays**: `key: [value1, value2]`
- **Comments**: `# comment` or `// comment`
- **Data types**: strings, numbers, booleans, arrays

### Example

```dan
name: "John Doe"
age: 30
active: true
tags: ["developer", "go"]

address {
  street: "123 Main St"
  city: "New York"
  zip: 10001
}

users: table(name, age, email) [
  "Alice", 25, "alice@example.com"
  "Bob", 30, "bob@example.com"
]
```

## Library Usage

You can also use the DAN library in your own Go programs:

```go
package main

import (
    "fmt"
    "dan-cli/dan"
)

func main() {
    // Decode DAN text
    data, err := dan.Decode(`name: "test"`)
    if err != nil {
        panic(err)
    }
    fmt.Println(data)

    // Encode to DAN
    obj := map[string]interface{}{
        "name": "test",
        "age": 30,
    }
    output := dan.Encode(obj)
    fmt.Println(output)
}
```

## Project Structure

```
dan-cli/
├── cmd/
│   └── dan/
│       └── main.go      # CLI application
├── dan/
│   └── dan.go           # DAN parser/encoder library
├── go.mod               # Go module file
└── README.md           # This file
```

## License

MIT
