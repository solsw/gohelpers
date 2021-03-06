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

// RoundInt64 returns the nearest to 'x' integer, rounding half away from zero.
func RoundInt64(x float64) int64 {
	return int64(math.Round(x))
}

// RoundToEvenInt64 returns the nearest to 'x' integer, rounding half to even.
func RoundToEvenInt64(x float64) int64 {
	return int64(math.RoundToEven(x))
}

// TruncInt64 returns the integer value of 'x'.
func TruncInt64(x float64) int64 {
	return int64(math.Trunc(x))
}

// Frac returns the fractional part of 'x'.
func Frac(x float64) float64 {
	_, f := math.Modf(x)
	return f
}

// NearlyEqual reports whether two float64 values are equal within 'tolerance'.
// 'tolerance' must be positive.
func NearlyEqual(v1, v2, tolerance float64) bool {
	return math.Abs(v1-v2) < tolerance
}

// Split4 splits sequence of ints (array indexes starting from 0 with length - 'len') into four parts.
// The result array contains start indexes of the second, third and fourth parts (start index of the first part is 0).
// Function is intended for splitting array for (parallel) processing of the separate parts.
func Split4(len int) [3]int {
	if len <= 0 {
		return [3]int{0, 0, 0}
	}
	switch len {
	case 1:
		return [3]int{1, 1, 1}
	case 2:
		return [3]int{1, 2, 2}
	case 3:
		return [3]int{1, 2, 3}
	default:
		i2 := len / 2
		i1 := i2 / 2
		i3 := i2 + (len-i2)/2
		return [3]int{i1, i2, i3}
	}
}
