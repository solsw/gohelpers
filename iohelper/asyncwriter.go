package iohelper

import (
	"io"
	"math"
)

// AsyncWriter implements io.WriteCloser, that does not block on Write call.
type AsyncWriter struct {
	wr io.Writer
	ch chan []byte
	dn chan struct{}
	nb int
	er error
}

// NewAsyncWriterSizeFunc returns a new AsyncWriter with 'size' capacity of the underlying channel.
// 'f' (if not nil) is used to process bytes before writing.
func NewAsyncWriterSizeFunc(w io.Writer, size int, f func([]byte) []byte) *AsyncWriter {
	r := AsyncWriter{wr: w, ch: make(chan []byte, size), dn: make(chan struct{})}
	go func(aw *AsyncWriter) {
		defer func() { aw.dn <- struct{}{} }()
		for bb := range aw.ch {
			if aw.er != nil {
				return
			}
			if f != nil {
				bb = f(bb)
			}
			aw.nb, aw.er = aw.wr.Write(bb)
		}
	}(&r)
	return &r
}

// NewAsyncWriterSize returns a new AsyncWriter with 'size' capacity of the underlying channel
// and nil processing function (see NewAsyncWriterSizeFunc).
func NewAsyncWriterSize(w io.Writer, size int) *AsyncWriter {
	return NewAsyncWriterSizeFunc(w, size, nil)
}

// NewAsyncWriterFunc returns a new AsyncWriter with math.MaxInt16 capacity of the underlying channel
// and 'f' processing function (see NewAsyncWriterSizeFunc).
func NewAsyncWriterFunc(w io.Writer, f func([]byte) []byte) *AsyncWriter {
	return NewAsyncWriterSizeFunc(w, math.MaxInt16, f)
}

// NewAsyncWriter returns a new AsyncWriter with math.MaxInt16 capacity of the underlying channel
// and nil processing function (see NewAsyncWriterSizeFunc).
func NewAsyncWriter(w io.Writer) *AsyncWriter {
	return NewAsyncWriterSizeFunc(w, math.MaxInt16, nil)
}

// Write implements io.Writer interface.
// Write returns len(p) and nil error.
func (aw *AsyncWriter) Write(p []byte) (int, error) {
	// since the same slice may be passed to this method in separate calls (e.g. as log.Println does),
	// the current contents of 'p' must be copied to new local slice
	// (fmt.Printf("%p\n", p) <- prints the same address when AsyncWriter is passed to log.SetOutput)
	loc := make([]byte, len(p))
	// (fmt.Printf("%p\n", loc) <- prints different addresses)
	copy(loc, p)
	aw.ch <- loc
	return len(loc), nil
}

// Result returns result of the last Write call on the underlying io.Writer.
func (aw *AsyncWriter) Result() (int, error) {
	return aw.nb, aw.er
}

// Close closes the underlying channel.
// Close must be called (typically by "defer" statement) to wait for the underlying io.Writer to finish writing.
// If the underlying io.Writer is io.Closer, Close calls its Close method
// and returns corresponding result error, otherwise Close returns nil.
// Close implements io.Closer interface.
func (aw *AsyncWriter) Close() error {
	close(aw.ch)
	<-aw.dn
	if cl, ok := aw.wr.(io.Closer); ok {
		return cl.Close()
	}
	return nil
}
