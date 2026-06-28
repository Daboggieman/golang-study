// Write a Go program to declare an array of integers and print the second element.
// Required Input:

// An array of integers (e.g., [1, 2, 3, 4, 5]).

// Expected Output:

// 20

package refrepo

import "fmt"

func DecArray() {
	var array = [6]int{10, 20, 30, 40, 50, 60}
	var slice = []int{10, 20, 30, 40, 50, 60}

	fmt.Println(array[1], slice[1])
}
