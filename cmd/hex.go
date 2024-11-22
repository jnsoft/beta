package main

import (
	"encoding/hex"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/jnsoft/beta/util/stringutil"
	"github.com/spf13/cobra"
)

func hexCmd() *cobra.Command {
	var hexCmd = &cobra.Command{
		Use:   "hex",
		Short: "Hex util.",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return IncorrectUsageErr()
		},
		Run: func(cmd *cobra.Command, args []string) {
		},
	}

	hexCmd.AddCommand(toHexCmd())
	hexCmd.AddCommand(fromHexCmd())
	hexCmd.AddCommand(fileToHexCmd())
	hexCmd.AddCommand(fileFromHex())

	return hexCmd
}

func toHexCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "to [string]",
		Short: "Encode a string to hex.",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			// Check if there is a string to encode
			if len(args) == 0 {
				fmt.Println("Please provide a string to encode")
				os.Exit(1)
			}

			n, err := strconv.Atoi(lines)
			if err != nil {
				fmt.Printf("Error converting lines to int: %v\n", err)
				os.Exit(1)
			}

			encoded := stringutil.ToHex([]byte(args[0]), n)
			upperEncoded := strings.ToUpper(encoded)
			fmt.Println(upperEncoded)
		},
	}
	insertLineBreakFlag(cmd)
	return cmd
}

func fromHexCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "from [hexstring]",
		Short: "Decode a hex string.",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			// Check if there is a string to decode
			if len(args) == 0 {
				fmt.Println("Please provide a string to decode")
				os.Exit(1)
			}

			decoded, err := stringutil.FromHex(args[0])
			if err != nil {
				fmt.Println("Error decoding string:", err)
				os.Exit(1)
			}
			fmt.Println(string(decoded))
		},
	}
	return cmd
}

func fileToHexCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "encode",
		Short: "Encode a file to hex.",
		Run: func(cmd *cobra.Command, args []string) {
			n, err := strconv.Atoi(lines)
			if err != nil {
				fmt.Printf("Error converting lines to int: %v\n", err)
				os.Exit(1)
			}

			if inputFile == "" {
				fmt.Println("Please provide a file to encode")
				os.Exit(1)
			}

			data, err := os.ReadFile(inputFile)
			if err != nil {
				fmt.Println("Error reading file:", err)
				os.Exit(1)
			}

			encoded := stringutil.ToHex(data, n)
			if outputFile != "" {
				err = os.WriteFile(outputFile, []byte(encoded), 0644)
				if err != nil {
					fmt.Println("Error writing file:", err)
					os.Exit(1)
				}
			} else {
				fmt.Println(encoded)
			}
		},
	}

	insertLineBreakFlag(cmd)
	addDefaultFileFlags(cmd)
	return cmd
}

func fileFromHex() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "decode",
		Short: "Decode a hex encoded file.",
		Run: func(cmd *cobra.Command, args []string) {
			// Check if there is a file to decode
			if inputFile == "" {
				fmt.Println("Please provide a file to decode")
				os.Exit(1)
			}

			data, err := os.ReadFile(inputFile)
			if err != nil {
				fmt.Println("Error reading file:", err)
				os.Exit(1)
			}

			decoded, err := hex.DecodeString(string(data))
			if err != nil {
				fmt.Println("Error decoding file:", err)
				os.Exit(1)
			}
			if outputFile != "" {
				err = os.WriteFile(outputFile, []byte(decoded), 0644)
				if err != nil {
					fmt.Println("Error writing file:", err)
					os.Exit(1)
				}

			} else {
				fmt.Println(string(decoded))
			}

		},
	}

	cmd.Flags().StringVarP(&inputFile, "input", "i", "", "Input file to decode")
	cmd.Flags().StringVarP(&outputFile, "output", "o", "", "Output file to write decoded data")

	return cmd
}
