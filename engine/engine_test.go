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

// Test the Trie and the dataset loading
func TestLoadFromFile(t *testing.T) {
	e := engine.NewSuggestionEngine()
	err := engine.LoadTestDataset(e)
	if err != nil {
		t.Fatal(err)
	}

	checks := []string{"python", "grafos", "programacao"}
	for _, word := range checks {
		if !e.Exists(word) {
			t.Errorf("expected Trie to contain %q, but it doesn't", word)
		}
	}
}

// Test the complete suggestion system
func TestSuggest(t *testing.T) {
	e := engine.NewSuggestionEngine()
	err := engine.LoadTestDataset(e)
	if err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		query string
		want  string
	}{
		{"pytohn", "python"},
		{"progra", "programacao"},
		{"dinam", "dinamica"},
	}

	for _, tt := range tests {
		results := e.Suggest(tt.query, 3)
		if len(results) == 0 {
			t.Errorf("Suggest(%q): expected suggestions, got 0", tt.query)
			continue
		}

		found := false
		for _, r := range results {
			if r.Word == tt.want {
				found = true
				break
			}
		}

		if !found {
			t.Errorf("Suggest(%q): expected contains %q, but dit not appear in the results", tt.query, tt.want)
		}
	}
}

// Benchmarck simple for suggest performace
func BenchmarkSuggest(b *testing.B) {
	e := engine.NewSuggestionEngine()
	err := engine.LoadTestDataset(e)
	if err != nil {
		b.Fatal(err)
	}

	for i := 0; i < b.N; i++ {
		e.Suggest("programacao", 5)
	}
}
