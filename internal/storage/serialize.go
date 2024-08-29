package storage

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/snaztoz/granary/internal/crypto"
	"github.com/snaztoz/granary/internal/data"
)

const (
	secretFileHeader = "gran-secret-file"
)

var (
	ErrInvalidHeader = errors.New("invalid Granary secret file header")
)

func toFileContent(keyString string, data data.T, key []byte) string {
	dataJson, _ := json.Marshal(data)
	encrypted := crypto.Encrypt(dataJson, key)
	encoded := base64.StdEncoding.EncodeToString(encrypted)
	return fmt.Sprintf("%s:%s:%s", secretFileHeader, keyString, encoded)
}

func toKeyStringAndData(fileContent string, key []byte) (keyString string, data data.T, err error) {
	data = make(map[string]string)
	splitted := strings.Split(fileContent, ":")

	if splitted[0] != secretFileHeader {
		return "", data, ErrInvalidHeader
	}

	keyString, encoded := splitted[1], splitted[2]

	encrypted, _ := base64.StdEncoding.DecodeString(encoded)
	dataJson := crypto.Decrypt(encrypted, key)
	json.Unmarshal(dataJson, &data)

	return keyString, data, nil
}
