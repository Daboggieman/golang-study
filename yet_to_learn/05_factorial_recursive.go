package yettolearn

// EXERCISE 5: Recursive Factorial
//
// INSTRUCTIONS:
// Calculate the factorial of `n` WITHOUT using a loop!
// You must use recursion (a function calling itself).
//
// HINTS ON HOW TO GO ABOUT IT:
// 1. Base case: If n == 0 or n == 1, return 1. (This stops the function from running forever).
// 2. Base case: If n < 0, return 0.
// 3. Recursive step: Return n * RecursiveFactorial(n - 1)
// 
// WHY THIS WORKS:
// Factorial(5) returns 5 * Factorial(4)
// Factorial(4) returns 4 * Factorial(3)... until it hits 1.

func RecursiveFactorial(n int) int {
	// TODO: implement
	if n == 0 || n == 1 {
		return 1
	}
	if n < 0 {
		return 0
	}
	return n * RecursiveFactorial(n - 1)
}
