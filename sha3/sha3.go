package sha3

// Drop-in replacement for golang.org/x/crypto/sha3
// Assumes openssl.Init() was called

import (
	"crypto"
	"github.com/golang-fips/openssl"
)

var (
	New224 = openssl.NewSHA3_224
	New256 = openssl.NewSHA3_256
	New384 = openssl.NewSHA3_384
	New512 = openssl.NewSHA3_512
	Sum224 = openssl.SHA3_224
	Sum256 = openssl.SHA3_256
	Sum384 = openssl.SHA3_384
	Sum512 = openssl.SHA3_512
)

func init() {
	crypto.RegisterHash(crypto.SHA3_224, New224)
	crypto.RegisterHash(crypto.SHA3_256, New256)
	crypto.RegisterHash(crypto.SHA3_384, New384)
	crypto.RegisterHash(crypto.SHA3_512, New512)
}
