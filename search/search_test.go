package search

import (
	"testing"
)

var tts = []struct {
	in     []int
	target int
	out    int
}{
	{in: []int{4, 5, 6, 7, 0, 1, 2}, target: 0, out: 4},
}

func TestSearch(t *testing.T) {
	for _, tt := range tts {
		got := search(tt.in, tt.target)
		if got != tt.out {
			t.Errorf("Input: %v | Expected: %d, Got: %d", tt.in, tt.out, got)
		}
	}
}
