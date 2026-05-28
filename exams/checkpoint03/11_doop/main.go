// =============================================================
// EXERCISE: doop
//
// QUESTION: Write a program that takes three arguments:
//             1. A value (number)
//             2. An operator (+, -, /, *, %)
//             3. Another value (number)
//           It should print the result.
//
// CONSTRAINTS:
//           - MUST handle division by zero (print nothing or "0",
//             check exam rules, usually "No division by 0\n").
//           - MUST handle modulo by zero (usually "No modulo by 0\n").
//           - MUST handle integer overflows. If a number doesn't
//             fit in a standard 64-bit int, do not panic.
//           - You CANNOT use `strconv` or `fmt` (often restricted).
//             (If `fmt` is allowed, use it. But typically you must
//              write your own `atoi` and `itoa`).
//
// EXAMPLE:
//           go run main.go 1 + 1      -> 2
//           go run main.go 10 / 0     -> No division by 0
//           go run main.go 1 +        -> (nothing, invalid args)
// =============================================================

package main

import (
	"fmt"
	"os"
)

var answer int

func SimpleAtoi(s string) int {
	result := 0
	l := len(s)
	for i := 0; i < l; i++ {
		ch := s[i]
		conv := int(ch - '0')
		result = (result * 10) + conv
	}
	return result
}

var x = os.Args[1]
var op = os.Args[2]
var y = os.Args[3]

func ClCalc() []string {

	if len(os.Args) != 4 {
		fmt.Println("invalid arguments")
	}
	num1 := SimpleAtoi(x)
	num2 := SimpleAtoi(y)

	if num1 == 0 || num2 == 0 {
		if op == "/" {
			fmt.Println("no division by 0")
		}
		if op == "%" {
			fmt.Println("no modulo by 0")
		}

	} else {
		switch op {
		case "+":
			answer = num1 + num2
		case "-":
			answer = num1 - num2
		case "*":
			answer = num1 * num2
		case "/":
			answer = num1 / num2
		case "%":
			answer = num1 % num2
		default:
			fmt.Println("invalid operator")
		}
		fmt.Println("answer is :", answer)
	}
	return nil
}

func main() {
	ClCalc()
}

// ---------------------------------------------------------
// TODO:
// 1. Check `len(os.Args)`. It must be exactly 4
//    (program name + 3 args).
// 2. Parse os.Args[1] and os.Args[3] using your own Atoi().
//    Ensure they are valid numbers.
// 3. Switch on os.Args[2] to determine the operator.
// 4. Perform the calculation. Catch divide/modulo by zero!
// 5. Print the result using your own Itoa() and os.Stdout.
// ---------------------------------------------------------
