package main

import "fmt"

func longestCommonPrefix(strs []string) string {
	strCount := len(strs)
	prefix := ""
	index := 0
	if strCount > 1 {
		for {
			for i := 0; i < strCount-1; i++ {
				if len(strs[i]) < 1 || len(strs[i+1]) < 1 {
					return ""
				}
				if index >= len(strs[i]) || index >= len(strs[i+1]) {
					return prefix
				}
				if strs[i][index] != strs[i+1][index] {
					return prefix
				}
			}
			prefix += string(strs[0][index])
			index += 1
		}
	}
	return strs[0]
}

func main() {
	strs := []string{"flower", "fl", "flight"}
	fmt.Println(longestCommonPrefix(strs))
}
