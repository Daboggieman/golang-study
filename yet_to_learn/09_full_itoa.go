package yet_to_learn

// EXERCISE 9: Full Itoa (No strconv)
//
// INSTRUCTIONS:
// Write a function that simulates the real Itoa.
// It must handle:
// 1. Negative numbers (e.g., -123 -> "-123").
// 2. The number 0.
//
// HINTS ON HOW TO GO ABOUT IT:
// 1. Check if `n == 0`. If so, return "0".
// 2. Keep track if the number is negative using a boolean `isNegative = false`.
// 3. If `n < 0`, set `isNegative = true`, and make `n` positive (`n = -n`).
// 4. Run the exact same modulo 10 loop you built in `simple_itoa` to build the string.
// 5. At the very end, if `isNegative` is true, add a "-" to the front of your result string!

func FullItoa(n int) string {
	// TODO: implement
	return ""
}
