package main

import "testing"

func Test_search(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   int
	}{
		{"case0", []int{-1, 0, 3, 5, 9, 12}, 9, 4},
		{"case1", []int{-1, 0, 3, 5, 9, 12}, 2, -1},
		{"case2", []int{4}, 4, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := search(tt.nums, tt.target); got != tt.want {
				t.Errorf("search() = %v, want %v", got, tt.want)
			}
		})
	}
}
