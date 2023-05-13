package cmd

import "github.com/spf13/cobra"

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "start",
	Long:  "start",
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

func init() {
	RootCmd.AddCommand(startCmd)
}