package main

// EXERCISE 3: Simple Itoa
//
// INSTRUCTIONS:
// Assume `n` is a positive number. Convert it to a string.
// This is the exact opposite of Atoi!
//
// HINTS ON HOW TO GO ABOUT IT:
// 1. If n is 0, just return "0".
// 2. Create an empty string `result = ""`.
// 3. Use a loop: while n > 0
// 4. Get the last digit using modulo: digit := n % 10
// 5. Convert that digit back to a character: char := string(rune(digit + '0'))
// 6. Prepend it to your string: result = char + result
// 7. Divide n by 10 to chop off the last digit: n = n / 10

func SimpleItoa(n int) string {
	if n == 0 {
		return "0"
	}
	result := ""
	for n > 0 {
		digit := n % 10
		char := string(rune(digit + '0'))
		result = char + result
		n = n / 10
	}
	return result
}
