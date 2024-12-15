package src

type Store struct {
	filename string
	days     []Day
}

// NewStore creates a new store with the given filename.
func NewStore(filename string) *Store {
	return &Store{filename: filename}
}

// LoadStore loads the store from the file.
func LoadStore(filename string) (*Store, error) {
	store := NewStore(filename)
	err := store.Read()
	return store, err
}

// Read reads the days from the file.
func (s *Store) Read() error {
	var err error
	s.days, err = ReadFile(s.filename)
	return err
}

// Days returns the days, reading them from the file if not done yet.
func (s *Store) Days() ([]Day, error) {
	if s.days == nil {
		if err := s.Read(); err != nil {
			return nil, err
		}
	}
	return s.days, nil
}

// Append appends a line to the file.
func (s *Store) Append(line string) error {
	if err := AppendDate(s.filename); err != nil {
		return err
	}
	return AppendTimeEntry(line, s.filename)
}

// Filter filters the tasks by the given tags.
func (s *Store) Filter(tags ...string) {
	for _, tag := range tags {
		for i, day := range s.days {
			s.days[i].Tasks = Filter(day.Tasks, tag)
		}
	}
}
