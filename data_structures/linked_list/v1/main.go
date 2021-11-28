package main

import (
	"fmt"
)

type node struct {
	data int
	next *node
}

type linkedList struct {
	head *node
	tail *node
	size int
}

func (ll *linkedList) addAtTheEnd(n *node) {
	if ll.head == nil {
		ll.head = n
		ll.tail = ll.head
	} else {
		ll.tail.next = n
	}

	ll.size++
}

func (ll *linkedList) prepend(n *node) {
	second := ll.head
	ll.head = n
	ll.head.next = second
	ll.size++
}

func (ll *linkedList) printLine() {
	for ll.size != 0 {
		fmt.Printf("%d\n", ll.head.data)
		ll.size--
	}
}

func (ll *linkedList) deleteWithValue(value int) {
	if ll.size == 0 {
		return
	}

	if ll.head.data == value {
		ll.head = ll.head.next
		ll.size--
		return
	}

	previous := ll.head
	for previous.next.data != value {
		if previous.next.next == nil {
			return
		}
		previous = previous.next
	}
	previous.next = previous.next.next
	ll.size--
}

//func lookupNode(t *Node, v int) bool {
//	if root == nil {
//		t = &Node{v, nil}
//		root = t
//		return false
//	}
//
//	if v == t.Value {
//		return true
//	}
//
//	if t.Next == nil {
//		return false
//	}
//
//	return lookupNode(t.Next, v)
//}

func main() {
	ll := linkedList{}
	node1 := &node{data: 10}
	node2 := &node{data: 20}
	node3 := &node{data: 30}

	ll.prepend(node1)
	ll.prepend(node2)
	ll.prepend(node3)

	//fmt.Println(ll)
	ll.printLine()

	//root = nil
	//traverse(root)
	//addNode(root, 1)
	//addNode(root, -1)
	//traverse(root)
	//addNode(root, 10)
	//addNode(root, 5)
	//addNode(root, 45)
	//addNode(root, 5)
	//addNode(root, 5)
	//traverse(root)
	//addNode(root, 100)
	//traverse(root)
	//
	//if lookupNode(root, 100) {
	//	fmt.Println("Node exists!")
	//} else {
	//	fmt.Println("Node does not exist!")
	//}
	//
	//if lookupNode(root, -100) {
	//	fmt.Println("Node exists!")
	//} else {
	//	fmt.Println("Node does not exist!")
	//}
}
