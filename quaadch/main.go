package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

func QuadA(x int, y int) string {
	if x < 1 || y < 1 {
		return ""
	}
	mw := x - 2
	if mw < 0 {
		mw = 0
	}
	top := "*" + strings.Repeat("*", mw)
	if x > 1 {
		top += "*"
	}
	mid := "*" + strings.Repeat(" ", mw)
	if x > 1 {
		mid += "*"
	}
	bot := "*" + strings.Repeat("*", mw)
	if x > 1 {
		bot += "*"
	}
	result := top + "\n"
	for i := 0; i < y-2; i++ {
		result += mid + "\n"
	}
	if y > 1 {
		result += bot + "\n"
	}
	return result
}

func QuadB(x int, y int) string {
	if x < 1 || y < 1 {
		return ""
	}
	mw := x - 2
	if mw < 0 {
		mw = 0
	}
	top := "/" + strings.Repeat("*", mw)
	if x > 1 {
		top += "\\"
	}
	mid := "*" + strings.Repeat(" ", mw)
	if x > 1 {
		mid += "*"
	}
	bot := "\\" + strings.Repeat("*", mw)
	if x >= 1 {
		bot += "/"
	}
	result := top + "\n"
	for i := 0; i < y-2; i++ {
		result += mid + "\n"
	}
	if y > 1 {
		result += bot + "\n"
	}
	return result
}

func QuadC(x int, y int) string {
	if x < 1 || y < 1 {
		return ""
	}
	mw := x - 2
	if mw < 0 {
		mw = 0
	}
	top := "A" + strings.Repeat("B", mw)
	if x > 1 {
		top += "A"
	}
	mid := "B" + strings.Repeat(" ", mw)
	if x > 1 {
		mid += "B"
	}
	bot := "C" + strings.Repeat("B", mw)
	if x >= 1 {
		bot += "C"
	}
	result := top + "\n"
	for i := 0; i < y-2; i++ {
		result += mid + "\n"
	}
	if y > 1 {
		result += bot + "\n"
	}
	return result
}

// func QuadD(x int, y int) string {
// 	implement me
// }

// func QuadE(x int, y int) string {
// 	implement me
// }

func runQuad(name string) {
	if len(os.Args) < 3 {
		fmt.Fprintf(os.Stderr, "Usage: %s <x> <y>\n", name)
		os.Exit(1)
	}
	x, _ := strconv.Atoi(os.Args[1])
	y, _ := strconv.Atoi(os.Args[2])

	var result string
	switch name {
	case "quadA":
		result = QuadA(x, y)
	case "quadB":
		result = QuadB(x, y)
	case "quadC":
		result = QuadC(x, y)
	// case "quadD":
	// 	result = QuadD(x, y)
	// case "quadE":
	// 	result = QuadE(x, y)
	default:
		fmt.Fprintf(os.Stderr, "Unknown quad: %s\n", name)
		os.Exit(1)
	}
	fmt.Print(result)
}

func runChecker() {
	input, err := io.ReadAll(os.Stdin)
	if err != nil || len(input) == 0 {
		fmt.Println("Not a quad function")
		return
	}
	inputStr := string(input)
	compare := []string{}

	for y := 1; y <= 100; y++ {
		for x := 1; x <= 100; x++ {
			if inputStr == QuadA(x, y) {
				compare = append(compare, fmt.Sprintf("[QuadA] [%d] [%d]", x, y))
			}
			if inputStr == QuadB(x, y) {
				compare = append(compare, fmt.Sprintf("[QuadB] [%d] [%d]", x, y))
			}
			if inputStr == QuadC(x, y) {
				compare = append(compare, fmt.Sprintf("[QuadC] [%d] [%d]", x, y))
			}
			// if inputStr == QuadD(x, y) {
			// 	compare = append(compare, fmt.Sprintf("[QuadD] [%d] [%d]", x, y))
			// }
			// if inputStr == QuadE(x, y) {
			// 	compare = append(compare, fmt.Sprintf("[QuadE] [%d] [%d]", x, y))
			// }
		}
	}
	if len(compare) == 0 {
		fmt.Println("Not a quad function")
		return
	}
	sort.Strings(compare)
	for i, m := range compare {
		if i > 0 {
			fmt.Print(" || ")
		}
		fmt.Print(m)
	}
	fmt.Println()
}

func main() {
	name := filepath.Base(os.Args[0])
	switch name {
	case "quadA", "quadB", "quadC", "quadD", "quadE":
		runQuad(name)
	default:
		runChecker()
	}
}