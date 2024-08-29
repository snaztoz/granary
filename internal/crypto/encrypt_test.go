package crypto_test

import (
	"bytes"
	"crypto/rand"
	"log"
	"testing"

	"github.com/snaztoz/granary/internal/crypto"
)

func TestEncrypt(t *testing.T) {
	plaintext := []byte("this is a super secret content")

	key := make([]byte, 32)
	if _, err := rand.Read(key); err != nil {
		log.Fatalln(err)
	}

	ciphertext := crypto.Encrypt(plaintext, key)

	if res := crypto.Decrypt(ciphertext, key); !bytes.Equal(res, plaintext) {
		t.Fatalf("plaintext mismatch: expecting %v, get %v instead", plaintext, res)
	}
}
