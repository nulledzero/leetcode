package main

import "testing"

var versions []int

func isBadVersion(version int) bool {
	return versions[version] != 1
}

func setVersions(vers []int) {
	versions = vers
}

func Test_firstBadVersion(t *testing.T) {
	tests := []struct {
		name string
		vers []int
		n    int
		want int
	}{
		{"case0", []int{1, 1, 1, 1, 0, 0}, 5, 4},
		{"case1", []int{1, 0}, 1, 1},
		{"case2", []int{1, 0, 0, 0, 0, 0, 0}, 6, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setVersions(tt.vers)
			if got := firstBadVersion(tt.n); got != tt.want {
				t.Errorf("firstBadVersion() = %v, want %v", got, tt.want)
			}
		})
	}
}
