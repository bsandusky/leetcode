package happyNumber

import (
	"testing"
)

var tts = []struct {
	in  int
	out bool
}{
	{in: 7, out: true},
	{in: 13, out: true},
	{in: 19, out: true},
	{in: -19, out: false},
	{in: 23, out: true},
	{in: 28, out: true},
	{in: 31, out: true},
	{in: 20, out: false},
	{in: 50, out: false},
}

func TestIsHappy(t *testing.T) {
	for _, tt := range tts {
		got := isHappy(tt.in)
		if got != tt.out {
			t.Errorf("Input: %d, Expected: %v, Got: %v", tt.in, tt.out, got)
		}
	}
}
