package main

import "testing"

func Test_romanToInt(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{"case0", "MCMXCIV", 1994},
		{"case1", "LXXI", 71},
		{"case2", "DCLXXI", 671},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := romanToInt(tt.s); got != tt.want {
				t.Errorf("romanToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
