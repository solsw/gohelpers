// Package stringhelper contains string-related helpers.
package stringhelper

import (
	"errors"
	"sort"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/solsw/gohelpers/oshelper"
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

// SubstrPrim retrieves a substring from 's' without error checking.
// The substring starts at a 'start' rune position and has a specified 'length'.
func SubstrPrim(s string, start, length uint) string {
	if length == 0 {
		return ""
	}
	rr := make([]rune, length)
	copy(rr, []rune(s)[start:start+length])
	return string(rr)
}

// Substr retrieves a substring from 's'.
// The substring starts at a 'start' rune position and has a specified 'length'.
func Substr(s string, start, length uint) (string, error) {
	if start+length > uint(len(s)) {
		return "", errors.New("start plus length is greater than string length")
	}
	return SubstrPrim(s, start, length), nil
}

// SubstrBeg retrieves a substring with a specified 'length' from the beginning of 's'.
func SubstrBeg(s string, length uint) (string, error) {
	if length > uint(len(s)) {
		return "", errors.New("length is greater than string length")
	}
	return Substr(s, 0, length)
}

// SubstrEnd retrieves a substring with a specified 'length' from the end of 's'.
func SubstrEnd(s string, length uint) (string, error) {
	if length > uint(len(s)) {
		return "", errors.New("length is greater than string length")
	}
	return Substr(s, uint(len(s))-length, length)
}

// SubstrToEnd retrieves a substring from 'start' rune position and to the end of 's'.
func SubstrToEnd(s string, start uint) (string, error) {
	if start > uint(len(s)) {
		return "", errors.New("start is greater than string length")
	}
	return Substr(s, start, uint(len(s))-start)
}

// Unique returns unique strings from 'ss', preserving order of strings in 'ss'.
func Unique(ss []string) []string {
	if len(ss) < 2 {
		return ss
	}
	var res []string
	var m = make(map[string]interface{})
	for i := range ss {
		_, ok := m[ss[i]]
		if !ok {
			res = append(res, ss[i])
			m[ss[i]] = nil
		}
	}
	return res
}

// UniqueSorted returns sorted unique strings from 'ss'.
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

// RemoveEscSGR removes SGR escape sequence from 's'
// (see https://en.wikipedia.org/wiki/ANSI_escape_code).
func RemoveEscSGR(s string) string {
	esc := false
	return strings.Map(func(r rune) rune {
		if r == '\x1B' {
			esc = true
			return -1
		}
		if esc {
			if r == 'm' {
				esc = false
			}
			return -1
		}
		return r
	}, s)
}

// AdjustNewLines replaces end of line sequences ("\r", "\n", "\r\n") within 's' with oshelper.NewLine.
func AdjustNewLines(s string) string {
	nn := strings.Split(strings.ReplaceAll(s, "\r\n", "\n"), "\n")
	var sb strings.Builder
	for i := range nn {
		rr := strings.Split(nn[i], "\r")
		for j := range rr {
			sb.WriteString(rr[j] + oshelper.NewLine)
		}
	}
	return strings.TrimSuffix(sb.String(), oshelper.NewLine)
}

// StringToStrings slices 's' into all substrings separated by end of line sequences ("\r", "\n", "\r\n").
func StringToStrings(s string) []string {
	return strings.Split(AdjustNewLines(s), oshelper.NewLine)
}

// RemoveLastStringIfEmpty removes last string from 'ss' if this string is empty.
func RemoveLastStringIfEmpty(ss []string) []string {
	if len(ss) == 0 {
		return ss
	}
	if ss[len(ss)-1] == "" {
		return ss[:len(ss)-1]
	}
	return ss
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
