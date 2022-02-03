package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func intMax(x, y int) int {
	if y > x {
		return y
	}
	return x
}

func maxDepth(root *TreeNode) int {
	if root != nil {
		return intMax(maxDepth(root.Left), maxDepth(root.Right)) + 1
	}
	return 0
}
