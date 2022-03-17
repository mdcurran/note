package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var newCmd = &cobra.Command{
	Use:   "new",
	Short: "create a new note",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("not implemented\n")
	},
}

func init() {
	rootCmd.AddCommand(newCmd)
}
