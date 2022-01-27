package main

import "testing"

func Test_removeDuplicates(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"case0", []int{1, 1, 1, 2}, 2},
		{"case1", []int{1, 2, 3, 3, 4, 5}, 5},
		{"case2", []int{6, 6, 6, 6}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeDuplicates(tt.nums); got != tt.want {
				t.Errorf("removeDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}
