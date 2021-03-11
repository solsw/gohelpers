package oshelper

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
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

// ReadFileStrings reads the file 'name' and returns the contents as []string.
func ReadFileStrings(name string) ([]string, error) {
	f, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	var r []string
	s := bufio.NewScanner(f)
	for s.Scan() {
		r = append(r, s.Text())
	}
	return r, s.Err()
}

// WriteFileAll writes all contents of 'r' to the named file.
// (See os.WriteFile for details.)
func WriteFileAll(name string, r io.Reader, perm os.FileMode) error {
	data, err := io.ReadAll(r)
	if err != nil {
		return err
	}
	return os.WriteFile(name, data, perm)
}

// WriteFileStringsNewLine writes 'ss' to the named file.
// Each string (including the last one) is followed by 'newLine'.
// (See os.WriteFile for details.)
func WriteFileStringsNewLine(name string, ss []string, newLine string, perm os.FileMode) error {
	return os.WriteFile(name, []byte(strings.Join(ss, newLine)+newLine), perm)
}

// WriteFileStrings writes 'ss' to the named file.
// Each string (including the last one) is followed by oshelper.NewLine.
// (See os.WriteFile for details.)
func WriteFileStrings(name string, ss []string, perm os.FileMode) error {
	return WriteFileStringsNewLine(name, ss, NewLine, perm)
}
