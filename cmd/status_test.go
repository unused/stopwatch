package cmd

import (
	"testing"
)

func TestStatusCmd(t *testing.T) {
	if statusCmd.Use != "status" {
		t.Errorf("Expected statusCmd.Use to be 'status', got %s", statusCmd.Use)
	}
}
