package stringhelper

import (
	"errors"
	"strings"
)

func wordFromWords(ww []string, n uint, last bool) (string, error) {
	if len(ww) == 0 {
		return "", errors.New("no words in string")
	}
	if last {
		return ww[len(ww)-1], nil
	}
	if n >= uint(len(ww)) {
		return "", errors.New("n is too large")
	}
	return ww[n], nil
}

// NthWord returns the 'n'-th (starting with 0) word from 's'.
// 's' is separated into words with white space characters (see strings.Fields).
func NthWord(s string, n uint) (string, error) {
	return wordFromWords(strings.Fields(s), n, false)
}

// LastWord returns the last word from 's'.
// 's' is separated into words with white space characters (see strings.Fields).
func LastWord(s string) (string, error) {
	return wordFromWords(strings.Fields(s), 0, true)
}

// NthWordFunc returns the 'n'-th (starting with 0) word from 's'.
// 's' is separated into words with 'f' (see strings.FieldsFunc).
func NthWordFunc(s string, n uint, f func(rune) bool) (string, error) {
	if f == nil {
		return "", errors.New("f is nil")
	}
	return wordFromWords(strings.FieldsFunc(s, f), n, false)
}

// LastWordFunc returns the last word from 's'.
// 's' is separated into words with 'f' (see strings.FieldsFunc).
func LastWordFunc(s string, f func(rune) bool) (string, error) {
	if f == nil {
		return "", errors.New("f is nil")
	}
	return wordFromWords(strings.FieldsFunc(s, f), 0, true)
}

// NthWordDelims returns the n-th (starting with 0) word from 's'.
// 'delims' - slice of word dilimeters.
// If 'delims' is empty, NthWord's result is returned.
func NthWordDelims(s string, n uint, delims []rune) (string, error) {
	if len(delims) == 0 {
		return NthWord(s, n)
	}
	return NthWordFunc(s, n, func(r rune) bool {
		for _, delim := range delims {
			if delim == r {
				return true
			}
		}
		return false
	})
}

// LastWordDelims returns the last word from 's'.
// 'delims' - slice of word dilimeters.
// If 'delims' is empty, LastWord's result is returned.
func LastWordDelims(s string, delims []rune) (string, error) {
	if len(delims) == 0 {
		return LastWord(s)
	}
	return LastWordFunc(s, func(r rune) bool {
		for _, delim := range delims {
			if delim == r {
				return true
			}
		}
		return false
	})
}
