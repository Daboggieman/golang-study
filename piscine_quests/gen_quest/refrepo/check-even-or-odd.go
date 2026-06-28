package refrepo

import "fmt"

func CheckOddEven(a int) {
	// Check if a number is even or odd
	if a <= 0 {
		return "invalid value to check, please insert a valid number to check odd or even"
	}
	if a%2 == 0 {
		fmt.Println("The Number is Even")
	} else {
		fmt.Println("The Number is Odd")
	}
}
