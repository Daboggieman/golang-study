package failed_questions

// Slice acts like a Python slice, taking a slice of strings and variadic integers.
// 
// THE RULES:
// 1. If no ints are passed, return nil.
// 2. If 1 int is passed, slice from that start index to the end.
// 3. If 2 ints are passed, slice from start to end index.
// 4. Handle negative indexes (they count from the back).
func Slice(a []string, nbrs ...int) []string {
	if len(nbrs) == 0 {
		return nil
	}
	
	start := nbrs[0]
	end := len(a) // default end is the length of the slice
	
	if len(nbrs) >= 2 {
		end = nbrs[1]
	}
	
	// Trick: Handle negative indexes by adding the length to them
	if start < 0 {
		start = len(a) + start
	}
	if end < 0 {
		end = len(a) + end
	}
	
	// Clamp variables to prevent "index out of range" panics
	if start < 0 { start = 0 }
	if start > len(a) { start = len(a) }
	
	if end < 0 { end = 0 }
	if end > len(a) { end = len(a) }
	
	// If start is somehow past the end, return nil
	if start > end {
		return nil
	}
	
	return a[start:end]
}
