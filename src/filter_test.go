package src

import "testing"

func TestFilter(t *testing.T) {
	tasks := []TimeEntry{
		{Tags: []string{"a", "b"}},
		{Tags: []string{"b", "c"}},
		{Tags: []string{"c", "d"}},
	}
	filtered := Filter(tasks, "b")
	if len(filtered) != 2 {
		t.Fatalf("expected 2 tasks, got %d", len(filtered))
	}
	if !Contains(filtered[0].Tags, "b") {
		t.Fatalf("expected first task to contain tag b")
	}
	if !Contains(filtered[1].Tags, "b") {
		t.Fatalf("expected second task to contain tag b")
	}
}

func TestContains(t *testing.T) {
	if Contains([]string{"a", "b"}, "a") != true {
		t.Fatalf("expected true")
	}
	if Contains([]string{"a", "b"}, "c") != false {
		t.Fatalf("expected false")
	}
}
