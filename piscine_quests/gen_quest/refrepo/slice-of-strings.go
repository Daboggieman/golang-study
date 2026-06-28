// Write a Go program to create a slice of strings and append a new string to it.
// Required Input:

// A slice of strings (e.g., ["Apple", "Banana"]) and a string to append (e.g., "Cherry").

// Expected Output:

// [Apple Banana Cherry]

package refrepo

import "fmt"

func SlcStr() {
	//set the name of the variable, then ':=', then append [] the variable, then call the type of the variable (string, int, float), then in a '{}' input the value of the array
	fruits := []string{"apple", "Banana", "peach"}

	fruits = append(fruits, ("mango", "cherry"))

	fmt.Println(fruits)
	fmt.Println(fruits[0], fruits[1], fruits[2], fruits[3], )
}
