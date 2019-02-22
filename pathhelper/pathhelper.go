package pathhelper

import (
	"path"
	"strings"
)

// SplitPath splits p (using '/' as seperator) into directories and filename.
// (E.g. "a/b/c.d" is splitted into {"a", "b", "c.d"} slice.)
func SplitPath(p string) []string {
	t := make([]string, 0)
	d := path.Clean(p)
	for len(d) > 0 {
		var f string
		d, f = path.Split(d)
		if len(f) > 0 {
			t = append(t, f)
		}
		d = strings.Trim(d, "/")
	}
	r := make([]string, 0, len(t))
	for i := len(t) - 1; i >= 0; i-- {
		r = append(r, t[i])
	}
	return r
}

// StartSlash ensures, that p starts with slash.
// If p is empty, empty string is returned.
func StartSlash(p string) string {
	if p == "" {
		return ""
	}
	if p[0] != '/' {
		return string('/') + p
	}
	return p
}

// EndSlash ensures, that p ends with slash.
// If p is empty, empty string is returned.
func EndSlash(p string) string {
	if p == "" {
		return ""
	}
	if p[len(p)-1] != '/' {
		return p + string('/')
	}
	return p
}
