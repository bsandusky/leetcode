package singleNumber

import (
	"testing"
)

var tts = []struct {
	in  []int
	out int
}{
	{in: []int{2, 2, 1}, out: 1},
	{in: []int{4, 1, 2, 1, 2}, out: 4},
}

func TestSingleNumber(t *testing.T) {
	for _, tt := range tts {
		got := singleNumber(tt.in)
		if got != tt.out {
			t.Errorf("Expected: %d, Got: %d", tt.out, got)
		}
	}
}
