package engine_test

import (
	"suggestion-engine/engine"
	"testing"
)

// Test the edit distance calculation
func TestEditDistance(t *testing.T) {
	tests := []struct {
		a, b string
		want int
	}{
		{"kitten", "sitting", 3},
		{"programacao", "progrmacao", 1},
		{"grafos", "graf", 2},
		{"python", "pytohn", 2},
		{"", "", 0},
		{"a", "", 1},
	}

	for _, tt := range tests {
		got := engine.EditDistance(tt.a, tt.b)
		if got != tt.want {
			t.Errorf("EditDistance(%q, %q) = %d; want %d", tt.a, tt.b, got, tt.want)
		}
	}
}

// Tests nomalization (accents, capitalization, spaces)
func TestNormalize(t *testing.T) {
	tests := []struct {
		in, want string
	}{
		{"Programação", "programacao"},
		{"  PYTHON", "python"},
		{"Ciência da Computação", "ciencia da computacao"},
	}

	for _, tt := range tests {
		got := engine.Normalize(tt.in)
		if got != tt.want {
			t.Errorf("Normalize(%q) = %q; want %q", tt.in, got, tt.want)
		}
	}
}
