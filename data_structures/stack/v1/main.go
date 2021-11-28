package main

type Int struct {
	Slice []int
	Len   int
}

func (s *Int) Push(i int) {
	s.Slice = append(s.Slice, i)
	s.Len++
}

func (s *Int) Peek() int {
	return s.Slice[s.Len-1]
}

func (s *Int) Pull() int {
	res := s.Peek()

	s.Slice = s.Slice[:s.Len]
	s.Len--
	return res
}

func (s Int) IsEmpty() bool {
	return s.Len == 0
}
