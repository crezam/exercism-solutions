/*
Package bob solves the exercism Go: bob problem
*/
package bob

import "strings"

// Hey returns lackadaisical teenager response to an input
func Hey(remark string) string {
	rem := strings.TrimSpace(remark)
	switch {
	case containsAlphabetCharacter(rem) && remark == strings.ToUpper(rem):
		if strings.HasSuffix(rem, "?") {
			return "Calm down, I know what I'm doing!"
		}
		return "Whoa, chill out!"
	case strings.HasSuffix(rem, "?"):
		return "Sure."
	case "" == rem:
		return "Fine. Be that way!"
	}
	return "Whatever."
}

func containsAlphabetCharacter(s string) bool {
	for _, r := range s {
		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') {
			return true
		}
	}
	return false
}
