// =============================================================
// EXERCISE: split
//
// QUESTION: Write a function that receives a string `s` and a
//           separator `sep` and returns a slice of strings that
//           results from splitting `s` by the separator `sep`.
//
// CONSTRAINTS:
//           - You are NOT allowed to use the `strings` package.
//           - If the separator is empty, behavior can vary, but
//             typically you split by character.
//
// EXAMPLE:
//           s := "HelloHAhowHAareHAyou?"
//           sep := "HA"
//           Result: ["Hello", "how", "are", "you?"]
// =============================================================

// package main

// import "fmt"

// func Split(s, sep string) []string {
// 	var result []string
// 	start := 0
// 	sepLen := len(sep)

// 	if sepLen == 0 {
// 		return []string{s}
// 	}

// 	for i := 0; i <= len(s)-sepLen; i++ {
// 		if s[i:i+sepLen] == sep {
// 			result = append(result, s[start:i])
// 			start = i + sepLen
// 			i += sepLen - 1 // Skip past the separator
// 		}
// 	}
// 	result = append(result, s[start:])
// 	return result
// }

//	func main() {
//		fmt.Print(Split("HelloHAhowHAareHAyou?", "HA"))
package main

import (
	"fmt"
	"strings"
)

func Split(s, sep string) []string {
	var result []string
	for {
		i := strings.Index(s, sep)
		if i == -1 {
			return append(result, s)
		}
		result = append(result, s[:i])
		s = s[i+len(sep):] // Move past the separator
	}
}

func main() {
	fmt.Println(Split("HelloHAhowHAareHAyou?", "H"))
}

// ---------------------------------------------------------
// TODO:
// 1. Find occurrences of `sep` inside `s`.
// 2. Extract substrings between separators.
// 3. Append them to a []string slice.
// ---------------------------------------------------------
