package iohelper

import (
	"io"
)

// AsyncWriter is the io.WriteCloser interface implementation, that does not block on Write call.
type AsyncWriter struct {
	writer    io.Writer
	ch        chan []byte
	done      chan struct{}
	isWriting bool
	nbytes    int
	er        error
}

// NewAsyncWriter returns a new AsyncWriter with 'size' capacity of the underlying channel.
// 'f' (if not nil) is used to process bytes before writing.
func NewAsyncWriter(w io.Writer, size int, f func([]byte) []byte) *AsyncWriter {
	newaw := AsyncWriter{writer: w, ch: make(chan []byte, size), done: make(chan struct{})}
	go func(aw *AsyncWriter) {
		defer func() { aw.done <- struct{}{} }()
		for bb := range aw.ch {
			if aw.er != nil {
				return
			}
			if f != nil {
				bb = f(bb)
			}
			aw.isWriting = true
			aw.nbytes, aw.er = aw.writer.Write(bb)
			aw.isWriting = false
		}
	}(&newaw)
	return &newaw
}

// Write implements the io.Writer interface.
// Write always returns len(p) and nil.
func (aw *AsyncWriter) Write(p []byte) (int, error) {
	// since the same slice may be passed to this method in separate calls (e.g. as log.Println does),
	// the current contents of 'p' must be copied to new local slice
	// fmt.Printf("%p\n", p) <- prints the same address when AsyncWriter is passed to log.SetOutput
	loc := make([]byte, len(p))
	// fmt.Printf("%p\n", loc) <- prints different addresses
	copy(loc, p)
	aw.ch <- loc
	return len(loc), nil
}

// IsWriting reports whether the underlying io.Writer is in the writing phase or not.
func (aw *AsyncWriter) IsWriting() bool {
	return aw.isWriting
}

// Result returns result of the last completed Write call on the underlying io.Writer.
func (aw *AsyncWriter) Result() (int, error) {
	return aw.nbytes, aw.er
}

// Close closes AsyncWriter by closing the underlying channel and waiting for the underlying io.Writer to finish writing.
// If the underlying io.Writer is io.Closer, it is closed and the result is returned, otherwise Close returns nil.
// Close implements the io.Closer interface.
func (aw *AsyncWriter) Close() error {
	close(aw.ch)
	<-aw.done
	if cl, ok := aw.writer.(io.Closer); ok {
		return cl.Close()
	}
	return nil
}
