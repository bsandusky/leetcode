package backspaceCompare

import (
	"reflect"
	"testing"
)

var tts = []struct {
	a   string
	b   string
	out bool
}{
	{
		a:   "ab#c",
		b:   "ad#c",
		out: true,
	},
	{
		a:   "ab##",
		b:   "c#d#",
		out: true,
	},
	{
		a:   "a##c",
		b:   "#a#c",
		out: true,
	},
	{
		a:   "a#c",
		b:   "b",
		out: false,
	},
}

func TestBackspaceCompare(t *testing.T) {
	for _, tt := range tts {
		got := backspaceCompare(tt.a, tt.b)
		if !reflect.DeepEqual(got, tt.out) {
			t.Errorf("Input: %s, %s | Expected: %t, Got: %t", tt.a, tt.b, tt.out, got)
		}
	}
}
