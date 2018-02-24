package scrabble

import (
	"unicode"
	"strings"
)

// Stores in contiguous memory the numeric value equivalent to the character as per its position in the sequence "abcdefghijklmnopqrstuvwxyz"
var charValues = [...]uint8{1, 3, 3, 2, 1, 4, 2, 4, 1, 8, 5, 1, 3, 1, 1, 3, 10, 1, 1, 1, 1, 4, 4, 8, 4, 10}

// Score returns Scrabble score of string
func Score(input string) int {
	input = strings.ToLower(input)
	var score int
	for _, r := range input {
		if unicode.IsLetter(r) {
			score += charValue(r)
		}
	}
	return score
}

func charValue(i rune) int {
	return int(charValues[i-'a'])
}
