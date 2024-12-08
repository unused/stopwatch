package cmd

import (
	"github.com/spf13/cobra"
	"github.com/unused/stopwatch/src"
)

// breakCmd represents the break command that adds a new task to the list. In
// case the current day is not today, then a new day entry is created before
// adding the task.
var breakCmd = &cobra.Command{
	Use:   "break",
	Short: "Insert a break or end the day",
	// Long:  `Insert a break or end the day.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return insertBreak(srcFile)
	},
}

func init() {
	rootCmd.AddCommand(breakCmd)
}

// insertBreak inserts a break entry in the source file. If the current day is
// not today, then a new day entry is created before adding the break.
func insertBreak(filename string) error {
	err := src.AppendDate(filename)
	if err != nil {
		return err
	}

	return src.AppendTimeEntry("#break", filename)
}
