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

	// add blank lines
	fmt.Fprint(os.Stderr, "\n\n")

	return string(passwordBytes), nil
}
