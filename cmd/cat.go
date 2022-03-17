package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

var catCmd = &cobra.Command{
	Use:   "cat",
	Short: "Print a note using GNU cat",
	Run: func(cmd *cobra.Command, args []string) {
		b, err := exec.Command("cat", getPath(args[0])).Output()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(b))
	},
}

func init() {
	rootCmd.AddCommand(catCmd)
}
