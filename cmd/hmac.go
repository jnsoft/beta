package main

import (
	"os"

	"github.com/jnsoft/beta/util/fs"
	"github.com/spf13/cobra"
)

func hmacCmd() *cobra.Command {
	var hmacCmd = &cobra.Command{
		Use:   "hmac",
		Short: "HMAC util.",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return IncorrectUsageErr()
		},
		Run: func(cmd *cobra.Command, args []string) {
		},
	}

	hmacCmd.AddCommand(sha256Cmd())
	hmacCmd.AddCommand(sha512Cmd())
	hmacCmd.AddCommand(sha3Cmd())

	return hmacCmd
}

func genericHMACCmd(hashFunc func(data []byte) string, cmd_String string) *cobra.Command {
	var cmd = &cobra.Command{
		Use:   cmd_String + " [string]",
		Short: cmd_String + "  hash a string or file.",
		Run: func(cmd *cobra.Command, args []string) {
			var hash string
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
				hash = hashFunc(data)
			} else {
				hash = hashFunc([]byte(args[0]))
			}
			cmd.Println(hash)
		},
	}
	addFileFlag(cmd)
	return cmd
}
