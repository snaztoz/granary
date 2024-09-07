package crypto_test

import (
	"testing"

	"github.com/snaztoz/granary/internal/crypto"
)

func TestKeyDerivation(t *testing.T) {
	passphrase := "foobarbat123"

	_, keyString := crypto.DeriveKey(passphrase)

	if _, err := crypto.MatchPassphrase(passphrase, keyString); err != nil {
		t.Fatal("expecting passphrase to match")
	}
}
