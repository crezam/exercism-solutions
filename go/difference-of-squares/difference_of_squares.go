package diffsquares

// SumOfSquares returns sum of squares
func SumOfSquares(n int) int {
	sum := 1
	for i := 2; i <= n; i++ {
		sum += i * i
	}
	return sum
}

// SquareOfSums returns square of sums
func SquareOfSums(n int) int {
	sum := n * (n + 1) / 2
	return sum * sum
}

// Difference returns difference of square of sums and sum of squares
func Difference(n int) int {
	return SquareOfSums(n) - SumOfSquares(n)
}
