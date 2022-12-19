package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

var dirCmd = &cobra.Command{
	Use:   "dir",
	Short: "Open the notes directory with default explorer",
	RunE: func(cmd *cobra.Command, args []string) error {
		b, err := exec.Command("open", noteDirectory).Output()
		if err != nil {
			return err
		}
		fmt.Print(string(b))
		return nil
	},
}

func init() {
	rootCmd.AddCommand(dirCmd)
}
