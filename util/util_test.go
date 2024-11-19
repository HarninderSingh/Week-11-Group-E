package util

import (
	"testing"
)

// TestIsPrime with table-driven tests
func TestIsPrime(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected bool
	}{
		{"Negative number", -1, false},
		{"Zero", 0, false},
		{"One", 1, false},
		{"Prime number", 7, true},
		{"Non-prime number", 8, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsPrime(tt.input)
			if result != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func BenchmarkIsPrime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsPrime(101)
	}
}