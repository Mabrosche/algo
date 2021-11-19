package main

import (
	"fmt"
)

type node struct {
	value int
	next  *node
}

type linkedList struct {
	head *node
	tail *node
	len  int
}

func (ll *linkedList) addAtTheEnd(n *node) {
	if ll.head == nil {
		ll.head = n
		ll.tail = ll.head
	} else {
		ll.tail.next = n
	}

	ll.len++
}

func (ll *linkedList) prepend(n *node) {
	temp := ll.head
	ll.head = n
	ll.head.next = temp

	if ll.tail == nil {
		ll.tail = ll.head
	}

	ll.len++
}

func (ll *linkedList) printLine() {
	temp := ll.head
	for ll.len != 0 {
		fmt.Printf("%d\n", temp.value)
		temp = temp.next
		ll.len--
	}
}

func (ll *linkedList) deleteWithValue(value int) {
	if ll.len == 0 {
		return
	}

	if ll.head.value == value {
		ll.head = ll.head.next
		ll.len--
		return
	}

	previous := ll.head
	for previous.next.value != value {
		if previous.next.next == nil {
			return
		}
		previous = previous.next
	}
	previous.next = previous.next.next
	ll.len--
}

func (ll linkedList) lookupNode(v int) int {
	if ll.len == 0 || ll.head.next == nil {
		return -1
	}

	if v == ll.head.value {
		return v
	}

	return ll.lookupNode(v)
}

func main() {
	ll := linkedList{}
	node1 := &node{value: 10}
	node2 := &node{value: 20}
	node3 := &node{value: 30}

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
