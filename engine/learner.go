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
	}

	// save file to the dataset
	l.appendToDataset(term)
}

// appendToDataset: Adds the term to the end of the dataset with incremental frequency
func (l *Learner) appendToDataset(term string) {
	f, err := os.OpenFile(l.dataset, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("[ERROR] could not save dataset:", err)
		return
	}
	defer f.Close()
	line := fmt.Sprintf("%s %d\n", term, 1)
	f.WriteString(line)
}
