package cmd

import (
	"os"
	"strings"
	"testing"
)

func TestBreakCmd(t *testing.T) {
	if breakCmd.Use != "break" {
		t.Errorf("Expected breakCmd.Use to be 'break', got %s", breakCmd.Use)
	}
}

func TestInsertBreak(t *testing.T) {
	tmpFile, err := os.CreateTemp("", "insert-break-test")
	if err != nil {
		t.Errorf("CreateTemp() failed: %v", err)
	}
	defer os.Remove(tmpFile.Name())

	err = insertBreak(tmpFile.Name())
	if err != nil {
		t.Errorf("insertBreak() failed: %v", err)
	}
	text, err := os.ReadFile(tmpFile.Name())
	if err != nil {
		t.Errorf("ReadFile() failed: %v", err)
	}
	if strings.Contains(string(text), "break") == false {
		t.Errorf("Expected '#break', got %q", text)
	}
}
