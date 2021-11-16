package main

import (
	"fmt"
)

type Int struct {
	Slice []int
}

func (q *Int) Enqueue(i int) {
	q.Slice = append(q.Slice, i)
}

func (q *Int) Dequeue() int {
	res := q.Slice[0]
	q.Slice = q.Slice[1:len(q.Slice)]
	return res
}

func (q Int) String() string {
	return fmt.Sprint(q.Slice)
}

////////////

type Node struct {
	Value int
	Next  *Node
}

type Queue struct {
	Root *Node
	Len  int
}

func main() {
	var q Queue
	q.traverse()
	//q.Push(q.Root, 10)
	//q.Push(q.Root, 11)
	//fmt.Println("Size:", q.Len)

	//v, b := Pop(queue)
	//if b {
	//	fmt.Println("Pop:", v)
	//}
	//fmt.Println("Size:", size)
	//
	//for i := 0; i < 5; i++ {
	//	Push(queue, i)
	//}
	//traverse(queue)
	//fmt.Println("Size:", size)
	//
	//v, b = Pop(queue)
	//if b {
	//	fmt.Println("Pop:", v)
	//}
	//fmt.Println("Size:", size)
	//
	//v, b = Pop(queue)
	//if b {
	//	fmt.Println("Pop:", v)
	//}
	//fmt.Println("Size:", size)
	//traverse(queue)
}

func (q *Queue) Push(node *Node, v int) bool {
	node = &Node{v, nil}
	if q != nil {
		node.Next = q.Root
	}

	q.Root = node
	q.Len++

	return true
}

func (q *Queue) Pop(t *Node) int {
	if q.Len == 0 {
		return -1
	}

	if q.Len == 1 {
		q.Root = nil
		q.Len--
		return t.Value
	}

	temp := t
	for (t.Next) != nil {
		temp = t
		t = t.Next
	}

	v := temp.Next.Value
	temp.Next = nil
	q.Len--

	return v
}

func (q *Queue) traverse() {
	res := q.Root
	for res != nil {
		fmt.Printf("%d -> \n", res.Value)
		res = res.Next
	}

	fmt.Println("Empty Queue!")
}
