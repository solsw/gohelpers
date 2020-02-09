// Package oshelper contains os helper functions.
package oshelper

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

// FSEntryExists checks if the file system entry (file or directory) 'entryName' exists.
func FSEntryExists(entryName string) (bool, error) {
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

// FileExistsFunc checks if the file 'fileName' exists.
//
// 'f' (if not nil) is used to transform 'fileName' before error returning.
// (E.g. 'f' may extract just file name from full path.)
func FileExistsFunc(fileName string, f func(string) string) (bool, error) {
	if len(fileName) == 0 {
		return false, errors.New("fileName is empty")
	}
	fi, err := os.Stat(fileName)
	if err == nil {
		if !fi.IsDir() {
			// it is file
			return true, nil
		}
		if f != nil {
			fileName = f(fileName)
		}
		return false, fmt.Errorf("'%s' is directory, not file", fileName)
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// FileExists checks if the file 'fileName' exists.
func FileExists(fileName string) (bool, error) {
	return FileExistsFunc(fileName, nil)
}

// FileExistsMust checks if the file 'fileName' exists.
// In case of error 'false' is returned.
func FileExistsMust(fileName string) bool {
	fe, err := FileExists(fileName)
	if err != nil {
		return false
	}
	return fe
}

// DirExistsFunc checks if the directory 'dirName' exists.
//
// 'f' (if not nil) is used to transform 'dirName' before error returning.
// (E.g. 'f' may shorten excessively long 'dirName'.)
func DirExistsFunc(dirName string, f func(string) string) (bool, error) {
	if len(dirName) == 0 {
		return false, errors.New("dirName is empty")
	}
	fi, err := os.Stat(dirName)
	if err == nil {
		if fi.IsDir() {
			// it is directory
			return true, nil
		}
		if f != nil {
			dirName = f(dirName)
		}
		return false, fmt.Errorf("'%s' is file, not directory", dirName)
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// DirExists checks if the directory 'dirName' exists.
func DirExists(dirName string) (bool, error) {
	return DirExistsFunc(dirName, nil)
}

// DirExistsMust checks if the directory 'dirName' exists.
// In case of error 'false' is returned.
func DirExistsMust(dirName string) bool {
	de, err := DirExists(dirName)
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
