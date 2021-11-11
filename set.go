package main

type set struct {
	m map[rune]bool
}

func NewSet() *set {
	s := &set{}
	s.m = make(map[rune]bool)
	return s
}

func (s *set) Add(value rune) {
	s.m[value] = true
}

func (s *set) Remove(value rune) {
	delete(s.m, value)
}

func (s *set) Contains(value rune) bool {
	_, val := s.m[value]
	return val
}
