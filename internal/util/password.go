package util

import (
	"fmt"
	"os"

	"golang.org/x/term"
)

func AskPassword() (string, error) {
	fmt.Fprintf(os.Stderr, "Enter a new passkey: ")

	passwordBytes, err := term.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		return "", err
	}

	fmt.Println() // change line

	return string(passwordBytes), nil
}
