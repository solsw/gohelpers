// Package delphi contains Delphi helpers.
package delphi

import (
	"fmt"
	"math"
	"time"

	"github.com/solsw/gohelpers/timehelper"
)

// TDateTime is Delphi's type to represent an instant in time.
type TDateTime = float64

// unixStart is TDateTime of 01/01/1970
const unixStart TDateTime = 25569.0

// TDateTimeToTime converts TDateTime to time.Time.
func TDateTimeToTime(dt TDateTime) time.Time {
	if dt < 0 {
		d, t := math.Modf(dt)
		dt = d - t
	}
	// Unix time in seconds
	usec := math.Round((dt - unixStart) * 86400)
	return time.Unix(int64(usec), 0)
}

// TimeToTDateTime converts time.Time to TDateTime.
func TimeToTDateTime(tt time.Time) TDateTime {
	// Unix time in seconds
	usec := tt.Unix()
	dt := (float64(usec) / 86400) + unixStart
	if dt >= 0 {
		return dt
	}
	d, t := math.Modf(dt)
	return d - 1 - (1 + t)
}

// EncodeDate returns a TDateTime value that represents a specified Year, Month and Day.
func EncodeDate(Year, Month, Day int) (TDateTime, error) {
	if !(0 <= Year && Year <= 9999 && 1 <= Month && Month <= 12 && 1 <= Day && Day <= timehelper.DaysInMonth(Year, time.Month(Month))) {
		return -math.MaxFloat64, fmt.Errorf("invalid date '%d-%d-%d'", Year, Month, Day)
	}
	return TimeToTDateTime(timehelper.DateYMD(Year, time.Month(Month), Day)), nil
}
