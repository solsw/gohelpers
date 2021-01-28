package timehelper

import (
	"time"
)

// AddMonth adds 'mons' months to 't'.
// If the resulting day is beyond the number of days in the month, the last day of the month is returned.
func AddMonth(t time.Time, mons int) time.Time {
	if mons == 0 {
		return t
	}
	year, month, day := t.Date()
	// zero-based offset (relative to initial year) of resulting month
	off2 := int(month) - 1 + mons
	year += off2 / 12
	if off2 < 0 {
		year--
	}
	// zero-based number (in resulting year) of resulting month
	mon2 := off2 % 12
	if mon2 < 0 {
		mon2 += 12
	}
	month = time.Month(mon2 + 1)
	dim := DaysInMonth(year, month)
	if day > dim {
		day = dim
	}
	hour, min, sec := t.Clock()
	return time.Date(year, month, day, hour, min, sec, t.Nanosecond(), t.Location())
}
