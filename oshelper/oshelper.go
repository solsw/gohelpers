package oshelper

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

// FileExistsFunc reports whether regular file 'filename' exists.
//
// 'f' (if not nil) is used to process 'filename' before own error returning.
// (E.g. 'f' may extract just file name from full path.)
func FileExistsFunc(filename string, f func(string) string) (bool, error) {
	if filename == "" {
		return false, errors.New("filename is empty")
	}
	fi, err := os.Stat(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}
	// 'filename' exists
	if !fi.Mode().IsRegular() {
		if f != nil {
			filename = f(filename)
		}
		return false, fmt.Errorf("'%s' is not a regular file", filename)
	}
	return true, nil
}

// FileExists reports whether regular file 'filename' exists.
func FileExists(filename string) (bool, error) {
	return FileExistsFunc(filename, nil)
}

// FileExistsMust is like FileExists but returns 'false' in case of error.
func FileExistsMust(filename string) bool {
	fe, err := FileExists(filename)
	if err != nil {
		return false
	}
	return fe
}

// DirExistsFunc reports whether directory 'dirname' exists.
//
// 'f' (if not nil) is used to process 'dirname' before own error returning.
// (E.g. 'f' may shorten excessively long 'dirname'.)
func DirExistsFunc(dirname string, f func(string) string) (bool, error) {
	if dirname == "" {
		return false, errors.New("dirname is empty")
	}
	fi, err := os.Stat(dirname)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}
	// 'dirname' exists
	if !fi.IsDir() {
		if f != nil {
			dirname = f(dirname)
		}
		return false, fmt.Errorf("'%s' is not a directory", dirname)
	}
	return true, nil
}

// DirExists reports whether directory 'dirname' exists.
func DirExists(dirname string) (bool, error) {
	return DirExistsFunc(dirname, nil)
}

// DirExistsMust is like DirExists but returns 'false' in case of error.
func DirExistsMust(dirname string) bool {
	de, err := DirExists(dirname)
	if err != nil {
		return false
	}
	return de
}

// ExeDir returns an absolute representation of the directory name
// of the executable that has started the current process.
func ExeDir() (string, error) {
	exe, err := os.Executable()
	if err != nil {
		return "", err
	}
	exeDir, err := filepath.Abs(filepath.Dir(exe))
	if err != nil {
		return "", err
	}
	return exeDir, nil
}

// GetenvDef retrieves the value of the environment variable named by the 'key'.
// If the value is empty 'def' is returned.
func GetenvDef(key, def string) string {
	r := os.Getenv(key)
	if r == "" {
		return def
	}
	return r
}
