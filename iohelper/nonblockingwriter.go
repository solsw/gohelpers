package iohelper

import (
	"io"
)

// nonBlockingWriter is an io.WriteCloser that does not block on write operations.
type nonBlockingWriter struct {
	wr io.Writer
	bb chan []byte
	dn chan bool
}

// Write implements io.Writer interface.
func (nbw *nonBlockingWriter) Write(b []byte) (n int, err error) {
	// since the same slice may be passed to this method in separate calls (e.g. as log.Println does),
	// the current contents of 'b' must be copied to new local slice
	// (fmt.Printf("%p\n", b) <- prints same address when NonBlockingWriter is used by log.SetOutput)
	l := make([]byte, len(b))
	// (fmt.Printf("%p\n", l) <- prints different addresses)
	copy(l, b)
	nbw.bb <- l
	return len(b), nil
}

// Close implements io.Closer interface.
// Must be called to stop nonBlockingWriter and free resources.
func (nbw *nonBlockingWriter) Close() error {
	close(nbw.bb)
	<-nbw.dn
	c, ok := nbw.wr.(io.Closer)
	if ok {
		return c.Close()
	}
	return nil
}

// NonBlockingWriter creates io.WriteCloser that does not block on Write operations.
// 'cap' is capacity of underlying channel used to hold []byte slices to Write.
// Results of individual Write calls are ignored.
func NonBlockingWriter(w io.Writer, cap int) io.WriteCloser {
	bb := make(chan []byte, cap)
	dn := make(chan bool)
	go func() {
		for b := range bb {
			w.Write(b)
		}
		dn <- true
	}()
	return &nonBlockingWriter{wr: w, bb: bb, dn: dn}
}
