package piscine

// ActiveBits returns the number of 1 bits in the binary representation of n.
func ActiveBits(n int) int {
	count := 0
	un := uint(n)
	for un > 0 {
		if un&1 == 1 {
			count++
		}
		un >>= 1
	}
	return count
}
