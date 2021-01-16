package timehelper

import (
	"time"
)

// Durations missing in time package.
const (
	Tick time.Duration = 100 * time.Nanosecond
	Day  time.Duration = 24 * time.Hour
	Week time.Duration = 7 * Day

	// https://en.wikipedia.org/wiki/Tropical_year
	TropicalYear time.Duration = time.Duration(365.24219 * float64(Day))
	// TropicalMonth time.Duration = time.Duration(float64(TropicalYear) / 12)
)

// FromDuration returns a time.Duration that represents a specified number of 'd'.
func FromDuration(v float64, d time.Duration) time.Duration {
	t := time.Date(0, 0, 0, 0, 0, 0, int(v*float64(d)), time.UTC)
	t0 := time.Date(0, 0, 0, 0, 0, 0, 0, time.UTC)
	return t.Sub(t0)
}

// FromWeeks returns a time.Duration that represents a specified number of weeks.
func FromWeeks(v float64) time.Duration {
	return FromDuration(v, Week)
}

// FromDays returns a time.Duration that represents a specified number of days.
func FromDays(v float64) time.Duration {
	return FromDuration(v, Day)
}

// FromHours returns a time.Duration that represents a specified number of hours.
func FromHours(v float64) time.Duration {
	return FromDuration(v, time.Hour)
}

// FromMinutes returns a time.Duration that represents a specified number of minutes.
func FromMinutes(v float64) time.Duration {
	return FromDuration(v, time.Minute)
}

// FromSeconds returns a time.Duration that represents a specified number of seconds.
func FromSeconds(v float64) time.Duration {
	return FromDuration(v, time.Second)
}

// FromMilliseconds returns a time.Duration that represents a specified number of milliseconds.
func FromMilliseconds(v float64) time.Duration {
	return FromDuration(v, time.Millisecond)
}

// FromMicroseconds returns a time.Duration that represents a specified number of microseconds.
func FromMicroseconds(v float64) time.Duration {
	return FromDuration(v, time.Microsecond)
}

// FromTicks returns a time.Duration that represents a specified number of ticks.
func FromTicks(v float64) time.Duration {
	return FromDuration(v, Tick)
}
