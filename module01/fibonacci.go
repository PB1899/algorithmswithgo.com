package module01

// Fibonacci returns the nth fibonacci number.
//
// A Fibonacci number N is defined as:
//
//	Fibonacci(N) = Fibonacci(N-1) + Fibonacci(N-2)
//
// Where the following base cases are used:
//
//	Fibonacci(0) => 0
//	Fibonacci(1) => 1
//
// Examples:
//
//	Fibonacci(0) => 0
//	Fibonacci(1) => 1
//	Fibonacci(2) => 1
//	Fibonacci(3) => 2
//	Fibonacci(4) => 3
//	Fibonacci(5) => 5
//	Fibonacci(6) => 8
//	Fibonacci(7) => 13
//	Fibonacci(14) => 377
func Fibonacci(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

// this solution is faster because the base values aren't recalculated like with recursion
func FibonacciLoop(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}
	nMinus1 := 1
	nMinus2 := 0
	var result int
	for i := 2; i <= n; i++ {
		result = nMinus1 + nMinus2
		nMinus2 = nMinus1
		nMinus1 = result
	}
	return result
}
