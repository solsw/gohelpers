package timehelper

import (
	"fmt"
	"time"
)

// DateYMD returns time.Time corresponding to the provided 'year', 'month' and 'day'.
// The returned value has zero time and time.UTC Location.
func DateYMD(year int, month time.Month, day int) time.Time {
	return time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
}

// IsLeapYear reports whether the 'year' is leap.
func IsLeapYear(year int) bool {
	return DateYMD(year, time.December, 31).YearDay() == 366
}

// IntAsMonth converts int into time.Month with range checking.
func IntAsMonth(month int) (time.Month, error) {
	if !(1 <= month && month <= 12) {
		return time.Month(0), fmt.Errorf("invalid month '%d'", month)
	}
	return time.Month(month), nil
}

// DaysInMonth returns number of days in the 'month' of the 'year'.
func DaysInMonth(year int, month time.Month) (int, error) {
	t, err := LastDayOfMonth(year, month)
	if err != nil {
		return 0, err
	}
	return t.Day(), nil
}

// YMD returns time.Time corresponding to 'year', 'month' and 'day'.
// The returned value has zero time and time.UTC Location.
// If 'year', 'month' and 'day' represent invalid date, error is not nil.
func YMD(year, month, day int) (time.Time, error) {
	m, err := IntAsMonth(month)
	if err != nil {
		return time.Time{}, err
	}
	d, _ := DaysInMonth(year, m)
	if !(1 <= day && day <= d) {
		return time.Time{}, fmt.Errorf("invalid day '%d'", day)
	}
	return DateYMD(year, m, day), nil
}

// YMDMust is like YMD but panics in case of error.
func YMDMust(year, month, day int) time.Time {
	t, err := YMD(year, month, day)
	if err != nil {
		panic(err)
	}
	return t
}

// FirstDayOfMonth returns the first day of 'month' in the 'year'.
// The returned value has zero time and time.UTC Location.
func FirstDayOfMonth(year int, month time.Month) (time.Time, error) {
	_, err := IntAsMonth(int(month))
	if err != nil {
		return time.Time{}, err
	}
	return DateYMD(year, month, 1), nil
}

// LastDayOfMonth returns the last day of 'month' in the 'year'.
// The returned value has zero time and time.UTC Location.
func LastDayOfMonth(year int, month time.Month) (time.Time, error) {
	_, err := IntAsMonth(int(month))
	if err != nil {
		return time.Time{}, err
	}
	return DateYMD(year, month+1, 1).AddDate(0, 0, -1), nil
}

// FirstWeekdayInYear returns the first occurrence of 'weekday' in the 'year'.
// The returned value has zero time and time.UTC Location.
func FirstWeekdayInYear(year int, weekday time.Weekday) (time.Time, error) {
	if weekday < time.Sunday || time.Saturday < weekday {
		return time.Time{}, fmt.Errorf("invalid weekday '%d'", weekday)
	}
	t := DateYMD(year, time.January, 1)
	for t.Weekday() != weekday {
		t = t.AddDate(0, 0, 1)
	}
	return t, nil
}

// FirstWeekdayInYear2 returns the first occurrence of 'weekday' in the 'year'.
// The returned value has zero time and time.UTC Location.
func FirstWeekdayInYear2(year int, weekday time.Weekday) (time.Time, error) {
	if weekday < time.Sunday || time.Saturday < weekday {
		return time.Time{}, fmt.Errorf("invalid weekday '%d'", weekday)
	}
	t1 := DateYMD(year, time.January, 1)
	d := int(weekday - t1.Weekday())
	if d == 0 {
		return t1, nil
	}
	if d < 0 {
		d += 7
	}
	return t1.AddDate(0, 0, d), nil
}

// WeekdaysInYear returns number of 'weekday's in the 'year'.
func WeekdaysInYear(year int, weekday time.Weekday) (int, error) {
	wd, err := FirstWeekdayInYear(year, weekday)
	if err != nil {
		return 0, err
	}
	n := 0
	for wd.Year() == year {
		n++
		wd = wd.AddDate(0, 0, 7)
	}
	return n, nil
}

// FridaysInYear returns number of Fridays in the 'year'.
func FridaysInYear(year int) int {
	jan1Weekday := time.Date(year, time.January, 1, 0, 0, 0, 0, time.UTC).Weekday()
	if (jan1Weekday == time.Friday) || (jan1Weekday == time.Thursday && IsLeapYear(year)) {
		return 53
	}
	return 52
}

// PrevClosestWeekday returns date of 'weekday' closest to 't' and not after 't'.
func PrevClosestWeekday(t time.Time, weekday time.Weekday) (time.Time, error) {
	if weekday < time.Sunday || weekday > time.Saturday {
		return time.Time{}, fmt.Errorf("invalid weekday '%d'", weekday)
	}
	d := int(weekday - t.Weekday())
	if d == 0 {
		return t, nil
	}
	if d > 0 {
		d -= 7
	}
	return t.AddDate(0, 0, d), nil
}

// NextClosestWeekday returns date of 'weekday' closest to 't' and not before 't'.
func NextClosestWeekday(t time.Time, weekday time.Weekday) (time.Time, error) {
	if weekday < time.Sunday || weekday > time.Saturday {
		return time.Time{}, fmt.Errorf("invalid weekday '%d'", weekday)
	}
	d := int(weekday - t.Weekday())
	if d == 0 {
		return t, nil
	}
	if d < 0 {
		d += 7
	}
	return t.AddDate(0, 0, d), nil
}
