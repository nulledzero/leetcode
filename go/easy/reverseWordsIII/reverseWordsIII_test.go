package main

import "testing"

func Test_reverseWords(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{"case0", "hello world", "olleh dlrow"},
		{"case1", "never gonna", "reven annog"},
		{"case2", "give you up", "evig uoy pu"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseWords(tt.s); got != tt.want {
				t.Errorf("reverseWords() = %v, want %v", got, tt.want)
			}
		})
	}
}
