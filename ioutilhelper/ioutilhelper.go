package ioutilhelper

import (
	"io/ioutil"
	"os"
)

// TempFileName returns filename of a temporary file.
// (See ioutil.TempFile for details about 'dir' and 'pattern'.)
func TempFileName(dir, pattern string) (string, error) {
	f, err := ioutil.TempFile(dir, pattern)
	if err != nil {
		return "", err
	}
	f.Close()
	os.Remove(f.Name())
	return f.Name(), nil
}

// TempFileName0 returns filename of a temporary file ignoring errors.
func TempFileName0() string {
	tfn, _ := TempFileName("", "")
	return tfn
}
