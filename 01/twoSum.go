package main

func twoSum(nums []int, target int) []int {
	hashMap := make(map[int]int)
	for index, elem := range nums {
		hashMap[elem] = index
	}
	for index, elem := range nums {
		if mapIndex, ok := hashMap[target-elem]; mapIndex != index && ok {
			return []int{index, mapIndex}
		}
	}
	return []int{}
}
