package oshelper

import (
	"io"
	"os"
)

// syncFileWriter is a file-based io.Writer that calls Sync on underlying file after every Write call.
type syncFileWriter struct {
	file *os.File
}

// Write implements io.Writer interface.
func (sfw *syncFileWriter) Write(b []byte) (n int, err error) {
	n, err = sfw.file.Write(b)
	if err != nil {
		return
	}
	err = sfw.file.Sync()
	return
}

// SyncFileWriter creates file-based io.Writer
// that calls Sync on underlying file 'f' after every Write call.
func SyncFileWriter(f *os.File) io.Writer {
	return &syncFileWriter{file: f}
}
