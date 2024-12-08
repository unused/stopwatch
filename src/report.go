package src

import (
	"fmt"
	"io"
	"os"
)

// PrintDays prints the last n days to the standard output.
func PrintDays(days []Day, n int) error {
	if n < len(days) {
		days = days[(len(days) - n):]
	}
	return ReportDays(days, os.Stdout)
}

// ReportDays writes a report of the days to a provided writer. Skip printing
// the date if there is only one day.
func ReportDays(days []Day, w io.Writer) error {
	for _, day := range days {
		summary, err := day.Report()
		if err != nil {
			return err
		}
		if _, err := io.WriteString(w, summary+"\n"); err != nil {
			return err
		}
		for _, task := range day.Tasks {
			duration, err := task.Duration()
			if err != nil {
				return err
			}
			line := fmt.Sprintf("  Start: %s, Duration: %d minutes\n", task.Start, duration)
			if _, err := io.WriteString(w, line); err != nil {
				return err
			}
		}
	}
	return nil
}
