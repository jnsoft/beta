package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/jnsoft/beta/util/security"
	"github.com/jnsoft/beta/util/stringutil"
	"github.com/spf13/cobra"
)

var outformat string

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

	cmd.AddCommand(newKeyCmd())

	return cmd
}

func newCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "new",
		Short: "Generate new secure key.",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return IncorrectUsageErr()
		},
		Run: func(cmd *cobra.Command, args []string) {
		},
	}

	cmd.AddCommand(newHexCmd())
	cmd.AddCommand(newB64Cmd())

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
				cmd.Println(stringKey)
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
				cmd.Println(stringKey)
			}
		},
	}

	addNFlag(cmd)
	cmd.Flags().StringVarP(&outputFile, "output", "o", "", "Output file")

	return cmd
}
