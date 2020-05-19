package timehelper

import (
	"errors"
	"time"
)

// IntAsMonth converts int into time.Month with range checking.
func IntAsMonth(month int) (time.Month, error) {
	if !(1 <= month && month <= 12) {
		return time.Month(0), errors.New("invalid month")
	}
	return time.Month(month), nil
}

func ymd(y int, m time.Month, d int) time.Time {
	return time.Date(y, m, d, 0, 0, 0, 0, time.UTC)
}

// IsLeapYear reports whether the 'year' is leap.
func IsLeapYear(year int) bool {
	return ymd(year, time.December, 31).YearDay() == 366
}

// FirstDayOfMonth returns the first day of the 'month' of the 'year'.
func FirstDayOfMonth(year int, month time.Month) (time.Time, error) {
	_, err := IntAsMonth(int(month))
	if err != nil {
		return time.Time{}, err
	}
	return ymd(year, month, 1), nil
}

// LastDayOfMonth returns the last day of the 'month' of the 'year'.
func LastDayOfMonth(year int, month time.Month) (time.Time, error) {
	_, err := IntAsMonth(int(month))
	if err != nil {
		return time.Time{}, err
	}
	return ymd(year, month+1, 1).AddDate(0, 0, -1), nil
}

// DaysInMonth returns number of days in the 'month' of the 'year'.
func DaysInMonth(year int, month time.Month) (int, error) {
	t, err := LastDayOfMonth(year, month)
	if err != nil {
		return 0, err
	}
	return t.Day(), nil
}

// DateYMD returns time.Time corresponding to 'year', 'month' and 'day'.
// If 'year', 'month' and 'day' represent invalid date, error is not nil.
func DateYMD(year, month, day int) (time.Time, error) {
	m, err := IntAsMonth(month)
	if err != nil {
		return time.Time{}, err
	}
	d, _ := DaysInMonth(year, m)
	if day < 1 || day > d {
		return time.Time{}, errors.New("invalid day")
	}
	return ymd(year, m, day), nil
}

// FirstWeekdayInYear returns first occurrence of the 'weekday' in the 'year'.
func FirstWeekdayInYear(year int, weekday time.Weekday) (time.Time, error) {
	if weekday < time.Sunday || time.Saturday < weekday {
		return time.Time{}, errors.New("invalid weekday")
	}
	t := ymd(year, time.January, 1)
	for t.Weekday() != weekday {
		t = t.AddDate(0, 0, 1)
	}
	return t, nil
}

// FirstWeekdayInYear2 returns first occurrence of the 'weekday' in the 'year'.
func FirstWeekdayInYear2(year int, weekday time.Weekday) (time.Time, error) {
	if weekday < time.Sunday || time.Saturday < weekday {
		return time.Time{}, errors.New("invalid weekday")
	}
	t1 := ymd(year, time.January, 1)
	d := int(weekday - t1.Weekday())
	if d == 0 {
		return t1, nil
	}
	if d < 0 {
		d += 7
	}
	return t1.AddDate(0, 0, d), nil
}

// WeekdaysInYear returns number of weekdays in the 'year'.
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
	res, _ := WeekdaysInYear(year, time.Friday)
	return res
}

// FridaysInYear2 returns number of Fridays in the 'year'.
func FridaysInYear2(year int) int {
	jan1Weekday := time.Date(year, time.January, 1, 0, 0, 0, 0, time.UTC).Weekday()
	if (jan1Weekday == time.Friday) || (jan1Weekday == time.Thursday && IsLeapYear(year)) {
		return 53
	}
	return 52
}

// PrevClosestWeekday returns date of 'weekday' closest to 't' and not after 't'.
func PrevClosestWeekday(t time.Time, weekday time.Weekday) (time.Time, error) {
	if weekday < time.Sunday || weekday > time.Saturday {
		return time.Time{}, errors.New("invalid weekday")
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
		return time.Time{}, errors.New("invalid weekday")
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
