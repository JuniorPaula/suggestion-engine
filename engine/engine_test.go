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
