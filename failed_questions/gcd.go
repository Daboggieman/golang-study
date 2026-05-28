package failed_questions

// Gcd returns the greatest common divisor of two unsigned integers.
// Uses the Euclidean algorithm.
func Gcd(a, b uint) uint {
	// Base case: if b is 0, a is the GCD
	if b == 0 {
		return a
	}
	// Recursive case: GCD(b, remainder of a/b)
	return Gcd(b, a%b)
}
