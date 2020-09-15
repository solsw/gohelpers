// Package mathhelper contains math helpers.
package mathhelper

import (
	"math"
)

// IsEven reports whether 'i' is even.
func IsEven(i int64) bool {
	return (i & 1) == 0
}

// AbsInt64 returns the absolute value of 'i'.
func AbsInt64(i int64) int64 {
	if i < 0 {
		return -i
	}
	return i
}

// ApproximatelyEquals reports whether two float64 values are equal within the provided 'tolerance'.
func ApproximatelyEquals(v1, v2, tolerance float64) bool {
	return math.Abs(v1-v2) <= math.Abs(tolerance)
}
