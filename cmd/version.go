package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const Version = "dev"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print note version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("note version %s\n", Version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
