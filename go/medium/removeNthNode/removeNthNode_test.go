package main

import (
	"reflect"
	"testing"
)

func createList(data []int) *ListNode {
	var head *ListNode = nil
	for _, elem := range data {
		newNode := &ListNode{Val: elem, Next: head}
		head = newNode
	}
	return head
}

func Test_removeNthFromEnd(t *testing.T) {
	tests := []struct {
		name string
		head *ListNode
		n    int
		want *ListNode
	}{
		{"case0", createList([]int{5, 4, 3, 2, 1}), 2, createList([]int{5, 3, 2, 1})},
		{"case1", createList([]int{1}), 1, createList([]int{})},
		{"case2", createList([]int{2, 1}), 1, createList([]int{1})},
		{"case3", createList([]int{2, 1}), 2, createList([]int{2})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeNthFromEnd(tt.head, tt.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeNthFromEnd() = %v, want %v", got, tt.want)
			}
		})
	}
}
