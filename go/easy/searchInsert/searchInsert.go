package main

func searchInsert(nums []int, target int) int {
	var min, mid int
	max := len(nums) - 1

	for min <= max {
		mid = (min + max) / 2
		if target < nums[mid] {
			max = mid - 1
		} else if target > nums[mid] {
			min = mid + 1
		} else {
			return mid
		}
	}
	if nums[mid] <= target {
		return mid + 1
	} else {
		return mid
	}
}
