package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	n := &node{1, nil, nil}
	n.left = &node{0, nil, nil}
	n.right = &node{2, nil, nil}

	b := bst{
		root: n,
		len:  3,
	}

	fmt.Println(b)

	b.appendNode(b.root, 5)
	b.appendNode(b.root, 4)
	b.appendNode(b.root, 6)

	fmt.Println(b)

	//fmt.Println(b.search(6))
	//fmt.Println(b.search(3))
	//
	//b.remove(6)
	//fmt.Println(b)
	//b.remove(3)
	//fmt.Println(b)
}

type node struct {
	value int
	left  *node
	right *node
}

type bst struct {
	root *node
	len  int
}

func (n node) String() string {
	return strconv.Itoa(n.value)
}

func (b bst) String() string {
	sb := strings.Builder{}
	b.inOrderTraversal(&sb, b.root)
	return sb.String()
}

func (b bst) inOrderTraversal(sb *strings.Builder, root *node) {
	if root == nil {
		return
	}

	b.inOrderTraversal(sb, root.left)
	sb.WriteString(fmt.Sprintf("%s ", root))
	b.inOrderTraversal(sb, root.right)
}

func (b *bst) appendNode(root *node, value int) *node {
	if root == nil {
		b.len++
		return &node{value: value}
	}

	if value < root.value {
		root.left = b.appendNode(root.left, value)
		b.len++

	} else {
		root.right = b.appendNode(root.right, value)
		b.len++
	}

	return root
}

func (b bst) search(value int) (*node, bool) {
	return b.searchByNode(b.root, value)
}

func (b bst) searchByNode(root *node, value int) (*node, bool) {
	if root == nil {
		return nil, false
	}

	if value == root.value {
		return root, true
	} else if value < root.value {
		return b.searchByNode(root.left, value)
	} else {
		return b.searchByNode(root.right, value)
	}
}

func (b *bst) remove(value int) {
	b.removeByNode(b.root, value)
}

func (b *bst) removeByNode(root *node, value int) *node {
	if root == nil {
		return root
	}

	if value > root.value {
		root.right = b.removeByNode(root.right, value)
	} else if value < root.value {
		root.left = b.removeByNode(root.left, value)
	} else {
		if root.left == nil {
			return root.right
		} else {
			temp := root.left
			for temp.right != nil {
				temp = temp.right
			}

			root.value = temp.value

			root.left = b.removeByNode(root.left, value)
		}
	}

	return root
}
