package main

func containsDuplicate(nums []int) bool {
	dict := make(map[int]int)
	for index, elem := range nums {
		if _, exists := dict[elem]; exists {
			return true
		}
		dict[elem] = index
	}
	return false
}
