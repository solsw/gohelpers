package timehelper

import (
	"fmt"
	"time"
)

// ValidMonth checks if 'mon' represents a valid month.
func ValidMonth(month time.Month) error {
	if !(time.January <= month && month <= time.December) {
		return fmt.Errorf("invalid month '%d'", month)
	}
	return nil
}

// ValidDate checks if 'year', 'mon', 'day' constitute a valid date.
func ValidDate(year int, month time.Month, day int) error {
	if err := ValidMonth(month); err != nil || !(1 <= day && day <= DaysInMonth(year, month)) {
		return fmt.Errorf("invalid date '%d-%d-%d'", year, month, day)
	}
	return nil
}

// ValidClockTime checks if 'hour', 'min', 'sec' constitute a valid clock time.
func ValidClockTime(hour, min, sec int) error {
	if !(0 <= hour && hour <= 23 && 0 <= min && min <= 59 && 0 <= sec && sec <= 59) {
		return fmt.Errorf("invalid clock time '%d:%d:%d'", hour, min, sec)
	}
	return nil
}

// ValidWeekday checks if 'weekday' is valid.
func ValidWeekday(weekday time.Weekday) error {
	if !(time.Sunday <= weekday && weekday <= time.Saturday) {
		return fmt.Errorf("invalid weekday '%d'", weekday)
	}
	return nil
}
