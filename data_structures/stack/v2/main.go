package main

import (
	"fmt"
)

func main() {
	var stack Stack
	v, b := stack.Pop(stack.Head)
	if b {
		fmt.Print(v, " ")
	} else {
		fmt.Println("Pop() failed!")
	}

	//Push(100)
	//traverse(stack)
	//Push(200)
	//traverse(stack)
	//
	//for i := 0; i < 10; i++ {
	//	Push(i)
	//}
	//
	//for i := 0; i < 15; i++ {
	//	v, b := Pop(stack)
	//	if b {
	//		fmt.Print(v, " ")
	//	} else {
	//		break
	//	}
	//}
	//fmt.Println()
	//traverse(stack)
}

//////////

type Stack struct {
	Head *Node
	Tail *Node
	Len  int
}

type Node struct {
	Next  *Node
	Value int
}

func (s *Stack) Push(v int) bool {
	if !s.IsEmpty() {
		s.Tail.Next.Value = v
	}

	s.Tail.Value = v

	if s.IsEmpty() {
		s.Head.Value = v
		s.Tail = s.Head
		s.Len++

		return true
	}

	temp := s.Head
	s.Head.Value = s.Head.Value
	s.Head = temp
	s.Len++

	return true
}

func (s Stack) Peek() *Node {
	return s.Head
}

func (s *Stack) Pop() *Node {
	if s.IsEmpty() {
		return s.Peek()
	}

	if s.Len == 1 {
		temp := s.Peek()
		s.Head = nil
		s.Tail = s.Head
		s.Len--

		return temp
	}

	temp := s.Head
	s.Head = s.Head.Next
	s.Len--

	return temp
}

func (s Stack) IsEmpty() bool {
	return s.Head == s.Tail
}

func (s *Stack) PrintAll() {
	if s.Len == 0 {
		fmt.Println("Empty Stack!")
		return
	}

	for s.Head.Next != nil {
		fmt.Printf("%d -> \n", s.Head.Next.Value)
		s.Head = s.Head.Next
	}
}
