package util

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/term"
)

// Get passphrase from two possible sources:
//
//  1. Check whether a passphrase file is exist. If it exists, read
//     and return the file content (whitespace-trimmed).
//  2. If the file is not exist, then resort back to manual prompt.
//
// It won't check the correctness of the passphrase itself.
func GetPassphrase(passphrasePath string, prompt string) (string, error) {
	passphrase, err := readPassphraseFromFile(passphrasePath)

	if err == nil {
		// passphrase file exist, return the content directly
		return passphrase, nil
	}

	passphrase, err = PromptPassphrase(prompt)
	if err != nil {
		return "", err
	}

	return passphrase, nil
}

func readPassphraseFromFile(passphrasePath string) (string, error) {
	content, err := os.ReadFile(passphrasePath)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(content)), nil
}

func PromptPassphrase(prompt string) (string, error) {
	fmt.Fprintf(os.Stderr, "%s: ", prompt)

	passphraseBytes, err := term.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		return "", err
	}

	// change line
	fmt.Fprint(os.Stderr, "\n")

	return string(passphraseBytes), nil
}
