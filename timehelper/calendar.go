package timehelper

import (
	"errors"
	"time"
)

// JulianToGregorian converts Julian date to Gregorian date.
func JulianToGregorian(julian time.Time) (time.Time, error) {
	switch {
	case julian.After(ymd(2400, time.February, 28)):
		return time0, errors.New("unsupported julian date")
	case julian.After(ymd(2300, time.February, 28)):
		return julian.AddDate(0, 0, 16), nil
	case julian.After(ymd(2200, time.February, 28)):
		return julian.AddDate(0, 0, 15), nil
	case julian.After(ymd(2100, time.February, 28)):
		return julian.AddDate(0, 0, 14), nil
	case julian.After(ymd(1900, time.February, 28)):
		return julian.AddDate(0, 0, 13), nil
	case julian.After(ymd(1800, time.February, 28)):
		return julian.AddDate(0, 0, 12), nil
	case julian.After(ymd(1700, time.February, 28)):
		return julian.AddDate(0, 0, 11), nil
	case julian.After(ymd(1500, time.February, 28)):
		return julian.AddDate(0, 0, 10), nil
	case julian.After(ymd(1400, time.February, 28)):
		return julian.AddDate(0, 0, 9), nil
	case julian.After(ymd(1300, time.February, 28)):
		return julian.AddDate(0, 0, 8), nil
	}
	return time0, errors.New("unsupported julian date")
}
