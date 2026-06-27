package solve

import "fmt"

func Loops() {
	for i := 1; i <= 100; i++ {
		fmt.Print(i, ",", " ")
	}
	fmt.Println()
	for i := 1; i <= 100; {
		if i%2 == 0 {
			fmt.Print(i, ",", " ")
		}
		i++ 
	}
	fmt.Println()
	for i := 1; i <= 100; {
		if i%2 != 0 {
			fmt.Print(i, ",", " ")
		}
		i++
	}
	fmt.Println('\n')
}

func FactLoop(n int) int {
	/*finding the factorial of a number 'n' is n! where n is multiplied by every number ranging from the value of n down till 1
	that is if n = 5, then the factoriial of 5 (5!) is 5x4x3x2x1 = 120 */
		if n < 0 {
			return -1
		}
		if n == 0 {
			return 0
		}
		if n == 1 {
			return 1
		}
		if n > 20 {
			return n
		}
		result := 1
		for i := 1; i <= n; i++ {
			result *= i
		}
		return result
}

func Fibonacci(a int) int {
	if a < 0 {
		return a
	}
	Fibonacci := (Fibonacci(a-2)) + (Fibonacci(a-1))
	return Fibonacci
}