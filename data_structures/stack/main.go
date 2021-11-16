package main

import (
	"fmt"
)

type Int struct {
	Slice []int
	Len   int
}

func (s *Int) Pushh(i int) {
	s.Slice = append(s.Slice, i)
	s.Len++
}

func (s *Int) Peek() int {
	return s.Slice[s.Len-1]
}

func (s *Int) Pulll() int {
	res := s.Peek()

	s.Slice = s.Slice[:s.Len]
	s.Len--
	return res
}

//////////

type Stack struct {
	Root *Node
	Len  int
}

type Node struct {
	Value int
	Next  *Node
}

func (s *Stack) Push(v int) bool {
	if s == nil {
		s.Root = &Node{v, nil}
		s.Len = 1
		return true
	}

	temp := &Node{v, nil}
	temp.Next = s.Root
	s.Root = temp
	s.Len++
	return true
}

func (s *Stack) Pop(t *Node) (int, bool) {
	if s.Len == 0 {
		return 0, false
	}

	if s.Len == 1 {
		s.Len = 0
		s.Root = nil
		return t.Value, true
	}

	s.Root = s.Root.Next
	s.Len--
	return t.Value, true
}

func (s *Stack) traverse(t *Node) {
	if s.Len == 0 {
		fmt.Println("Empty Stack!")
		return
	}

	for t != nil {
		fmt.Printf("%d -> ", t.Value)
		t = t.Next
	}
	fmt.Println()
}

func main() {
	var stack Stack
	v, b := stack.Pop(stack.Root)
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
