package main

func removeDuplicates(nums []int) int {
	k := 0
	dupes := 0
	for i := 0; i < len(nums); i++ {
		if i != 0 {
			if nums[i] == nums[i-1] && dupes > 0 {
				dupes += 1
				continue
			} else if nums[i] == nums[i-1] {
				dupes = 1
			} else {
				dupes = 0
			}
			nums[k] = nums[i]
		}
		k += 1
	}
	return k
}
