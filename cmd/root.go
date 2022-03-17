package cmd

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

var noteDirectory string

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

func getPath(filename string) string {
	return strings.Join([]string{noteDirectory, filename}, "/")
}

func openFile(path string) error {
	c := exec.Command("open", "-e", path)
	return c.Run()
}

func init() {
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	noteDirectory = fmt.Sprintf("%s/.local/opt/note", home)

	cobra.OnInitialize(func() {
		_, err := os.Stat(noteDirectory)
		if errors.Is(err, os.ErrNotExist) {
			err = os.MkdirAll(noteDirectory, 0755)
		}
		if err != nil {
			panic(err)
		}
	})
}
