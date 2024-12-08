package src

import "testing"

func TestParseTaskLineTime(t *testing.T) {
	line := "  09:00  Meeting with John #work #important"
	day := Day{}
	ParseTaskLine(line, &day)

	if len(day.Tasks) != 1 {
		t.Errorf("Expected 1 task, got %d", len(day.Tasks))
	}

	task := day.Tasks[0]
	if task.Start != "09:00" {
		t.Errorf("Expected start time 09:00, got %s", task.Start)
	}
	expectation := "Meeting with John #work #important"
	if task.Comment != expectation {
		t.Errorf("Expected description %s, got '%s'", expectation, task.Comment)
	}
	if task.Tags[0] != "work" {
		t.Errorf("Expected tag 'work', got '%s'", task.Tags[0])
	}
	if task.Tags[1] != "important" {
		t.Errorf("Expected tag 'important', got '%s'", task.Tags[1])
	}
}
