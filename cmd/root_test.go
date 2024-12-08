package cmd

import (
	"testing"
)

func TestRootCmd(t *testing.T) {
	if rootCmd.Use != "stpw" {
		t.Errorf("Expected rootCmd.Use to be 'stpw', got %s", rootCmd.Use)
	}
}
