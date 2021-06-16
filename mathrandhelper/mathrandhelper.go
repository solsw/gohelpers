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

// cryptoRandSource implements the crypto/rand.Source interface.
type cryptoRandSource struct {
	// inspired by github.com/andrew-d/csmrand
	byteReader io.ByteReader
}

// Int63 implements the crypto/rand.Source.Int63 method.
func (s cryptoRandSource) Int63() int64 {
	u64, _ := binary.ReadUvarint(s.byteReader)
	return int64(u64 >> 1)
}

// Seed implements the crypto/rand.Source.Seed method.
func (s cryptoRandSource) Seed(int64) {}
