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

func Test_mergeTwoLists(t *testing.T) {
	type args struct {
		list1 *ListNode
		list2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{"case0", args{createList([]int{3, 2, 1}), createList([]int{4, 2, 1})}, createList([]int{4, 3, 2, 2, 1, 1})},
		{"case1", args{createList([]int{}), createList([]int{1})}, createList([]int{1})},
		{"case2", args{createList([]int{2}), createList([]int{3})}, createList([]int{3, 2})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeTwoLists(tt.args.list1, tt.args.list2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeTwoLists() = %v, want %v", got, tt.want)
			}
		})
	}
}
