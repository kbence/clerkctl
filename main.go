package main

import (
	"os"

	"github.com/clerk/clerk-sdk-go/v2"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "clerkctl",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

func main() {
	clerk.SetKey(os.Getenv("CLERK_SECRET_KEY"))

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
