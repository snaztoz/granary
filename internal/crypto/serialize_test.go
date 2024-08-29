package crypto

import (
	"bytes"
	"testing"
)

func TestSerialization(t *testing.T) {
	// These are just a fake salt and hash bytes as we only
	// need them for testing the serialization results
	salt := randBytes(16)
	hash := randBytes(32)

	keyString := toKeyString(salt, hash)
	saltResult, hashResult, err := toSaltAndHash(keyString)
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(saltResult, salt) || !bytes.Equal(hashResult, hash) {
		t.Fatal("salt/hash results are not the same")
	}
}
