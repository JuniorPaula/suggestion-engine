package engine

type TrieNode struct {
	children map[rune]*TrieNode
	isEnd    bool
	freq     int // use frequence (for ranking)
}

type Trie struct {
	root *TrieNode
}

func NewTrie() *Trie {
	return &Trie{
		root: &TrieNode{children: make(map[rune]*TrieNode)},
	}
}

// Insert a word into Trie
func (t *Trie) Insert(word string) {
	word = Normalize(word)

	node := t.root
	for _, ch := range word {
		if node.children[ch] == nil {
			node.children[ch] = &TrieNode{children: make(map[rune]*TrieNode)}
		}
		node = node.children[ch]
	}
	node.isEnd = true
	node.freq++
}

// Find all words with the prefix
func (t *Trie) SearchPrefix(prefix string) []string {
	prefix = Normalize(prefix)

	node := t.root
	for _, ch := range prefix {
		if node.children[ch] == nil {
			return []string{} // no prefix
		}
		node = node.children[ch]
	}

	var results []string
	t.collectWords(node, prefix, &results)
	return results
}

// Collect words recusively
func (t *Trie) collectWords(node *TrieNode, prefix string, results *[]string) {
	if node.isEnd {
		*results = append(*results, prefix)
	}

	for ch, child := range node.children {
		t.collectWords(child, prefix+string(ch), results)
	}
}

// GetFrequency returns a word frequency if exists
func (t *Trie) GetFrequency(word string) int {
	word = Normalize(word)

	node := t.root
	for _, ch := range word {
		if node.children[ch] == nil {
			return 0
		}
		node = node.children[ch]
	}

	if node.isEnd {
		return node.freq
	}

	return 0
}

// TrieNodeFromWord returns the node of the word if it exists
func (t *Trie) TrieNodeFromWord(word string) *TrieNode {
	word = Normalize(word)
	node := t.root

	for _, ch := range word {
		next := node.children[ch]
		if next == nil {
			return nil
		}
		node = next
	}

	if node.isEnd {
		return node
	}

	return nil
}

// TrieWords return all words in the Trie
func (t *Trie) TrieWords() []string {
	var results []string
	t.collectWords(t.root, "", &results)
	return results
}

// TrieFreq return the word frequency
func (t *Trie) TrieFreq(word string) int {
	node := t.TrieNodeFromWord(word)
	if node != nil {
		return node.freq
	}
	return 0
}
