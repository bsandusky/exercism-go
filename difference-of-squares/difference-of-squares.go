package diffsquares

import "math"

const testVersion = 1

// SquareOfSums takes an int, calculates the sum of natural numbers between 0
// and the int, then returns an int with that sum to the second power
func SquareOfSums(a int) int {

	var sum int
	for i := a; i > 0; i-- {
		sum += i
	}
	return int(math.Pow(float64(sum), float64(2)))
}

// SumOfSquares takes an int, calculates the sum of all natural numbers between 0
// and that int to the second power, and returns an int of the sum
func SumOfSquares(a int) int {

	var sum int
	for i := a; i > 0; i-- {
		sum += int(math.Pow(float64(i), float64(2)))
	}
	return sum
}

// Difference takes an int, calls SquareOfSums and SumOfSquares, then returns
// an int representing the difference between the two
func Difference(a int) int {
	return SquareOfSums(a) - SumOfSquares(a)
}
