package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Create a new note",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		filename := args[0]
		path := getPath(filename)

		_, err := os.Stat(path)
		if err == nil {
			return fmt.Errorf("%s already exists", filename)
		}

		_, err = os.Create(path)
		if err != nil {
			return fmt.Errorf("error creating %s: %s", filename, err)
		}

		err = openFile(path)
		if err != nil {
			return err
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(newCmd)
}
