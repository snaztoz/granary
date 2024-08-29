package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"log"
)

func Encrypt(plaintext []byte, key []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		log.Fatalln(err)
	}

	aesGcm, err := cipher.NewGCM(block)
	if err != nil {
		log.Fatalln(err)
	}

	nonce := make([]byte, aesGcm.NonceSize())
	if _, err = rand.Read(nonce); err != nil {
		log.Fatalln(err)
	}

	// Append the resulting ciphertext to nonce
	return aesGcm.Seal(nonce, nonce, plaintext, nil)
}

func Decrypt(ciphertext []byte, key []byte) []byte {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		panic(err)
	}

	aesGcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err)
	}

	// Since we know the ciphertext is actually nonce+ciphertext
	// And len(nonce) == NonceSize(). We can separate the two.
	nonceSize := aesGcm.NonceSize()
	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]

	plaintext, err := aesGcm.Open(nil, []byte(nonce), []byte(ciphertext), nil)
	if err != nil {
		panic(err)
	}

	return plaintext
}
