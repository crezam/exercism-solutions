package grains

import (
	"errors"
	"math"
)

// Square returns the number of grains in the input square position (starting by !)
func Square(input int) (actualVal uint64, actualErr error) {
	if input < 1 || input > 64 {
		return 0, errors.New("Input should be between 1 and 64 inclusive")
	}
	return uint64(math.Pow(float64(2), float64(input-1))), nil
}

// Total returns the total number of grains in a chess board
func Total() (sum uint64) {
	for i := 1; i <= 64; i++ {
		v, _ := Square(i)
		sum += v
	}
	return
}
