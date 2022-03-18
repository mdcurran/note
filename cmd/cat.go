package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

var catCmd = &cobra.Command{
	Use:   "cat",
	Short: "Print a note using GNU cat",
	RunE: func(cmd *cobra.Command, args []string) error {
		b, err := exec.Command("cat", getPath(args[0])).Output()
		if err != nil {
			return err
		}
		fmt.Println(string(b))
		return nil
	},
}

func init() {
	rootCmd.AddCommand(catCmd)
}
