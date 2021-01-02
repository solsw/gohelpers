// Package ioutilhelper contains ioutil helpers.
package ioutilhelper

import (
	"bufio"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/solsw/gohelpers/oshelper"
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

// ReadFileStrings reads the file 'filename' and returns the contents as []string.
func ReadFileStrings(filename string) ([]string, error) {
	f, err := os.Open(filename)
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

// WriteFileStringsNewLine writes 'ss' to a file named by 'filename'.
// Each string (including the last one) is followed by 'newLine'.
// (See ioutil.WriteFile for 'filename' and 'perm' usage.)
func WriteFileStringsNewLine(filename string, ss []string, newLine string, perm os.FileMode) error {
	return ioutil.WriteFile(filename, []byte(strings.Join(ss, newLine)+newLine), perm)
}

// WriteFileStrings writes 'ss' to a file named by 'filename'.
// Each string (including the last one) is followed by oshelper.NewLine.
// (See ioutil.WriteFile for 'filename' and 'perm' usage.)
func WriteFileStrings(filename string, ss []string, perm os.FileMode) error {
	return WriteFileStringsNewLine(filename, ss, oshelper.NewLine, perm)
}
