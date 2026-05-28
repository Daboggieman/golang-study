package failed_questions

import "strings"

// --- HOW TO CONVERT A STRING TO A SLICE ---
//
// There are two tools from the "strings" package:
//
// TOOL 1: strings.Fields(s)
//   - Splits by ANY whitespace (spaces, tabs, newlines)
//   - Automatically ignores extra/leading/trailing spaces
//   - Best choice for splitting sentences into words
//
// TOOL 2: strings.Split(s, sep)
//   - Splits by a SPECIFIC separator you choose (e.g. ",", "/", " ")
//   - Does NOT ignore extra spaces — "a  b" gives ["a", "", "b"]
//   - Best for structured data like "a,b,c"
//
// QUICK RULE:
//   Splitting a sentence into words?  → Use strings.Fields
//   Splitting by a known character?   → Use strings.Split

// Example 1: Sentence into words (ignores extra spaces)
func SplitIntoWords(s string) []string {
	// "  hello   world  " → ["hello", "world"]
	return strings.Fields(s)
}

// Example 2: Split by a specific separator
func SplitBySeparator(s string, sep string) []string {
	// "a,b,c" split by "," → ["a", "b", "c"]
	return strings.Split(s, sep)
}
