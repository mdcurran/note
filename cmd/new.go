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
	Run: func(cmd *cobra.Command, args []string) {
		filename := args[0]
		fullPath := getPath(filename)

		_, err := os.Stat(fullPath)
		if err == nil {
			fmt.Printf("%s already exists\n", filename)
			return
		}

		_, err = os.Create(fullPath)
		if err != nil {
			fmt.Printf("error creating %s: %s\n", filename, err)
			return
		}

		err = openFile(fullPath)
		if err != nil {
			fmt.Println(err)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(newCmd)
}
