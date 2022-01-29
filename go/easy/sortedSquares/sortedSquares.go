package main

func findMid(array []int) int {
	var min, mid, result int
	max := len(array) - 1

	for min <= max {
		mid = (min + max) / 2
		if array[mid] > 0 {
			max = mid - 1
		} else if array[mid] < 0 {
			result = mid
			min = mid + 1
		} else {
			return mid
		}
	}
	return result
}

func intAbs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func sortedSquares(nums []int) []int {
	ln := len(nums)
	result := make([]int, ln)
	var left, right int
	left = findMid(nums)
	right = left + 1
	for i := 0; i < ln; i++ {
		if left > -1 && right < ln {
			if intAbs(nums[left]) < intAbs(nums[right]) {
				result[i] = nums[left] * nums[left]
				left -= 1
			} else {
				result[i] = nums[right] * nums[right]
				right += 1
			}
		} else if left > -1 {
			result[i] = nums[left] * nums[left]
			left -= 1
		} else {
			result[i] = nums[right] * nums[right]
			right += 1
		}
	}
	return result
}
