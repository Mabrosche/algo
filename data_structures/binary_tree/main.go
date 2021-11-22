package main

import (
	"fmt"
)

func main() {
	test := &Node{
		Left:  nil,
		Value: 100,
		Right: nil,
	}

	fmt.Println(test)
}

type Node struct {
	Left  *Node
	Value int
	Right *Node
}

func (n *Node) String() string {
	var leftValue, rightValue int

	if n.Left != nil {
		leftValue = n.Left.Value
	}
	if n.Left != nil {
		rightValue = n.Left.Value
	}

	return fmt.Sprintf("Val=%d L=%d R=%d\n", n.Value, leftValue, rightValue)
}

//type Tree struct {
//	Root *Node
//	Len  int
//}
//
//func traverse(t *Tree) {
//	if t == nil {
//		return
//	}
//
//	traverse(t.Left)
//	fmt.Println(t.Value, " ")
//	traverse(t.Right)
//}
//
//func create(n int) *Tree {
//	var t *Tree
//	rand.Seed(time.Now().Unix())
//	for i := 0; i < 2*n; i++ {
//		temp := rand.Intn(n * 2)
//		t = insert(t, temp)
//	}
//
//	return t
//}
//
//func insert(t *Tree, val int) *Tree {
//	if t == nil {
//		return &Tree{nil, val, nil}
//	}
//
//	if val == t.Value {
//		return t
//	}
//
//	if val < t.Value {
//		t.Left = insert(t.Left, val)
//		return t
//	}
//
//	t.Right = insert(t.Right, val)
//	return t
//}
