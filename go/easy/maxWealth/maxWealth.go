package main

func maximumWealth(accounts [][]int) int {
	maxWealth := 0
	for _, account := range accounts {
		wealth := 0
		for _, elem := range account {
			wealth += elem
		}
		if wealth > maxWealth {
			maxWealth = wealth
		}
	}
	return maxWealth
}
