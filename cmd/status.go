package cmd

import (
	"github.com/spf13/cobra"
	"github.com/unused/stopwatch/src"
)

// statusCmd represents the status command
var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Print the status of the current day",
	Long: `Print the status of the current day. This includes the date and all
tasks`,
	RunE: func(cmd *cobra.Command, args []string) error {
		days, err := src.ReadFile(srcFile)
		if err != nil {
			return err
		}
		return src.PrintDays(days, 1)
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)
}
