package main

import (
	"fmt"
	"suggestion-engine/engine"
)

func main() {
	e := engine.NewSuggestionEngine()

	// insert words
	words := []string{
		"programação", "programador", "produto", "processo",
		"python", "projeto", "professor", "dinâmica", "código", "golang",
	}

	for _, w := range words {
		e.AddWord(w)
	}

	testInputs := []string{
		"prog", "progrmação", "proje",
		"pytohn", "dinamica",
	}

	for _, input := range testInputs {
		fmt.Printf("\n Entrada: %s\n", input)
		suggestions := e.Suggest(input, 5)

		for _, s := range suggestions {
			fmt.Printf(" → %s (score %.2f)\n", s.Word, s.Score)
		}
	}
}
