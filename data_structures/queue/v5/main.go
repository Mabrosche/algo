package main

import (
	"fmt"
	"strconv"
	"strings"
)

<<<<<<< HEAD:data_structures/queue/v2/main.go
=======
type Int struct {
	Slice []int
}

func (q *Int) Enqueue(i int) {
	q.Slice = append(q.Slice, i)
}

func (q *Int) Dequeue() int {
	res := q.Slice[0]
	q.Slice = q.Slice[1:]
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

>>>>>>> 0ff0ee8acbf8f89afbd48976fe26d182c4c3aebc:data_structures/queue/v5/main.go
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

////////////

type Node struct {
	Next  *Node
	Value int
}

type Queue struct {
	Head *Node
	Tail *Node
	Len  int
}

func (q *Queue) Push(v int) bool {
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

func (q Queue) IsEmpty() bool {
	return q.Head == q.Tail
}

func (n Node) String() string {
	return strconv.Itoa(n.Value)
}

func (q Queue) String() string {
	sb := strings.Builder{}
	q.traverse(sb)
	return sb.String()
}

func (q *Queue) traverse() {
	res := q.Root
	for res != nil {
		fmt.Printf("%d -> \n", res.Value)
		res = res.Next
	}

	fmt.Println("Empty Queue!")
}
