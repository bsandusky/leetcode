package countingElements

import (
	"testing"
)

var tts = []struct {
	in  []int
	out int
}{
	{in: []int{1, 2, 3}, out: 2},
	{in: []int{1, 1, 3, 3, 5, 5, 7, 7}, out: 0},
	{in: []int{1, 3, 2, 3, 5, 0}, out: 3},
	{in: []int{1, 1, 2, 2}, out: 2},
}

func TestCountElements(t *testing.T) {
	for _, tt := range tts {
		got := countElements(tt.in)
		if got != tt.out {
			t.Errorf("Input: %v | Expected: %d, Got: %d", tt.in, tt.out, got)
		}
	}
}
