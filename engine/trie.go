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
