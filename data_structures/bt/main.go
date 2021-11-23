package main

import (
	"fmt"
)

//https://www.youtube.com/watch?v=fAAZixBzIAI
func main() {
	a := newNode(1)
	b := newNode(2)
	c := newNode(3)

	d := newNode(7)
	e := newNode(5)
	f := newNode(8)

	t := &bst{}
	t.node = b
	b.left = a
	b.right = c
	c.right = d

	d.left = e
	d.right = f

	t.len = 6

	res := t.dftTraversalIterative()
	for v := range res {
		fmt.Print(res[v].data)
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

func (b bst) dfSearch(root *node) []*node {
	if root == nil {
		return nil
	}

	var res []*node

	res = append(res, b.dftTraversalRecursive(root.left)...)
	res = append(res, root)
	res = append(res, b.dftTraversalRecursive(root.right)...)

	return res
}

func (b bst) bftTraversalIterative() []*node {
	if b.node == nil {
		return nil
	}

	var res []*node

	var qu []*node
	qu = append(qu, b.node)
	for b.len > 0 {
		current := qu[0]
		fmt.Println("---> ", current)
		res = append(res, current)

		b.len--
		qu = qu[1:]

		if current.left != nil {
			qu = append(qu, current.left)
		}
		if current.right != nil {
			qu = append(qu, current.right)
		}
	}

	return res
}

//https://www.youtube.com/watch?v=fAAZixBzIAI
func (b bst) dftTraversalRecursive(root *node) []*node {
	if root == nil {
		return nil
	}

	var res []*node

	res = append(res, b.dftTraversalRecursive(root.left)...)
	res = append(res, root)
	res = append(res, b.dftTraversalRecursive(root.right)...)

	return res
}

func (b bst) dftTraversalIterative() []*node {
	if b.len == 0 {
		return nil
	}

	var stk []*node
	stk = append(stk, b.node)

	var res []*node
	for b.len > 0 {
		current := stk[len(stk)-1]
		fmt.Println("---> ", current)
		res = append(res, current)

		b.len--
		stk = stk[:len(stk)-1]

		if current.right != nil {
			stk = append(stk, current.right)
		}
		if current.left != nil {
			stk = append(stk, current.left)
		}
	}

	return res
}

//https://www.youtube.com/watch?v=5cPbNCrdotA&list=PL2_aWCzGMAwI3W_JlcBbtYTwiQSsOTa6P&index=37
//func (b bst) getSuccessor(data int) *node {
//	if b.len == 0 {
//		return nil
//	}
//
//	// search the Node - O(h)
//	current := b.find(b.node, data)
//	if current == nil {
//		return nil
//	}
//
//	// case 1: Node has right subtree
//	if current.right != nil {
//		b.findMin(current.right)
//	} else { // case
//		var successor *node
//		ancestor := b.node
//
//		for ancestor != current {
//			if current.data < ancestor.data {
//				// so far this is the deepest node for which current node is in left
//				successor = ancestor
//				ancestor = ancestor.left
//			} else {
//				ancestor = ancestor.right
//			}
//		}
//		return successor
//	}
//
//}

func (b bst) findMin(root *node) *node {
	if root == nil {
		return nil
	}

	for root.left != nil {
		root = root.left
	}
	return root
}

//func (b bst) find(root *node, data int) *node {
//	if(root == NULL) return NULL;
//	else if(root->data == data) return root;
//	else if(root->data < data) return Find(root->right,data);
//	else return Find(root->left,data);
//}

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
