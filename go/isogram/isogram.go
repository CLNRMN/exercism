// Package isogram , contains all functions regarding Isograms
package isogram

import (
	"unicode"
)

// IsIsogram returns true or false
func IsIsogram(iso string) bool {
	// Initialize Map
	runeChecker := map[rune]bool{}
	for _, c := range iso {
		// Ignore all special characters
		if !unicode.IsLetter(c) {
			continue
		}
		// Set character c to lowercase
		c = unicode.ToLower(c)
		if runeChecker[c] {
			return false
		}
		runeChecker[c] = true
	}
	return true
}
