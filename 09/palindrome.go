package main

import (
	"fmt"
	"math"
)

func isPalindrome(x int) bool {
	str := fmt.Sprint(x)
	midLn := int(math.Ceil(float64(len(str)) / 2))
	ln := len(str) - 1
	for i := 0; i < midLn; i++ {
		if str[i] != str[ln-i] {
			return false
		}
	}
	return true
}

func main() {
	palindrome := 121
	fmt.Println(isPalindrome(palindrome))
}
