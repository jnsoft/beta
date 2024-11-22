package main

import (
	"bytes"
	"os"
	"strings"
	"testing"

	"github.com/spf13/cobra"
)

const PLAIN = "Hello, World!"
const ENCODED = "SGVsbG8sIFdvcmxkIQ=="

func TestB64Cmd(t *testing.T) {
	rootCmd := &cobra.Command{Use: "beta"}
	rootCmd.AddCommand(b64Cmd())

	// Test the `b64` command with no arguments
	output := executeCommand(rootCmd, "b64")
	firstLine := strings.Split(output, "\n")[0]
	if firstLine != "Error: incorrect usage" {
		t.Errorf("Expected incorrect usage error, got %s", firstLine)
	}
}

func TestToB64Cmd(t *testing.T) {
	rootCmd := &cobra.Command{Use: "beta"}
	rootCmd.AddCommand(b64Cmd())

	// Test the `b64 encode` command with a string argument
	output := executeCommand(rootCmd, "b64", "to", PLAIN)
	expectedOutput := ENCODED + "\n"
	if output != expectedOutput {
		t.Errorf("Expected %s, got %s", expectedOutput, output)
	}
}
func TestFromB64Cmd(t *testing.T) {
	rootCmd := &cobra.Command{Use: "beta"}
	rootCmd.AddCommand(b64Cmd())

	// Test the `b64 decode` command with a base64 string argument
	output := executeCommand(rootCmd, "b64", "from", ENCODED)
	expectedOutput := PLAIN + "\n"
	if output != expectedOutput {
		t.Errorf("Expected %s, got %s", expectedOutput, output)
	}
}

func TestFileToB64Cmd(t *testing.T) {
	rootCmd := &cobra.Command{Use: "beta"}
	rootCmd.AddCommand(b64Cmd())

	// Create a temporary file to act as input
	tempInputFile, err := os.CreateTemp("", "input.txt")
	if err != nil {
		t.Fatal(err)
	}

	defer os.Remove(tempInputFile.Name()) // Clean up

	// Write some data to the temp file
	_, err = tempInputFile.Write([]byte(PLAIN))
	if err != nil {
		t.Fatal(err)
	}
	tempInputFile.Close()

	// Create a temporary file to act as output
	tempOutputFile, err := os.CreateTemp("", "output.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tempOutputFile.Name()) // Clean up

	output := executeCommand(rootCmd, "b64", "encode", "--input", tempInputFile.Name(), "--output", tempOutputFile.Name(), "--lines", "0")
	expectedOutput := ENCODED
	// Read the output from the temp file
	outputData, err := os.ReadFile(tempOutputFile.Name())
	if err != nil {
		t.Fatal(err)
	}
	if string(outputData) != expectedOutput {
		t.Errorf("Expected %s, got %s", expectedOutput, output)
	}
}

func TestFileFromB64Cmd(t *testing.T) {
	rootCmd := &cobra.Command{Use: "beta"}
	rootCmd.AddCommand(b64Cmd())

	// Create a temporary file to act as input
	tempInputFile, err := os.CreateTemp("", "input.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tempInputFile.Name()) // Clean up

	// Write some base64 data to the temp file
	_, err = tempInputFile.Write([]byte(ENCODED))
	if err != nil {
		t.Fatal(err)
	}
	tempInputFile.Close()

	// Create a temporary file to act as output
	tempOutputFile, err := os.CreateTemp("", "output.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tempOutputFile.Name()) // Clean up

	// Run the command to decode the file
	output := executeCommand(rootCmd, "b64", "decode", "--input", tempInputFile.Name(), "--output", tempOutputFile.Name())
	expectedOutput := PLAIN
	// Read the output from the temp file
	outputData, err := os.ReadFile(tempOutputFile.Name())
	if err != nil {
		t.Fatal(err)
	}
	if string(outputData) != expectedOutput {
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
