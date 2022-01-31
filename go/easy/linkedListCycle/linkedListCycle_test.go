package main

import (
	"testing"
)

func createList(data []int, pos int) *ListNode {
	var head *ListNode = nil
	var loopNode, tail *ListNode
	pos2 := len(data) - 1 - pos
	for index, elem := range data {
		newNode := &ListNode{Val: elem, Next: head}
		if index == pos2 {
			loopNode = newNode
		}
		if index == 0 {
			tail = newNode
		}
		head = newNode
	}
	if pos != -1 {
		tail.Next = loopNode
	}
	return head
}

func Test_hasCycle(t *testing.T) {
	tests := []struct {
		name string
		head *ListNode
		want bool
	}{
		{"case0", createList([]int{-4, 0, 2, 3}, 1), true},
		{"case1", createList([]int{2, 1}, 0), true},
		{"case2", createList([]int{1}, -1), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasCycle(tt.head); got != tt.want {
				t.Errorf("hasCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}
