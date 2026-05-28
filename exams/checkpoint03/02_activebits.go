// =============================================================
// EXERCISE: activebits
//
// QUESTION: Write a function that returns the number of active bits
//           (1s) in the binary representation of an integer number.
//
// CONSTRAINTS:
//           - You cannot use the `math/bits` package.
//           - You must use bitwise operators (e.g., >>, &, |, ^)
//             or division/modulo by 2.
//
// EXAMPLE:
//           activebits(7)  -> Returns 3 (because 7 in binary is 111)
//           activebits(15) -> Returns 4 (because 15 in binary is 1111)
// =============================================================

package piscine

// ActiveBits returns the number of active bits in n.
func ActiveBits(n int) int {
	// ---------------------------------------------------------
	// TODO:
	// 1. Loop while n > 0.
	// 2. Use modulo 2 or bitwise AND (& 1) to check the last bit.
	// 3. Shift the number right (n >>= 1) or divide by 2.
	// 4. Count the active bits.
	// ---------------------------------------------------------
	return 0
}
