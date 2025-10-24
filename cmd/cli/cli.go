package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"suggestion-engine/engine"
)

func main() {
	e := engine.NewSuggestionEngine()
	err := engine.LoadFromFile("data/searches.txt", e)
	if err != nil {
		fmt.Println("[ERROR] could not load dataset:", err)
		return
	}

	reader := bufio.NewReader(os.Stdin)

	fmt.Println(" Modo interativo do motor de sugestão")
	fmt.Println(" Digite algo e veja as sugestões em rempo real (CRTL+C pra sair)")
	fmt.Println("-----------------------------------------------------------")

	var in string

	for {
		fmt.Print("\nBuscar: ")
		text, _ := reader.ReadString('\n')
		in = strings.TrimSpace(text)

		if in == "" {
			continue
		}

		start := time.Now()
		suggestions := e.Suggest(in, 5)
		elapsed := time.Since(start)

		fmt.Printf("→ %d sugestão em %v\n", len(suggestions), elapsed)
		for _, s := range suggestions {
			fmt.Printf("  🞄 %-40s (score %.2f)\n", s.Word, s.Score)
		}
	}
}
