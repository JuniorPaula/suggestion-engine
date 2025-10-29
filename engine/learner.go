package engine

import (
	"fmt"
	"os"
	"sync"
)

// Learner makes the engine learn from searches
type Learner struct {
	engine  *SuggestionEngine
	dataset string
	mu      sync.Mutex
}

func NewLearner(e *SuggestionEngine, datasetPath string) *Learner {
	return &Learner{
		engine:  e,
		dataset: datasetPath,
	}
}

// Learn: Increments the frequency of a term and saves it in the dataset
func (l *Learner) Learn(term string) {
	l.mu.Lock()
	defer l.mu.Unlock()

	term = Normalize(term)
	if term == "" {
		return
	}

	// Increments the frequency in the Trie
	node := l.engine.trie.TrieNodeFromWord(term)
	if node != nil {
		node.freq++
	} else {
		// if not exists, add to Trie
		l.engine.AddWord(term)

		// update frequency
		node = l.engine.trie.TrieNodeFromWord(term)
		node.freq = 1
	}

}

// Save: overwrite dataset with the current frequencies
func (l *Learner) Save() error {
	l.mu.Lock()
	defer l.mu.Unlock()

	f, err := os.Create(l.dataset)
	if err != nil {
		return fmt.Errorf("[ERROR] could not save dataset: %v", err)
	}
	defer f.Close()

	all := l.engine.trie.TrieWords()
	for _, w := range all {
		freq := l.engine.trie.TrieFreq(w)
		if freq > 0 {
			fmt.Println("[INFO] updating dataset...")
			fmt.Fprintf(f, "%s %d\n", w, freq)
		}
	}

	return nil
}
