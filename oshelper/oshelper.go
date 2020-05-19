package oshelper

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

// EntryExists checks if file system entry (file or directory) 'entryName' exists.
func EntryExists(entryName string) (bool, error) {
	if len(entryName) == 0 {
		return false, errors.New("entryName is empty")
	}
	_, err := os.Stat(entryName)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// FileExistsFunc checks if file 'filename' exists.
//
// 'f' (if not nil) is used to process 'filename' before own error returning.
// (E.g. 'f' may extract just file name from full path.)
func FileExistsFunc(filename string, f func(string) string) (bool, error) {
	if len(filename) == 0 {
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
	if fi.Mode().IsRegular() {
		// it is a regular file
		return true, nil
	}
	if f != nil {
		filename = f(filename)
	}
	return false, fmt.Errorf("'%s' is not a regular file", filename)
}

// FileExists checks if file 'filename' exists.
func FileExists(filename string) (bool, error) {
	return FileExistsFunc(filename, nil)
}

// FileExistsMust checks if file 'filename' exists.
// In case of error 'false' is returned.
func FileExistsMust(filename string) bool {
	fe, err := FileExists(filename)
	if err != nil {
		return false
	}
	return fe
}

// DirExistsFunc checks if directory 'dirname' exists.
//
// 'f' (if not nil) is used to process 'dirname' before own error returning.
// (E.g. 'f' may shorten excessively long 'dirname'.)
func DirExistsFunc(dirname string, f func(string) string) (bool, error) {
	if len(dirname) == 0 {
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
	if fi.IsDir() {
		// it is a directory
		return true, nil
	}
	if f != nil {
		dirname = f(dirname)
	}
	return false, fmt.Errorf("'%s' is not a directory", dirname)
}

// DirExists checks if directory 'dirname' exists.
func DirExists(dirname string) (bool, error) {
	return DirExistsFunc(dirname, nil)
}

// DirExistsMust checks if directory 'dirname' exists.
// In case of error 'false' is returned.
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
