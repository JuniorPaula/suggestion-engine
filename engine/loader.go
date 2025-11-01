package engine

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func LoadFromFile(path string, e *SuggestionEngine) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		if len(parts) < 2 {
			continue
		}

		// Last part is the frequency
		freqStr := parts[len(parts)-1]
		freq, _ := strconv.Atoi(freqStr)
		word := strings.Join(parts[:len(parts)-1], " ")

		for i := 0; i < freq; i++ {
			e.AddWord(word)
		}
	}

	return scanner.Err()
}

func loadFromBytes(data []byte, e *SuggestionEngine) error {
	scanner := bufio.NewScanner(bytes.NewReader(data))
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		parts := strings.Split(line, " ")
		if len(parts) < 2 {
			continue
		}

		word := strings.Join(parts[:len(parts)-1], " ")
		freqStr := parts[len(parts)-1]

		freq, err := strconv.Atoi(freqStr)
		if err != nil {
			continue
		}

		e.AddWord(word)
		node := e.trie.TrieNodeFromWord(word)
		if node != nil {
			node.freq = freq
			node.isEnd = true
		}
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("[ERROR] on load dataset: %v", err)
	}

	return nil
}
