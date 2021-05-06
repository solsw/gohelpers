// Package pathhelper contains path helpers.
package pathhelper

import (
	"path"
	"strings"
)

// SplitPath splits path 'p' (using slash as seperator) into directories and filename.
// (E.g. "a/b/c.d" is splitted into {"a", "b", "c.d"} slice.)
func SplitPath(p string) []string {
	p = path.Clean(p)
	if p == "." || p == "/" {
		return []string{}
	}
	return strings.Split(strings.Trim(p, "/"), "/")
}

// StartSlash returns path 'p' guaranteed to start with slash.
// If 'p' is empty, slash is returned.
func StartSlash(p string) string {
	if len(p) > 0 && p[0] == '/' {
		return p
	}
	return "/" + p
}

// EndSlash returns path 'p' guaranteed to end with slash.
// If 'p' is empty, slash is returned.
func EndSlash(p string) string {
	if len(p) > 0 && p[len(p)-1] == '/' {
		return p
	}
	return p + "/"
}
