package encodinghelper

import (
	"crypto/rand"
	"crypto/rsa"
	"testing"
)

var (
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
)

func init() {
	privateKey, _ = rsa.GenerateKey(rand.Reader, 1024)
	publicKey = &privateKey.PublicKey
}

func TestRsaPublicKeyMarshalBinary(t *testing.T) {
	tests := []struct {
		name string
		got  *rsa.PublicKey
	}{
		{
			name: "1",
			got: func() *rsa.PublicKey {
				bb := RsaPublicKeyMarshalBinary(publicKey)
				pub, _ := RsaPublicKeyUnmarshalBinary(bb)
				return pub
			}(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !tt.got.Equal(publicKey) {
				t.Errorf("RsaPublicKeyMarshalBinary() error")
			}
		})
	}
}

func TestRsaPrivateKeyMarshalBinary(t *testing.T) {
	tests := []struct {
		name string
		got  *rsa.PrivateKey
	}{
		{
			name: "1",
			got: func() *rsa.PrivateKey {
				bb := RsaPrivateKeyMarshalBinary(privateKey)
				priv, _ := RsaPrivateKeyUnmarshalBinary(bb)
				return priv
			}(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !tt.got.Equal(privateKey) {
				t.Errorf("RsaPrivateKeyMarshalBinary() error")
			}
		})
	}
}
