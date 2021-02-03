package stringhelper

import (
	"errors"
)

// SubstrPrim retrieves a substring from 's' without error checking.
// The substring starts at a 'start' rune position and has a specified 'length'.
func SubstrPrim(s string, start, length int) string {
	if length == 0 {
		return ""
	}
	rr := make([]rune, length)
	copy(rr, []rune(s)[start:start+length])
	return string(rr)
}

// Substr retrieves a substring from 's'.
// The substring starts at a 'start' rune position and has a specified 'length'.
func Substr(s string, start, length int) (string, error) {
	if start < 0 || length < 0 {
		return "", errors.New("start and/or length is negative")
	}
	if start+length > len(s) {
		return "", errors.New("start plus length is greater than string length")
	}
	return SubstrPrim(s, start, length), nil
}

// SubstrBeg retrieves a substring with a specified 'length' from the beginning of 's'.
func SubstrBeg(s string, length int) (string, error) {
	if length > len(s) {
		return "", errors.New("length is greater than string length")
	}
	return Substr(s, 0, length)
}

// SubstrEnd retrieves a substring with a specified 'length' from the end of 's'.
func SubstrEnd(s string, length int) (string, error) {
	if length > len(s) {
		return "", errors.New("length is greater than string length")
	}
	return Substr(s, len(s)-length, length)
}

// SubstrToEnd retrieves a substring from 'start' rune position and to the end of 's'.
func SubstrToEnd(s string, start int) (string, error) {
	if start > len(s) {
		return "", errors.New("start is greater than string length")
	}
	return Substr(s, start, len(s)-start)
}
