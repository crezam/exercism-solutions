package isogram

import (
	"strings"
	"unicode"
)

// IsIsogram returns true if s is an isogram: no repeating letters
func IsIsogram(s string) bool {
	s = strings.ToLower(s)
	contents := map[rune]bool{}
	for _, r := range s {
		if unicode.IsLetter(r) {
			if contents[r] {
				return false
			}
			contents[r] = true
		}
	}
	return true
}
