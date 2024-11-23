package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/jnsoft/beta/util/httputil"
	"github.com/jnsoft/beta/util/security"
	"github.com/spf13/cobra"
)

var inputFile string  // i
var outputFile string // o
var fileName string   // f
var lines string      // l
var proxyUrl string   // x
var portNumber string // p
var key string        // k
var hexString string  // h
var n string          // n

func main() {
	var betaCmd = &cobra.Command{
		Use:   "beta",
		Short: "The beta CLI tool.",
		Run: func(cmd *cobra.Command, args []string) {
		},
	}

	betaCmd.AddCommand(versionCmd)
	betaCmd.AddCommand(testConnectionCmd())
	betaCmd.AddCommand(passCmd())
	betaCmd.AddCommand(b64Cmd())
	betaCmd.AddCommand(hexCmd())
	betaCmd.AddCommand(uuidCmd())
	betaCmd.AddCommand(httpCmd())
	betaCmd.AddCommand(hashCmd())
	betaCmd.AddCommand(hmacCmd())
	betaCmd.AddCommand(keyCmd())

	err := betaCmd.Execute()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func testConnectionCmd() *cobra.Command {
	var testConnectionCmd = &cobra.Command{
		Use:   "connect [address]",
		Short: "Test network connection.",
		Args:  cobra.ExactArgs(1),
		//PreRunE: func(cmd *cobra.Command, args []string) error {
		//	return IncorrectUsageErr()
		//},
		Run: func(cmd *cobra.Command, args []string) {
			address := args[0]

			port, err := strconv.Atoi(portNumber)
			if err != nil {
				fmt.Printf("Error converting repitions to int: %v\n", err)
				os.Exit(1)
			}
			res, duration := httputil.TestConnection(address, port, 3*time.Second, proxyUrl)
			strRes := fmt.Sprintf("Connected: %t, Time: %d", res, duration)
			cmd.Println(strRes)
		},
	}

	addProxyFlag(testConnectionCmd)
	testConnectionCmd.Flags().StringVarP(&portNumber, "port", "p", "80", "Port number to use")

	return testConnectionCmd
}

func passCmd() *cobra.Command {
	var passCmd = &cobra.Command{
		Use:   "pass [length]",
		Short: "Generate random password.",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			len, err := strconv.Atoi(args[0])
			if err != nil {
				fmt.Printf("Error converting length to int: %v\n", err)
				os.Exit(1)
			}
			pw, _ := security.GeneratePassword(len, true)
			cmd.Println(pw)
		},
	}
	return passCmd
}

func insertLineBreakFlag(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&lines, "lines", "l", "0", "Number of characters per line (-l 76)")
}

func addNFlag(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&n, "n", "n", "32", "Length of key in bytes")
}

func addProxyFlag(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&proxyUrl, "proxy", "x", "", "proxy url")
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
