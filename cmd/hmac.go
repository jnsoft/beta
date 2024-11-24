package main

import (
	"fmt"
	"os"

	"github.com/jnsoft/beta/util/fs"
	"github.com/jnsoft/beta/util/security"
	"github.com/jnsoft/beta/util/stringutil"
	"github.com/spf13/cobra"
)

func hmacCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "hmac",
		Short: "HMAC util.",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return IncorrectUsageErr()
		},
		Run: func(cmd *cobra.Command, args []string) {
		},
	}

	cmd.AddCommand(sha256HmacCmd())
	cmd.AddCommand(sha512HmacCmd())
	cmd.AddCommand(sha3HmacCmd())
	cmd.AddCommand(verifyCmd())

	return cmd
}

func sha256HmacCmd() *cobra.Command {
	return genericHMACCmd(security.HmacSHA256_hex, "sha256")
}

func sha512HmacCmd() *cobra.Command {
	return genericHMACCmd(security.HmacSHA512_hex, "sha512")
}

func sha3HmacCmd() *cobra.Command {
	return genericHMACCmd(security.HmacSHA3_hex, "sha3")
}

func verifyCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "verify",
		Short: "Verify HMAC.",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return IncorrectUsageErr()
		},
		Run: func(cmd *cobra.Command, args []string) {
		},
	}

	cmd.AddCommand(verifySha256HmacCmd())
	cmd.AddCommand(verifySha512HmacCmd())
	cmd.AddCommand(verifySha3HmacCmd())

	return cmd
}

func verifySha256HmacCmd() *cobra.Command {
	return genericHVerifyMACCmd(security.HmacSHA256_verify_hex, "sha256")
}

func verifySha512HmacCmd() *cobra.Command {
	return genericHVerifyMACCmd(security.HmacSHA512_verify_hex, "sha512")
}

func verifySha3HmacCmd() *cobra.Command {
	return genericHVerifyMACCmd(security.HmacSHA3_verify_hex, "sha3")
}

func genericHMACCmd(hmacFunc func(data, key []byte) string, cmd_String string) *cobra.Command {
	var cmd = &cobra.Command{
		Use:   cmd_String + " [string]",
		Short: "Compute " + cmd_String + "  HMAC for string or file.",
		Run: func(cmd *cobra.Command, args []string) {
			if key == "" {
				cmd.Println("Please provide a key")
				os.Exit(1)
			}
			byte_key, err := stringutil.FromHex(key)
			if err != nil {
				cmd.Println("Error reading key:", err)
				os.Exit(1)
			}

			var hmac_hex string
			if len(args) == 0 {
				if !fs.IsValidFile(fileName, true) {
					cmd.Println("Missing input")
					os.Exit(1)
				}
				data, err := os.ReadFile(fileName)
				if err != nil {
					cmd.Println("Error reading file:", err)
					os.Exit(1)
				}

				hmac_hex = hmacFunc(data, byte_key)
			} else {
				hmac_hex = hmacFunc([]byte(args[0]), byte_key)
			}
			//cmd.Println(hmac_hex)
			fmt.Println(hmac_hex)
		},
	}
	addFileFlag(cmd)
	cmd.Flags().StringVarP(&key, "key", "k", "", "Key to use in hex")
	return cmd
}

func genericHVerifyMACCmd(hmacVerifyFunc func(data, key []byte, hex string) bool, cmd_String string) *cobra.Command {
	var cmd = &cobra.Command{
		Use:   cmd_String + " [string]",
		Short: "Verify " + cmd_String + " HMAC for string or file.",
		Run: func(cmd *cobra.Command, args []string) {

			if key == "" {
				cmd.Println("Please provide a key")
				os.Exit(1)
			}

			if hexString == "" {
				cmd.Println("Please provide a HMAC to verify")
				os.Exit(1)
			}

			byte_key, err := stringutil.FromHex(key)
			if err != nil {
				cmd.Println("Error reading key:", err)
				os.Exit(1)
			}

			var res bool

			if len(args) == 0 {
				if !fs.IsValidFile(fileName, true) {
					cmd.Println("Missing input")
					os.Exit(1)
				}
				data, err := os.ReadFile(fileName)
				if err != nil {
					cmd.Println("Error reading file:", err)
					os.Exit(1)
				}

				res = hmacVerifyFunc(data, byte_key, hexString)
			} else {
				res = hmacVerifyFunc([]byte(args[0]), byte_key, hexString)
			}
			cmd.Println(res)
		},
	}
	addFileFlag(cmd)
	cmd.Flags().StringVarP(&key, "key", "k", "", "Key to use in hex")
	cmd.Flags().StringVarP(&hexString, "hmac", "", "", "HMAC to verify")
	return cmd
}
