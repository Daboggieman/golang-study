package failed_questions

// FishAndChips returns specific strings based on whether the number is divisible by 2, 3, or both.
// 
// THE TRICK: Always check the "divisible by 2 AND 3" case FIRST! 
// If you check "divisible by 2" first, it will return "fish" instead of "fish and chips" for the number 6.
func FishAndChips(n int) string {
	if n < 0 {
		return "error: number is negative"
	}
	
	// Check the hardest condition first!
	if n%2 == 0 && n%3 == 0 {
		return "fish and chips"
	}
	
	if n%2 == 0 {
		return "fish"
	}
	
	if n%3 == 0 {
		return "chips"
	}
	
	return "error: non divisible"
}
