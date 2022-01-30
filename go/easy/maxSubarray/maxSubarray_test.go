package main

import "testing"

func Test_maxSubArray(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"case0", []int{1}, 1},
		{"case1", []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, 6},
		{"case2", []int{5, 4, -1, 7, 8}, 23},
		{"case3", []int{-1}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSubArray(tt.nums); got != tt.want {
				t.Errorf("maxSubarray() = %v, want %v", got, tt.want)
			}
		})
	}
}
