package ioutilhelper

import (
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

// WriteStrings writes 'ss' to a file named by 'filename'.
// Each string is followed by oshelper.NewLine.
// (See ioutil.WriteFile for 'filename' and 'perm' usage.)
func WriteStrings(filename string, ss []string, perm os.FileMode) error {
	var sb strings.Builder
	for i := range ss {
		sb.WriteString(ss[i] + oshelper.NewLine)
	}
	err := ioutil.WriteFile(filename, []byte(sb.String()), perm)
	if err != nil {
		return err
	}
	return nil
}
