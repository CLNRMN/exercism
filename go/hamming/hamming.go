package hamming

// This package calculates the the Hamming Distance

import "errors"

// Distance returns the number of diffrenreces between two DNA
func Distance(a, b string) (int, error) {

	if len(a) != len(b) {
		return 0, errors.New("DNA are not equal")
	}
	// Counter for the Distance
	var distance int
	for n := 0; n < len(a); n++ {
		if a[n] != b[n] {
			distance++
		}
	}
	return distance, nil
}
