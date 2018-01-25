package hamming

import (
	"errors"
)

// Distance calculates Hamming distance of two DNA strands of same length
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return -1, errors.New("Strands have different lengths")
	}
	dist := 0
	bRunes := []rune(b)
	for i, val := range a {
		if val != bRunes[i] {
			dist++
		}
	}
	return dist, nil
}
