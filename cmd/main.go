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

	testInputs := []string{"programa", "progrmação", "dinam", "pytohn"}

	for _, input := range testInputs {
		fmt.Printf("\n Entrada: %s\n", input)
		suggestions := e.Suggest(input, 5)

		for _, s := range suggestions {
			fmt.Printf(" → %s (score %.2f)\n", s.Word, s.Score)
		}
	}
}
