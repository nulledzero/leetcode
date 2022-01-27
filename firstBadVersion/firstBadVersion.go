package main

func firstBadVersion(n int) int {
	min, mid, max := 1, 1, n
	result := n
	for min <= max {
		mid = (min + max) / 2
		if isBadVersion(mid) {
			result = mid
			max = mid - 1
		} else {
			min = mid + 1
		}
	}
	return result
}
