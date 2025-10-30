package engine

import (
	"strings"
)

type SuggestionEngine struct {
	trie *Trie
}

func NewSuggestionEngine() *SuggestionEngine {
	return &SuggestionEngine{trie: NewTrie()}
}

func (e *SuggestionEngine) AddWord(word string) {
	e.trie.Insert(word)
}

// Exist wrapper method to call trie.Exists
func (e *SuggestionEngine) Exists(word string) bool {
	return e.trie.Exists(word)
}

// Suggest return the most relevants suggestion for the typing term
func (e *SuggestionEngine) Suggest(input string, limit int) []RankedWord {
	input = Normalize(input)

	// Find the prefix
	prefixMatches := e.trie.SearchPrefix(input)

	// if didn't find anything with the prefix, try typing correction
	if len(prefixMatches) == 0 {

		// the bigger the word, the greater the tolerance
		maxDist := 2
		if len(input) > 7 {
			maxDist = 3
		}

		prefixMatches = e.findByEditDistance(input, maxDist) // tolerates up 3 errors
	}

	ranked := rankWords(input, prefixMatches, e.trie)
	if len(ranked) > limit {
		return ranked[:limit]
	}
	return ranked
}

// findByEditDistance search by approximate edit distance
func (e *SuggestionEngine) findByEditDistance(input string, baseMax int) []string {
	var results []string
	all := e.trie.SearchPrefix("") // find all

	// maximum addaptive distance (in % of size)
	maxDist := baseMax
	if len(input) > 8 {
		maxDist = 4
	}

	for _, w := range all {
		mainWord := strings.Fields(w)[0]
		dist := EditDistance(input, mainWord)

		// tolerates up to 40% difference in word size
		if dist <= maxDist || dist <= len(mainWord)/3 {
			results = append(results, w)
		}
	}

	return results
}
