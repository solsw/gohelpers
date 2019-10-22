package ioutilhelper

import (
	"io/ioutil"
	"os"
	"strings"
)

// TempFileName returns filename of a temporary file.
// (See ioutil.TempFile for dir and pattern usage.)
func TempFileName(dir, pattern string) (string, error) {
	f, err := ioutil.TempFile(dir, pattern)
	if err != nil {
		return "", err
	}
	f.Close()
	os.Remove(f.Name())
	return f.Name(), nil
}

// TempFileName0 returns filename of a temporary file.
// In case of any error empty string is returned.
func TempFileName0() string {
	tfn, err := TempFileName("", "")
	if err != nil {
		return ""
	}
	return tfn
}

// WriteStrings writes ss to a file named by filename.
// Each string is followed by '\n' rune.
// (See ioutil.WriteFile for filename and perm usage.)
func WriteStrings(filename string, ss []string, perm os.FileMode) error {
	var sb strings.Builder
	for _, s := range ss {
		sb.WriteString(s)
		sb.WriteRune('\n')
	}
	err := ioutil.WriteFile(filename, []byte(sb.String()), perm)
	if err != nil {
		return err
	}
	return nil
}
