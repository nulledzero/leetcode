package main

import "testing"

func Test_maxProfit(t *testing.T) {
	tests := []struct {
		name   string
		prices []int
		want   int
	}{
		{"case0", []int{7, 1, 5, 3, 6, 4}, 5},
		{"case1", []int{7, 6, 4, 3, 1}, 0},
		{"case2", []int{1, 2, 3, 8, 1}, 7},
		{"case3", []int{1}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfit(tt.prices); got != tt.want {
				t.Errorf("maxProfit() = %v, want %v", got, tt.want)
			}
		})
	}
}
