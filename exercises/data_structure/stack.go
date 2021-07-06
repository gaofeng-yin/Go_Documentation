package data_structure

import "log"

type Stack struct {
	items []int
}

func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *Stack) Push(i int) {
	s.items = append(s.items, i)
}

func (s *Stack) Pop() int {
	if s.IsEmpty() {
		log.Panic("Empty Stack")
	}
	last := len(s.items) - 1

	toRemove := s.items[last]

	s.items = s.items[:last]

	return toRemove
}
