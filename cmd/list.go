package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all notes",
	Run: func(cmd *cobra.Command, args []string) {
		b, err := exec.Command("ls", noteDirectory).Output()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Print(string(b))
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
