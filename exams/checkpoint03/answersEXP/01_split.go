package piscine

// Split returns a slice of substrings from s using sep as the separator.
// This implementation does not rely on the strings package.
func Split(s, sep string) []string {
	if len(sep) == 0 {
		return []string{s}
	}

	var result []string
	start := 0
	sepLen := len(sep)

	for i := 0; i+sepLen <= len(s); i++ {
		if s[i:i+sepLen] == sep {
			result = append(result, s[start:i])
			start = i + sepLen
			i += sepLen - 1
		}
	}

	result = append(result, s[start:])
	return result
}
