package cmd

import (
	"testing"
)

func TestStartCmd(t *testing.T) {
	if startCmd.Use != "start" {
		t.Errorf("Expected startCmd.Use to be 'start', got %s", startCmd.Use)
	}
}
