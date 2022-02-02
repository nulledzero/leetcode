package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	index := 1
	ptr := head
	nth := head
	prev := head
	for ptr != nil {
		if index-n-1 >= 0 {
			nth = nth.Next
		}
		if index-n-2 >= 0 {
			prev = prev.Next
		}
		ptr = ptr.Next
		index += 1
	}
	if prev != nth {
		prev.Next = nth.Next
	} else {
		prev = prev.Next
		return prev
	}
	return head
}
