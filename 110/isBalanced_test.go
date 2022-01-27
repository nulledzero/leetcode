package main

import (
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

func Test_isBalanced(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		want bool
	}{
		{"case0", createNode([]string{"3", "9", "20", "null", "null", "15", "7"}, 0), true},
		{"case1", createNode([]string{}, 0), true},
		{"case2", createNode([]string{"1", "2", "2", "3", "3", "null", "null", "4", "4"}, 0), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isBalanced(tt.root); got != tt.want {
				t.Errorf("isBalanced() = %v, want %v", got, tt.want)
			}
		})
	}
}
