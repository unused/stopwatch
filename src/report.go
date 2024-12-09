package src

import (
	"encoding/json"
	"io"
	"os"
	"text/template"
)

var TextReport = template.Must(template.New("report").Parse(`{{ range .}}Date: {{.Date}}, Duration: {{.DurationSum}} minutes
{{ range .Tasks}}  Start: {{.Start}}, Duration: {{.Duration}} minutes {{.Comment}}
{{ end }}
{{ end }}`))

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
