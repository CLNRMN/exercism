// Package isogram , contains all functions regarding Isograms
package isogram

import (
	"strings"
	"unicode"
)

// IsIsogram returns true or false
func IsIsogram(iso string) bool {
	// Set variable iso to lowercase
	iso = strings.ToLower(iso)
	// Set default to "false"
	var check bool
	// Empty strings are also isograms
	if iso == "" {
		check = true
	}
	for _, c := range iso {
		// Check if the character IsLetter and if the count is bigger than 1
		if strings.Count(iso, string(c)) > 1 && unicode.IsLetter(c) {
			check = false
			break
		} else {
			check = true
		}
	}
	return check
}
