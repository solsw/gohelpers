// Package ioutilhelper contains ioutil helper functions.
package ioutilhelper

import (
	"bufio"
	"io/ioutil"
	"os"
	"strings"

	"github.com/solsw/gohelpers/oshelper"
)

// TempFileName returns filename of a temporary file.
// (See ioutil.TempFile for 'dir' and 'pattern' usage.)
func TempFileName(dir, pattern string) (string, error) {
	f, err := ioutil.TempFile(dir, pattern)
	if err != nil {
		return "", err
	}
	f.Close()
	os.Remove(f.Name())
	return f.Name(), nil
}

// TempFileNameMust returns filename of a temporary file.
// In case of error empty string is returned.
func TempFileNameMust() string {
	tfn, err := TempFileName("", "")
	if err != nil {
		return ""
	}
	return tfn
}

// ReadFileStrings reads the file 'filename' and returns its contents as []string.
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

// WriteFileStrings writes 'ss' to a file 'filename'.
// Each string is followed by oshelper.NewLine.
// (See ioutil.WriteFile for 'filename' and 'perm' usage.)
func WriteFileStrings(filename string, ss []string, perm os.FileMode) error {
	var b strings.Builder
	for i := range ss {
		b.WriteString(ss[i] + oshelper.NewLine)
	}
	err := ioutil.WriteFile(filename, []byte(b.String()), perm)
	if err != nil {
		return err
	}
	return nil
}
