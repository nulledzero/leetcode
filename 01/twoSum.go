package main

import "fmt"

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

func main() {
	array := []int{-1, -2, -3, -4, -5}
	target := -8

	fmt.Println(twoSum(array, target))
}
