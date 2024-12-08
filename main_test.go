package main

import (
	"io"
	"os"
	"testing"
)

func captureOutput(f func() error) (string, error) {
	orig := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	err := f()
	os.Stdout = orig
	w.Close()
	out, _ := io.ReadAll(r)
	return string(out), err
}

func TestRun(t *testing.T) {
	_, err := captureOutput(func() error {
		Run()
		return nil
	})
	if err != nil {
		t.Errorf("Run() failed: %v", err)
	}
}
