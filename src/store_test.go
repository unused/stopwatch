package src

import (
	"os"
	"testing"
)

func TestNewStore(t *testing.T) {
	store := NewStore("test")
	if store.filename != "test" {
		t.Errorf("Expected filename 'test', got '%s'", store.filename)
	}
	if store.days != nil {
		t.Error("Expected days to be initialized")
	}
}

func TestLoadStore(t *testing.T) {
	tmpfile, err := os.CreateTemp("", "testfile")
	if err != nil {
		t.Errorf("TempFile() failed: %v", err)
	}
	defer os.Remove(tmpfile.Name())
	_, err = tmpfile.WriteString("2024-12-12\n  12:23 sent a letter to santa\n")
	if err != nil {
		t.Errorf("Failed to write to file: %v", err)
	}

	store, err := LoadStore(tmpfile.Name())
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if store.filename != tmpfile.Name() {
		t.Errorf("Expected filename 'test', got '%s'", store.filename)
	}
	if store.days == nil {
		t.Error("Expected days to be initialized")
	}
}

func TestRead(t *testing.T) {
	tmpfile, err := os.CreateTemp("", "testfile")
	if err != nil {
		t.Errorf("TempFile() failed: %v", err)
	}
	_, err = tmpfile.WriteString("2024-12-12\n  12:23 sent a letter to santa\n")
	if err != nil {
		t.Errorf("Failed to write to file: %v", err)
	}

	store := NewStore(tmpfile.Name())
	if err := store.Read(); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if store.days == nil {
		t.Error("Expected days to be initialized")
	}
}
