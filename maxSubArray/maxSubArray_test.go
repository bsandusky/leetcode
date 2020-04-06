package maxSubArray

import (
	"testing"
)

var tts = []struct {
	in  []int
	out int
}{
	{in: []int{1, 2, 3, -4}, out: 6},
	{in: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, out: 6},
	{in: []int{-1}, out: -1},
	{in: []int{-2, -1}, out: -1},
	{in: []int{-1, -2}, out: -1},
	{in: []int{-1, 0, -2}, out: 0},
}

func TestMaxSubArray(t *testing.T) {
	for _, tt := range tts {
		got := maxSubArray(tt.in)
		if got != tt.out {
			t.Errorf("Input: %v | Expected: %v, Got: %v", tt.in, tt.out, got)
		}
	}
}
