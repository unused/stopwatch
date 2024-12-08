package cmd

import (
	"github.com/spf13/cobra"
	"github.com/unused/stopwatch/src"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Print the list of the last 5 days",
	Long: `Print the list of the last five days. This includes the date and all
tasks`,
	RunE: func(cmd *cobra.Command, args []string) error {
		days, err := src.ReadFile(srcFile)
		if err != nil {
			return err
		}
		return src.PrintDays(days, 5)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
