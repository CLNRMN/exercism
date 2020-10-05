package pangram

import (
	"strings"
)

// IsPangram checks if a given string is a Pangram
func IsPangram(s string) bool {
	stringLowercase := strings.ToLower(s)
	for char := 'a'; char <= 'z'; char++ {
		if strings.IndexRune(stringLowercase, char) == -1 {
			return false
		}
	}
	return true
}
