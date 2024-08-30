package util

import (
	"fmt"
	"os"

	"golang.org/x/term"
)

func AskPassword(prompt string) (string, error) {
	fmt.Fprintf(os.Stderr, "%s: ", prompt)

	passwordBytes, err := term.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		return "", err
	}

	// change line
	fmt.Fprint(os.Stderr, "\n")

	return string(passwordBytes), nil
}
