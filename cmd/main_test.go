package main

import (
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/spf13/cobra"
)

func TestB64Cmd(t *testing.T) {
	rootCmd := &cobra.Command{Use: "beta"}
	rootCmd.AddCommand(b64Cmd())

	// Test the `b64` command with no arguments
	output := executeCommand(rootCmd, "b64")
	firstLine := strings.Split(output, "\n")[0]
	if firstLine != "Error: incorrect usage" {
		t.Errorf("Expected incorrect usage error, got %s", firstLine)
	}

	// Test the `b64 encode` command with a string argument
	output = executeCommand(rootCmd, "b64", "to", "Hello, World!")
	//fmt.Println("Actual out: " + output)
	// print length of output:
	fmt.Println("Length of output: ", len(output))
	expectedOutput := "SGVsbG8sIFdvcmxkIQ==\n"
	if output != expectedOutput {
		t.Errorf("Expected %s, got %s", expectedOutput, output)
	}

	// Test the `b64 decode` command with a base64 string argument
	output = executeCommand(rootCmd, "b64", "from", "SGVsbG8sIFdvcmxkIQ==")
	expectedOutput = "Hello, World!\n"
	if output != expectedOutput {
		t.Errorf("Expected %s, got %s", expectedOutput, output)
	}
}

func executeCommand(root *cobra.Command, args ...string) string {
	buf := new(bytes.Buffer)
	root.SetOut(buf)
	root.SetErr(buf)
	root.SetArgs(args)
	_ = root.Execute()
	return buf.String()
}
