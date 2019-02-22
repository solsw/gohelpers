package timehelper

import (
	"errors"
	"time"
)

var time0 time.Time

// Time0 returns zero Time.
func Time0() time.Time {
	return time0
}

// IntAsMonth converts int into time.Month with range checking.
func IntAsMonth(month int) (time.Month, error) {
	if month < 1 || month > 12 {
		return time.Month(0), errors.New("invalid month")
	}
	return time.Month(month), nil
}

// IsLeapYear reports whether the 'year' is leap.
func IsLeapYear(year int) bool {
	return time.Date(year, time.December, 31, 0, 0, 0, 0, time.UTC).YearDay() == 366
}

// FirstDayOfMonth returns the first day of the 'month' of the 'year'.
func FirstDayOfMonth(year int, month time.Month) (time.Time, error) {
	_, err := IntAsMonth(int(month))
	if err != nil {
		return time0, err
	}
	return dateYMDPrim(year, month, 1), nil
}

// LastDayOfMonth returns the last day of the 'month' of the 'year'.
func LastDayOfMonth(year int, month time.Month) (time.Time, error) {
	_, err := IntAsMonth(int(month))
	if err != nil {
		return time0, err
	}
	return dateYMDPrim(year, month+1, 1).AddDate(0, 0, -1), nil
}

// DaysInMonth returns number of days in the 'month' of the 'year'.
func DaysInMonth(year int, month time.Month) (int, error) {
	t, err := LastDayOfMonth(year, month)
	if err != nil {
		return 0, err
	}
	return t.Day(), nil
}

// dateYMDPrim returns Time corresponding to 'year', 'month' and 'day'.
func dateYMDPrim(year int, month time.Month, day int) time.Time {
	return time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
}

// DateYMDEr returns Time corresponding to 'year', 'month' and 'day'.
// If year, month and day represent invalid date, error is not nil.
func DateYMDEr(year, month, day int) (time.Time, error) {
	mon, err := IntAsMonth(month)
	if err != nil {
		return time0, err
	}
	d, _ := DaysInMonth(year, mon)
	if day < 1 || day > d {
		return time0, errors.New("invalid day")
	}
	return dateYMDPrim(year, mon, day), nil
}

// DateYMD returns (with the help of DateYMDEr) Time corresponding to 'year', 'month' and 'day',
// but panics in case of error.
func DateYMD(year, month, day int) time.Time {
	r, err := DateYMDEr(year, month, day)
	if err != nil {
		panic(err)
	}
	return r
}

// FirstWeekdayInYear returns first occurrence of the weekday in the 'year'.
func FirstWeekdayInYear(year int, weekday time.Weekday) (time.Time, error) {
	if weekday < time.Sunday || weekday > time.Saturday {
		return time0, errors.New("invalid weekday")
	}
	t := dateYMDPrim(year, time.January, 1)
	for t.Weekday() != weekday {
		t = t.AddDate(0, 0, 1)
	}
	return t, nil
}

// FirstWeekdayInYear2 returns first occurrence of the weekday in the 'year'.
func FirstWeekdayInYear2(year int, weekday time.Weekday) (time.Time, error) {
	if weekday < time.Sunday || weekday > time.Saturday {
		return time0, errors.New("invalid weekday")
	}
	t1 := dateYMDPrim(year, time.January, 1)
	days := int(weekday - t1.Weekday())
	if days == 0 {
		return t1, nil
	}
	if days < 0 {
		days = days + 7
	}
	return t1.AddDate(0, 0, days), nil
}

// WeekdaysInYear returns number of weekdays in the 'year'.
func WeekdaysInYear(year int, weekday time.Weekday) (int, error) {
	wd, err := FirstWeekdayInYear(year, weekday)
	if err != nil {
		return 0, err
	}
	res := 0
	for wd.Year() == year {
		res++
		wd = wd.AddDate(0, 0, 7)
	}
	return res, nil
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

// PrevClosestWeekday returns 'weekday' closest to 't' and not after 't'.
func PrevClosestWeekday(t time.Time, weekday time.Weekday) (time.Time, error) {
	if weekday < time.Sunday || weekday > time.Saturday {
		return time0, errors.New("invalid weekday")
	}
	days := int(weekday - t.Weekday())
	if days > 0 {
		days = days - 7
	}
	return t.AddDate(0, 0, days), nil
}

// NextClosestWeekday returns 'weekday' closest to 't' and not before 't'.
func NextClosestWeekday(t time.Time, weekday time.Weekday) (time.Time, error) {
	if weekday < time.Sunday || weekday > time.Saturday {
		return time0, errors.New("invalid weekday")
	}
	days := int(weekday - t.Weekday())
	if days < 0 {
		days = days + 7
	}
	return t.AddDate(0, 0, days), nil
}
