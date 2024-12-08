package src

import (
	"fmt"
	"time"
)

// Day represents a day with a date and a list of tasks.
type Day struct {
	Date  string
	Tasks []TimeEntry
}

func (d *Day) Report() (string, error) {
	sum, err := d.DurationSum()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("Date: %s, Duration: %d minutes", d.Date, sum), nil
}

func (d *Day) DurationSum() (int, error) {
	sum := 0
	for _, task := range d.Tasks {
		duration, err := task.Duration()
		if err != nil {
			return 0, err
		}
		sum += duration
	}
	return sum, nil
}

// TimeEntry represents a task with a start time, comment and tags.
type TimeEntry struct {
	Start   string
	Comment string
	Tags    []string
	Next    *TimeEntry
}

// Duration returns the duration of the task in minutes.
func (t *TimeEntry) Duration() (int, error) {
	if t.Next == nil {
		return 0, nil
	}
	current, err := t.StartTime()
	if err != nil {
		return 0, err
	}
	next, err := t.Next.StartTime()
	if err != nil {
		return 0, err
	}
	return int(next.Sub(current).Minutes()), nil
}

// StartTime returns the start time of the task.
func (t *TimeEntry) StartTime() (time.Time, error) {
	return time.Parse("15:04", t.Start)
}

// NewTask creates a new task with the given start time, comment and tags.
func NewTask(start string, comment string, tags []string) TimeEntry {
	return TimeEntry{Start: start, Comment: comment, Tags: tags}
}
