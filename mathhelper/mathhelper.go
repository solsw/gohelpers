// Package mathhelper contains various math helpers.
package mathhelper

import (
	"math"
)

// IsEven checks if 'i' is even.
func IsEven(i int64) bool {
	return (i & 1) == 0
}

// AbsInt returns the absolute value of 'i'.
func AbsInt(i int64) int64 {
	if i < 0 {
		return -i
	}
	return i
}

// ApproximatelyEquals reports whether two values are equal within provided 'tolerance'.
func ApproximatelyEquals(value1, value2, tolerance float64) bool {
	return math.Abs(value1-value2) <= math.Abs(tolerance)
}
