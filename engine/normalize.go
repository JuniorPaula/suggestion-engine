package engine

import (
	"strings"
	"unicode"

	"golang.org/x/text/unicode/norm"
)

// Normalize converts text to lowercase, remove accents and extra spaces
func Normalize(s string) string {
	s = strings.ToLower(strings.TrimSpace(s))
	t := norm.NFD.String(s)

	// combining marks
	result := make([]rune, 0, len(t))
	for _, r := range t {
		if unicode.Is(unicode.Mn, r) { // Mn = Mark, Nonspacing
			continue
		}
		result = append(result, r)
	}
	return string(result)
}
