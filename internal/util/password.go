package util

import (
	"fmt"
	"os"

	"golang.org/x/term"
)

func AskPassphrase(prompt string) (string, error) {
	fmt.Fprintf(os.Stderr, "%s: ", prompt)

	passphraseBytes, err := term.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		return "", err
	}

	// change line
	fmt.Fprint(os.Stderr, "\n")

	return string(passphraseBytes), nil
}
