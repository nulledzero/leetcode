package main

import (
	"reflect"
	"testing"
)

func Test_rotate(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want []int
	}{
		{"case0", []int{1, 2, 3, 4, 5, 6, 7}, 3, []int{5, 6, 7, 1, 2, 3, 4}},
		{"case1", []int{-1, -100, 3, 99}, 2, []int{3, 99, -1, 100}},
		{"case2", []int{1, 2, 3}, 3, []int{1, 2, 3}},
		{"case3", []int{1, 2, 3, 4, 5}, 2, []int{4, 5, 1, 2, 3}},
		{"case4", []int{-1}, 2, []int{-1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rotate(tt.nums, tt.k)
			if !reflect.DeepEqual(tt.nums, tt.want) {
				t.Errorf("rotate() = %v, want %v", tt.nums, tt.want)
			}
		})
	}
}
