package stringhelper

import (
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

// IsEmptyOrWhite reports whether 's' is empty or contains only white spaces.
func IsEmptyOrWhite(s string) bool {
	return strings.TrimSpace(s) == ""
}

// IsDigital reports whether 's' consists only of digits.
func IsDigital(s string) bool {
	if s == "" || !utf8.ValidString(s) {
		return false
	}
	for _, r := range s {
		if !unicode.IsDigit(r) {
			return false
		}
	}
	return true
}

// IsNumeric reports whether 's' represents a number.
func IsNumeric(s string) bool {
	if s == "" || !utf8.ValidString(s) {
		return false
	}
	if utf8.RuneCountInString(s) == 1 {
		r, _ := utf8.DecodeRuneInString(s)
		return unicode.IsDigit(r)
	}
	_, err := strconv.ParseFloat(s, 32)
	return err == nil
}

// IsUpper reports whether 's' is upper case.
func IsUpper(s string) bool {
	return s == strings.ToUpper(s)
}

// IsUpperRune reports whether 'r' is upper case.
func IsUpperRune(r rune) bool {
	return IsUpper(string(r))
}

// IsLower reports whether 's' is lower case.
func IsLower(s string) bool {
	return s == strings.ToLower(s)
}

// IsLowerRune reports whether 'r' is lower case.
func IsLowerRune(r rune) bool {
	return IsLower(string(r))
}
