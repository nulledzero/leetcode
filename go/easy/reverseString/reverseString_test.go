package main

import (
	"reflect"
	"testing"
)

func Test_reverseString(t *testing.T) {
	tests := []struct {
		name string
		s    []byte
		want []byte
	}{
		{"case0", []byte{'h', 'e', 'l', 'l', 'o'}, []byte{'o', 'l', 'l', 'e', 'h'}},
		{"case1", []byte{'t', 'e', 's', 't'}, []byte{'t', 's', 'e', 't'}},
		{"case2", []byte{'r'}, []byte{'r'}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reverseString(tt.s)
			if !reflect.DeepEqual(tt.s, tt.want) {
				t.Errorf("reverseString() = %v, want %v", tt.s, tt.want)
			}
		})
	}
}
