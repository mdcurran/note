package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "note",
	Short: "note manages text files from the command line",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	noteDirectory := fmt.Sprintf("%s/.local/opt/note", home)

	cobra.OnInitialize(func() {
		err := os.MkdirAll(noteDirectory, 0755)
		if err != nil {
			panic(err)
		}
	})
}
