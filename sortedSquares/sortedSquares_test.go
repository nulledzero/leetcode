package main

import (
	"reflect"
	"testing"
)

func Test_sortedSquares(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{"case0", []int{-4, -1, 0, 3, 10}, []int{0, 1, 9, 16, 100}},
		{"case1", []int{-7, -3, 2, 3, 11}, []int{4, 9, 9, 49, 121}},
		{"case2", []int{-2, -1, 3}, []int{1, 4, 9}},
		{"case3", []int{-1, 2, 2}, []int{1, 4, 4}},
		{"case4", []int{-5, -3, -2, -1}, []int{1, 4, 9, 25}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortedSquares(tt.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortedSquares() = %v, want %v", got, tt.want)
			}
		})
	}
}
