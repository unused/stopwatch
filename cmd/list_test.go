package cmd

import (
	"testing"
)

func TestListCmd(t *testing.T) {
	if listCmd.Use != "list" {
		t.Errorf("Expected listCmd.Use to be 'list', got %s", listCmd.Use)
	}
}
