package cmd

import "github.com/spf13/cobra"

var openCmd = &cobra.Command{
	Use:   "open",
	Short: "Open an existing note",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		path := getPath(args[0])

		err := openFile(path)
		if err != nil {
			return err
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(openCmd)
}
