package timehelper

import (
	"time"
)

// DateYMD returns time.Time corresponding to 'year', 'month' and 'day' (see time.Date).
// The returned value has zero clock and time.UTC Location.
func DateYMD(year int, month time.Month, day int) time.Time {
	return time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
}

// IsLeapYear reports whether the 'year' is leap.
func IsLeapYear(year int) bool {
	// return DateYMD(year, time.December, 31).YearDay() == 366 :-)
	return year%4 == 0 && year%100 != 0 || year%400 == 0
}

// DaysInMonth returns number of days in the 'month' of the 'year'.
// If 'month' is invalid -1 is returned. Use ValidMonth if need month validation.
func DaysInMonth(year int, month time.Month) int {
	switch month {
	case time.February:
		if IsLeapYear(year) {
			return 29
		}
		return 28
	case time.April, time.June, time.September, time.November:
		return 30
	case time.January, time.March, time.May, time.July, time.August, time.October, time.December:
		return 31
	default:
		return -1
	}
}
