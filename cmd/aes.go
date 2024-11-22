package main

import "github.com/spf13/cobra"

func aesCmd() *cobra.Command {
	var aesCmd = &cobra.Command{
		Use:   "aes",
		Short: "AES util.",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return IncorrectUsageErr()
		},
		Run: func(cmd *cobra.Command, args []string) {
		},
	}

	//aesCmd.AddCommand(toB64Cmd())

	return aesCmd
}
