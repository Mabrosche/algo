package main

import (
	"fmt"
)

func main() {
	var t Tree

	t.insert('F')
	t.insert('B')
	t.insert('A')
	t.insert('D')
	t.insert('C')
	t.insert('E')
	t.insert('G')
	t.insert('I')
	t.insert('H')

	fmt.Printf("Inorder: ")
	printInOrder(t.rootNode)
	fmt.Println()

	fmt.Printf("Preorder: ")
	printPreOrder(t.rootNode)
	fmt.Println()

	fmt.Printf("Postorder: ")
	printPostOrder(t.rootNode)
	fmt.Println()
}

type Tree struct {
	rootNode *Node
}

type Node struct {
	data  int
	left  *Node
	right *Node
}

func (t *Tree) find(data byte) int {
	if t.rootNode == nil {
		return 0
	}
	return t.find(data)
}

func (n *Node) find(k int) int {
	if n == nil {
		return 0
	}

	switch {
	case k == n.data:
		return n.data
	case k < n.data:
		return n.left.find(k)
	default:
		return n.right.find(k)
	}
}

func (t *Tree) insertIterative(data int) {
	if t.rootNode == nil {
		t.rootNode = &Node{data: data}
	}

	var parent *Node
	var current = t.rootNode

	for current != nil {
		parent = current
		if current.data <= data {
			current = current.right
		} else {
			current = current.left
		}
	}

	if parent.data <= data {
		parent.right = &Node{data: data}
	} else {
		parent.left = &Node{data: data}
	}
}

func (t *Tree) insert(data int) {
	if t.rootNode == nil {
		t.rootNode = &Node{data: data}
	} else {
		t.rootNode.insert(data)
	}
}

func (n *Node) insert(data int) {
	if data <= n.data {
		if n.left == nil {
			n.left = &Node{data: data}
		} else {
			n.left.insert(data)
		}
	} else {
		if n.right == nil {
			n.right = &Node{data: data}
		} else {
			n.right.insert(data)
		}
	}
}

func printInOrder(n *Node) {
	if n == nil {
		return
	} else {
		printInOrder(n.left)
		fmt.Printf("%c ", n.data)
		printInOrder(n.right)
	}
}

func printPreOrder(n *Node) {
	if n == nil {
		return
	} else {
		fmt.Printf("%c ", n.data)
		printPreOrder(n.left)
		printPreOrder(n.right)
	}
}

func printPostOrder(n *Node) {
	if n == nil {
		return
	} else {
		printPostOrder(n.left)
		printPostOrder(n.right)
		fmt.Printf("%c ", n.data)
	}
}

func (t *Tree) Traverse(n *Node, f func(*Node)) {
	if n == nil {
		return
	}
	t.Traverse(n.left, f)
	f(n)
	t.Traverse(n.right, f)
}

func sameTree(root1 *Node, root2 *Node) bool {
	if root1 == nil && root2 == nil {
		return true
	}

	if root1 == nil || root2 == nil {
		return false
	}

	return root1.data == root2.data && sameTree(root1.left, root2.left) && sameTree(root1.right, root2.right)
}
