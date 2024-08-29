package crypto

import (
	"encoding/hex"
	"errors"
	"strings"
)

var (
	ErrInvalidKeyString = errors.New("invalid key string")
)

func toKeyString(salt, hash []byte) string {
	return hex.EncodeToString(salt) + "$" + hex.EncodeToString(hash)
}

func toSaltAndHash(keyString string) (salt []byte, hash []byte, err error) {
	splitted := strings.Split(keyString, "$")
	if len(splitted) != 2 {
		return nil, nil, ErrInvalidKeyString
	}

	salt, _ = hex.DecodeString(splitted[0])
	hash, _ = hex.DecodeString(splitted[1])

	return
}
