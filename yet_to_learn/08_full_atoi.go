package yet_to_learn

// EXERCISE 8: Full Atoi (No strconv)
//
// INSTRUCTIONS:
// Write a function that simulates the real Atoi. You already did the simple version!
// Now, you must handle:
// 1. A leading '+' or '-' sign.
// 2. Invalid strings (if it contains letters or spaces, you must return 0).
//
// HINTS ON HOW TO GO ABOUT IT:
// 1. Keep a `sign` variable (start it at 1).
// 2. Check the very first character `s[0]`. If it is '-', set `sign = -1` and skip to the next character. 
//    If it's '+', just skip to the next character.
// 3. Loop through the rest of the string.
// 4. If any character is NOT between '0' and '9', return 0 immediately.
// 5. Build the number just like you did in the simple version.
// 6. At the end, return `result * sign`.

func FullAtoi(s string) int {
	// TODO: implement
	return 0
}
