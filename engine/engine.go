package engine

type SuggestionEngine struct {
	trie *Trie
}

func NewSuggestionEngine() *SuggestionEngine {
	return &SuggestionEngine{trie: NewTrie()}
}

func (e *SuggestionEngine) AddWord(word string) {
	e.trie.Insert(word)
}

// Suggest return the most relevants suggestion for the typing term
func (e *SuggestionEngine) Suggest(input string, limit int) []RankedWord {
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
func (e *SuggestionEngine) findByEditDistance(input string, maxDist int) []string {
	var results []string
	all := e.trie.SearchPrefix("") // find all

	for _, w := range all {
		if EditDistance(input, w) <= maxDist {
			results = append(results, w)
		}
	}

	return results
}
