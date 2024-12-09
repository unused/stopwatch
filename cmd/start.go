package cmd

import (
	"strings"

	"github.com/spf13/cobra"
	"github.com/unused/stopwatch/src"
)

// startCmd represents the start command that adds a new task to the list. In
// case the current day is not today, then a new day entry is created before
// adding the task.
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start a new task",
	Long: `Start a new task. This command will create a new task and record the
current time as the start time. Note that #break is a reserved keyword.`,
	Args: cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		err := src.AppendDate(srcFile)
		if err != nil {
			return err
		}

		line := strings.Join(args, " ")
		return src.AppendTimeEntry(line, srcFile)
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
