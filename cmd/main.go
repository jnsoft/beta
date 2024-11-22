package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/jnsoft/beta/util/httputil"
	"github.com/spf13/cobra"
)

var inputFile string
var outputFile string
var fileName string
var lines string
var proxyUrl string
var portNumber string

func main() {
	var betaCmd = &cobra.Command{
		Use:   "beta",
		Short: "The beta CLI tool.",
		Run: func(cmd *cobra.Command, args []string) {
		},
	}

	betaCmd.AddCommand(versionCmd)
	betaCmd.AddCommand(testConnectionCmd())
	betaCmd.AddCommand(b64Cmd())
	betaCmd.AddCommand(hexCmd())
	betaCmd.AddCommand(uuidCmd())
	betaCmd.AddCommand(httpCmd())
	betaCmd.AddCommand(hashCmd())
	betaCmd.AddCommand(hmacCmd())

	err := betaCmd.Execute()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func testConnectionCmd() *cobra.Command {
	var testConnectionCmd = &cobra.Command{
		Use:   "connect",
		Short: "Test network connection.",
		Args:  cobra.ExactArgs(1),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return IncorrectUsageErr()
		},
		Run: func(cmd *cobra.Command, args []string) {
			address := args[0]

			port, err := strconv.Atoi(portNumber)
			if err != nil {
				fmt.Printf("Error converting repitions to int: %v\n", err)
				os.Exit(1)
			}
			httputil.TestConnection(address, port, 3*time.Second, proxyUrl)
		},
	}

	addProxyFlag(testConnectionCmd)
	testConnectionCmd.Flags().StringVarP(&repitions, "port", "p", "80", "Port number to use")

	return testConnectionCmd
}

func insertLineBreakFlag(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&lines, "lines", "l", "0", "Number of characters per line (-l 76)")
}

func addProxyFlag(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&proxyUrl, "proxy", "p", "", "proxy url")
}

func addFileFlag(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&fileName, "filename", "f", "", "Path to file")
}

func addDefaultFileFlags(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&inputFile, "input", "i", "", "Input file to encode")
	cmd.Flags().StringVarP(&outputFile, "output", "o", "", "Output file to write decoded data")
}

func IncorrectUsageErr() error {
	return fmt.Errorf("incorrect usage")
}
