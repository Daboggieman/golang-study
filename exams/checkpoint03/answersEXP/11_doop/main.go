package main

import (
	"os"

	"github.com/01-edu/z01"
)

func PrintStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

func PrintLn(s string) {
	PrintStr(s)
	z01.PrintRune('\n')
}

func Atoi(s string) (int64, bool) {
	if s == "" {
		return 0, false
	}

	sign := int64(1)
	start := 0
	if s[0] == '-' {
		sign = -1
		start = 1
	} else if s[0] == '+' {
		start = 1
	}

	if start == len(s) {
		return 0, false
	}

	var value uint64
	const maxPos uint64 = 1<<63 - 1
	const maxNeg uint64 = 1 << 63
	limit := maxPos
	if sign < 0 {
		limit = maxNeg
	}

	for i := start; i < len(s); i++ {
		ch := s[i]
		if ch < '0' || ch > '9' {
			return 0, false
		}
		digit := uint64(ch - '0')
		if value > limit/10 || (value == limit/10 && digit > limit%10) {
			return 0, false
		}
		value = value*10 + digit
	}

	if sign < 0 {
		return -int64(value), true
	}

	return int64(value), true
}

func Itoa(n int64) string {
	if n == 0 {
		return "0"
	}

	negative := n < 0
	if negative {
		n = -n
	}

	buf := make([]byte, 0, 20)
	for n > 0 {
		digit := byte(n%10) + '0'
		buf = append(buf, digit)
		n /= 10
	}

	if negative {
		buf = append(buf, '-')
	}

	for i, j := 0, len(buf)-1; i < j; i, j = i+1, j-1 {
		buf[i], buf[j] = buf[j], buf[i]
	}

	return string(buf)
}

func main() {
	if len(os.Args) != 4 {
		return
	}

	num1, ok1 := Atoi(os.Args[1])
	num2, ok2 := Atoi(os.Args[3])
	if !ok1 || !ok2 {
		return
	}

	op := os.Args[2]
	switch op {
	case "+":
		PrintLn(Itoa(num1 + num2))
	case "-":
		PrintLn(Itoa(num1 - num2))
	case "*":
		PrintLn(Itoa(num1 * num2))
	case "/":
		if num2 == 0 {
			PrintLn("No division by 0")
			return
		}
		PrintLn(Itoa(num1 / num2))
	case "%":
		if num2 == 0 {
			PrintLn("No modulo by 0")
			return
		}
		PrintLn(Itoa(num1 % num2))
	default:
		return
	}
}
