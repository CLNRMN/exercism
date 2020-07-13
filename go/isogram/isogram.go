// Package isogram , contains all functions regarding Isograms
package isogram

import (
	"unicode"
)

// IsIsogram returns true or false
func IsIsogram(iso string) bool {
	// Set default to "false"
	var check bool
	// Initialize Map
	runeChecker := map[rune]bool{}
	// Empty strings are also isograms
	if iso == "" {
		check = true
	}
	for _, c := range iso {
		// Set character c to lowercase
		c = unicode.ToLower(c)
		// Check if the character IsLetter and if the count is bigger than 1
		if unicode.IsLetter(c) && runeChecker[c] {
			check = false
			break
		} else {
			runeChecker[c] = true
			check = true
		}
	}
	return check
}
