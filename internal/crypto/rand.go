package crypto

import (
	"crypto/rand"
	"log"
)

func randBytes(n int) []byte {
	b := make([]byte, n)
	if _, err := rand.Read(b); err != nil {
		log.Fatalln(err)
	}
	return b
}
