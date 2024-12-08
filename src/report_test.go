package src

import (
	"bytes"
	"testing"
)

func TestReportDays(t *testing.T) {
	buf := new(bytes.Buffer)
	if err := ReportDays([]Day{}, buf); err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}
