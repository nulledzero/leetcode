package main

func intMax(x, y int) int {
	if y > x {
		return y
	}
	return x
}

func maxSubArray(nums []int) int {
	maxSum := nums[0]
	maxCur := nums[0]
	for i := 1; i < len(nums); i++ {
		maxCur = intMax(maxCur+nums[i], nums[i])
		maxSum = intMax(maxSum, maxCur)
	}
	return maxSum
}
