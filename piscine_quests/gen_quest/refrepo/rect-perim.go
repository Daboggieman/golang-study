package refrepo

import "fmt"

func RectPerimeter(a, b int) int{
	warn := " - Error: can't have a negative/null perimeter value"
	if a <= 0 || b <= 0 {
		return warn
	}
	result := (a + b) * 2
	return result
}