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

func Test_reverseList(t *testing.T) {
	tests := []struct {
		name string
		head *ListNode
		want *ListNode
	}{
		{"case0", createList([]int{1, 2, 3}), createList([]int{3, 2, 1})},
		{"case1", createList([]int{3, 4, 5, 6}), createList([]int{6, 5, 4, 3})},
		{"case2", createList([]int{1, 10, 100}), createList([]int{100, 10, 1})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseList(tt.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseList() = %v, want %v", got, tt.want)
			}
		})
	}
}
