package longestCommonSubsequence

import "testing"

var tts = []struct {
	a, b string
	out  int
}{
	{
		a:   "abcde",
		b:   "abc",
		out: 3,
	},
	{
		a:   "abcde",
		b:   "abcdfghijklmn",
		out: 4,
	},
	{
		a:   "abc",
		b:   "abc",
		out: 3,
	},
	{
		a:   "abc",
		b:   "def",
		out: 0,
	},
	{
		a:   "psnw",
		b:   "vozsh",
		out: 1,
	},
}

func TestLongestCommonSubsequence(t *testing.T) {
	for _, tt := range tts {
		got := longestCommonSubsequence(tt.a, tt.b)
		if got != tt.out {
			t.Errorf("Input: %s, %s| Expected: %d, Got: %d", tt.a, tt.b, tt.out, got)
		}
	}
}
