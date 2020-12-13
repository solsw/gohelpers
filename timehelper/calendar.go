package timehelper

import (
	"errors"
	"time"
)

// JulianToGregorian converts Julian date to Gregorian date.
func JulianToGregorian(julian time.Time) (time.Time, error) {
	switch {
	case julian.After(DateYMD(2400, time.February, 28)):
		return time.Time{}, errors.New("unsupported julian date")
	case julian.After(DateYMD(2300, time.February, 28)):
		return julian.AddDate(0, 0, 16), nil
	case julian.After(DateYMD(2200, time.February, 28)):
		return julian.AddDate(0, 0, 15), nil
	case julian.After(DateYMD(2100, time.February, 28)):
		return julian.AddDate(0, 0, 14), nil
	case julian.After(DateYMD(1900, time.February, 28)):
		return julian.AddDate(0, 0, 13), nil
	case julian.After(DateYMD(1800, time.February, 28)):
		return julian.AddDate(0, 0, 12), nil
	case julian.After(DateYMD(1700, time.February, 28)):
		return julian.AddDate(0, 0, 11), nil
	case julian.After(DateYMD(1500, time.February, 28)):
		return julian.AddDate(0, 0, 10), nil
	case julian.After(DateYMD(1400, time.February, 28)):
		return julian.AddDate(0, 0, 9), nil
	case julian.After(DateYMD(1300, time.February, 28)):
		return julian.AddDate(0, 0, 8), nil
	}
	return time.Time{}, errors.New("unsupported julian date")
}
