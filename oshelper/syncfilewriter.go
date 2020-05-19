package oshelper

import (
	"os"
)

// SyncFileWriter implements file-based io.Writer,
// that calls Sync on underlying file after every Write call.
type SyncFileWriter struct {
	fl *os.File
}

// NewSyncFileWriter returns a new SyncFileWriter based on file 'f'.
func NewSyncFileWriter(f *os.File) *SyncFileWriter {
	return &SyncFileWriter{fl: f}
}

// Write implements io.Writer interface.
func (sfw *SyncFileWriter) Write(p []byte) (n int, err error) {
	n, err = sfw.fl.Write(p)
	if err != nil {
		return
	}
	err = sfw.fl.Sync()
	return
}
