package contiguousArray

import "testing"

var tts = []struct {
	in  []int
	out int
}{
	{in: []int{0, 1}, out: 2},
	{in: []int{0, 1, 0}, out: 2},
	{in: []int{0, 0, 0, 1, 1, 1}, out: 6},
	{in: []int{0, 0, 1, 1, 1, 0}, out: 6},
	{in: []int{0, 0, 1, 1, 0}, out: 4},
	{in: []int{0, 1, 1, 0, 1, 1, 1, 0}, out: 4},
}

func TestFindMaxLength(t *testing.T) {
	for _, tt := range tts {
		got := findMaxLength(tt.in)
		if got != tt.out {
			t.Errorf("Input: %v | Expected: %d, Got: %d", tt.in, tt.out, got)
		}
	}
}
