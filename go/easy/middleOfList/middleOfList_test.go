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

func Test_middleNode(t *testing.T) {
	tests := []struct {
		name string
		head *ListNode
		want *ListNode
	}{
		{"case0", createList([]int{5, 4, 3, 2, 1}), createList([]int{5, 4, 3})},
		{"case1", createList([]int{6, 5, 4, 3, 2, 1}), createList([]int{6, 5, 4})},
		{"case2", createList([]int{3, 2, 1}), createList([]int{3, 2})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := middleNode(tt.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("middleNode() = %v, want %v", got, tt.want)
			}
		})
	}
}
