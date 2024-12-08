package src

import "testing"

func TestNewTask(t *testing.T) {
	task := NewTask("09:00", "Meeting with John", []string{"#work", "#important"})

	if task.Start != "09:00" {
		t.Errorf("Expected start time 09:00, got %s", task.Start)
	}
	expectation := "Meeting with John"
	if task.Comment != expectation {
		t.Errorf("Expected description %s, got '%s'", expectation, task.Comment)
	}
	if task.Tags[0] != "#work" {
		t.Errorf("Expected tag 'work', got '%s'", task.Tags[0])
	}
	if task.Tags[1] != "#important" {
		t.Errorf("Expected tag 'important', got '%s'", task.Tags[1])
	}
}
