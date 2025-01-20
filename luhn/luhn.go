package luhn

import (
	"strings"
	"unicode"
)

func Valid(id string) bool {

	id = strings.ReplaceAll(id, " ", "")
	if len(id) <= 1 {
		return false
	}

	sum := 0
	doubleDigit := false

	for i := len(id) - 1; i >= 0; i-- {
		if !unicode.IsDigit(rune(id[i])) {
			return false
		}
		digit := int(id[i] - '0')
		if doubleDigit {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}
		sum += digit
		doubleDigit = !doubleDigit
	}
	return sum%10 == 0
}
