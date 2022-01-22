package main

import "fmt"

func romanToInt(s string) int {
	roman := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000}

	sum := 0
	prev := 0

	for _, elem := range s {
		value := roman[elem]
		if value > prev {
			sum += value - 2*prev
		} else {
			sum += value
		}
		prev = value
	}
	return sum
}

func main() {
	fmt.Println(romanToInt("MCMXCIV"))
}
