package main

import (
	"fmt"
	"os"
)

// Union is a PROGRAM (not just a function) that takes two strings as arguments.
// It prints characters that appear in EITHER string, but strictly NO DOUBLES.
// 
// THE TRICK: Use a map (or array) to keep track of characters you have already printed.
func main() {
	// If it doesn't have exactly 2 arguments (plus the program name itself = 3), just print newline
	if len(os.Args) != 3 {
		fmt.Print("\n")
		return
	}

	str1 := os.Args[1]
	str2 := os.Args[2]
	
	// This map will act as our memory to remember if we've seen a character
	seen := make(map[rune]bool)
	
	// Loop over first string
	for _, char := range str1 {
		if !seen[char] {
			fmt.Print(string(char))
			seen[char] = true // Mark as seen!
		}
	}
	
	// Loop over second string
	for _, char := range str2 {
		if !seen[char] {
			fmt.Print(string(char))
			seen[char] = true // Mark as seen!
		}
	}
	
	fmt.Print("\n")
}
