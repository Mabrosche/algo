package main

import "fmt"

//https://www.youtube.com/watch?v=9PHkM0Jri_4&list=PLrmLmBdmIlpv_jNDXtJGYTPNQ2L1gdHxu&index=9
func (t Tree) bftLevelOrderTraversal(root *Node) {
	if root == nil {
		return
	}

	var q []*Node
	q = append(q, root)

	for len(q) == 0 {
		root = q[len(q)-1]
		fmt.Println(root.data)

		if root.left != nil {
			q = append(q, root.left)
		}
		if root.right != nil {
			q = append(q, root.right)
		}
	}
}

func (t *Tree) isBST(rootNode *Node, min, max int) bool {
	if t.rootNode == nil {
		return true
	}

	if rootNode.data <= min || rootNode.data > max {
		return false
	}

	return t.isBST(rootNode.left, min, rootNode.data) && t.isBST(rootNode.right, rootNode.data, max)
}

func (t *Tree) rootToLeafSum(rootNode *Node, sum int, res []int) bool {
	if rootNode == nil {
		return false
	}

	if rootNode.left == nil && rootNode.right == nil {
		if rootNode.data == sum {
			res = append(res, rootNode.data)
			return true
		} else {
			return false
		}
	}

	if t.rootToLeafSum(rootNode.left, sum-rootNode.data, res) {
		res = append(res, rootNode.data)
		return true
	}

	if t.rootToLeafSum(rootNode.right, sum-rootNode.data, res) {
		res = append(res, rootNode.data)
		return true
	}

	return false
}

func (t *Tree) maxHeight(root *Node) int {
	if t.rootNode == nil {
		return 0
	}

	leftHeight := t.maxHeight(root.left)
	rightHeight := t.maxHeight(root.right)

	if leftHeight < rightHeight {
		return rightHeight + 1
	}

	return leftHeight + 1
}

func (t *Tree) nodesCount() int {
	if t.rootNode == nil {
		return 0
	}

	leftCount := t.nodesCount()
	rightCount := t.nodesCount()

	return leftCount + rightCount + 1
}
