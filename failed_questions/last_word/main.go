package main

import (
	"fmt"
	"os"
)

// Last Word is a program that takes a string and prints its last word.
// A word is a sequence of characters separated by spaces or tabs.
// 
// THE TRICK: Don't read from left-to-right! Start at the END of the string 
// and read backwards right-to-left!
func main() {
	if len(os.Args) != 2 {
		fmt.Print("\n")
		return
	}

	str := os.Args[1]
	
	// Step 1: Start at the very end of the string
	end := len(str) - 1
	
	// Step 2: Skip over any trailing spaces or tabs at the very end
	for end >= 0 && (str[end] == ' ' || str[end] == '\t') {
		end--
	}
	
	// Step 3: If the string was entirely empty or just spaces, end here.
	if end < 0 {
		fmt.Print("\n")
		return
	}
	
	// Step 4: Keep moving backwards until we hit a space (which means we found the start of the word)
	start := end
	for start >= 0 && str[start] != ' ' && str[start] != '\t' {
		start--
	}
	
	// Step 5: Slice the string from the start of the word to the end of the word
	// (We do start+1 because `start` is currently sitting on the space character)
	word := str[start+1 : end+1]
	
	fmt.Print(word + "\n")
}
