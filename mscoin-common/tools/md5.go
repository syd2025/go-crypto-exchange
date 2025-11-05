package tools

import (
	"crypto/sha512"
	"encoding/hex"
	"hash"
	"math/rand"

	"golang.org/x/crypto/pbkdf2"
)

const (
	defaultSaltLen    = 64
	defaultIterations = 10000
	defaultKeyLen     = 128
)

var defaultHashFunction = sha512.New

type Options struct {
	SaltLen      int
	Iterations   int
	KeyLen       int
	HashFunction func() hash.Hash
}

func generateSalt(saltLen int) []byte {
	const alphanum = "0123456789ABDCEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	salt := make([]byte, saltLen)
	rand.Read(salt)
	for key, val := range salt {
		salt[key] = alphanum[val%byte(len(alphanum))]
	}
	return salt
}

func Encode(rawPwd string, options *Options) (string, string) {
	if options == nil {
		salt := generateSalt(defaultSaltLen)
		encodedPwd := pbkdf2.Key([]byte(rawPwd), salt, defaultIterations, defaultKeyLen, defaultHashFunction)
		return string(salt), hex.EncodeToString(encodedPwd)
	}
	saltLen := options.SaltLen
	if saltLen == 0 {
		saltLen = defaultSaltLen
	}
	iterations := options.Iterations
	if iterations == 0 {
		iterations = defaultIterations
	}
	keyLen := options.KeyLen
	if keyLen == 0 {
		keyLen = defaultKeyLen
	}
	hashFn := options.HashFunction
	if hashFn == nil {
		hashFn = defaultHashFunction
	}
	salt := generateSalt(saltLen)
	encodedPwd := pbkdf2.Key([]byte(rawPwd), salt, iterations, keyLen, hashFn)
	return string(salt), hex.EncodeToString(encodedPwd)
}

func Verify(rawPwd string, salt string, encodedPwd string, options *Options) bool {
	if options == nil {
		return encodedPwd == hex.EncodeToString(pbkdf2.Key([]byte(rawPwd), []byte(salt), defaultIterations, defaultKeyLen, defaultHashFunction))
	}
	return encodedPwd == hex.EncodeToString(pbkdf2.Key([]byte(rawPwd), []byte(salt), options.Iterations, options.KeyLen, options.HashFunction))
}
