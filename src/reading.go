package src

import (
	"bufio"
	"os"
	"regexp"
	"strings"
)

var dateRegex = regexp.MustCompile(`^\d{4}-\d{2}-\d{2}$`)
var taskRegex = regexp.MustCompile(`^\s+(\d{2}:\d{2})\s+`)
var commentRegex = regexp.MustCompile(`^\s*\d{2}:\d{2}\s+(.*)`)
var tagRegex = regexp.MustCompile(`#(\w+)`)

// ReadFile reads file and returns a slice of days.
func ReadFile(filename string) ([]Day, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var days []Day
	var currentDay Day

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if dateRegex.MatchString(line) {
			// Start a new day
			if currentDay.Date != "" {
				days = append(days, currentDay)
			}
			currentDay = Day{Date: line}
		} else if taskRegex.MatchString(line) {
			ParseTaskLine(line, &currentDay)
		} // else ignore content
	}

	// Add the last day
	if currentDay.Date != "" {
		days = append(days, currentDay)
	}

	return days, scanner.Err()
}

// ParseTaskLine parses a line of text and adds a task to the current day.
func ParseTaskLine(line string, currentDay *Day) error {
	// Extract time
	timeMatches := taskRegex.FindStringSubmatch(line)
	if timeMatches == nil {
		return nil
	}
	start := timeMatches[1]

	// Extract comment and tags
	var comment string
	commentMatches := commentRegex.FindStringSubmatch(line)
	if commentMatches != nil {
		comment = strings.TrimSpace(commentMatches[1])
	}

	tags := tagRegex.FindAllStringSubmatch(line, -1)
	// Extract tag strings from the tags slice
	var tagStrings []string
	for _, tag := range tags {
		tagStrings = append(tagStrings, tag[1])
	}
	currentTask := NewTask(start, comment, tagStrings)
	if len(currentDay.Tasks) > 0 {
		startTime, err := currentTask.StartTime()
		if err != nil {
			return err
		}
		currentDay.Tasks[len(currentDay.Tasks)-1].EndTime = startTime
	}
	currentDay.Tasks = append(currentDay.Tasks, currentTask)
	return nil
}
