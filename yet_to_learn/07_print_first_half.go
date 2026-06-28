package yet_to_learn

// EXERCISE 7: Print First Half
//
// INSTRUCTIONS:
// Write a function that takes a string and returns the first half of it.
// If the string is empty, you usually return a specific error string like "Invalid Output\n" 
// (or whatever your specific quest specifies).
//
// HINTS ON HOW TO GO ABOUT IT:
// 1. Check if the length of the string `len(s)` is 0. If so, handle the error.
// 2. Calculate the halfway point: `half := len(s) / 2`. 
//    (Note: If the piscine requires you to round up for odd lengths, you can use: 
//    `half := (len(s) + 1) / 2`).
// 3. Use Go's slicing feature to return the first half: `return s[:half]`.
import "os"

func PrintFirstHalf(s string) string {
	// TODO: implement
	//this nigga caused me problems during my piscine checkpoint, very messed up something
	if len(s) == 0 {
		fmt.Println("the string is empty, please insert valid sstring")
		os.Exit
	}
	halfstr := len(s) / 2
	if len(s)%2 != 0 {
		halfstr = (len(s) + 1) / 2
		return halfstr
	}
	return s[:halfstr]
}
