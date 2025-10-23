package engine

import (
	"bufio"
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
