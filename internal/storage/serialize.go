package storage

import (
	"encoding/base64"
	"errors"
	"fmt"
	"strings"
)

const (
	SecretFileHeader = "gran-secret-file"
)

var (
	ErrInvalidHeader = errors.New("invalid Granary secret file header")
)

func toFileContent(keyString string, ciphertext []byte) string {
	encoded := base64.StdEncoding.EncodeToString(ciphertext)
	return fmt.Sprintf("%s:%s:%s", SecretFileHeader, keyString, encoded)
}

func toKeyStringAndData(fileContent string) (keyString string, ciphertext []byte, err error) {
	splitted := strings.Split(fileContent, ":")
	if splitted[0] != SecretFileHeader {
		return "", nil, ErrInvalidHeader
	}

	keyString = splitted[1]
	ciphertext, err = base64.StdEncoding.DecodeString(splitted[2])
	if err != nil {
		return "", nil, err
	}

	return keyString, ciphertext, nil
}
