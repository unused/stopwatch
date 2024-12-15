package src

import (
	_ "embed"
	"encoding/json"
	"io"
	"os"
	"text/template"
)

//go:embed templates/report.txt
var TextReportContent string
var TextReport = template.Must(template.New("report").Parse(TextReportContent))

// PrintDays prints the last n days to the standard output.
func PrintDays(days []Day, n int, format string) error {
	if n < len(days) {
		days = days[(len(days) - n):]
	}
	return ReportDays(days, format, os.Stdout)
}

// ReportDays writes a report of the days to a provided writer. Skip printing
// the date if there is only one day.
func ReportDays(days []Day, format string, w io.Writer) error {
	if format == "json" {
		jsonDays, err := json.Marshal(days)
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, string(jsonDays))
		return err
	}
	return TextReport.Execute(w, days)
}
