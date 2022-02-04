package main

func search(nums []int, target int) int {
	min := 0
	max := len(nums) - 1

	for min <= max {
		mid := (min + max) / 2
		if target < nums[mid] {
			max = mid - 1
		} else if target > nums[mid] {
			min = mid + 1
		} else {
			return mid
		}
	}

	return -1
}
