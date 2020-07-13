// Package accumulate contains all functions regarding accumulating strings
package accumulate

type converter func(string) string

// Accumulate converts slices with the given converter
func Accumulate(s []string, conv converter) []string {
	var result []string
	for _, e := range s {
		e = conv(e)
		result = append(result, e)
	}
	return result
}
