package main

import "testing"

func Test_maximumWealth(t *testing.T) {
	tests := []struct {
		name     string
		accounts [][]int
		want     int
	}{
		{"case0", [][]int{{1, 2, 3}, {3, 2, 1}}, 6},
		{"case1", [][]int{{1, 5}, {7, 3}, {3, 5}}, 10},
		{"case2", [][]int{{2, 8, 7}, {7, 1, 3}, {1, 9, 5}}, 17},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumWealth(tt.accounts); got != tt.want {
				t.Errorf("maximumWealth() = %v, want %v", got, tt.want)
			}
		})
	}
}
