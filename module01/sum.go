package module01

// Sum will sum up all of the numbers passed
// in and return the result
func Sum(numbers []int) int {
	if len(numbers) > 0 {
		var result int
		for _, number := range numbers {
			result += number
		}
		return result
	}
	return 0
}

// with a very long list this may result in stack overflow!
func SumRecursive(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}
	return numbers[0] + SumRecursive(numbers[1:])
}
