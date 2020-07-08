// Package triangle determine the type of an trangle
package triangle

import "math"

// Kind Notice KindFromSides() returns this type. Pick a suitable data type.
type Kind int

const (
	// NaT not a triangle
	NaT Kind = iota
	// Equ equilateral
	Equ
	// Iso isosceles
	Iso
	// Sca scalene
	Sca
)

// KindFromSides should have a comment documenting it.
func KindFromSides(a, b, c float64) Kind {
	var k Kind
	// all sides have to be of length > 0
	if a > 0 && b > 0 && c > 0 && !(math.IsInf(a, 0)) && !(math.IsInf(b, 0)) && !(math.IsInf(c, 0)) {
		if a+b >= c && b+c >= a && a+c >= b {
			if a == b && a == c {
				k = Equ
			}
			if a == b && a != c || b == c && b != a || a == c && a != b {
				k = Iso
			}
			if a != b && a != c && b != c {
				k = Sca
			}
		}
	}
	return k
}
