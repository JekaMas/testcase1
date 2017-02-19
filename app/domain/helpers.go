package domain

func IsLeadingZero(s string, startPosition int) bool {
	if len(s) > startPosition+1 && s[startPosition] == '0' && s[startPosition+1] == '0' {
		return true
	}

	return false
}

func IsLatinUpper(s string, pos int) bool {
	if pos > len(s)-1 {
		return false
	}

	if s[pos] < 'A' || s[pos] > 'Z' {
		return false
	}

	return true
}

func IsLatinNumber(s string, startPosition int) bool {
	if startPosition > len(s)-1 {
		return false
	}

	for i := startPosition; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return false
		}
	}

	return true
}
