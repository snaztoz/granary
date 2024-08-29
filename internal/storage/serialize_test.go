package storage

import (
	"reflect"
	"testing"

	"github.com/snaztoz/granary/internal/crypto"
	"github.com/snaztoz/granary/internal/data"
)

func TestSerialization(t *testing.T) {
	password := "my-password"
	data := make(data.T)
	key, keyString := crypto.DeriveKey(password)

	fileContent := toFileContent(keyString, data, key)
	keyStringResult, dataResult, err := toKeyStringAndData(fileContent, key)
	if err != nil {
		t.Fatal(err)
	}

	if keyStringResult != keyString || !reflect.DeepEqual(dataResult, data) {
		t.Fatal("keystring/data are not the same")
	}
}
