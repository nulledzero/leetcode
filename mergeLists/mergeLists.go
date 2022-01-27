package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var head, tail *ListNode

	if list1 == nil {
		return list2
	} else if list2 == nil {
		return list1
	}

	if list1.Val < list2.Val {
		head = list1
		list1 = list1.Next
	} else {
		head = list2
		list2 = list2.Next
	}
	tail = head
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			tail.Next = list1
			tail = tail.Next
			list1 = list1.Next
		} else {
			tail.Next = list2
			tail = tail.Next
			list2 = list2.Next
		}
	}
	for list1 != nil {
		tail.Next = list1
		tail = tail.Next
		list1 = list1.Next
	}
	for list2 != nil {
		tail.Next = list2
		tail = tail.Next
		list2 = list2.Next
	}
	return head
}
