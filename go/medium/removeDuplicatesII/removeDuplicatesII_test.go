package main

import (
	"reflect"
	"testing"
)

func Test_removeDuplicates(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		want   int
		output []int
	}{
		{"case0", []int{1, 1, 1, 2, 2, 3}, 5, []int{1, 1, 2, 2, 3, 3}},
		{"case1", []int{0, 0, 1, 1, 1, 1, 2, 3, 3}, 7, []int{0, 0, 1, 1, 2, 3, 3, 3, 3}},
		{"case2", []int{}, 0, []int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeDuplicates(tt.nums); got != tt.want || !reflect.DeepEqual(tt.nums, tt.output) {
				t.Errorf("removeDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}
