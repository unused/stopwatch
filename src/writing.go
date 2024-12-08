package src

import (
	"os"
	"time"
)

func AppendDate(filename string) error {
	date := time.Now().Format("2006-01-02")
	days, err := ReadFile(filename)
	if err != nil {
		return err
	}

	if len(days) > 0 && days[len(days)-1].Date == date {
		return nil
	}
	return AppendLine("\n"+date, filename)
}

func AppendTimeEntry(text, filename string) error {
	currentTime := time.Now()
	start := currentTime.Format("15:04")
	return AppendLine("  "+start+" "+text, filename)
}

// AppendLine appends a line to a file
func AppendLine(line, filename string) error {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(line + "\n")
	return err
}
