package main

import (
	"fmt"
	"suggestion-engine/engine"
)

func main() {
	trie := engine.NewTrie()
	words := []string{
		"programação", "programador", "produto", "processo",
		"dinâmica", "python",
	}

	for _, w := range words {
		trie.Insert(w)
	}

	fmt.Println("Prefixo: pro")
	fmt.Println(trie.SearchPrefix("pro"))

	fmt.Println("Distância entre 'progrmação' e programação:",
		engine.EditDistance("progrmação", "programação"))
}
