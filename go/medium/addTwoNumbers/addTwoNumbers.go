package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var prev, curr, next *ListNode
	curr = head
	for curr != nil {
		next = curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var rem int
	var l3 *ListNode = nil
	for l1 != nil && l2 != nil {
		value := l1.Val + l2.Val + rem
		node := &ListNode{Val: value % 10, Next: l3}
		rem = value / 10
		l1, l2, l3 = l1.Next, l2.Next, node
	}
	remList := l1
	if l1 == nil {
		remList = l2
	}
	for remList != nil || rem != 0 {
		if remList != nil {
			value := remList.Val + rem
			node := &ListNode{Val: value % 10, Next: l3}
			rem = value / 10
			remList, l3 = remList.Next, node
		} else {
			node := &ListNode{Val: rem % 10, Next: l3}
			rem /= 10
			l3 = node
		}
	}
	return reverseList(l3)
}
