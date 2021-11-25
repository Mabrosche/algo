package main

//https://www.youtube.com/watch?v=IJDJ0kBx2LM
func reverseLinkedList(head *node) *node {
	if head == nil || head.next == nil {
		return head
	}

	res := reverseLinkedList(head)
	head.next.next = head
	head.next = nil

	return res
}
