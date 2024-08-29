package storage

import (
	"encoding/json"
	"log"
	"os"

	"github.com/snaztoz/granary/internal/crypto"
	"github.com/snaztoz/granary/internal/data"
)

func NewFile(path, password string) {
	f, err := os.Create(path)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	key, keyString := crypto.DeriveKey(password)
	jsonData, _ := json.Marshal(make(data.T))
	ciphertext := crypto.Encrypt(jsonData, key)

	fileContent := toFileContent(keyString, ciphertext)
	if _, err := f.Write([]byte(fileContent)); err != nil {
		log.Fatalln(err)
	}
}

func ReadFile(path, password string) (data data.T, err error) {
	return make(map[string]string), nil
}

func WriteFile(path, password string, data data.T) error {
	return nil
}
