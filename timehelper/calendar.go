package timehelper

import (
	"errors"
	"time"
)

// JulianToGregorian converts Julian date to Gregorian date.
func JulianToGregorian(julian time.Time) (time.Time, error) {
	if julian.After(dateYMDPrim(2400, time.February, 28)) {
		return time0, errors.New("unsupported julian date")
	}
	if julian.After(dateYMDPrim(2300, time.February, 28)) {
		return julian.AddDate(0, 0, 16), nil
	}
	if julian.After(dateYMDPrim(2200, time.February, 28)) {
		return julian.AddDate(0, 0, 15), nil
	}
	if julian.After(dateYMDPrim(2100, time.February, 28)) {
		return julian.AddDate(0, 0, 14), nil
	}
	if julian.After(dateYMDPrim(1900, time.February, 28)) {
		return julian.AddDate(0, 0, 13), nil
	}
	if julian.After(dateYMDPrim(1800, time.February, 28)) {
		return julian.AddDate(0, 0, 12), nil
	}
	if julian.After(dateYMDPrim(1700, time.February, 28)) {
		return julian.AddDate(0, 0, 11), nil
	}
	if julian.After(dateYMDPrim(1500, time.February, 28)) {
		return julian.AddDate(0, 0, 10), nil
	}
	if julian.After(dateYMDPrim(1400, time.February, 28)) {
		return julian.AddDate(0, 0, 9), nil
	}
	if julian.After(dateYMDPrim(1300, time.February, 28)) {
		return julian.AddDate(0, 0, 8), nil
	}
	return time0, errors.New("unsupported julian date")
}
