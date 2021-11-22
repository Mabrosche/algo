package main

import (
	"fmt"
)

//https://www.youtube.com/watch?v=fAAZixBzIAI
func main() {
	a := newNode(1)
	b := newNode(2)
	c := newNode(3)

	d := newNode(4)

	t := &bst{}
	t.node = b
	b.left = a
	b.right = c
	c.right = d

	t.len = 4

	res := t.dftTraversalIterative()
	for i := range res {
		fmt.Println(res[i])
	}
}

type bst struct {
	node *node
	len  int
}

type node struct {
	data  int
	left  *node
	right *node
}

func (b bst) dftTraversalIterative() []*node {
	if b.len == 0 {
		return nil
	}

	var res []*node
	var stk []*node
	stk = append(stk, b.node)

	for b.len > 0 {
		node := stk[len(stk)-1]
		fmt.Println("---> ", node)
		res = append(res, node)

		b.len--
		stk = stk[:len(stk)-1]

		if node.right != nil {
			stk = append(stk, node.right)
		}
		if node.left != nil {
			stk = append(stk, node.left)
		}
	}

	return res
}

func (n *node) BreadthFirstSearch(array []int) []int {
	queue := []*node{n}
	for len(queue) > 0 {
		current := queue[0]
		queue := queue[1:]

		array = append(array, current.Value)

		for _, child := range n.Children {
			queue := append(queue, child)
		}
	}
}

func newNode(val int) *node {
	return &node{
		data: val,
	}
}

//func (n *node) insert(val int) {
//
//}
//
//func (b *bst) insert(val int) *bst {
//	n := newNode(val)
//	if b.node == nil {
//		b.node = n
//		b.len++
//	}
//
//	if b.node.data == val {
//		return b
//	}
//
//	if val > b.node.data {
//		b.node.right = n
//		b.len++
//	} else {
//		b.node.left = n
//		b.len++
//	}
//
//	return b
//}

func (n node) String() string {
	return fmt.Sprintf("[Val: %d -> Left: %s -> Right: %s]", n.data, n.left, n.right)
}

//func (b bst) String() string {
//	sb := strings.Builder{}
//
//	return sb.String()
//}
