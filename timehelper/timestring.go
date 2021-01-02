package timehelper

import (
	"time"
)

// TimeString is a helper type to store time.Time as string.
//
// TimeString may be used instead of time.Time for encoding/decoding time.Time as JSON,
// since time.Time does not respect encoding/json 'omitempty' option.
type TimeString string

// NewTimeString creates TimeString from 't'.
func NewTimeString(t time.Time) (TimeString, error) {
	if t.IsZero() {
		return "", nil
	}
	bb, err := t.MarshalText()
	if err != nil {
		return "", err
	}
	return TimeString(bb), nil
}

// Time converts 'ts' to time.Time.
func (ts TimeString) Time() (time.Time, error) {
	if len(ts) == 0 {
		return time.Time{}, nil
	}
	var t time.Time
	if err := t.UnmarshalText([]byte(ts)); err != nil {
		return time.Time{}, err
	}
	return t, nil
}
