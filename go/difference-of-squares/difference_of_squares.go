package diffsquares

// SquareOfSum square of the sum of the first x natural numbers
func SquareOfSum(num int) int {
	var squareofsum int
	for i := 0; i <= num; i++ {
		squareofsum = squareofsum + i
	}
	return squareofsum * squareofsum
}

// SumOfSquares sum of the squares of the first x natural numbers is
func SumOfSquares(num int) int {
	var sumofsquare int
	for i := 0; i <= num; i++ {
		sumofsquare = sumofsquare + (i * i)
	}
	return sumofsquare
}

// Difference returns the difference between SquareOfSum() and SumOfSquares()
func Difference(num int) int {
	return SquareOfSum(num) - SumOfSquares(num)
}
