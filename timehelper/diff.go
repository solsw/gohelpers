package timehelper

import (
	"time"
)

// Missing in time package Durations.
const (
	Day  time.Duration = 24 * time.Hour
	Week time.Duration = 7 * Day
)

// diff returns difference (t2 - t1) truncated to whole number of 'dur's.
func diff(t1, t2 time.Time, dur time.Duration) int {
	// inspired by https://github.com/ribice/dt/blob/c4e011852071395c6d984ef7b3b579f5951c88de/date.go#L81
	deltaUnix := t2.In(time.UTC).Unix() - t1.In(time.UTC).Unix()
	return int(deltaUnix / int64(dur/time.Second))
}

// DaysDiff returns difference (t2 - t1) as a number of whole days.
func DaysDiff(t1, t2 time.Time) int {
	return diff(t1, t2, Day)
}

// WeeksDiff returns difference (t2 - t1) as a number of whole weeks.
func WeeksDiff(t1, t2 time.Time) int {
	return diff(t1, t2, Week)
}

// MonthsDiff returns difference (t2 - t1) as a number of whole months.
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

// YearsDiff returns difference (t2 - t1) as a number of whole years.
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
