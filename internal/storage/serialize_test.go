package storage

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/snaztoz/granary/internal/crypto"
	"github.com/snaztoz/granary/internal/data"
)

func TestSerialization(t *testing.T) {
	passphrase := "my-passphrase"
	key, keyString := crypto.DeriveKey(passphrase)

	data := make(data.T)
	jsonData, _ := json.Marshal(data)

	ciphertext, err := crypto.Encrypt(jsonData, key)
	if err != nil {
		t.Fatal(err)
	}

	fileContent := toFileContent(keyString, ciphertext)
	keyStringResult, ciphertextResult, err := toKeyStringAndData(fileContent)
	if err != nil {
		t.Fatal(err)
	}

	if keyStringResult != keyString || !reflect.DeepEqual(ciphertextResult, ciphertext) {
		t.Fatal("keystring/data are not the same")
	}
}
