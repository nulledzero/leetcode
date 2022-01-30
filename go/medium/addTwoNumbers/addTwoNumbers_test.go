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

func Test_addTwoNumbers(t *testing.T) {
	tests := []struct {
		name string
		l1   *ListNode
		l2   *ListNode
		want *ListNode
	}{
		{"case0", createList([]int{3, 4, 2}), createList([]int{4, 6, 5}), createList([]int{8, 0, 7})},
		{"case1", createList([]int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}), createList([]int{5, 6, 4}), createList([]int{6, 6, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addTwoNumbers(tt.l1, tt.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addTwoNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
