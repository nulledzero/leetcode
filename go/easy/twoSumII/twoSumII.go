package main

func twoSum(array []int, target int) []int {
	left := 0
	right := len(array) - 1
	for left <= right {
		sum := array[left] + array[right]
		if sum < target {
			left += 1
		} else if sum > target {
			right -= 1
		} else {
			return []int{left + 1, right + 1}
		}
	}
	return []int{}
}
