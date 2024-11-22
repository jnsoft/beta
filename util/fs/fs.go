package fs

import (
	"fmt"
	"os"
)

func IsValidFile(filename string, verbose bool) bool {
	// Check if the string is empty
	if filename == "" {
		if verbose {
			fmt.Println("Filename is empty")
		}
		return false
	}

	// Check if the file exists
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		if verbose {
			fmt.Println("File does not exist")
		}
		return false
	} else if err != nil {
		// Handle other potential errors
		if verbose {
			fmt.Printf("Error checking file: %v\n", err)
		}
		return false
	}
	// File exists
	return true
}
