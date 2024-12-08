package src

import (
	"os"
	"strings"
	"testing"
)

func TestAppendLine(t *testing.T) {
	tmpfile, err := os.CreateTemp("", "testfile")
	if err != nil {
		t.Errorf("TempFile() failed: %v", err)
	}
	defer os.Remove(tmpfile.Name())

	line := "This is a test"
	err = AppendLine(line, tmpfile.Name())
	if err != nil {
		t.Errorf("AppendLine() failed: %v", err)
	}
	data, err := os.ReadFile(tmpfile.Name())
	if err != nil {
		t.Errorf("Failed to read file: %v", err)
	}
	if strings.TrimSpace(string(data)) != line {
		t.Errorf("Incorrect data written to file. Got: %q, Want: %q", data, line)
	}
}
