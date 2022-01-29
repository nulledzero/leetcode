package main

func reverse(nums []int, left, right int) {
	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left += 1
		right -= 1
	}
}

func rotate(nums []int, k int) {
	k = k % len(nums)
	left, right := 0, len(nums)-1
	reverse(nums, left, right)
	reverse(nums, left, k-1)
	reverse(nums, k, right)
}
