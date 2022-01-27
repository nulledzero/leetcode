package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func intMax(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func intAbs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func checkBalance(root *TreeNode) (bool, int) {
	rState, lState := true, true
	var rHeight, lHeight int
	if root.Right != nil {
		rState, rHeight = checkBalance(root.Right)
	}
	if root.Left != nil {
		lState, lHeight = checkBalance(root.Left)
	}
	diff := intAbs(rHeight - lHeight)
	return rState && lState && diff < 2, intMax(rHeight, lHeight) + 1
}

func isBalanced(root *TreeNode) bool {
	if root != nil {
		result, _ := checkBalance(root)
		return result
	}
	return true
}
