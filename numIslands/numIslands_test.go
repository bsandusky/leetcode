package numIslands

import "testing"

var tts = []struct {
	in  [][]byte
	out int
}{
	{in: [][]byte{[]byte{1, 1, 1, 1, 0}, []byte{1, 1, 0, 1, 0}, []byte{1, 1, 0, 0, 0}, []byte{0, 0, 0, 0, 0}}, out: 1},
	{in: [][]byte{[]byte{1, 1, 0, 0, 0}, []byte{1, 1, 0, 0, 0}, []byte{0, 0, 1, 0, 0}, []byte{0, 0, 0, 1, 1}}, out: 3},
}

func TestNumIslands(t *testing.T) {
	for _, tt := range tts {
		got := numIslands(tt.in)
		if got != tt.out {
			t.Errorf("Input: %v | Expected: %d, Got: %d", tt.in, tt.out, got)
		}
	}
}
