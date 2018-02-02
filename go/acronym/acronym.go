package acronym

import (
	"strings"
	"unicode"
)

// Abbreviate should have a comment documenting it.
func Abbreviate(s string) string {
	acronym := ""
	fields := strings.Fields(s)
	for _, f := range fields {
		for _, t := range strings.Split(f, "-") {
			appendFirstCharacter(&acronym, t)
		}
	}
	return acronym
}

func appendFirstCharacter(s *string, toAppend string) {
	var firstChar = []rune(toAppend)[0]
	if unicode.IsLetter(firstChar) {
		*s += string(unicode.ToUpper(firstChar))
	}
}
