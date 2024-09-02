package crypto

import (
	"crypto/sha256"
	"errors"

	"golang.org/x/crypto/argon2"
	"golang.org/x/crypto/pbkdf2"
)

var (
	ErrPasswordMismatch = errors.New("incorrect password")
)

func DeriveKey(password string) (key []byte, keyString string) {
	salt := randBytes(16)
	return deriveKeyWithSalt(password, salt)
}

func MatchPassword(password, keyString string) (key []byte, err error) {
	salt, _, err := toSaltAndHash(keyString)
	if err != nil {
		return nil, err
	}

	key, calculatedKeyString := deriveKeyWithSalt(password, salt)
	if calculatedKeyString != keyString {
		return nil, ErrPasswordMismatch
	}

	return key, nil
}

func deriveKeyWithSalt(password string, salt []byte) (key []byte, keyString string) {
	key = pbkdf2.Key([]byte(password), salt, 4096, 32, sha256.New)
	hash := argon2.IDKey(key, salt, 1, 64*1024, 4, 32)

	keyString = toKeyString(salt, hash[:])
	return key, keyString
}
