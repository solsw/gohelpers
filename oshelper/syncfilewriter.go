package oshelper

import (
	"os"
)

// SyncFileWriter is the os.File-based io.Writer interface implementation,
// that calls Sync on underlying file after every Write call.
type SyncFileWriter struct {
	fl *os.File
}

// NewSyncFileWriter returns a new SyncFileWriter based on file 'f'.
func NewSyncFileWriter(f *os.File) *SyncFileWriter {
	return &SyncFileWriter{fl: f}
}

// Write implements the io.Writer interface.
func (sfw *SyncFileWriter) Write(p []byte) (n int, err error) {
	n, err = sfw.fl.Write(p)
	if err != nil {
		return
	}
	err = sfw.fl.Sync()
	return
}
