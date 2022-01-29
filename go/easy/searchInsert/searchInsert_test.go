package main

import "testing"

func Test_searchInsert(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   int
	}{
		{"case0", []int{1, 3, 5, 6}, 5, 2},
		{"case1", []int{1, 3, 5, 6}, 2, 1},
		{"case2", []int{1, 3, 5, 6}, 7, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchInsert(tt.nums, tt.target); got != tt.want {
				t.Errorf("searchInsert() = %v, want %v", got, tt.want)
			}
		})
	}
}
