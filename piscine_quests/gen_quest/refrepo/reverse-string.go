// package main

// func revStr(s string) string {
// 	a := []rune(s)
// 	b := len(a)
// 	for i, j := 0, b-1; i < j; i, j = i+1, j-1 {
// 		a[i], a[j] = a[j], a[i]
// 	}
// 	return string(a)
// }

//not yet working, still try to fix

package refrepo

func RevStr(s string) string {
	a := []rune(s)
	b := len(a)
	for i, j := 0, b-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
	// Convert the rune slice back to a string
	return string(a)
}
