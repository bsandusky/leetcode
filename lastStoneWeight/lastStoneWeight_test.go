package lastStoneWeight

import "testing"

var tts = []struct {
	in  []int
	out int
}{
	{in: []int{2, 7, 4, 1, 8, 1}, out: 1},
	{in: []int{2, 2}, out: 0},
}

func TestLastStoneWeight(t *testing.T) {
	for _, tt := range tts {
		got := lastStoneWeight(tt.in)
		if got != tt.out {
			t.Errorf("Input: %v | Expected: %d, Got: %d", tt.in, tt.out, got)
		}
	}
}
