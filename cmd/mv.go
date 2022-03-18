package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var mvCmd = &cobra.Command{
	Use:   "mv",
	Short: "Rename a note",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		oldpath := getPath(args[0])
		newfile := args[1]
		newpath := getPath(newfile)

		_, err := os.Stat(newpath)
		if err == nil {
			fmt.Printf("cannot rename as '%s' already exists\n", newfile)
			return
		}

		err = os.Rename(oldpath, newpath)
		if err != nil {
			fmt.Println(err)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(mvCmd)
}
