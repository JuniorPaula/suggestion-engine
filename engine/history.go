package engine

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

type History struct {
	Entries []string
	Path    string
}

func NewHistory(path string) *History {
	return &History{
		Path:    path,
		Entries: []string{},
	}
}

// Add: Record the search in the memory history and in the file
func (h *History) Add(term string) {
	term = Normalize(term)

	// Add on history in memory
	h.Entries = append([]string{term}, h.Entries...)
	if len(h.Entries) > 10 { // keep the last then
		h.Entries = h.Entries[:10]
	}

	// Write log file
	file, err := os.OpenFile(h.Path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("[ERROR] could not save history:", err)
		return
	}
	defer file.Close()

	timestamp := time.Now().Format("2006-01-02 15:04:05")
	line := fmt.Sprintf("%s | %s\n", timestamp, term)
	file.WriteString(line)
}

// Load: load the previus history
func (h *History) Load() {
	file, err := os.Open(h.Path)
	if err != nil {
		fmt.Println("[INFO] file history is empty:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, " | ", 2)
		if len(parts) == 2 {
			h.Entries = append([]string{parts[1]}, h.Entries...)
		}
	}

	if len(h.Entries) > 10 {
		h.Entries = h.Entries[:10]
	}
}
