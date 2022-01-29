package main

import "testing"

func Test_longestCommonPrefix(t *testing.T) {
	tests := []struct {
		name string
		strs []string
		want string
	}{
		{"case0", []string{"flower", "flow", "flight"}, "fl"},
		{"case1", []string{"dog", "race", "car"}, ""},
		{"case2", []string{"tree", "trunk", "tribrid"}, "tr"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestCommonPrefix(tt.strs); got != tt.want {
				t.Errorf("longestCommonPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}
