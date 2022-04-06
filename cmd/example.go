package cmd

import (
	"github.com/spf13/cobra"
)

var print string
var exampleCmd = &cobra.Command{
	Use:   "Example",
	Short: "Example Short",
	Run: func(cmd *cobra.Command, args []string) {
		println("Running Example..." + print)
	},
}

func init() {
	rootCmd.AddCommand(exampleCmd)

	exampleCmd.Flags().StringVarP(&print, "print", "p", "Test", "Example to print")
}
