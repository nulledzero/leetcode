package main

func moveZeroes(nums []int) {
	k := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			continue
		}
		if nums[k] == 0 {
			nums[k], nums[i] = nums[i], 0
		}
		k += 1
	}
}
