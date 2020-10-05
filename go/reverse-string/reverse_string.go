package reverse

import "strings"

// Reverse returns the string in reverse
func Reverse(s string) string {
	reversedString := strings.Builder{}
	runes := []rune(s)
	for i := len(runes) - 1; i >= 0; i-- {
		reversedString.WriteRune(runes[i])
	}
	return reversedString.String()
}
