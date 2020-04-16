package checkValidString

import (
	"testing"
)

var tts = []struct {
	in  string
	out bool
}{
	{
		in:  "()",
		out: true,
	},
	{
		in:  "(*)",
		out: true,
	},
	{
		in:  "(*))",
		out: true,
	},
	{
		in:  "(()()",
		out: false,
	},
	{
		in:  ")(",
		out: false,
	},
	{
		in:  "())(",
		out: false,
	},
	{
		in:  "()((",
		out: false,
	},
	{
		in:  "*(",
		out: false,
	},
	{
		in:  "(*",
		out: true,
	},
}

func TestCheckValidString(t *testing.T) {
	for _, tt := range tts {
		got := checkValidString(tt.in)
		if got != tt.out {
			t.Errorf("Input: %s | Expected: %v, Got: %v", tt.in, tt.out, got)
		}
	}
}
