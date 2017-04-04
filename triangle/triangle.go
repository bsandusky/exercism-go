package triangle

import (
	"math"
	"time"

	"../trackTime"
)

const testVersion = 3

// Kind asserts triangle type
type Kind int

// const block identifies possible triangle types
const (
	NaT = iota // not a triangle
	Equ = iota // equilateral
	Iso = iota // isosceles
	Sca = iota // scalene
)

// KindFromSides takes three float64 values and returns triangle type or
// NaT if not mathematically a triangle
func KindFromSides(a, b, c float64) Kind {
	defer trackTime.TrackTime(time.Now())

	switch {

	// Not a triangle
	case a+b < c || a+c < b || b+c < a,
		a == 0 || b == 0 || c == 0,
		math.IsNaN(a) || math.IsNaN(b) || math.IsNaN(c),
		math.IsInf(a, 0) || math.IsInf(b, 0) || math.IsInf(c, 0):
		return NaT

	// Equilateral
	case a == b && a == c:
		return Equ

	// Isosceles
	case a == b || b == c || a == c:
		return Iso

	// Scalene
	default:
		return Sca
	}
}
