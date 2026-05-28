package failed_questions

func Atoi(s string) int {
	if len(s) == 0 {
		return 0
	}

	sign := 1

	// Trick 1: If there's a sign, record it, then CHOP IT OFF the string!
	if s[0] == '-' || s[0] == '+' {
		if s[0] == '-' {
			sign = -1
		}
		s = s[1:] // This slices the string, removing the first character
	}

	result := 0

	// Trick 2: Loop over the remaining clean string
	for _, char := range s {
		if char < '0' || char > '9' {
			return 0
		}
		result = result*10 + int(char-'0')
	}

	return result * sign
}
