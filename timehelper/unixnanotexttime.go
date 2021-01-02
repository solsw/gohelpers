package timehelper

import (
	"strconv"
	"time"
)

// UnixNanoTextTime is a helper type to marshal/unmarshal time.Time to/from text as a Unix time in nanoseconds.
type UnixNanoTextTime time.Time

// MarshalText implements the encoding.TextMarshaler interface.
func (t UnixNanoTextTime) MarshalText() ([]byte, error) {
	return []byte(strconv.FormatInt(time.Time(t).UnixNano(), 10)), nil
}

func unixNanoTextTimeUnmarshalStr(str string) (UnixNanoTextTime, error) {
	nsecs, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return UnixNanoTextTime{}, err
	}
	return UnixNanoTextTime(time.Unix(0, nsecs).In(time.UTC)), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
// Location of the unmarshaled value is time.UTC.
func (t *UnixNanoTextTime) UnmarshalText(text []byte) error {
	untt, err := unixNanoTextTimeUnmarshalStr(string(text))
	if err != nil {
		return err
	}
	*t = untt
	return nil
}
