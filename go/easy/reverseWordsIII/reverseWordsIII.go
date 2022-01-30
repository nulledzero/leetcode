package main

import (
	"strings"
)

func reverseString(str []byte) []byte {
	left := 0
	right := len(str) - 1
	for left < right {
		str[left], str[right] = str[right], str[left]
		left += 1
		right -= 1
	}
	return str
}

func reverseWords(str string) string {
	arr := strings.Fields(str)
	for i := 0; i < len(arr); i++ {
		arr[i] = string(reverseString([]byte(arr[i])))
	}
	return strings.Join(arr, " ")
}
