package main

import (
	"fmt"
)

type Int struct {
	Slice []int
	Len   int
}

func main() {
	m := &Int{}
	fmt.Println(m)
	buildHeap := []int{10, 20, 30, 5, 7, 9, 11, 13, 15, 17}
	for _, v := range buildHeap {
		m.MaxInsert(v)
		fmt.Println(m)
	}

	for i := 0; i < 5; i++ {
		m.MaxExtract()
		fmt.Println(m)
	}
}

func (s *Int) MaxInsert(i int) {
	s.Slice = append(s.Slice, i)
	s.Len++
	s.MaxHeapifyUp()
}

func (s *Int) MaxExtract() int {
	if s.Len == 0 {
		return -1
	}

	res := s.Slice[0]

	s.Slice[0] = s.Slice[s.Len-1]
	s.Slice = s.Slice[:s.Len-1]
	s.Len--

	s.MaxHeapifyDown()

	return res
}

func (s *Int) MaxHeapifyUp() {
	index := s.Len - 1
	for s.Slice[parent(index)] < s.Slice[index] {
		s.Swap(parent(index), index)
		index = parent(index)
	}
}

func parent(i int) int {
	return (i - 1) / 2
}

func left(i int) int {
	return (i * 2) + 1
}

func right(i int) int {
	return (i * 2) + 2
}
func (s *Int) Swap(i1, i2 int) {
	s.Slice[i1], s.Slice[i2] = s.Slice[i2], s.Slice[i1]
}

func (s *Int) MaxHeapifyDown() {
	index := 0
	l, r := left(index), right(index)
	childToCompare := 0

	for l <= s.Len-1 {
		if l == s.Len-1 {
			childToCompare = l
		} else if s.Slice[l] > s.Slice[r] {
			childToCompare = l
		} else {
			childToCompare = r
		}

		if s.Slice[index] < s.Slice[childToCompare] {
			s.Swap(index, childToCompare)
			index = childToCompare
			l, r = left(index), right(index)
		} else {
			return
		}
	}
}
