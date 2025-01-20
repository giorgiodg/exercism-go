package isogram

import "strings"

func IsIsogram(word string) bool {
	word = strings.ToLower(word)

	for i, c := range word {
		if c == ' ' || c == '-' {
			continue
		}
		for j, d := range word {
			if c == d && i != j {
				return false
			}
		}
	}
	return true
}
