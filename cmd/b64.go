package main

import (
	"os"
	"strconv"

	"github.com/jnsoft/beta/util/stringutil"
	"github.com/spf13/cobra"
)

func b64Cmd() *cobra.Command {
	var b64Cmd = &cobra.Command{
		Use:   "b64",
		Short: "Base64 util.",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return IncorrectUsageErr()
		},
		Run: func(cmd *cobra.Command, args []string) {
		},
	}

	b64Cmd.AddCommand(toB64Cmd())
	b64Cmd.AddCommand(fromB64Cmd())
	b64Cmd.AddCommand(fileToB64Cmd())
	b64Cmd.AddCommand(fileFromB64())

	return b64Cmd
}

func toB64Cmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "to [string]",
		Short: "Encode a string to base64.",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {

			if len(args) == 0 {
				cmd.Println("Please provide a string to encode")
				os.Exit(1)
			}

			n, err := strconv.Atoi(lines)
			if err != nil {
				cmd.Printf("Error converting lines to int: %v\n", err)
				os.Exit(1)
			}
			encoded := stringutil.ToBase64([]byte(args[0]), n)
			cmd.Println(encoded)
		},
	}
	insertLineBreakFlag(cmd)
	return cmd
}

func fromB64Cmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "from [b64string]",
		Short: "Decode a base64 string.",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			// Check if there is a string to decode
			if len(args) == 0 {
				cmd.Println("Please provide a string to decode")
				os.Exit(1)
			}

			decoded, err := stringutil.FromBase64(args[0])
			if err != nil {
				cmd.Println("Error decoding string:", err)
				os.Exit(1)
			}
			cmd.Println(string(decoded))
		},
	}
	return cmd
}

func fileToB64Cmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "encode",
		Short: "Encode a file to base64.",
		Run: func(cmd *cobra.Command, args []string) {
			n, err := strconv.Atoi(lines)
			if err != nil {
				cmd.Printf("Error converting lines to int: %v\n", err)
				os.Exit(1)
			}

			if inputFile == "" {
				cmd.Println("Please provide a file to encode")
				os.Exit(1)
			}

			data, err := os.ReadFile(inputFile)
			if err != nil {
				cmd.Println("Error reading file:", err)
				os.Exit(1)
			}

			encoded := stringutil.ToBase64(data, n)
			if outputFile != "" {
				err = os.WriteFile(outputFile, []byte(encoded), 0644)
				if err != nil {
					cmd.Println("Error writing file:", err)
					os.Exit(1)
				}
			} else {
				cmd.Println(encoded)
			}
		},
	}

	insertLineBreakFlag(cmd)
	addDefaultFileFlags(cmd)
	return cmd
}

func fileFromB64() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "decode",
		Short: "Decode a base64 file.",
		Run: func(cmd *cobra.Command, args []string) {

			if inputFile == "" {
				cmd.Println("Please provide a file to decode")
				os.Exit(1)
			}

			data, err := os.ReadFile(inputFile)
			if err != nil {
				cmd.Println("Error reading file:", err)
				os.Exit(1)
			}

			decoded, err := stringutil.FromBase64(string(data))
			if err != nil {
				cmd.Println("Error decoding file:", err)
				os.Exit(1)
			}
			if outputFile != "" {
				err = os.WriteFile(outputFile, []byte(decoded), 0644)
				if err != nil {
					cmd.Println("Error writing file:", err)
					os.Exit(1)
				}

			} else {
				cmd.Println(string(decoded))
			}

		},
	}

	addDefaultFileFlags(cmd)
	return cmd
}
