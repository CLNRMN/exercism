// Package raindrops contains functions to play with raindrops
package raindrops

import "strconv"

// Convert returns a String for each Raindrop
func Convert(number int) string {
	var output string
	if number%3 == 0 {
		output += "Pling"
	}
	if number%5 == 0 {
		output += "Plang"
	}
	if number%7 == 0 {
		output += "Plong"
	}
	// if no factor check matched (output is empty), use the inital number as string
	if output == "" {
		output = strconv.Itoa(number)
	}
	return output
}
