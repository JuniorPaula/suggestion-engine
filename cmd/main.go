package main

import (
	"fmt"
	"log"
	"suggestion-engine/engine"
)

func main() {
	e := engine.NewSuggestionEngine()

	err := engine.LoadFromFile("data/searches.txt", e)
	if err != nil {
		log.Fatal(err)
	}

	testInputs := []string{
		"programa",
		"progrmação",
		"programacao",
		"programação",
		"dinam",
		"dinamica",
		"dinâmica",
		"pytohn",
	}

	for _, in := range testInputs {
		fmt.Printf("\n Entrada: %s\n", in)
		suggestions := e.Suggest(in, 5)

		for _, s := range suggestions {
			fmt.Printf(" → %s (score %.2f)\n", s.Word, s.Score)
		}
	}
}
