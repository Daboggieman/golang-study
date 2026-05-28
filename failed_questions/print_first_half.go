package failed_questions

import "fmt"

// PrintFirstHalf prints the first half of a string.
// If the string is empty, it prints "Invalid Output".
func PrintFirstHalf(str string) {
	if len(str) == 0 {
		fmt.Print("Invalid Output\n")
		return
	}

	// Calculate the halfway point (rounds up for odd-length strings)
	half := (len(str) + 1) / 2

	// Print characters up to the halfway point
	for i := 0; i < half; i++ {
		fmt.Print(string(str[i]))
	}
	fmt.Print("\n")
}
