package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	slow := head
	fast := head
	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
		if fast != nil {
			if fast.Next != nil {
				fast = fast.Next
			}
		}
	}
	return slow
}
