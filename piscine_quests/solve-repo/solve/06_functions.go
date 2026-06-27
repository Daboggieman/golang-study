package solve

// import "fmt"

// func CheckPrime(n int) {
// 	var isPrime string

// 	if n <= 0 {
// 		fmt.Println(n, " is invalid, use a valid integer")
// 	} else if n == 1 {
// 		isPrime := "is a Prime Number"
// 		fmt.Println(isPrime)
// 	} else if n < 6 {
// 		if n != 4 {
// 			isPrime := "is a Prime Number"
// 			fmt.Println(isPrime)
// 		}
// 	} else if n%2 == 0 || n%3 == 0 || n%5 == 0 {
// 		isPrime := "is not a Prime Number"
// 		fmt.Println(isPrime)
// 	} else {
// 		isPrime := "is a Prime Number"
// 		fmt.Println(isPrime)
// 	}
// 	fmt.Println(isPrime)
// }

func RecurFactorial(n int) int {
	if n <= 1 {
		return n
	}
	//what we are supposed to design is taht if the if n is greater than 1 then n is multiplied by the value of results which 
	// result *= (n-1)
	// return results
	for i := n; i > 1 && i <= n; i-- {
	n *= (i - 1)
	}
	return n
}

/* what is the difference between an itterative factorial and a recursive factorial, well the difference is the method used, answers are the same but the method used are different, in the iterative factorial the inital int/number to factor is multiplied by the base factorial -1 over and over again until the new base factorial is now 1, then it returns the result, otherwise it continues using the loop again and again*/