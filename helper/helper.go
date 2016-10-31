package helper

import "fmt"

// TruncateString returns a truncated string with a given length
func TruncateString(content string, length int) string {
	var numRunes = 0
	for index, _ := range content {
		numRunes++
		if numRunes > length {
			return fmt.Sprintf("%s...", content[:index])
		}
	}
	return content
}
