package main

import (
	"reflect"
	"strconv"
	"testing"
)

func createNode(arr []string, index int) *TreeNode {
	if index >= len(arr) {
		return nil
	}
	value, err := strconv.Atoi(arr[index])
	if err != nil {
		return nil
	}
	node := &TreeNode{Val: value}
	node.Right = createNode(arr, index*2+2)
	node.Left = createNode(arr, index*2+1)
	return node
}

func Test_inorderTraversal(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		want []int
	}{
		{"case0", createNode([]string{"1", "n", "2", "n", "n", "3"}, 0), []int{1, 3, 2}},
		{"case1", createNode([]string{}, 0), []int{}},
		{"case2", createNode([]string{"1"}, 0), []int{1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := inorderTraversal(tt.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("inorderTraversal() = %v, want %v", got, tt.want)
			}
		})
	}
}
