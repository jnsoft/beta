package main

import (
	"fmt"
	"os"

	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

func uuidCmd() *cobra.Command {
	var uuidCmd = &cobra.Command{
		Use:   "uuid",
		Short: "A UUID is a 128 bit (16 byte) Universal Unique IDentifier as defined in RFC 4122.",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return IncorrectUsageErr()
		},
		Run: func(cmd *cobra.Command, args []string) {
		},
	}

	uuidCmd.AddCommand(newUuidCmd())

	return uuidCmd
}

func newUuidCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "new",
		Short: "Generate a new version 4 uuid.",
		Run: func(cmd *cobra.Command, args []string) {
			guid, err := uuid.NewRandom()
			if err != nil {
				fmt.Println("Error generating UUID: ", err)
				os.Exit(1)
			}
			fmt.Println(guid)
		},
	}
	return cmd
}
