package timehelper

import (
	"strconv"
	"time"
)

// UnixTextTime is a helper type to marshal/unmarshal time.Time to/from text as a Unix time in seconds.
type UnixTextTime time.Time

// MarshalText implements the encoding.TextMarshaler interface.
// Location associated with 't' is ignored.
func (t UnixTextTime) MarshalText() ([]byte, error) {
	return []byte(strconv.FormatInt(time.Time(t).Unix(), 10)), nil
}

func unixTextTimeUnmarshalStr(str string) (UnixTextTime, error) {
	secs, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return UnixTextTime{}, err
	}
	return UnixTextTime(time.Unix(secs, 0).In(time.UTC)), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
// Location of the unmarshaled value is time.UTC.
func (t *UnixTextTime) UnmarshalText(text []byte) error {
	utt, err := unixTextTimeUnmarshalStr(string(text))
	if err != nil {
		return err
	}
	*t = utt
	return nil
}
