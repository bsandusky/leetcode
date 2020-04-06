package reverseInt

import (
	"testing"
)

var tts = []struct {
	in  int
	out int
}{
	{in: 123, out: 321},
	{in: -123, out: -321},
	{in: 120, out: 21},
	{in: 1534236469, out: 0},
}

func TestReverse(t *testing.T) {
	for _, tt := range tts {
		got := reverse(tt.in)
		if got != tt.out {
			t.Errorf("Expected: %d, Got: %d", tt.out, got)
		}
	}
}
