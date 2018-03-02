package luhn

import (
	"strings"
	"unicode"
)
// Valid determines if a string is valid as per the Luhn algorithm
func Valid(s string) bool {
	// strip spaces
	s = strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, s)
	if len(s) < 2 {
		return false
	}
	runes := []rune(s)
	shouldDouble := false
	var sum, current int
	for i := len(runes) - 1; i >= 0; i-- {
		if !unicode.IsDigit(runes[i]) {
			return false
		}
		if shouldDouble {
			current = 2 * int(runes[i]-'0')
			if current > 9 {
				current -= 9
			}
		} else {
			current = int(runes[i] - '0')
		}
		sum += current
		shouldDouble = !shouldDouble
	}
	return sum%10 == 0
}
