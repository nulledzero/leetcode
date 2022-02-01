package main

func intMax(x, y int) int {
	if y > x {
		return y
	}
	return x
}

func intMin(x, y int) int {
	if y < x {
		return y
	}
	return x
}

func maxProfit(prices []int) int {
	if len(prices) > 1 {
		price := prices[0]
		profit := 0
		for i := 1; i < len(prices); i++ {
			price = intMin(price, prices[i])
			profit = intMax(profit, prices[i]-price)
		}
		if profit > 0 {
			return profit
		}
	}
	return 0
}
