package failed_questions

func Itoa(n int) string {
	if n == 0 {
		return "0"
	}

	sign := ""
	
	// Trick 1: If negative, save the "-", then make the number positive
	if n < 0 {
		sign = "-"
		n = -n 
	}

	result := ""

	// Trick 2: Build the string backwards using math
	for n > 0 {
		// Get last digit, turn to letter, stick it on the FRONT of result
		result = string(rune((n%10)+'0')) + result
		
		// Chop off the last digit
		n /= 10
	}

	return sign + result
}
