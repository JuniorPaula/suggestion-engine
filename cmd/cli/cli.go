package main

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"suggestion-engine/engine"
)

// clearScreen clear the terminal
func clearScreen() {
	fmt.Print("\003[H\033[2J")
}

// colorText mapped the color with the text
func colorText(text, color string) string {
	colors := map[string]string{
		"green":  "\033[1;32m",
		"cyan":   "\033[1;36m",
		"yellow": "\033[1;33m",
		"gray":   "\033[0;37m",
		"reset":  "\033[0m",
	}
	return colors[color] + text + colors["reset"]
}

func main() {
	e := engine.NewSuggestionEngine()
	err := engine.LoadFromFile("data/searches.txt", e)
	if err != nil {
		fmt.Println("[ERROR] could not load dataset:", err)
		return
	}

	history := engine.NewHistory("data/search_log.txt")
	history.Load()

	learner := engine.NewLearner(e, "data/searches.txt")

	reader := bufio.NewReader(os.Stdin)

	clearScreen()
	fmt.Println(colorText("🧠 Modo interativo do motor de sugestão", "cyan"))
	fmt.Println(colorText("Digite algo e veja as sugestões em rempo real (CRTL+C pra sair)", "gray"))
	fmt.Println(strings.Repeat("—", 55))

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-sig
		fmt.Println("\n[INFO] closing interactive mode...")
		fmt.Println("[INFO] saving dataset and history...")

		if err := learner.Save(); err != nil {
			fmt.Println(colorText("[ERROR] could not save dataset:", "red"))
		} else {
			fmt.Println(colorText("Dataset save on success!", "green"))
		}

		fmt.Println("See you soon! Bye...")
		os.Exit(0)
	}()

	in := ""

	for {
		if len(history.Entries) > 0 {
			fmt.Println(colorText("Últimas buscas:", "gray"))
			for i, h := range history.Entries {
				fmt.Printf("  %d. %s\n", i+1, colorText(h, "yellow"))
			}
			fmt.Println(strings.Repeat("—", 55))
		}

		fmt.Print(colorText("\nBuscar: ", "yellow"))
		text, _ := reader.ReadString('\n')
		in = strings.TrimSpace(text)

		if in == "" {
			continue
		}

		start := time.Now()
		suggestions := e.Suggest(in, 5)
		elapsed := time.Since(start)

		// real-time learning
		history.Add(in)
		learner.Learn(in)

		clearScreen()

		fmt.Printf("%s %s\n", colorText("🔍 Entrada:", "cyan"), colorText(in, "yellow"))
		fmt.Printf("%s %d sugestão em %s\n", colorText("🕑 Tempo:", "cyan"), len(suggestions), colorText(fmt.Sprintf("(%v)", elapsed), "gray"))
		fmt.Println(strings.Repeat("—", 55))

		if len(suggestions) == 0 {
			fmt.Println(colorText("Nenhuma sugestão encontrada.", "gray"))
			continue
		}

		for i, s := range suggestions {
			fmt.Printf("%s %-40s %s %.2f\n",
				colorText(fmt.Sprintf("%2d.", i+1), "green"),
				colorText(s.Word, "yellow"),
				colorText("score", "gray"),
				s.Score,
			)
		}

		fmt.Println(strings.Repeat("—", 55))
	}
}
