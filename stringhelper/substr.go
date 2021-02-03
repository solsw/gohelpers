package stringhelper

import (
	"fmt"
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
		return "", fmt.Errorf("start (%d) and/or length (%d) is negative", start, length)
	}
	if start+length > len(s) {
		return "", fmt.Errorf("start (%d) plus length (%d) is greater than string length (%d)", start, length, len(s))
	}
	return SubstrPrim(s, start, length), nil
}

// SubstrBeg retrieves a substring with a specified 'length' from the beginning of 's'.
func SubstrBeg(s string, length int) (string, error) {
	if length > len(s) {
		return "", fmt.Errorf("length (%d) is greater than string length (%d)", length, len(s))
	}
	return Substr(s, 0, length)
}

// SubstrEnd retrieves a substring with a specified 'length' from the end of 's'.
func SubstrEnd(s string, length int) (string, error) {
	if length > len(s) {
		return "", fmt.Errorf("length (%d) is greater than string length (%d)", length, len(s))
	}
	return Substr(s, len(s)-length, length)
}

// SubstrToEnd retrieves a substring from 'start' rune position and to the end of 's'.
func SubstrToEnd(s string, start int) (string, error) {
	if start > len(s) {
		return "", fmt.Errorf("start (%d) is greater than string length (%d)", start, len(s))
	}
	return Substr(s, start, len(s)-start)
}
