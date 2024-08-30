package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
)

func Encrypt(plaintext []byte, key []byte) (ciphertext []byte, err error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	aesGcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, aesGcm.NonceSize())
	if _, err = rand.Read(nonce); err != nil {
		return nil, err
	}

	// Append the resulting ciphertext to nonce
	return aesGcm.Seal(nonce, nonce, plaintext, nil), nil
}

func Decrypt(ciphertext []byte, key []byte) (plaintext []byte, err error) {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return nil, err
	}

	aesGcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	// Since we know the ciphertext is actually nonce+ciphertext
	// And len(nonce) == NonceSize(). We can separate the two.
	nonceSize := aesGcm.NonceSize()
	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]

	plaintext, err = aesGcm.Open(nil, []byte(nonce), []byte(ciphertext), nil)
	if err != nil {
		return nil, err
	}

	return
}
