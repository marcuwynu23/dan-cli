package tests

import (
	"testing"

	"dan-cli/dan"
)

var benchmarkInput = `app {
  name: "MyApp"
  version: "1.0.0"
  settings {
    debug: true
    port: 8080
    hosts: ["localhost", "127.0.0.1"]
  }
}

users: table(id, name, email, active) [
  1, "Alice", "alice@example.com", true
  2, "Bob", "bob@example.com", true
  3, "Charlie", "charlie@example.com", false
  4, "David", "david@example.com", true
  5, "Eve", "eve@example.com", false
]

features: ["auth", "logging", "metrics", "caching"]
enabled: true
`

func BenchmarkDecode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := dan.Decode(benchmarkInput)
		if err != nil {
			b.Fatalf("Decode() error = %v", err)
		}
	}
}

func BenchmarkEncode(b *testing.B) {
	decoded, _ := dan.Decode(benchmarkInput)
	b.ResetTimer()
	
	for i := 0; i < b.N; i++ {
		_ = dan.Encode(decoded)
	}
}

func BenchmarkRoundTrip(b *testing.B) {
	for i := 0; i < b.N; i++ {
		decoded, err := dan.Decode(benchmarkInput)
		if err != nil {
			b.Fatalf("Decode() error = %v", err)
		}
		encoded := dan.Encode(decoded)
		_, err = dan.Decode(encoded)
		if err != nil {
			b.Fatalf("Second Decode() error = %v", err)
		}
	}
}

func BenchmarkDecodeLargeTable(b *testing.B) {
	// Generate a large table
	largeInput := "users: table(id, name, email) [\n"
	for i := 0; i < 1000; i++ {
		largeInput += "  " + string(rune('0'+i%10)) + ", \"User" + string(rune('0'+i%10)) + "\", \"user" + string(rune('0'+i%10)) + "@example.com\"\n"
	}
	largeInput += "]"
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := dan.Decode(largeInput)
		if err != nil {
			b.Fatalf("Decode() error = %v", err)
		}
	}
}
