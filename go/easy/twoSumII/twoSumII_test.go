package main

import (
	"reflect"
	"testing"
)

func Test_twoSum(t *testing.T) {
	tests := []struct {
		name   string
		array  []int
		target int
		want   []int
	}{
		{"case0", []int{2, 7, 11, 15}, 9, []int{1, 2}},
		{"case1", []int{2, 3, 4}, 6, []int{1, 3}},
		{"case2", []int{-1, 0}, -1, []int{1, 2}},
		{"case3", []int{1, 2, 3, 4, 4, 9, 56, 90}, 8, []int{4, 5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoSum(tt.array, tt.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
