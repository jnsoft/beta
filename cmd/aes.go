package main

import (
	"fmt"
	"os"

	"github.com/jnsoft/beta/util/aesutil"
	"github.com/jnsoft/beta/util/stringutil"
	"github.com/spf13/cobra"
)

func aesCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "aes",
		Short: "AES util.",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return IncorrectUsageErr()
		},
		Run: func(cmd *cobra.Command, args []string) {
		},
	}

	cmd.AddCommand(encryptCmd())
	cmd.AddCommand(decryptCmd())

	return cmd
}

func encryptCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "encrypt [string]",
		Short: "AES GCM encrypt a string or file.",
		Run: func(cmd *cobra.Command, args []string) {

			var plain []byte

			if key == "" {
				cmd.Println("Please provide a key")
				os.Exit(1)
			}

			key, err := stringutil.FromHex(key)
			if err != nil {
				cmd.Println("Error reading key:", err)
				os.Exit(1)
			}

			stringToString := false
			stringToFile := false

			if len(args) > 0 && args[0] != "" {
				if outputFile == "" {
					stringToString = true
					cmd.Println("stringToSting")
				} else {
					stringToFile = true
					cmd.Println("stringToFile")
				}
			} else {
				if inputFile == "" {
					cmd.Println("Please provide a string or file to encrypt")
					os.Exit(1)
				} else { // fileToFile
					cmd.Println("fileToFile")
					if outputFile == "" {
						cmd.Println("Please provide a target file for encrypted data")
						os.Exit(1)
					}
				}
			}

			if stringToString || stringToFile {
				plain = []byte(args[0])
			} else { // fileToFile
				plain, err = os.ReadFile(inputFile)
				if err != nil {
					cmd.Println("Error reading file:", err)
					os.Exit(1)
				}
			}

			cipher, err := aesutil.GcmEncrypt(plain, key)
			if err != nil {
				cmd.Println("Error during encryption:", err)
				os.Exit(1)
			}

			if stringToString {
				cmd.Println("here")
				fmt.Println(stringutil.ToBase64(cipher, 76))
			} else {
				err = os.WriteFile(outputFile, cipher, 0644)
				if err != nil {
					cmd.Println("Error writing file:", err)
					os.Exit(1)
				}
			}
		},
	}
	addDefaultFileFlags(cmd)
	insertKeyFlag(cmd)
	return cmd
}

func decryptCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "decrypt [b64 string]",
		Short: "AES GCM decrypt a string or file.",
		Run: func(cmd *cobra.Command, args []string) {

			var cipher []byte

			if key == "" {
				cmd.Println("Please provide a key")
				os.Exit(1)
			}

			key, err := stringutil.FromHex(key)
			if err != nil {
				cmd.Println("Error reading key:", err)
				os.Exit(1)
			}

			stringToString := false
			stringToFile := false

			if len(args) > 0 && args[0] != "" {
				if outputFile == "" {
					stringToString = true
				} else {
					stringToFile = true
				}
			} else {
				if inputFile == "" {
					cmd.Println("Please provide a string or file to decrypt")
					os.Exit(1)
				} else { // fileToFile
					if outputFile == "" {
						cmd.Println("Please provide a target file for decrypted data")
						os.Exit(1)
					}
				}
			}

			if stringToString || stringToFile {
				cipher, err = stringutil.FromBase64(args[0])
				if err != nil {
					cmd.Println("Error reading b64 string:", err)
					os.Exit(1)
				}
			} else { // file input
				cipher, err = os.ReadFile(inputFile)
				if err != nil {
					cmd.Println("Error reading file:", err)
					os.Exit(1)
				}
			}

			plain, err := aesutil.GcmDecrypt(cipher, key)
			if err != nil {
				cmd.Println("Error during decryption:", err)
				os.Exit(1)
			}

			if stringToString {
				fmt.Println(string(plain))
			} else {
				err = os.WriteFile(outputFile, plain, 0644)
				if err != nil {
					cmd.Println("Error writing file:", err)
					os.Exit(1)
				}
			}
		},
	}
	addDefaultFileFlags(cmd)
	insertKeyFlag(cmd)
	return cmd
}
