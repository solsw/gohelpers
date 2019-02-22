package filepathhelper

import (
	"path/filepath"

	"github.com/vechabus/go/pathhelper"
)

// BaseNoExt returns the file name of the specified path string without the extension.
func BaseNoExt(p string) string {
	if p == "" || p[len(p)-1] == filepath.Separator {
		return ""
	}
	b := filepath.Base(p)
	e := filepath.Ext(b)
	le := len(e)
	if le == 0 {
		return b
	}
	return b[:len(b)-le]
}

// SplitFilePath splits p (using filepath.Separator as seperator) into directories and filename.
// (E.g. on Windows "a\b\c.d" is splitted into {"a", "b", "c.d"} slice.)
func SplitFilePath(p string) []string {
	return pathhelper.SplitPath(filepath.ToSlash(p))
}

// StartSeparator ensures, that p starts with filepath.Separator.
// If p is empty, empty string is returned.
func StartSeparator(p string) string {
	if p == "" {
		return ""
	}
	if p[0] != filepath.Separator {
		return string(filepath.Separator) + p
	}
	return p
}

// EndSeparator ensures, that p ends with filepath.Separator.
// If p is empty, empty string is returned.
func EndSeparator(p string) string {
	if p == "" {
		return ""
	}
	if p[len(p)-1] != filepath.Separator {
		return p + string(filepath.Separator)
	}
	return p
}
