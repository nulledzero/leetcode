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

func Test_maxDepth(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		want int
	}{
		{"case0", createNode([]string{"3", "9", "20", "null", "null", "15", "7"}, 0), 3},
		{"case1", createNode([]string{"1", "null", "2"}, 0), 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxDepth(tt.root); got != tt.want {
				t.Errorf("maxDepth() = %v, want %v", got, tt.want)
			}
		})
	}
}
