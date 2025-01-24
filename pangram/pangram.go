package pangram

import "strings"

func IsPangram(input string) bool {
	input = strings.ToUpper(input)
	for chr := 'A'; chr <= 'Z'; chr++ {
		if !strings.ContainsRune(input, chr) {
			return false
		}
	}

	return true
}
