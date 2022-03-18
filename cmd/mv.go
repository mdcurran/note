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
	RunE: func(cmd *cobra.Command, args []string) error {
		oldpath := getPath(args[0])
		newfile := args[1]
		newpath := getPath(newfile)

		_, err := os.Stat(newpath)
		if err == nil {
			return fmt.Errorf("cannot rename as '%s' already exists", newfile)
		}

		err = os.Rename(oldpath, newpath)
		if err != nil {
			return err
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(mvCmd)
}
