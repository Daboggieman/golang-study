package main

import (
	"fmt"
	"os"
)

// Print Last String is a program that prints the very last argument passed to it.
//
// THE TRICK: The arguments are stored in the slice `os.Args`.
// The length is `len(os.Args)`. To get the last item in ANY slice, 
// the index is always `length - 1`.
func main() {
	// If no arguments are passed (length is 1 because the program name is os.Args[0]), print newline
	if len(os.Args) < 2 {
		fmt.Print("\n")
		return
	}
	
	// Get the very last string in the array!
	lastIndex := len(os.Args) - 1
	fmt.Print(os.Args[lastIndex] + "\n")
}
