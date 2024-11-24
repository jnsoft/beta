package main

import (
	"os"

	"github.com/jnsoft/beta/util/fs"
	"github.com/jnsoft/beta/util/security"
	"github.com/spf13/cobra"
)

const MISSING_INPUT = "Please provide a string or file to hash"

func hashCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "hash",
		Short: "Hash util.",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return IncorrectUsageErr()
		},
		Run: func(cmd *cobra.Command, args []string) {
		},
	}

	cmd.AddCommand(md5Cmd())
	cmd.AddCommand(sha1Cmd())
	cmd.AddCommand(sha256Cmd())
	cmd.AddCommand(sha512Cmd())
	cmd.AddCommand(sha3Cmd())

	return cmd
}

func md5Cmd() *cobra.Command {
	return genericHashCmd(security.HashMD5, "md5")
}

func sha1Cmd() *cobra.Command {
	return genericHashCmd(security.HashSHA1, "sha1")
}

func sha256Cmd() *cobra.Command {
	return genericHashCmd(security.HashSHA256, "sha256")
}

func sha512Cmd() *cobra.Command {
	return genericHashCmd(security.HashSHA512, "sha512")
}

func sha3Cmd() *cobra.Command {
	return genericHashCmd(security.HashSHA3, "sha3")
}

func genericHashCmd(hashFunc func(data []byte) string, cmd_String string) *cobra.Command {
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
