package etl

import (
	"strings"
)

//Transform moves the old scoring system in the new one.
func Transform(m map[int][]string) map[string]int {
	output := make(map[string]int)
	for points, letters := range m {
		for _, letter := range letters {
			output[strings.ToLower(letter)] = points
		}
	}
	return output
}
