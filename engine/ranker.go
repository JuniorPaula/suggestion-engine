package engine

type RankedWord struct {
	Word  string
	Score float64
}

// rankwords calculates the relevance of words on distance and frequency
func rankWords(input string, words []string, trie *Trie) []RankedWord {
	results := make([]RankedWord, 0, len(words))

	for _, w := range words {
		dist := EditDistance(input, w)
		freq := trie.GetFrequency(w)

		// Simple Score: freq / (1 + distance)
		score := float64(freq) / (1.0 + float64(dist))
		results = append(results, RankedWord{Word: w, Score: score})
	}

	// Sort by largest to smallest
	for i := 0; i < len(results); i++ {
		for j := i + 1; j < len(results); j++ {
			if results[j].Score > results[i].Score {
				results[i], results[j] = results[j], results[i]
			}
		}
	}

	return results
}
