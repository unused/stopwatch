package src

import (
	"bytes"
	"testing"
)

func TestReportEmptyDays(t *testing.T) {
	buf := new(bytes.Buffer)
	if err := ReportDays([]Day{}, "", buf); err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

func TestReportDays(t *testing.T) {
	buf := new(bytes.Buffer)
	tasks := []TimeEntry{{Start: "10:00", Comment: "meeting"},
		{Start: "11:00", Comment: "break"},
		{Start: "12:00", Comment: "lunch"}}
	endTime, err := tasks[1].StartTime()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	tasks[0].EndTime = endTime
	endTime, err = tasks[2].StartTime()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	tasks[1].EndTime = endTime
	days := []Day{{Date: "2020-01-01", Tasks: tasks}}
	if err := ReportDays(days, "", buf); err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	expectedResult := `Date: 2020-01-01, Duration: 120 minutes
  Start: 10:00, Duration: 60 minutes meeting
  Start: 11:00, Duration: 60 minutes break
  Start: 12:00, Duration: 0 minutes lunch


`
	if buf.String() != expectedResult {
		t.Errorf("Expected %q, got %q", expectedResult, buf.String())
	}
}
