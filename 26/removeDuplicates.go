package main

func removeDuplicates(nums []int) int {
	k := 0
	ln := len(nums)
	for i := 0; i < ln; i++ {
		if i != 0 {
			if nums[i] == nums[i-1] {
				continue
			}
			nums[k] = nums[i]
		}
		k += 1
	}
	return k
}
