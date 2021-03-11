// Package ioutilhelper contains ioutil helpers.
package ioutilhelper

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

// TempFileJustName returns just name of a temporary file.
// (See ioutil.TempFile for 'pattern' usage.)
func TempFileJustName(pattern string) (string, error) {
	f, err := ioutil.TempFile("", pattern)
	if err != nil {
		return "", err
	}
	defer func() {
		f.Close()
		os.Remove(f.Name())
	}()
	return filepath.Base(f.Name()), nil
}
