package main

func singleNumber(nums []int) int {
	var num int
	for i := 0; i < len(nums); i++ {
		num ^= nums[i]
	}
	return num
}
