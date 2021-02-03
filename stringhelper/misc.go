package stringhelper

import (
	"strings"

	"github.com/solsw/gohelpers/oshelper"
)

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

// RemoveSGREsc removes SGR escape sequence from 's'
// (see https://en.wikipedia.org/wiki/ANSI_escape_code).
func RemoveSGREsc(s string) string {
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
