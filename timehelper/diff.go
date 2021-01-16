package timehelper

import (
	"time"
)

// diff returns difference (t2 - t1) truncated to whole number of 'dur's.
func diff(t1, t2 time.Time, dur time.Duration) int {
	// inspired by https://github.com/ribice/dt/blob/c4e011852071395c6d984ef7b3b579f5951c88de/date.go#L81
	deltaUnix := t2.In(time.UTC).Unix() - t1.In(time.UTC).Unix()
	return int(deltaUnix / int64(dur/time.Second))
}

// DaysDiff returns number of whole days containing in difference (t2 - t1).
// If 't2' is before 't1', result is negative.
func DaysDiff(t1, t2 time.Time) int {
	return diff(t1, t2, Day)
}

// WeeksDiff returns number of whole weeks containing in difference (t2 - t1).
// If 't2' is before 't1', result is negative.
func WeeksDiff(t1, t2 time.Time) int {
	return diff(t1, t2, Week)
}

// MonthsDiff returns number of whole months containing in difference (t2 - t1).
// If 't2' is before 't1', result is negative.
func MonthsDiff(t1, t2 time.Time) int {
	var (
		after, before time.Time
		negative      bool
	)
	if t1.After(t2) {
		after = t1
		before = t2
		negative = true
	} else {
		after = t2
		before = t1
		negative = false
	}
	r := 12*(after.Year()-before.Year()) + int(after.Month()) - int(before.Month())
	if r == 0 {
		return 0
	}
	if after.Day() < before.Day() {
		r--
	}
	if negative {
		return -r
	}
	return r
}

// YearsDiff returns number of whole years containing in difference (t2 - t1).
// If 't2' is before 't1', result is negative.
func YearsDiff(t1, t2 time.Time) int {
	var (
		after, before time.Time
		negative      bool
	)
	if t1.After(t2) {
		after = t1
		before = t2
		negative = true
	} else {
		after = t2
		before = t1
		negative = false
	}
	r := after.Year() - before.Year()
	if r == 0 {
		return 0
	}
	if after.YearDay() < before.YearDay() {
		r--
	}
	if negative {
		return -r
	}
	return r
}
