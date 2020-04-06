package maxProfit

import (
	"testing"
)

var tts = []struct {
	in  []int
	out int
}{
	{in: []int{7, 1, 5, 3, 6, 4}, out: 7},
	{in: []int{1, 2, 3, 4, 5}, out: 4},
	{in: []int{7, 6, 4, 3, 1}, out: 0},
}

func TestMaxProfit(t *testing.T) {
	for _, tt := range tts {
		got := maxProfit(tt.in)
		if got != tt.out {
			t.Errorf("Input: %v | Expected: %d, Got: %d", tt.in, tt.out, got)
		}
	}
}
