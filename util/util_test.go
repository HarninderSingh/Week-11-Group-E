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

// TestReverseString with edge cases
func TestReverseString(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"Empty string", "", ""},
		{"Single character", "a", "a"},
		{"Palindrome", "madam", "madam"},
		{"Regular string", "hello", "olleh"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ReverseString(tt.input)
			if result != tt.expected {
				t.Errorf("expected %s, got %s", tt.expected, result)
			}
		})
	}
}

// Benchmark ReverseString
func BenchmarkReverseString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ReverseString("Benchmarking")
	}
}

// TestSumOfSlice with edge cases
func TestSumOfSlice(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{"Empty slice", []int{}, 0},
		{"Single element", []int{5}, 5},
		{"Multiple elements", []int{1, 2, 3, 4}, 10},
		{"Negative numbers", []int{-1, -2, -3}, -6},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SumOfSlice(tt.input)
			if result != tt.expected {
				t.Errorf("expected %d, got %d", tt.expected, result)
			}
		})
	}
}

// Benchmark SumOfSlice
func BenchmarkSumOfSlice(b *testing.B) {
	slice := make([]int, 1000)
	for i := range slice {
		slice[i] = i
	}
	for i := 0; i < b.N; i++ {
		SumOfSlice(slice)
	}
}