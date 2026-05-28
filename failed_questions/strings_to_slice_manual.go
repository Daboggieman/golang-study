package failed_questions

// --- HOW TO CONVERT A STRING TO A SLICE (NO IMPORTS NEEDED) ---
//
// THE TRICK: Loop through the string one character at a time.
// You track two states:
//   - "inside a word"  → keep building the current word
//   - "inside a space" → if you just finished a word, SAVE it and reset
//
// GOTCHA: After the loop ends, there is no final space to trigger the save!
// You MUST manually append the last word AFTER the loop.
//
// STEP BY STEP:
//   1. Create an empty result slice and an empty word string
//   2. Loop over each character
//   3. If it's a space/tab/newline → save word (if not empty), reset word to ""
//   4. Otherwise → add the character to word
//   5. After the loop → save the last word (if not empty)

func StringsToSliceManual(s string) []string {
	var result []string
	word := ""

	for _, ch := range s {
		if ch == ' ' || ch == '\t' || ch == '\n' {
			// Hit whitespace — if we were building a word, save it
			if word != "" {
				result = append(result, word)
				word = "" // reset for the next word
			}
		} else {
			// Still inside a word — keep building
			word += string(ch)
		}
	}

	// IMPORTANT: Save the very last word.
	// There's no trailing space to trigger the save above!
	if word != "" {
		result = append(result, word)
	}

	return result
}

// Example:
//   StringsToSliceManual("  hello   world  foo  ")
//   → ["hello", "world", "foo"]
