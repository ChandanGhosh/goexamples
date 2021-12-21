package main

import (
	"testing"
)

func Test_factorial(t *testing.T) {
	tests := []struct {
		name string
		give int64
		want int64
	}{
		// TODO: Add test cases.
		{name: "0!", give: 0, want: 1},
		{name: "1!", give: 1, want: 1},
		{name: "3!", give: 3, want: 6},
		{name: "6!", give: 6, want: 720},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := factorial(int64(tt.give)); got != tt.want {
				t.Errorf("factorial(%v) = %v, want %v", tt.give, got, tt.want)
			}
		})
	}
}

// Benchmark_factorial tests recursive factorial function
func Benchmark_factorial(b *testing.B) {
	for n := 0; n < b.N; n++ {
		factorial(int64(n))
	}
}

func Test_factorial_improved(t *testing.T) {
	tests := []struct {
		name string
		give int
		want int64
	}{
		// TODO: Add test cases.
		{name: "0!", give: 0, want: 1},
		{name: "1!", give: 1, want: 1},
		{name: "3!", give: 3, want: 6},
		{name: "6!", give: 6, want: 720},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := factorial_improved(tt.give); got != tt.want {
				t.Errorf("factorial_improved(%v) = %v, want %v", tt.give, got, tt.want)
			}
		})
	}
}

// Benchmark_factorial_improved tests non-recursive factorial function
func Benchmark_factorial_improved(b *testing.B) {
	for n := 0; n < b.N; n++ {
		factorial_improved(n)
	}
}
