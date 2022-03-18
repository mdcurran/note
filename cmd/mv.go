package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var mvCmd = &cobra.Command{
	Use:   "mv",
	Short: "Rename a note",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("not implemented\n")
	},
}

func init() {
	rootCmd.AddCommand(mvCmd)
}
