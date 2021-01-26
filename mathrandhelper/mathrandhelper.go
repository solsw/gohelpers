// Package mathrandhelper contains math/rand helpers.
package mathrandhelper

import (
	"bufio"
	cryptorand "crypto/rand"
	"encoding/binary"
	"io"
	mathrand "math/rand"
)

// NewCryptoRand returns math/rand.Rand instance backed by crypto/rand.
func NewCryptoRand() *mathrand.Rand {
	s := cryptoRandSource{bufio.NewReader(cryptorand.Reader)}
	return mathrand.New(s)
}

// inspired by github.com/andrew-d/csmrand
type cryptoRandSource struct {
	byteReader io.ByteReader
}

func (r cryptoRandSource) Int63() int64 {
	u64, _ := binary.ReadUvarint(r.byteReader)
	return int64(u64 >> 1)
}

func (r cryptoRandSource) Seed(int64) {}
