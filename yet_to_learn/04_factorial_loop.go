package yettolearn

// EXERCISE 4: Iterative Factorial (Using a Loop)
//
// INSTRUCTIONS:
// Calculate the factorial of `n` using a standard `for` loop.
// Example: if n = 5, result is 5 * 4 * 3 * 2 * 1 = 120
//
// EDGE CASES (Standard Piscine Rules):
// 1. The factorial of 0 is 1.
// 2. If n is negative, return 0.
// 3. You can use a loop that counts down from `n` to 1, multiplying a `result` variable.

func IterativeFactorial(n int) int {
	if n == 0 {
		return 1
	}
	if n < 0 {
		return 0
	}
	var result int

	for i := n; i > 0; i-- {
		result *= i
		return result
	}
	return result
}
