package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/jnsoft/beta/util/cmdutil"
	"github.com/jnsoft/beta/util/security"
	"github.com/jnsoft/beta/util/stringutil"
	"github.com/spf13/cobra"
)

func keyCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "key",
		Short: "Generate secure key.",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return IncorrectUsageErr()
		},
		Run: func(cmd *cobra.Command, args []string) {
		},
	}

	cmd.AddCommand(newHexCmd())
	cmd.AddCommand(newB64Cmd())
	cmd.AddCommand(deriveCmd())

	return cmd
}

func newHexCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "hex",
		Short: "Generate new secure key of length n bytes in hex format.",
		Run: func(cmd *cobra.Command, args []string) {

			n, err := strconv.Atoi(n)
			if err != nil {
				cmd.Printf("Error converting n to int: %v\n", err)
				os.Exit(1)
			}

			key, err := security.RandomBytes(n)
			if err != nil {
				fmt.Println("Error generating key: ", err)
				os.Exit(1)
			}

			stringKey := stringutil.ToHex(key, 0)

			if outputFile != "" {
				err = os.WriteFile(outputFile, []byte(stringKey), 0644)
				if err != nil {
					cmd.Println("Error writing file:", err)
					os.Exit(1)
				}
			} else {
				fmt.Println(stringKey)
			}
		},
	}

	addNFlag(cmd)
	cmd.Flags().StringVarP(&outputFile, "output", "o", "", "Output file")
	return cmd
}

func newB64Cmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "b64",
		Short: "Generate new secure key of length n bytes in b64 format.",
		Run: func(cmd *cobra.Command, args []string) {

			n, err := strconv.Atoi(n)
			if err != nil {
				cmd.Printf("Error converting n to int: %v\n", err)
				os.Exit(1)
			}

			key, err := security.RandomBytes(n)
			if err != nil {
				fmt.Println("Error generating key: ", err)
				os.Exit(1)
			}

			stringKey := stringutil.ToBase64(key, 0)

			if outputFile != "" {
				err = os.WriteFile(outputFile, []byte(stringKey), 0644)
				if err != nil {
					cmd.Println("Error writing file:", err)
					os.Exit(1)
				}
			} else {
				fmt.Println(stringKey)
			}
		},
	}

	addNFlag(cmd)
	cmd.Flags().StringVarP(&outputFile, "output", "o", "", "Output file")
	return cmd
}

func deriveCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "derive",
		Short: "Derive a secure key of length n bytes in hex format from given password.",
		Run: func(cmd *cobra.Command, args []string) {

			n, err := strconv.Atoi(n)
			if err != nil {
				cmd.Printf("Error converting n to int: %v\n", err)
				os.Exit(1)
			}

			if n < 1 {
				fmt.Println("Key length to short")
				os.Exit(1)
			}

			pass, err := cmdutil.ReadPassword(false)
			if err != nil {
				fmt.Println("Error reading password: ", err)
				os.Exit(1)
			}

			var key []byte

			if salt == "" { // no salt
				key = security.DeriveKeyWithoutSalt(pass, n, security.SCRYPT_N)
			} else if len(salt) != 16 || !stringutil.IsHexString(salt) { // invalid salt
				fmt.Println("Salt must be 16 bytes, given as a hex string")
				os.Exit(1)
			} else { // with salt
				salt, err := stringutil.FromHex(salt)
				if err != nil {
					fmt.Println("Error reading hex: ", err)
					os.Exit(1)
				}
				key = security.DeriveKey(pass, salt, n, security.SCRYPT_N)
			}

			stringKey := stringutil.ToHex(key, 0)

			if outputFile != "" {
				err = os.WriteFile(outputFile, []byte(stringKey), 0644)
				if err != nil {
					cmd.Println("Error writing file:", err)
					os.Exit(1)
				}
			} else {
				fmt.Println(stringKey)
			}
		},
	}

	addNFlag(cmd)
	cmd.Flags().StringVarP(&salt, "salt", "s", "", "Provided salt (16 byte hex string)")
	cmd.Flags().StringVarP(&outputFile, "output", "o", "", "Output file")
	return cmd
}
