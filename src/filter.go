package src

// Filter filters the tasks by the given tag.
func Filter(tasks []TimeEntry, tag string) (filtered []TimeEntry) {
	for _, task := range tasks {
		if Contains(task.Tags, tag) {
			filtered = append(filtered, task)
		}
	}
	return filtered
}

// Contains checks if a list contains an element.
func Contains(list []string, element string) bool {
	for _, e := range list {
		if e == element {
			return true
		}
	}
	return false
}
