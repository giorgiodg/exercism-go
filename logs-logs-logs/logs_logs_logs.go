package logs

import "unicode/utf8"

// Application identifies the application emitting the given log.
func Application(log string) string {

	for _, char := range log {
		if char == '❗' { // U+2757
			return "recommendation"
		} else if char == '🔍' { // U+1F50D
			return "search"
		} else if char == '☀' { // U+2600
			return "weather"
		}
	}
	return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	newLog := ""
	for _, char := range log {
		if char == oldRune {
			newLog = newLog + string(newRune)
		} else {
			newLog = newLog + string(char)
		}
	}
	return newLog
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	return (utf8.RuneCountInString(log) <= limit)
}
