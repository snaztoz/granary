package crypto_test

import (
	"testing"

	"github.com/snaztoz/granary/internal/crypto"
)

func TestKeyDerivation(t *testing.T) {
	password := "foobarbat123"

	_, keyString := crypto.DeriveKey(password)

	if _, err := crypto.MatchPassword(password, keyString); err != nil {
		t.Fatal("expecting password to match")
	}
}
