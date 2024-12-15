package cmd

import (
	"strings"

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
		store, err := src.LoadStore(srcFile)
		if err != nil {
			return err
		}
		if filter != "" {
			store.Filter(strings.Split(filter, ",")...)
		}
		days, err := store.Days()
		if err != nil {
			return err
		}
		return src.PrintDays(days, 5, format)
	},
}

func init() {
	listCmd.Flags().StringVarP(&format, "format", "f", "", "Output format")
	listCmd.Flags().StringVarP(&filter, "filter", "", "", "Filter by tags")
	rootCmd.AddCommand(listCmd)
}
