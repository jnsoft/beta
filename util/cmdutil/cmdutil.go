package cmdutil

import (
	"fmt"
	"syscall"

	"golang.org/x/term"
)

func ReadPassword(verification bool) (string, error) {
	for {
		fmt.Print("Enter password: ")
		bytePassword, err := term.ReadPassword(int(syscall.Stdin))
		if err != nil {
			return "", err
		}
		password := string(bytePassword)
		fmt.Println()

		if verification {
			fmt.Print("Confirm password: ")
			bytePasswordConfirm, err := term.ReadPassword(int(syscall.Stdin))
			if err != nil {
				return "", err
			}
			confirmPassword := string(bytePasswordConfirm)
			fmt.Println()
			if password == confirmPassword {
				return password, nil
			}
			fmt.Println("Passwords do not match. Please try again.")
		} else {
			return password, nil
		}
	}
}
