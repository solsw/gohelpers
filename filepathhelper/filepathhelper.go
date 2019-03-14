package filepathhelper

import (
	"path/filepath"

	"github.com/solsw/gohelpers/pathhelper"
)

// NoExt returns path p without extension.
// If p is empty, empty string is returned. If there is no extension, p is returned.
func NoExt(p string) string {
	if p == "" {
		return ""
	}
	e := filepath.Ext(p)
	le := len(e)
	if le == 0 {
		return p
	}
	return p[:len(p)-le]
}

// BaseNoExt returns filename without extension from path p.
// If p is empty, empty string is returned. See TestBaseNoExt for usage exanples.
func BaseNoExt(p string) string {
	ne := NoExt(p)
	if ne == "" || ne[len(ne)-1] == filepath.Separator {
		return ""
	}
	return filepath.Base(ne)
}

// ChangeExt changes extension of path p to ext.
// If p or ext is empty, p is returned. ext may or may not start with dot.
func ChangeExt(p, ext string) string {
	if p == "" || ext == "" || ext == "." {
		return p
	}
	if ext[0] == '.' {
		return NoExt(p) + ext
	}
	return NoExt(p) + "." + ext
}

// SplitFilePath splits path p (using filepath.Separator as seperator) into directories and filename.
// (E.g. on Windows "a\b\c.d" is splitted into {"a", "b", "c.d"} slice.)
func SplitFilePath(p string) []string {
	return pathhelper.SplitPath(filepath.ToSlash(p))
}

// StartSeparator ensures, that path p starts with filepath.Separator.
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

// EndSeparator ensures, that path p ends with filepath.Separator.
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
