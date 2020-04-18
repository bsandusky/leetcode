package minPathSum

import "testing"

var tts = []struct {
	in  [][]int
	out int
}{
	{
		in:  [][]int{[]int{1, 3, 1}, []int{1, 5, 1}, []int{4, 2, 1}},
		out: 7,
	},
}

func TestMinPathSum(t *testing.T) {
	for _, tt := range tts {
		got := minPathSum(tt.in)
		if got != tt.out {
			t.Errorf("Input: %v | Expected: %d, Got: %d", tt.in, tt.out, got)
		}
	}
}
