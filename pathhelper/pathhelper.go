package pathhelper

import (
	"path"
	"strings"
)

// SplitPath splits path p (using slash as seperator) into directories and filename.
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

// StartSlash returns path p guaranteed to start with slash.
// If p is empty, slash is returned.
func StartSlash(p string) string {
	if p == "" {
		return "/"
	}
	if p[0] != '/' {
		return string('/') + p
	}
	return p
}

// EndSlash returns path p guaranteed to end with slash.
// If p is empty, slash is returned.
func EndSlash(p string) string {
	if p == "" {
		return "/"
	}
	if p[len(p)-1] != '/' {
		return p + string('/')
	}
	return p
}
