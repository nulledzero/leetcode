package main

import "testing"

func Test_isPalindrome(t *testing.T) {
	tests := []struct {
		name string
		data int
		want bool
	}{
		{"case0", 121, true},
		{"case1", 122, false},
		{"case2", 3330333, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.data); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
