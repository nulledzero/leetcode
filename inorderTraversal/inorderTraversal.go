package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Node struct {
	Data *TreeNode
	Next *Node
}

type Stack struct {
	Last *Node
}

func (stack *Stack) push(val *TreeNode) {
	node := &Node{Data: val, Next: stack.Last}
	stack.Last = node
}

func (stack *Stack) pop() *TreeNode {
	value := stack.Last.Data
	stack.Last = stack.Last.Next
	return value
}

func inorderTraversal(root *TreeNode) []int {
	stack := &Stack{}
	resultStack := &Stack{}
	treeSize := 0
	for {
		if root != nil {
			stack.push(root)
			root = root.Left
		} else if stack.Last != nil {
			root = stack.pop()
			resultStack.push(root)
			treeSize += 1
			root = root.Right
		} else {
			break
		}
	}
	results := make([]int, treeSize)
	for i := treeSize - 1; i >= 0; i-- {
		results[i] = resultStack.pop().Val
	}
	return results
}
