package main

import (
	"fmt"
	"strconv"
)

// EXERCISE 2: Simple Atoi
//
// INSTRUCTIONS:
// In the real Atoi quest, you must handle negative signs, plus signs, and
// invalid letters. Here, we assume the string ONLY contains positive digits
// (e.g., "456"). Let's understand the core logic first.
//
// HINTS ON HOW TO GO ABOUT IT:
// 1. Create a `result` variable starting at 0.
// 2. Loop through the string `s`.
// 3. To turn a rune character like '4' into the number 4, you do: int(ch - '0')
// 4. To build the final number, multiply your current result by 10 and add the new digit.
//    (e.g., result *= 10 + digit)

func SimpleAtoi(s string) int {
	result := 0
	l := len(s)
	for i := 0; i < l; i++ {
		ch := s[i]
		conv := int(ch - '0')
		result = (result * 10) + conv
	}
	fmt.Println(result)
	return result
}

// using the strconv package to convert string to int, but it is not allowed in piscine exercises
func SimpleAtoi2(s string) int {
	conv, _ := strconv.Atoi(s)
	fmt.Println(conv)
	return conv
}

func main() {

	SimpleAtoi("123") // Output: 123
	SimpleAtoi2("432")
}
