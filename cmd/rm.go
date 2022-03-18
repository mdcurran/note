package cmd

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Delete a note permanently",
	Args:  cobra.ExactArgs(1),
	PreRunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("delete? [Y/n]")

		reader := bufio.NewReader(os.Stdin)
		for {
			s, _ := reader.ReadString('\n')
			s = strings.TrimSuffix(s, "\n")
			s = strings.ToLower(s)
			if len(s) > 1 {
				fmt.Fprintln(os.Stderr, "Please enter Y or N")
				continue
			}
			if strings.Compare(s, "n") == 0 {
				return fmt.Errorf("cancelling...")
			} else if strings.Compare(s, "y") == 0 {
				break
			} else {
				continue
			}
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		filename := args[0]
		path := getPath(filename)

		_, err := os.Stat(path)
		if errors.Is(err, os.ErrNotExist) {
			fmt.Printf("%s does not exist\n", filename)
			return
		}

		err = os.Remove(path)
		if err != nil {
			fmt.Printf("error deleting %s: %s\n", filename, err)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(rmCmd)
}
