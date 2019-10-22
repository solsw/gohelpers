package strhelper

import (
	"errors"
	"sort"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

// strhelper-related errors
var (
	ErrEmptyString   = errors.New("empty string")
	ErrInvalidString = errors.New("invalid string")
)

// IsEmptyOrWhite reports whether the string is empty or contains only white spaces.
func IsEmptyOrWhite(s string) bool {
	return len(strings.TrimSpace(s)) == 0
}

// IsDigital reports whether the string consists only of digits.
func IsDigital(s string) bool {
	if len(s) == 0 || !utf8.ValidString(s) {
		return false
	}
	for _, r := range s {
		if !unicode.IsDigit(r) {
			return false
		}
	}
	return true
}

// IsNumeric reports whether the string represents a number.
func IsNumeric(s string) bool {
	if len(s) == 0 || !utf8.ValidString(s) {
		return false
	}
	if utf8.RuneCountInString(s) == 1 {
		r, _ := utf8.DecodeRuneInString(s)
		return unicode.IsDigit(r)
	}
	_, err := strconv.ParseFloat(s, 32)
	return err == nil
}

func nthRunePrim(s string, n uint, strict bool) (rune, error) {
	if len(s) == 0 {
		return utf8.RuneError, ErrEmptyString
	}
	if strict && !utf8.ValidString(s) {
		return utf8.RuneError, ErrInvalidString
	}
	var count uint
	for _, r := range s {
		if !strict && r == utf8.RuneError {
			return utf8.RuneError, ErrInvalidString
		}
		if count == n {
			return r, nil
		}
		count++
	}
	return utf8.RuneError, errors.New("n is too large")
}

// NthRuneStrict returns n-th (starting with 0) rune from s.
// The string must be not empty and valid.
func NthRuneStrict(s string, n uint) (rune, error) {
	return nthRunePrim(s, n, true)
}

// NthRuneAny returns the 'n'-th (starting with 0) rune from the string.
// The string must be not empty, but may be not valid.
// If the string is invalid and the required rune is situated before an invalid UTF-8 sequence,
// the rune is returned without error.
func NthRuneAny(s string, n uint) (rune, error) {
	return nthRunePrim(s, n, false)
}

func nthWordFromWords(ww []string, n uint) (string, error) {
	if len(ww) == 0 {
		return "", errors.New("no words in string")
	}
	if uint(len(ww)) <= n {
		return "", errors.New("n is too large")
	}
	return ww[n], nil
}

// NthWord returns the 'n'-th (starting with 0) word from the string.
func NthWord(s string, n uint) (string, error) {
	if len(s) == 0 {
		return "", ErrEmptyString
	}
	return nthWordFromWords(strings.Fields(s), n)
}

// LastWord returns the last word from the string.
func LastWord(s string) (string, error) {
	if len(s) == 0 {
		return "", ErrEmptyString
	}
	ww := strings.Fields(s)
	return nthWordFromWords(ww, uint(len(ww)-1))
}

func wordByDelims(s string, n uint, delims []rune, last bool) (string, error) {
	if len(s) == 0 {
		return "", ErrEmptyString
	}
	if len(delims) == 0 {
		if last {
			return LastWord(s)
		}
		return NthWord(s, n)
	}
	ww := strings.FieldsFunc(s, func(r rune) bool {
		for _, delim := range delims {
			if delim == r {
				return true
			}
		}
		return false
	})
	if last {
		return nthWordFromWords(ww, uint(len(ww)-1))
	}
	return nthWordFromWords(ww, n)
}

// NthWordDelims returns the 'n'-th (starting with 0) word from the string.
// 'delims' - slice of word dilimeters.
// If 'delims' is empty, NthWord's result is returned.
func NthWordDelims(s string, n uint, delims []rune) (string, error) {
	return wordByDelims(s, n, delims, false)
}

// LastWordDelims returns the last word from the string.
// 'delims' - slice of word dilimeters.
// If 'delims' is empty, LastWord's result is returned.
func LastWordDelims(s string, delims []rune) (string, error) {
	return wordByDelims(s, 0, delims, true)
}

// SubstrPrim retrieves a substring from the string without error checking.
// The substring starts at a 'start' rune position and has a specified 'length'.
func SubstrPrim(s string, start, length int) string {
	if length == 0 {
		return ""
	}
	rr := make([]rune, length)
	copy(rr, []rune(s)[start:start+length])
	return string(rr)
}

// Substr retrieves a substring from the string.
// The substring starts at a 'start' rune position and has a specified 'length'.
func Substr(s string, start, length int) (string, error) {
	if start < 0 {
		return "", errors.New("start is less than zero")
	}
	if length < 0 {
		return "", errors.New("length is less than zero")
	}
	if start+length > len(s) {
		return "", errors.New("start plus length is greater than string length")
	}
	return SubstrPrim(s, start, length), nil
}

// SubstrBeg retrieves a substring from the string.
// The substring starts at the beginning of the string and has a specified 'length'.
func SubstrBeg(s string, length int) (string, error) {
	if length > len(s) {
		return "", errors.New("length is greater than string length")
	}
	return Substr(s, 0, length)
}

// SubstrEnd retrieves a substring from the string.
// The substring starts at a 'start' rune position and continues to the end of the string.
func SubstrEnd(s string, start int) (string, error) {
	if start > len(s) {
		return "", errors.New("start is beyond len(s)")
	}
	return Substr(s, start, len(s)-start)
}

// LastByte returns the last byte from the string.
func LastByte(s string) (byte, error) {
	if len(s) == 0 {
		return 0, ErrEmptyString
	}
	return s[len(s)-1], nil
}

// LastRune returns the last rune from the string.
func LastRune(s string) (rune, error) {
	if len(s) == 0 {
		return utf8.RuneError, ErrEmptyString
	}
	rr := []rune(s)
	return rr[len(rr)-1], nil
}

// Unique returns unique strings from ss preserving order.
func Unique(ss []string) []string {
	if len(ss) < 2 {
		return ss
	}
	var res []string
	var m = make(map[string]interface{})
	for _, s := range ss {
		_, ok := m[s]
		if !ok {
			res = append(res, s)
			m[s] = nil
		}
	}
	return res
}

// UniqueSorted returns sorted unique strings from ss.
// (May be up to two times faster than Unique. Subject for benchmarking.)
func UniqueSorted(ss []string) []string {
	if len(ss) < 2 {
		return ss
	}
	sort.Strings(ss)
	res := []string{ss[0]}
	for i := 1; i < len(ss); i++ {
		if ss[i] != ss[i-1] {
			res = append(res, ss[i])
		}
	}
	return res
}
