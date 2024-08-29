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
	fileContent, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	keyString, ciphertext, err := toKeyStringAndData(string(fileContent))
	if err != nil {
		return nil, err
	}

	key, err := crypto.MatchPassword(password, keyString)
	if err != nil {
		return nil, err
	}

	plaintext := crypto.Decrypt(ciphertext, key)
	err = json.Unmarshal(plaintext, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func WriteFile(path, password string, data data.T) error {
	return nil
}
