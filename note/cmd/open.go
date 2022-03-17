package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var openCmd = &cobra.Command{
	Use:   "open",
	Short: "open an existing note",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("not implemented\n")
	},
}

func init() {
	rootCmd.AddCommand(openCmd)
}
